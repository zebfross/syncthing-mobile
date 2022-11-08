// Copyright (C) 2022 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !windows
// +build !windows

package fs

import (
	"github.com/zebfross/syncthing-mobile/lib/protocol"
)

func (f *BasicFilesystem) PlatformData(name string) (protocol.PlatformData, error) {
	return unixPlatformData(f, name)
}
