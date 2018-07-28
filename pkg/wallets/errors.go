package wallets

type WErr struct {
	code    string
	message string
}

func (e *WErr) Error() string {
	return e.code
}

const (
	// There has been an error in the storage layer:
	ErrStorageFailed = "STORAGE_FAILED"
	// User can not perform this action:
	ErrNotAllowed = "NOT_ALLOWED"
	// When the user has reached the maximum numer of wallets:
	ErrMaxWallets = "MAX_WALLETS"
	// Not enough money in the wallet:
	ErrNotEnoughMoney = "NOT_ENOUGH_MONEY"
	// The value provided as amount is not valid
	ErrInvalidAmount = "INVALID_AMOUNT"
	// When the wallet can not be found
	ErrWalletNotFound = "WALLET_NOT_FOUND"
)
