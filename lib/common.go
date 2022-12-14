package syncthing

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/zebfross/syncthing-mobile/lib/config"
	"github.com/zebfross/syncthing-mobile/lib/protocol"
)

func (o *SyncthingLib) GetDeviceId() (string, error) {
	// Ensure that we have a certificate and key.

		myID := protocol.NewDeviceID(o.cert.Certificate[0])

		return myID.String(), nil;
}


func (o *SyncthingLib) deviceIdFromString(deviceId string) (config.DeviceConfiguration, bool) {
	id, err := protocol.DeviceIDFromString(deviceId)
	if err != nil {
		return config.DeviceConfiguration{}, false
	}
	device, ok := o.cfg.Device(id)
	if !ok {
		return config.DeviceConfiguration{}, false
	}
	return device, true
}

func (o *SyncthingLib) LinkDeviceId(deviceId string) (bool, error) {
	device := o.cfg.DefaultDevice()
	err := device.DeviceID.UnmarshalText([]byte(deviceId));
	if(err != nil) {
		return false, errors.ErrNotFound
	}
	waiter, err := o.cfg.Modify(func(cfg *config.Configuration) {
			cfg.SetDevice(device)
	})
	if err != nil {
		return false, errors.ErrNotFound
	}
	err = o.waitAndSave(waiter)
	return true, err
}

func (o *SyncthingLib) waitAndSave(waiter config.Waiter) (error) {
	waiter.Wait()
	return o.cfg.Save();
}