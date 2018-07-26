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
	NewWallet(owner string) (*Wallet, error)
	ListWallets(owner string) ([]Wallet, error)
	FetchWallet(walletId string) (*Wallet, error)
	DeleteWallet(walletId string) error
	SaveTransfer(t *Transfer) (*Transfer, error)
}

type WalletService struct {
	storage WalletStorage
}

func NewWalletService(ws WalletStorage) *WalletService {
    var wservice = WalletService{ws}
    return &wservice
}

func (ws *WalletService) Transfer(from *Wallet, to *Wallet, currency *string, amount float64) {

}
