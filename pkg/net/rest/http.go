package rest

import (
	enc "bitbucket.org/dhontecillas/gowallet/pkg/encoding/json"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func writeError(w http.ResponseWriter, httpCode int, code string, message string) {
	je := enc.JErrorDesc{code, message}
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(je)
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
	if userId, err = authS.AuthorizeUser(bearerT); err != nil {
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

var authS AuthService

func listWallets(w http.ResponseWriter, r *http.Request, userId string) {
	writeError(w, 200, "OK", "List wallets")
}

func createWallet(w http.ResponseWriter, r *http.Request, userId string) {
	writeError(w, 200, "OK", "Create Wallet")
}

func walletInfo(w http.ResponseWriter, r *http.Request, userId string, walletId string) {
	writeError(w, 200, "OK", "Wallet Info")
}

func deleteWallet(w http.ResponseWriter, r *http.Request, userId string, walletId string) {
	writeError(w, 200, "OK", "Delete wallet")
}

func transferMoney(w http.ResponseWriter, r *http.Request, userId string, walletId string) {
	writeError(w, 200, "OK", "Transfer money")
}

func allWalletsEndpoint(w http.ResponseWriter, r *http.Request, userId string) {
	switch r.Method {
	case "GET":
		listWallets(w, r, userId)
	case "POST":
		createWallet(w, r, userId)
	default:
		writeError(w, 405, "CANT_DO_THAT", "Method not allowed")
	}
}

func singleWalletEndpoint(w http.ResponseWriter, r *http.Request, userId string, walletId string) {
	switch r.Method {
	case "GET":
		walletInfo(w, r, userId, walletId)
	case "PUT":
		transferMoney(w, r, userId, walletId)
	case "DELETE":
		deleteWallet(w, r, userId, walletId)
	default:
		writeError(w, 405, "CANT_DO_THAT", "Method not allowed")
	}
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
	if wId := extractWalletId(r); wId != "" {
		singleWalletEndpoint(w, r, userId, wId)
	} else {
		allWalletsEndpoint(w, r, userId)
	}
}

func NewServer(auth AuthService) {
	authS = auth
	http.HandleFunc("/v1/wallets/", WalletsEndpoint)
}
