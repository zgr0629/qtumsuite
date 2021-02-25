// multi sign address test
package qtumutil_test

import (
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"strings"
	"testing"

	"github.com/zgr0629/qtumsuite/qtumd/chaincfg"
)

func TestMultiAddrTestNet(t *testing.T) {
	privKey1 := "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	privKey2 := "cUe2AEjCFeyd3U43CeWRhwJNKVVSzhhp6n9VXSLCYcUTkihTVdfy"
	wif1, err := btcutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("decode privkey1 %v err %v", privKey1, err)
		return
	}
	pubkey1 := wif1.PrivKey.PubKey().SerializeCompressed()

	wif2, err := btcutil.DecodeWIF(privKey2)
	if err != nil {
		t.Errorf("decode privKey2 %v err %v", privKey2, err)
		return
	}
	pubkey2 := wif2.PrivKey.PubKey().SerializeCompressed()

	var chainParams = &chaincfg.TestNet3AddrParams
	addressHashKeys := make([]*btcutil.AddressPubKey, 2)
	addressHashKeys[0], err = btcutil.NewAddressPubKey(pubkey1, chainParams)
	if err != nil {
		t.Errorf("create new pub key1 err %v", err)
		return
	}

	addressHashKeys[1], err = btcutil.NewAddressPubKey(pubkey2, chainParams)
	if err != nil {
		t.Errorf("create new pub key2 err %v", err)
		return
	}

	pubScript, err := txscript.MultiSigScript(addressHashKeys, 2)
	if err != nil {
		t.Errorf("create MultiSigScript err %v", err)
		return
	}

	addrScriptHash, err := btcutil.NewAddressScriptHash(pubScript, chainParams)
	if err != nil {
		t.Errorf("create NewAddressScriptHash err %v", err)
		return
	}
	expectAddr := "mHJEeR2UXHsGjMn6hVyQ1Si4pJymWpG7cp"
	addr := addrScriptHash.EncodeAddress()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}

}

func TestMultiAddrMainNet(t *testing.T) {
	privKey1 := "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	privKey2 := "cUe2AEjCFeyd3U43CeWRhwJNKVVSzhhp6n9VXSLCYcUTkihTVdfy"
	wif1, err := btcutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("decode privkey1 %v err %v", privKey1, err)
		return
	}
	pubkey1 := wif1.PrivKey.PubKey().SerializeCompressed()

	wif2, err := btcutil.DecodeWIF(privKey2)
	if err != nil {
		t.Errorf("decode privKey2 %v err %v", privKey2, err)
		return
	}
	pubkey2 := wif2.PrivKey.PubKey().SerializeCompressed()

	var chainParams = &chaincfg.MainNetAddrParams
	addressHashKeys := make([]*btcutil.AddressPubKey, 2)
	addressHashKeys[0], err = btcutil.NewAddressPubKey(pubkey1, chainParams)
	if err != nil {
		t.Errorf("create new pub key1 err %v", err)
		return
	}

	addressHashKeys[1], err = btcutil.NewAddressPubKey(pubkey2, chainParams)
	if err != nil {
		t.Errorf("create new pub key2 err %v", err)
		return
	}

	pubScript, err := txscript.MultiSigScript(addressHashKeys, 2)
	if err != nil {
		t.Errorf("create MultiSigScript err %v", err)
		return
	}

	addrScriptHash, err := btcutil.NewAddressScriptHash(pubScript, chainParams)
	if err != nil {
		t.Errorf("create NewAddressScriptHash err %v", err)
		return
	}
	expectAddr := "M923Zu9BvU4hfMQuEKzGuvNs43XAqj7n3q"
	addr := addrScriptHash.EncodeAddress()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}

}
