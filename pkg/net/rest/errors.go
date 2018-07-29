package rest

import (
	"errors"
	"net/http"
	"encoding/json"
	enc "github.com/dhontecillas/gowallet/pkg/encoding/json"
	"github.com/dhontecillas/gowallet/pkg/wallets"
)

const (
	ErrForbidden     = "FORBIDDEN"
	ErrInvalidMethod = "INVALID_METHOD"
)

func writeError(w http.ResponseWriter, httpCode int, code string, message string) error {
	je := enc.JErrorDesc{code, message}
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(je)
	return errors.New(code)
}

type HErr struct {
	Code int
	Msg  string
}

var errToHErr = map[string]HErr{
	ErrForbidden:               HErr{403, "Forbiden"},
	ErrInvalidMethod:          HErr{405, "Invalid method"},
	wallets.ErrStorageFailed:  HErr{503, "Storage service not available"},
	wallets.ErrNotAllowed:     HErr{403, "Not allowed"},
	wallets.ErrMaxWallets:     HErr{403, "Max wallets"},
	wallets.ErrNotEnoughMoney: HErr{409, "Not enough money"},
	wallets.ErrInvalidAmount:  HErr{409, "Invalid amount"},
	wallets.ErrWalletNotFound: HErr{404, "Wallet id not found"},
}

// mapError converts from wallet errors to http errors
func mapError(w http.ResponseWriter, err error) error {
	strErr := err.Error()
	if hErr, ok := errToHErr[strErr]; ok {
		writeError(w, hErr.Code, strErr, hErr.Msg)
	}
	writeError(w, 400, "ERROR", err.Error())
	return err
}
