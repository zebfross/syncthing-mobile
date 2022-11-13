package syncthing

import (
	"context"
	"fmt"

	"github.com/zebfross/syncthing-mobile/lib/config"
	"github.com/zebfross/syncthing-mobile/lib/events"
	"github.com/zebfross/syncthing-mobile/lib/locations"
	"github.com/zebfross/syncthing-mobile/lib/syncthing"
)

type SyncthingLib struct {
	cfg config.Wrapper
	cancel context.CancelFunc
}

func NewSyncthingLib() (*SyncthingLib, error) {
	lib := &SyncthingLib{}

	var err error
	cert, err := syncthing.LoadOrGenerateCertificate(
		locations.Get(locations.CertFile),
		locations.Get(locations.KeyFile),
	)
	if(err != nil) {
		println(fmt.Sprintf("Error loading certificate %v", err));
		return nil, err
	}
	evLogger := events.NewLogger()
	lib.cfg, err = syncthing.LoadConfigAtStartup(locations.Get(locations.ConfigFile), cert, evLogger, true /*AllowNewerConfig*/, false /*NoDefaultFolder*/, true /*SkipPortProbing*/)
	if err != nil {
		return nil, err
	}

	var ctx context.Context;
	ctx, lib.cancel = context.WithCancel(context.Background())
	go lib.cfg.Serve(ctx)

	return lib, nil
}
