package json

import (
	// "encoding/json"
	"bitbucket.org/dhontecillas/gowallet/pkg/wallets"
)

type JWallet struct {
	Id      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type JTransferOrder struct {
	From   string  `json:"from_wallet"`
	Amount float64 `json:"amount"`
}

type JTransfer struct {
	Id        string  `json:"id"`
	From      string  `json:"from_wallet"`
	To        string  `json:"to_wallet"`
	Amount    float64 `json:"amount"`
	Completed int64   `json:"completed"`
}

type JErrorDesc struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type JWalletList struct {
	Wallets []*JWallet `json:"wallets"`
}

func EncodeWallet(inW *wallets.Wallet) *JWallet {
	jw := JWallet{inW.Id, inW.Owner, inW.Balance}
	return &jw
}

func EncodeWalletList(inWL []*wallets.Wallet) *JWalletList {
	jwl := JWalletList{make([]*JWallet, len(inWL))}
	for idx, w := range inWL {
		jwl.Wallets[idx] = EncodeWallet(w)
	}
	return &jwl
}

func EncodeTransfer(inT *wallets.Transfer) *JTransfer {
	jt := JTransfer{inT.Id, inT.From, inT.To, inT.Amount, inT.Completed}
	return &jt
}
