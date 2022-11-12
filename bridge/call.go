package syncthingBridge

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zebfross/syncthing-mobile/bridge/model"
	syncthing "github.com/zebfross/syncthing-mobile/lib"
)

// Call ...
func Call(name string, payload []byte) ([]byte, error) {

	instance := NewInstance()
	var output []byte = nil
	switch name {
	case "getDeviceId":
		output = instance.getDeviceId()
	case "linkDeviceId":
		output = instance.linkDeviceId(payload);
	default:
		return output, fmt.Errorf("not implemented: %s", name)
	}

	return output, nil;
}

type instance struct {
	instance *syncthing.SyncthingLib
}

func NewInstance() *instance {
	return &instance{instance: syncthing.NewSyncthingLib()}
}

func (m instance) getDeviceId() []byte {
	response := flatbuffers.NewBuilder(0)
	output, err := m.instance.GetDeviceId()
	return m._stringResponse(response, output, err)
}

func (m instance) linkDeviceId(payload []byte) []byte {
	response := flatbuffers.NewBuilder(0)
	request := model.GetRootAsLinkDeviceIdRequest(payload, 0)

	output, err := m.instance.LinkDeviceId(m.toString(request.DeviceId()))
	return m._boolResponse(response, output, err)
}

func (m instance) _boolResponse(response *flatbuffers.Builder, output bool, err error) []byte {
	if err != nil {
		outputOffset := response.CreateString(err.Error())
		model.BoolResponseStart(response)
		model.BoolResponseAddError(response, outputOffset)
		response.Finish(model.BoolResponseEnd(response))
		return response.FinishedBytes()
	}
	model.BoolResponseStart(response)
	model.BoolResponseAddOutput(response, output)
	response.Finish(model.BoolResponseEnd(response))
	return response.FinishedBytes()
}

func (m instance) _bytesResponse(response *flatbuffers.Builder, output []byte, err error) []byte {
	if err != nil {
		outputOffset := response.CreateString(err.Error())
		model.BytesResponseStart(response)
		model.BytesResponseAddError(response, outputOffset)
		response.Finish(model.BytesResponseEnd(response))
		return response.FinishedBytes()
	}
	outputOffset := response.CreateByteVector(output)
	model.BytesResponseStart(response)
	model.BytesResponseAddOutput(response, outputOffset)
	response.Finish(model.BytesResponseEnd(response))
	return response.FinishedBytes()
}

func (m instance) _intResponse(response *flatbuffers.Builder, output int64, err error) []byte {
	if err != nil {
		outputOffset := response.CreateString(err.Error())
		model.IntResponseStart(response)
		model.IntResponseAddError(response, outputOffset)
		response.Finish(model.IntResponseEnd(response))
		return response.FinishedBytes()
	}
	model.IntResponseStart(response)
	model.IntResponseAddOutput(response, output)
	response.Finish(model.IntResponseEnd(response))
	return response.FinishedBytes()
}

func (m instance) _stringResponse(response *flatbuffers.Builder, output string, err error) []byte {
	if err != nil {
		outputOffset := response.CreateString(err.Error())
		model.StringResponseStart(response)
		model.StringResponseAddError(response, outputOffset)
		response.Finish(model.StringResponseEnd(response))
		return response.FinishedBytes()
	}
	outputOffset := response.CreateString(output)
	model.StringResponseStart(response)
	model.StringResponseAddOutput(response, outputOffset)
	response.Finish(model.StringResponseEnd(response))
	return response.FinishedBytes()
}

func (m instance) toString(input []byte) string {
	if input == nil {
		return ""
	}

	return string(input)
}