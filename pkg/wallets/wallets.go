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
	Transfer(from string, to string, amount float64) (*Transfer, error)
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

func (ws walletService) Transfer(from string, to string, amount float64) (*Transfer, error) {
	s := ws.storage
	var wFrom, wTo *Wallet
	var err error
	if wFrom, err = s.FetchWallet(from); wFrom != nil {
	}
	if wTo, err = s.FetchWallet(to); wTo != nil {

	}
	if wFrom.Balance < amount {
		return nil, errors.New("Not enough money in source wallet")
	}
	wFrom.Balance -= amount
	wTo.Balance += amount
	t := Transfer{From: from, To: to, Amount: amount, Completed: time.Now().Unix()}
	if _, err = s.SaveWallet(wFrom); err != nil {
		return nil, err
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
		return nil, errors.New("Amount must be positive")
	}
	w.Balance += amount
	if w, err = s.SaveWallet(w); err != nil {
		return nil, err
	}
	return w, nil
}

func (ws walletService) NewWallet(owner string) (*Wallet, error) {
	s := ws.storage
	var l []*Wallet
	var err error
	if l, err = s.ListWallets(owner); err != nil {
		return nil, err
	}
	if len(l) >= 1 {
		return nil, errors.New("Only one wallet per user allowed")
	}
	w := Wallet{Owner: owner, Balance: 0}
	var dbW *Wallet
	if dbW, err = s.SaveWallet(&w); err != nil {
		return nil, err
	}
	return dbW, nil
}
