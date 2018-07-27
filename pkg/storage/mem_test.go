package storage_test

import (
	. "bitbucket.org/dhontecillas/gowallet/pkg/storage"
	"bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	"testing"
)

var userA = "a497fe7e-afd1-447a-bc48-55a5a5cb69fc"

func TestMemoryStorage(t *testing.T) {
	var err error
	ms := NewMemStorage()

	emptyW := wallets.Wallet{}
	if _, err = ms.SaveWallet(&emptyW); err == nil {
		t.Errorf("Should no save wallet without owner")
	}

	w0 := wallets.Wallet{Id: "", Owner: "OWNER_ID", Currency: "eur", Balance: 0}
	var dbW0 *wallets.Wallet
	if dbW0, err = ms.SaveWallet(&w0); err != nil {
		t.Errorf("Can not save valid wallet w0")
		return
	}
	if len(dbW0.Id) < 0 {
		t.Errorf("Missing created Id")
		return
	}

	if _, err = ms.SaveWallet(&w0); err == nil {
		t.Errorf("Should not be able to save another wallet for the same user")
		return
	}

	var dbW1 *wallets.Wallet
	w0 = *dbW0
	w0.Balance = 1.0
	w0.Currency = "usd"
	if dbW1, err = ms.SaveWallet(&w0); err != nil {
		t.Errorf("Must be able to update an existing wallet for the user")
		return
	}
	if dbW1.Id != dbW0.Id {
		t.Errorf("The ID should not have changed")
		return
	}
	if dbW1.Balance == dbW0.Balance {
		t.Errorf("The balance should be updated")
		return
	}

	var list []*wallets.Wallet
	if list, err = ms.ListWallets(userA); err != nil {
		t.Errorf("Should be able to list the wallets")
		return
	}
	if len(list) != 1 {
		t.Errorf("List should contain created wallet")
	}
}
