build:
	go build ./...

test: pkg/wallets/*_test.go
	go test ./... -cover -coverprofile=coverprof.out
	go tool cover -html=coverprof.out -o coverprof.html
