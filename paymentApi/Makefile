

.PHONY: proto
proto:
	sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=5CB0645D0EDFC5DA  cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/paymentApi/paymentApi.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o paymentApi-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t paymentApi-api:latest
