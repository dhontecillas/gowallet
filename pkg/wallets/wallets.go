package wallets

import (
	"errors"
	"time"
)

type Wallet struct {
	Id      string
	Owner   string
	Balance float64
}

type Transfer struct {
	Id        string
	From      string
	To        string
	Amount    float64
	Completed int64
}

// Repository tha stores information about wallets and transfers
type WalletStorage interface {
	SaveWallet(wallet *Wallet) (*Wallet, error)
	ListWallets(owner string) ([]*Wallet, error)
	FetchWallet(walletId string) (*Wallet, error)
	DeleteWallet(walletId string) error
	SaveTransfer(t *Transfer) (*Transfer, error)
}

type WalletService interface {
	List(owner string) ([]*Wallet, error)
	Info(owner string, walletId string) (*Wallet, error)
	Delete(owner string, walletId string) error
	Transfer(owner string, from string, to string, amount float64) (*Transfer, error)
	NewWallet(owner string) (*Wallet, error)
	Load(to string, amount float64) (*Wallet, error)
}

type walletService struct {
	storage WalletStorage
}

func NewWalletService(ws WalletStorage) WalletService {
	w := walletService{storage: ws}
	return w
}

func (ws walletService) List(owner string) ([]*Wallet, error) {
	s := ws.storage
	return s.ListWallets(owner)
}

func (ws walletService) Info(owner string, walletId string) (*Wallet, error) {
	var err error
	var w *Wallet
	s := ws.storage
	if w, err = s.FetchWallet(walletId); err != nil {
		return nil, &WErr{ErrStorageFailed, err.Error()}
	}
	if w.Owner != owner {
		return nil, &WErr{ErrNotAllowed, ""}
	}
	return w, nil
}

func (ws walletService) Delete(owner string, walletId string) error {
	var err error
	var w *Wallet
	s := ws.storage
	if w, err = s.FetchWallet(walletId); err != nil {
		return err
	}
	if w.Owner != owner {
		return errors.New("Not owned")
	}
	// TODO: we can add a check that the wallet is empty
	return s.DeleteWallet(walletId)
}

func (ws walletService) Transfer(owner string, from string, to string, amount float64) (*Transfer, error) {
	s := ws.storage
	var wFrom, wTo *Wallet
	var err error
	if amount <= 0 {
		return nil, &WErr{ErrInvalidAmount, ""}
	}
	if wFrom, err = s.FetchWallet(from); wFrom == nil {
		return nil, &WErr{ErrWalletNotFound, from}
	}
	if wTo, err = s.FetchWallet(to); wTo == nil {
		return nil, &WErr{ErrWalletNotFound, to}
	}
	if wFrom.Balance < amount {
		return nil, &WErr{ErrNotEnoughMoney, ""}
	}
	wFrom.Balance -= amount
	wTo.Balance += amount
	t := Transfer{From: from, To: to, Amount: amount, Completed: time.Now().Unix()}
	if _, err = s.SaveWallet(wFrom); err != nil {
		return nil, &WErr{ErrStorageFailed, err.Error()}
	}
	if _, err = s.SaveWallet(wTo); err != nil {
		return nil, err
	}
	var completedTransfer *Transfer
	if completedTransfer, err = s.SaveTransfer(&t); err != nil {
		return nil, err
	}
	return completedTransfer, nil
}

// Load puts some amount of money into wallet
func (ws walletService) Load(to string, amount float64) (*Wallet, error) {
	s := ws.storage
	var w *Wallet
	var err error
	if w, err = s.FetchWallet(to); err != nil {
		return nil, err
	}
	if amount <= 0 {
		return nil, &WErr{ErrInvalidAmount, ""}
	}
	w.Balance += amount
	if w, err = s.SaveWallet(w); err != nil {
		return nil, &WErr{ErrStorageFailed, err.Error()}
	}
	return w, nil
}

func (ws walletService) NewWallet(owner string) (*Wallet, error) {
	s := ws.storage
	var l []*Wallet
	var err error
	if l, err = s.ListWallets(owner); err != nil {
		return nil, &WErr{ErrStorageFailed, err.Error()}
	}
	if len(l) >= 1 {
		return nil, &WErr{ErrMaxWallets, "Only one wallet per user allowed"}
	}
	w := Wallet{Owner: owner, Balance: 0}
	var dbW *Wallet
	if dbW, err = s.SaveWallet(&w); err != nil {
		return nil, &WErr{ErrStorageFailed, err.Error()}
	}
	return dbW, nil
}
