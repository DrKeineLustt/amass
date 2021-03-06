// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package amass

import (
	"testing"
)

func TestNetblockService(t *testing.T) {
	in := make(chan *AmassRequest)
	out := make(chan *AmassRequest)
	srv := NewNetblockService(in, out, DefaultConfig())

	srv.Start()
	in <- &AmassRequest{Address: "104.244.42.65"}

	req := <-out
	if req.Netblock.String() != "104.244.42.0/24" {
		t.Errorf("Address %s instead belongs in netblock %s\n", req.Address, req.Netblock.String())
	}

	srv.Stop()
}
