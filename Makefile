default: deps
	go install
	cp ./config.toml "${GOPATH}/bin/config.toml"

deps:
	go mod tidy
	go mod vendor

clean:
	rm -rf vendor/
	rm -rf go.sum

.PHONY: go test