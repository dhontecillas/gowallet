package wallets

import (
	"testing"
	// "time"
)

type WalletStorageStub struct{}

func (ws *WalletStorageStub) NewWallet(owner string) (*Wallet, error) {
    var w = Wallet{}
    return &w, nil
}

func (ws *WalletStorageStub) ListWallets(owner string) ([]Wallet, error) {
    var wallets = make([]Wallet, 1)
	return wallets, nil
}

func (ws *WalletStorageStub) FetchWallet(walletId string) (*Wallet, error) {
    var w = Wallet{}
    return &w, nil
}

func (ws *WalletStorageStub) DeleteWallet(walletId string) error {
    return nil
}

func (ws *WalletStorageStub) SaveTransfer(t *Transfer) (*Transfer, error) {
    return nil, nil
}

func TestWalletServiceCreation(t *testing.T) {
	wstorage := WalletStorageStub{}
	var ws = NewWalletService(&wstorage)
    if ws == nil {
		t.Errorf("Can not create WalletService")
    }
}
