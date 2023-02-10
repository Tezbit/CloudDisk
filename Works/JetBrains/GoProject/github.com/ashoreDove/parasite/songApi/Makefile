

.PHONY: proto
proto:
	sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=xxxxxx cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/songApi/songApi.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o songApi-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t songApi-api:latest
