
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/micro/api/proto/api.proto=github.com/micro/micro/v3/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/post.proto
	@rsync -a github.com/ygpark2/mboard/service/board/posts/proto/ proto && rm -Rf github.com

.PHONY: build
build: proto

	go build -o post-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t post-service:latest
