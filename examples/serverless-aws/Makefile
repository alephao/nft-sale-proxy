.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/proxy proxy/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean gomodgen build
	npm run deploy

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
	go mod tidy
