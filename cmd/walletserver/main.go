package main

import (
	"bitbucket.org/dhontecillas/gowallet/pkg/net/rest"
	"net/http"
)

func main() {
	auth := rest.AllowAllAuthService{}
	rest.NewServer(&auth)
	http.ListenAndServe(":8000", nil)
}
