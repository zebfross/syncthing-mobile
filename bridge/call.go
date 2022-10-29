package syncthingBridge

import (
	"fmt"

	syncthing "github.com/zebfross/syncthing-mobile/syncthing"
)

// Call ...
func Call(name string, payload []byte) ([]byte, error) {

	NewInstance()
	var output []byte = nil
	switch name {
	default:
		return output, fmt.Errorf("not implemented: %s", name)
	}
}

type instance struct {
	instance *syncthing.SyncthingLib
}

func NewInstance() *instance {
	return &instance{instance: syncthing.NewSyncthingLib()}
}

/*func (m instance) decrypt(payload []byte) []byte {
	response := flatbuffers.NewBuilder(0)
	request := model.GetRootAsDecryptRequest(payload, 0)

	output, err := m.instance.Decrypt(m.toString(request.Message()), m.toString(request.PrivateKey()), m.toString(request.Passphrase()), m.parseKeyOptions(request.Options(nil)))
	return m._stringResponse(response, output, err)
}*/
