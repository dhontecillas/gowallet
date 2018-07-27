package wallets_test

import (
	. "bitbucket.org/dhontecillas/gowallet/pkg/storage"
	. "bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	"testing"
    "time"
    "fmt"
)

var userA = "a497fe7e-afd1-447a-bc48-55a5a5cb69fc"
var userB = "df703dbc-74fa-49d9-8759-ad53da71db02"

func TestWalletCreation(t *testing.T) {
	var wA, wB *Wallet
	var err error

	s := NewMemStorage()
	var ws = NewWalletService(s)
	if ws == nil {
		t.Errorf("Can not create WalletService")
        return
	}

	if wA, _ = ws.NewWallet(userA); wA == nil {
		t.Errorf("Can not create a wallet for user A")
        return
	}

    var wA2 *Wallet
    if wA2, _ = ws.NewWallet(userA); wA2 != nil {
        t.Errorf("User can only have a single account")
        return
    }

	if wB, err = ws.NewWallet(userB); err != nil || wB == nil {
		t.Errorf("Can not create a wallet for user B")
        return
	}

    if wA, err = ws.Load(wA.Id, 20.0); err != nil {
        t.Errorf("Can not load money to wallet")
        return
    }

    var wt *Transfer
    var before = time.Now().Unix()
    fmt.Printf("%v %v", wA, wB)
    if wt, err = ws.Transfer(wA.Id, wB.Id, 10.0); err != nil || wt == nil{
        t.Errorf("Can not transfer money")
        return
    }
    var after = time.Now().Unix()

    if wt.From != wA.Id {
        t.Errorf("Wrong From field in transfer record")
        return
    }
    if wt.To != wB.Id {
        t.Errorf("Wrong To field in transfer record")
        return
    }
    if wt.Amount != 10.0 {
        t.Errorf("Wrong amount in transfer")
        return
    }
    if wt.Completed < before || wt.Completed > after {
        t.Errorf("Wrong completion time")
        return
    }

    // Test we can not transfer more than the existing amount
}
