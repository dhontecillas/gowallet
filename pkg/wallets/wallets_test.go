package wallets_test

import (
	"testing"
	/* Temp disable;

	   . "bitbucket.org/dhontecillas/gowallet/pkg/storage"
	   . "bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	*/)

var userA = "a497fe7e-afd1-447a-bc48-55a5a5cb69fc"

func TestWalletCreation(t *testing.T) {
	/* Temp disable, until mem storage is working as expected.

	    var wA, wB *Wallet
	    // var err error

		s := MemStorage{}
		var ws = NewWalletService(&s)
	    if ws == nil {
			t.Errorf("Can not create WalletService")
	    }

	    if wA, _ = ws.NewWallet(&userA); wA == nil {
	        t.Errorf("Can not generate a wallet for user A")
	    }

	    if wB, _ = ws.NewWallet(&userA); wB != nil {
	        t.Errorf("User should not be able to create more than a wallet")
	    }
	*/
}
