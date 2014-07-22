#!/usr/bin/env ruby
# Generated by the protocol buffer compiler. DO NOT EDIT!

require 'protocol_buffers'

begin; require 'spec.pb'; rescue LoadError; end

module Main
  # forward declarations
  class GetProjectRootRequest < ::ProtocolBuffers::Message; end
  class GetProjectRootResponse < ::ProtocolBuffers::Message; end
  class GetInstallationRootRequest < ::ProtocolBuffers::Message; end
  class GetInstallationRootResponse < ::ProtocolBuffers::Message; end
  class GetAllStepsRequest < ::ProtocolBuffers::Message; end
  class GetAllStepsResponse < ::ProtocolBuffers::Message; end
  class GetAllSpecsRequest < ::ProtocolBuffers::Message; end
  class GetAllSpecsResponse < ::ProtocolBuffers::Message; end
  class GetStepValueRequest < ::ProtocolBuffers::Message; end
  class GetStepValueResponse < ::ProtocolBuffers::Message; end
  class ErrorResponse < ::ProtocolBuffers::Message; end
  class APIMessage < ::ProtocolBuffers::Message; end

  class GetProjectRootRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetProjectRootRequest"

  end

  class GetProjectRootResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetProjectRootResponse"

    required :string, :projectRoot, 1
  end

  class GetInstallationRootRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetInstallationRootRequest"

  end

  class GetInstallationRootResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetInstallationRootResponse"

    required :string, :installationRoot, 1
  end

  class GetAllStepsRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllStepsRequest"

  end

  class GetAllStepsResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllStepsResponse"

    repeated :string, :steps, 1
  end

  class GetAllSpecsRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllSpecsRequest"

  end

  class GetAllSpecsResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllSpecsResponse"

    repeated ::Main::ProtoSpec, :specs, 1
  end

  class GetStepValueRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetStepValueRequest"

    required :string, :stepText, 1
  end

  class GetStepValueResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetStepValueResponse"

    required :string, :stepValue, 1
    repeated :string, :parameters, 2
  end

  class ErrorResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ErrorResponse"

    required :string, :error, 1
  end

  class APIMessage < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module APIMessageType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.APIMessage.APIMessageType"

      GetProjectRootRequest = 1
      GetProjectRootResponse = 2
      GetInstallationRootRequest = 3
      GetInstallationRootResponse = 4
      GetAllStepsRequest = 5
      GetAllStepResponse = 6
      GetAllSpecsRequest = 7
      GetAllSpecsResponse = 8
      GetStepValueRequest = 9
      GetStepValueResponse = 10
      ErrorResponse = 11
    end

    set_fully_qualified_name "main.APIMessage"

    required ::Main::APIMessage::APIMessageType, :messageType, 1
    required :int64, :messageId, 2
    optional ::Main::GetProjectRootRequest, :projectRootRequest, 3
    optional ::Main::GetProjectRootResponse, :projectRootResponse, 4
    optional ::Main::GetInstallationRootRequest, :installationRootRequest, 5
    optional ::Main::GetInstallationRootResponse, :installationRootResponse, 6
    optional ::Main::GetAllStepsRequest, :allStepsRequest, 7
    optional ::Main::GetAllStepsResponse, :allStepsResponse, 8
    optional ::Main::GetAllSpecsRequest, :allSpecsRequest, 9
    optional ::Main::GetAllSpecsResponse, :allSpecsResponse, 10
    optional ::Main::GetStepValueRequest, :stepValueRequest, 11
    optional ::Main::GetStepValueResponse, :stepValueResponse, 12
    optional ::Main::ErrorResponse, :error, 13
  end

end
