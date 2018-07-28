package json

import (
// "encoding/json"
)

type JWallet struct {
	Id      string `json:"id"`
	Owner   string `json:"owner"`
	Balance string `json:"balance"`
}

type JTransferOrder struct {
	From   string  `json:"from_wallet"`
	Amount float64 `json:"amount"`
}

type JTransfer struct {
	Id        string  `json:"id"`
	From      string  `json:"from_wallet"`
	To        string  `json:"to_waller"`
	Amount    float64 `json:"amount"`
	Completed int64   `json:"completed"`
}

type JErrorDesc struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
