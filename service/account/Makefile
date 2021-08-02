
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/micro/api/proto/api.proto=github.com/micro/micro/v3/api/proto

.PHONY: proto
proto:
    
	# protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/post/post.proto
	for f in ./proto/**/*.proto; do \
		protoc -I vendor/github.com/envoyproxy/protoc-gen-validate \
				-I vendor/github.com/infobloxopen/protoc-gen-gorm \
				--proto_path=.:${GOPATH}/src \
				--go_out=paths=source_relative:. \
				--micro_out=paths=source_relative:. \
				--gorm_out=engine=postgres,enums=string,paths=source_relative:. \
				--validate_out=lang=go,paths=source_relative:. $$f; \
				echo ✓ compiled: $$f; \
	done
	@rsync -a github.com/ygpark2/mboard/service/account/proto/ proto && rm -Rf github.com

.PHONY: build
build: proto

	go build -mod=vendor -o account-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t account-service:latest
