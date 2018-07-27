package wallets

type Wallet struct {
	Id       string
	Owner    string
	Currency string
	Balance  float64
}

type Transfer struct {
	Id       string
	Ordered  string
	From     string
	To       string
	Currency string
	Amount   float64
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
	Transfer(from *Wallet, to *Wallet, currency *string, amount float64) (*Transfer, error)
	NewWallet(owner *string) (*Wallet, error)
}

type walletService struct {
	storage WalletStorage
}

func NewWalletService(ws WalletStorage) WalletService {
	w := walletService{storage: ws}
	return w
}

func (ws walletService) Transfer(from *Wallet, to *Wallet, currency *string, amount float64) (*Transfer, error) {
	return nil, nil
}

func (ws walletService) NewWallet(owner *string) (*Wallet, error) {
	return nil, nil
}
