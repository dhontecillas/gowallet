package rest

import (
	enc "bitbucket.org/dhontecillas/gowallet/pkg/encoding/json"
	"bitbucket.org/dhontecillas/gowallet/pkg/storage"
	"bitbucket.org/dhontecillas/gowallet/pkg/wallets"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var serviceAuth AuthService
var serviceStorage storage.TransactionalStorage

func writeError(w http.ResponseWriter, httpCode int, code string, message string) error {
	je := enc.JErrorDesc{code, message}
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(je)
	return errors.New(code)
}

func extractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("No auth header")
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) <= 1 {
		return "", errors.New("No bearer token")
	}
	return splitToken[1], nil
}

// authRequests returns a userId from a request
func authRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	var err error
	var userId string
	var bearerT string
	if bearerT, err = extractBearerToken(r); err != nil {
		writeError(w, 403, "FORBIDDEN", "Can not authorize user")
		return "", err
	}
	if userId, err = serviceAuth.AuthorizeUser(bearerT); err != nil {
		writeError(w, 403, "FORBIDDEN", "Can not authorize user")
		return "", err
	}
	return userId, nil
}

// TODO: Check if already exists an Auth interface
type AuthService interface {
	// authorizeUser returns a user id from an authToken
	AuthorizeUser(authToken string) (string, error)
}

type AllowAllAuthService struct {
}

func (aas *AllowAllAuthService) AuthorizeUser(authToken string) (string, error) {
	return authToken, nil
}

func listWallets(w http.ResponseWriter, r *http.Request, userId string) error {
	var err error
	var userWallets []*wallets.Wallet
	ws := wallets.NewWalletService(serviceStorage)
	if userWallets, err = ws.List(userId); err != nil {
		writeError(w, 400, "OK", "List wallets")
		return err
	}
	je := enc.EncodeWalletList(userWallets)
	json.NewEncoder(w).Encode(je)
	return err
}

func createWallet(w http.ResponseWriter, r *http.Request, userId string) error {
	var err error
	var userW *wallets.Wallet
	ws := wallets.NewWalletService(serviceStorage)
	if userW, err = ws.NewWallet(userId); err != nil {
		writeError(w, 400, "CANT_CREATE", err.Error())
		return err
	}
	je := enc.EncodeWallet(userW)
	json.NewEncoder(w).Encode(je)
	return nil
}

func walletInfo(w http.ResponseWriter, r *http.Request, userId string, walletId string) error {
	writeError(w, 200, "OK", "Wallet Info")
	return nil
}

func deleteWallet(w http.ResponseWriter, r *http.Request, userId string, walletId string) error {
	writeError(w, 200, "OK", "Delete wallet")
	return nil
}

func transferMoney(w http.ResponseWriter, r *http.Request, userId string, walletId string) error {
	writeError(w, 200, "OK", "Transfer money")
	return nil
}

func allWalletsEndpoint(w http.ResponseWriter, r *http.Request, userId string) error {
	var err error
	switch r.Method {
	case "GET":
		err = listWallets(w, r, userId)
	case "POST":
		err = createWallet(w, r, userId)
	default:
		err = writeError(w, 405, "CANT_DO_THAT", "Method not allowed")
	}
	return err
}

func singleWalletEndpoint(w http.ResponseWriter, r *http.Request, userId string, walletId string) error {
	var err error
	switch r.Method {
	case "GET":
		err = walletInfo(w, r, userId, walletId)
	case "PUT":
		err = transferMoney(w, r, userId, walletId)
	case "DELETE":
		err = deleteWallet(w, r, userId, walletId)
	default:
		err = writeError(w, 405, "CANT_DO_THAT", "Method not allowed")
	}
	return err
}

func extractWalletId(r *http.Request) string {
	// this is a dirty way of selecting the route, a package for
	// url regex should be used instead.(Perhaps with gin-gonic?).
	path := strings.SplitAfter(r.URL.Path, "/v1/wallets/")
	if len(path) <= 1 {
		return ""
	}
	return path[1]
}

func WalletsEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error
	var userId string
	if userId, err = authRequest(w, r); err != nil {
		return
	}

	serviceStorage.Begin()
	if wId := extractWalletId(r); wId != "" {
		err = singleWalletEndpoint(w, r, userId, wId)
	} else {
		err = allWalletsEndpoint(w, r, userId)
	}

	if err != nil {
		serviceStorage.Rollback()
	} else {
		serviceStorage.Commit()
	}
}

func NewServer(auth AuthService, storage storage.TransactionalStorage) {
	serviceAuth = auth
	serviceStorage = storage
	http.HandleFunc("/v1/wallets/", WalletsEndpoint)
}
