package syncthing

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/zebfross/syncthing-mobile/lib/locations"
	"github.com/zebfross/syncthing-mobile/lib/protocol"
	"github.com/zebfross/syncthing-mobile/lib/syncthing"
)

func (o *SyncthingLib) GetDeviceId() (string, error) {
	// Ensure that we have a certificate and key.
	cert, err := syncthing.LoadOrGenerateCertificate(
		locations.Get(locations.CertFile),
		locations.Get(locations.KeyFile),
	)

		myID := protocol.NewDeviceID(cert.Certificate[0])

		return myID.String(), err;
}

func (o *SyncthingLib) LinkDeviceId(deviceId string) (bool, error) {

	return false, errors.ErrNotFound
}