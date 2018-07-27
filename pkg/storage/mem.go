package storage

import (
	"bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	"errors"
)

type MemStorage struct {
	wallets       map[string]*wallets.Wallet
	transfersFrom map[string]*wallets.Transfer
	walletIds     int64
	transferIds   int64
}

func NewMemStorage() *MemStorage {
	ms := MemStorage{map[string]*wallets.Wallet{},
		map[string]*wallets.Transfer{},
		1, 1}
	return &ms
}

func (ms *MemStorage) SaveWallet(inW *wallets.Wallet) (*wallets.Wallet, error) {
	if len(inW.Owner) == 0 {
		return nil, errors.New("No owner id for wallet")
	}
	var existing *wallets.Wallet
	var ok bool
	w := *inW
	if existing, ok = ms.wallets[inW.Id]; ok {
		w.Id = existing.Id
	} else {
		w.Id = string(ms.walletIds)
		ms.walletIds += 1
	}
	ms.wallets[w.Id] = &w
	return &w, nil
}

func (ms *MemStorage) ListWallets(owner string) ([]*wallets.Wallet, error) {
	var wallets = make([]*wallets.Wallet, 0)
	for _, w := range ms.wallets {
		if w.Owner == owner {
			wallets = append(wallets, w)
		}
	}
	return wallets, nil
}

func (ms *MemStorage) FetchWallet(walletId string) (*wallets.Wallet, error) {
	w, ok := ms.wallets[walletId]
	if !ok {
		return nil, nil
	}
	return w, nil
}

func (ms *MemStorage) DeleteWallet(walletId string) error {
	if w, _ := ms.FetchWallet(walletId); w == nil {
		return errors.New("Not found")
	}
	delete(ms.wallets, walletId)
	return nil
}

func (ms *MemStorage) SaveTransfer(inT *wallets.Transfer) (*wallets.Transfer, error) {
	t := *inT
	t.Id = string(ms.transferIds)
	ms.transferIds += 1
	ms.transfersFrom[t.From] = &t
	return &t, nil
}