default: deps
	go install
	cp ./config.toml "${GOPATH}/bin/config.toml"

deps:
	go mod tidy
	go mod vendor

run: default
	cd ${GOPATH}/bin;./go-live

docker: deps
	go install
	cp ${GOPATH}/bin/go-live ./main
	sudo docker build -t go-live .

clean:
	rm -rf vendor/
	rm -rf go.sum

.PHONY: go test