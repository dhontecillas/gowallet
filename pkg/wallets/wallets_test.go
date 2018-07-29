package wallets_test

import (
	. "github.com/dhontecillas/gowallet/pkg/storage"
	. "github.com/dhontecillas/gowallet/pkg/wallets"
	"fmt"
	"testing"
	"time"
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
	if wA2, err = ws.NewWallet(userA); wA2 != nil {
		t.Errorf("User can only have a single account")
		return
	}
	if err.Error() != ErrMaxWallets {
		t.Errorf("Expected error %s (err=%s)", ErrMaxWallets, err.Error())
	}

	if wB, err = ws.NewWallet(userB); err != nil || wB == nil {
		t.Errorf("Can not create a wallet for user B")
		return
	}

	if _, err = ws.Load(wA.Id, -10.0); err == nil {
		t.Errorf("Can not load negative money")
		return
	}

	if wA, err = ws.Load(wA.Id, 20.0); err != nil {
		t.Errorf("Can not load money to wallet")
		return
	}

	var wt *Transfer
	var before = time.Now().Unix()
	fmt.Printf("%v %v", wA, wB)
	if wt, err = ws.Transfer(userA, wA.Id, wB.Id, 10.0); err != nil || wt == nil {
		t.Errorf("Can not transfer money %s", err.Error())
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

	// Test we not transfer negative amount (steal from other wallet)
	if wt, err = ws.Transfer(userA, wA.Id, wB.Id, -10.0); err == nil || wt != nil {
		t.Errorf("Should not be able to transfer negative amount")
		return
	}
	if err.Error() != ErrInvalidAmount {
		t.Errorf("Expected error %s (err=%s)", ErrInvalidAmount, err.Error())
		return
	}

	// Test we can not transfer more than the existing amount
	if wt, err = ws.Transfer(userA, wA.Id, wB.Id, 200.0); err == nil || wt != nil {
		t.Errorf("Should not be able to transfer more than what's in the wallet")
		return
	}
	if err.Error() != ErrNotEnoughMoney {
		t.Errorf("Expected error %s (err=%s)", ErrNotEnoughMoney, err.Error())
		return
	}

	// Get info by id
	if _, err = ws.Info(userA, "non_existing_id"); err == nil {
		t.Errorf("Expected error when can not find wallet")
		return
	}
	if _, err = ws.Info(userA, wB.Id); err == nil {
		t.Errorf("Expected error when trying to get info on other user wallet")
		return
	}
	if _, err = ws.Info(userA, wA.Id); err != nil {
		t.Errorf("User should be able to get info about his wallet")
		return
	}

	// List wallets
	var allW []*Wallet
	if allW, err = ws.List(userA); err != nil {
		t.Errorf("User should be able to get info about all his wallet")
		return
	}
	if len(allW) != 1 {
		t.Errorf("Expected one wallet in the list")
		return
	}

	// Test we can delete an existing wallet
	if err = ws.Delete(userA, wA.Id); err != nil {
		t.Errorf("Cant delete account")
		return
	}

	if err = ws.Delete(userA, wA.Id); err == nil {
		t.Errorf("Should not be able to delete existing account")
		return
	}

}
