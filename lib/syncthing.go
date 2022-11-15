package syncthing

import (
	"context"
	"crypto/tls"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/thejerf/suture/v4"
	"github.com/zebfross/syncthing-mobile/lib/config"
	"github.com/zebfross/syncthing-mobile/lib/db/backend"
	"github.com/zebfross/syncthing-mobile/lib/events"
	"github.com/zebfross/syncthing-mobile/lib/fs"
	"github.com/zebfross/syncthing-mobile/lib/locations"
	"github.com/zebfross/syncthing-mobile/lib/logger"
	"github.com/zebfross/syncthing-mobile/lib/svcutil"
	"github.com/zebfross/syncthing-mobile/lib/syncthing"
)

type SyncthingLib struct {
	app *syncthing.App
	cfg config.Wrapper
	ldb backend.Backend
	cert tls.Certificate
	cancel context.CancelFunc
}

var (
	l = logger.DefaultLogger.NewFacility("main", "Main package")
)

func NewSyncthingLib() (*SyncthingLib, error) {	
	// Ensure that we have a certificate and key.
	cert, err := syncthing.LoadOrGenerateCertificate(
		locations.Get(locations.CertFile),
		locations.Get(locations.KeyFile),
	)
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// earlyService is a supervisor that runs the services needed for or
	// before app startup; the event logger, and the config service.
	spec := svcutil.SpecWithDebugLogger(l)
	earlyService := suture.New("early", spec)
	earlyService.ServeBackground(ctx)

	evLogger := events.NewLogger()
	earlyService.Add(evLogger)

	cfgWrapper, err := syncthing.LoadConfigAtStartup(locations.Get(locations.ConfigFile), cert, evLogger, true, true, true)
	if err != nil {
		os.Exit(svcutil.ExitError.AsInt())
	}
	earlyService.Add(cfgWrapper)

	dbFile := locations.Get(locations.Database)
	ldb, err := syncthing.OpenDBBackend(dbFile, cfgWrapper.Options().DatabaseTuning)
	if err != nil {
		os.Exit(1)
	}

	appOpts := syncthing.Options{
	}

	if t := os.Getenv("STDEADLOCKTIMEOUT"); t != "" {
		secs, _ := strconv.Atoi(t)
		appOpts.DeadlockTimeoutS = secs
	}
	if dur, err := time.ParseDuration(os.Getenv("STRECHECKDBEVERY")); err == nil {
		appOpts.DBRecheckInterval = dur
	}
	if dur, err := time.ParseDuration(os.Getenv("STGCINDIRECTEVERY")); err == nil {
		appOpts.DBIndirectGCInterval = dur
	}

	app, err := syncthing.New(cfgWrapper, ldb, evLogger, cert, appOpts)
	if err != nil {
		os.Exit(svcutil.ExitError.AsInt())
	}

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := app.Start(); err != nil {
		return nil, errors.ErrNotFound
	}

	cleanConfigDirectory()

	lib := &SyncthingLib{
		app: 	app,
		cfg: 	cfgWrapper,
		ldb: 	ldb,
		cert:	cert,
		cancel: cancel,
	}

	return lib, nil
}


// cleanConfigDirectory removes old, unused configuration and index formats, a
// suitable time after they have gone out of fashion.
func cleanConfigDirectory() {
	patterns := map[string]time.Duration{
		"panic-*.log":        7 * 24 * time.Hour,  // keep panic logs for a week
		"audit-*.log":        7 * 24 * time.Hour,  // keep audit logs for a week
		"index":              14 * 24 * time.Hour, // keep old index format for two weeks
		"index-v0.11.0.db":   14 * 24 * time.Hour, // keep old index format for two weeks
		"index-v0.13.0.db":   14 * 24 * time.Hour, // keep old index format for two weeks
		"index*.converted":   14 * 24 * time.Hour, // keep old converted indexes for two weeks
		"config.xml.v*":      30 * 24 * time.Hour, // old config versions for a month
		"*.idx.gz":           30 * 24 * time.Hour, // these should for sure no longer exist
		"backup-of-v0.8":     30 * 24 * time.Hour, // these neither
		"tmp-index-sorter.*": time.Minute,         // these should never exist on startup
		"support-bundle-*":   30 * 24 * time.Hour, // keep old support bundle zip or folder for a month
	}

	for pat, dur := range patterns {
		fs := fs.NewFilesystem(fs.FilesystemTypeBasic, locations.GetBaseDir(locations.ConfigBaseDir))
		files, err := fs.Glob(pat)
		if err != nil {
			continue
		}

		for _, file := range files {
			info, err := fs.Lstat(file)
			if err != nil {
				continue
			}

			if time.Since(info.ModTime()) > dur {
				if err = fs.RemoveAll(file); err != nil {
				} else {
				}
			}
		}
	}
}

