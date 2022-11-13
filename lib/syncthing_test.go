package syncthing

import (
	"testing"
)

func TestSyncthingLib_Complete(t *testing.T) {

	lib, err := NewSyncthingLib()

	if(err != nil) {
		t.Errorf("%v", err);
	}

	deviceId, err := lib.GetDeviceId();

	if (err != nil && deviceId != "") {
		t.Errorf("%v", err);
	}

	success, err := lib.LinkDeviceId("XFEJ326-ETMMXHG-YVHQVEY-GR4Y5RG-CEV3LNK-F7NLFVO-PYWDDC7-G2HIZAP");

	if (err != nil && success) {
		t.Errorf("%v", err);
	}


}