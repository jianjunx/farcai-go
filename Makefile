.PHONY: build run image


build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main ./main.go

image:
	docker build -t farcai-serve .

run:
	air -d
