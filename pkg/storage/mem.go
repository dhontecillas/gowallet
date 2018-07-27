package storage

import (
	"bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	"errors"
)

type MemStorage struct {
	wallets       map[string]*wallets.Wallet
	transfersFrom map[string]*wallets.Transfer
	transfersTo   map[string]*wallets.Transfer
}

func NewMemStorage() *MemStorage {
	ms := MemStorage{map[string]*wallets.Wallet{},
		map[string]*wallets.Transfer{},
		map[string]*wallets.Transfer{}}
	return &ms
}

func (ms *MemStorage) SaveWallet(inW *wallets.Wallet) (*wallets.Wallet, error) {
	if len(inW.Owner) == 0 {
		return nil, errors.New("No owner id for wallet")
	}
	var existing *wallets.Wallet
	var ok bool
	if existing, ok = ms.wallets[inW.Owner]; ok && existing.Id != inW.Id {
		return nil, errors.New("Only one wallet per user allowed")
	}
	w := *inW
	if existing == nil {
		w.Id = string(len(ms.wallets))
	} else {
		w.Id = existing.Id
	}
	ms.wallets[w.Owner] = &w
	return &w, nil
}

func (ms *MemStorage) ListWallets(owner string) ([]*wallets.Wallet, error) {
	var wallets = make([]*wallets.Wallet, 1)
	for _, w := range ms.wallets {
		if w.Owner == owner {
			wallets = append(wallets, w)
		}
	}
	return wallets, nil
}

func (ms *MemStorage) FetchWallet(walletId string) (*wallets.Wallet, error) {
	for _, w := range ms.wallets {
		if w.Id == walletId {
			return w, nil
		}
	}
	return nil, nil
}

func (ms *MemStorage) DeleteWallet(walletId string) error {
	return errors.New("Not implemented")
}

func (ms *MemStorage) SaveTransfer(t *wallets.Transfer) (*wallets.Transfer, error) {
	return nil, errors.New("Not implemented")
}
