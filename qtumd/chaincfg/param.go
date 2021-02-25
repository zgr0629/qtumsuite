package chaincfg

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/zgr0629/qtumsuite/qtumd/wire"
)

// MainNetAddrParams defines the network parameters for qutm addr of the main network .
var MainNetAddrParams = chaincfg.Params{
	Name: "mainnet",
	Net:  wire.MainNet,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "qc", // always bc for main net

	// Address encoding magics
	PubKeyHashAddrID: 58,   //
	ScriptHashAddrID: 50,   //
	PrivateKeyID:     0x80, // starts with 5 (uncompressed) or K (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 0,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3AddrParams = chaincfg.Params{
	Name: "testnet3",
	Net:  wire.TestNet3,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tq", // always tb for test net

	// Address encoding magics
	PubKeyHashAddrID: 120, //
	ScriptHashAddrID: 110, //
	PrivateKeyID:     239, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}
