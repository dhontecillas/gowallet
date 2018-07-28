build:
	go build ./...

test: pkg/wallets/*_test.go
	go test ./... -cover -coverprofile=coverprof.out
	go tool cover -html=coverprof.out -o coverprof.html

dockerimg:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o walletserver ./cmd/walletserver
	docker build -t dhontecillas/gowallet .

dockerrun:
	docker run -p 8000:8000 -d dhontecillas/gowallet
