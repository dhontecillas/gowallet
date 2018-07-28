package main

import (
	"bitbucket.org/dhontecillas/gowallet/pkg/net/rest"
	"bitbucket.org/dhontecillas/gowallet/pkg/storage"
	"log"
	"net/http"
)

func main() {
	auth := rest.AllowAllAuthService{}
	storage := storage.NewMemStorage()
	rest.NewServer(&auth, storage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
