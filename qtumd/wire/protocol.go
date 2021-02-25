// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"github.com/btcsuite/btcd/wire"
)

// Constants used to indicate the message bitcoin network.  They can also be
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
	// pchMessageStart[0] = 0xf1;
	// pchMessageStart[1] = 0xcf;
	// pchMessageStart[2] = 0xa6;
	// pchMessageStart[3] = 0xd3;
	// MainNet represents the main qtum network.
	MainNet wire.BitcoinNet = 0xd3a6cff1

	// pchMessageStart[0] = 0x0d;
	// pchMessageStart[1] = 0x22;
	// pchMessageStart[2] = 0x15;
	// pchMessageStart[3] = 0x06;
	// TestNet3 represents the test network (version 3).
	TestNet3 wire.BitcoinNet = 0x0615220d
)
