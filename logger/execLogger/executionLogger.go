// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package execLogger

import (
	"fmt"
	"github.com/getgauge/gauge/formatter"
	"github.com/getgauge/gauge/gauge_messages"
	"github.com/getgauge/gauge/logger"
	"github.com/getgauge/gauge/parser"
	"github.com/wsxiaoys/terminal"
	"os"
	"strings"
)

type ExecutionLogger interface {
	Write([]byte) (int, error)
	Text(string)
	PrintError(string)
	SpecHeading(string)
	ScenarioHeading(string)
	Comment(*parser.Comment)
	Step(*parser.Step)
	StepStarting(*parser.Step)
	StepFinished(*parser.Step, bool)
	Table(*parser.Table)
	Critical(string, ...interface{})
	Warning(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	ConceptStarting(*gauge_messages.ProtoConcept)
	ConceptFinished(*gauge_messages.ProtoConcept)
}

var currentLogger ExecutionLogger

var SimpleConsoleOutput = false

type coloredLogger struct {
	linesAfterLastStep int
	isInsideStep       bool
	indentation        int
}

func newColoredConsoleWriter() *coloredLogger {
	return &coloredLogger{linesAfterLastStep: 0, isInsideStep: false, indentation: 0}
}

func Current() ExecutionLogger {
	if currentLogger == nil {
		if SimpleConsoleOutput {
			currentLogger = newSimpleConsoleWriter()
		} else {
			currentLogger = newColoredConsoleWriter()
		}
	}
	return currentLogger
}

func (writer *coloredLogger) Write(b []byte) (int, error) {
	message := indent(string(b), writer.indentation)
	if writer.isInsideStep {
		writer.linesAfterLastStep += strings.Count(message, "\n")
	}
	fmt.Print(message)
	return len(b), nil
}

func (writer *coloredLogger) Text(value string) {
	writer.Write([]byte(value))
}

func (writer *coloredLogger) PrintError(value string) {
	if writer.isInsideStep {
		writer.linesAfterLastStep += strings.Count(value, "\n")
	}
	terminal.Stdout.Colorf("@r%s", value)
}

func (writer *coloredLogger) Critical(formatString string, args ...interface{}) {
	logger.Log.Critical(formatString, args...)
}

func (writer *coloredLogger) Info(formatString string, args ...interface{}) {
	logger.Log.Info(formatString, args...)
}

func (writer *coloredLogger) Warning(formatString string, args ...interface{}) {
	logger.Log.Warning(formatString, args...)
}

func (writer *coloredLogger) Debug(formatString string, args ...interface{}) {
	logger.Log.Debug(formatString, args...)
}

func (writer *coloredLogger) Error(formatString string, args ...interface{}) {
	logger.Log.Error(formatString, args...)
}

func (writer *coloredLogger) SpecHeading(heading string) {
	formattedHeading := formatter.FormatSpecHeading(heading)
	writer.Write([]byte(formattedHeading))
}

func (writer *coloredLogger) Comment(comment *parser.Comment) {
	writer.Write([]byte(formatter.FormatComment(comment)))
}

func (writer *coloredLogger) ScenarioHeading(scenarioHeading string) {
	formattedHeading := formatter.FormatScenarioHeading(scenarioHeading)
	writer.Write([]byte(fmt.Sprintf("\n%s", formattedHeading)))
}

func (writer *coloredLogger) writeContextStep(step *parser.Step) {
	writer.Step(step)
}

func (writer *coloredLogger) Step(step *parser.Step) {
	stepText := formatter.FormatStep(step)
	terminal.Stdout.Colorf("@b%s", stepText)
	writer.isInsideStep = true
	writer.linesAfterLastStep = 0
}

func (writer *coloredLogger) ConceptStarting(protoConcept *gauge_messages.ProtoConcept) {
	conceptText := indent(formatter.FormatConcept(protoConcept), writer.indentation)
	terminal.Stdout.Colorf("@b%s", conceptText)
	writer.indentation += 4
}

func (writer *coloredLogger) ConceptFinished(protoConcept *gauge_messages.ProtoConcept) {
	writer.indentation -= 4
}

func (writer *coloredLogger) StepStarting(step *parser.Step) {
	stepText := formatter.FormatStep(step)
	terminal.Stdout.Colorf("@b%s", stepText)
	writer.isInsideStep = true
	writer.linesAfterLastStep = 0
}

//todo: pass protostep instead
func (writer *coloredLogger) StepFinished(step *parser.Step, failed bool) {
	stepText := indent(formatter.FormatStep(step), writer.indentation)
	linesInStepText := strings.Count(stepText, "\n")
	if linesInStepText == 0 {
		linesInStepText = 1
	}
	linesToMoveUp := writer.linesAfterLastStep + linesInStepText
	terminal.Stdout.Up(linesToMoveUp)
	if failed {
		terminal.Stdout.Colorf("@r%s", stepText)
	} else {
		terminal.Stdout.Colorf("@g%s", stepText)
	}
	terminal.Stdout.Down(linesToMoveUp)
	writer.isInsideStep = false
}

func (writer *coloredLogger) Table(table *parser.Table) {
	formattedTable := indent(formatter.FormatTable(table), writer.indentation)
	terminal.Stdout.Colorf("@m%s", formattedTable)
}

func indent(message string, indentation int) string {
	if indentation == 0 {
		return message
	}
	lines := strings.Split(message, "\n")
	prefixedLines := make([]string, 0)
	spaces := getEmptySpacedString(indentation)
	for i, line := range lines {
		if (i == len(lines)-1) && line == "" {
			prefixedLines = append(prefixedLines, line)
		} else {
			prefixedLines = append(prefixedLines, spaces+line)
		}
	}
	return strings.Join(prefixedLines, "\n")
}

func getEmptySpacedString(numOfSpaces int) string {
	text := ""
	for i := 0; i < numOfSpaces; i++ {
		text += " "
	}
	return text
}

func CriticalError(err error) {
	Current().Critical(err.Error())
	os.Exit(1)
}
