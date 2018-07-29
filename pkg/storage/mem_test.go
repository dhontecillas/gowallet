package storage_test

import (
	. "github.com/dhontecillas/gowallet/pkg/storage"
	"github.com/dhontecillas/gowallet/pkg/wallets"
	"testing"
)

var userA = "a497fe7e-afd1-447a-bc48-55a5a5cb69fc"

func TestMemoryWalletStorage(t *testing.T) {
	var err error
	ms := NewMemStorage()

	emptyW := wallets.Wallet{}
	if _, err = ms.SaveWallet(&emptyW); err == nil {
		t.Errorf("Should no save wallet without owner")
	}

	w0 := wallets.Wallet{Id: "", Owner: userA, Balance: 0}
	var dbW0 *wallets.Wallet
	if dbW0, err = ms.SaveWallet(&w0); err != nil {
		t.Errorf("Can not save valid wallet w0")
		return
	}
	if len(dbW0.Id) < 0 {
		t.Errorf("Missing created Id")
		return
	}

	var dbW1 *wallets.Wallet
	w0 = *dbW0
	w0.Balance = 1.0
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
		t.Errorf("Should be able to list the wallets %d", len(list))
		return
	}
	if len(list) != 1 || list[0] == nil {
		t.Errorf("List should contain created wallet")
		return
	}
	var fw *wallets.Wallet
	if fw, err = ms.FetchWallet(list[0].Id); err != nil || fw == nil {
		t.Errorf("Should be able to fetch walled by id")
		return
	}
	if fw.Id != list[0].Id {
		t.Errorf("Id Mismatch")
		return
	}

	if err = ms.DeleteWallet(fw.Id); err != nil {
		t.Errorf("Wallet should be deletable")
		return
	}

	if err = ms.DeleteWallet(fw.Id); err == nil {
		t.Errorf("Non existing wallet should not be deleted")
		return
	}
}

func TestMemoryTransferStorage(t *testing.T) {
	var err error
	var wt *wallets.Transfer
	ms := NewMemStorage()

	emptyT := wallets.Transfer{From: "FROM"}

	if wt, err = ms.SaveTransfer(&emptyT); wt == nil || err != nil {
		t.Errorf("Can not save transfer")
		return
	}
	if wt.From != emptyT.From {
		t.Errorf("Data from transfer is not saved")
		return
	}
}
