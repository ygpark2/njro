# mboard
micro board


service/board/pkged.go:9:2: cannot find package "." in:
	/home/ygpark2/pjt/go/github.com/mboard/vendor/github.com/markbates/pkger/pkging/mem

"github.com/markbates/pkger/pkging/mem"
This package is not generated due to the "pkged.go" file not existing.

## 기본 셑업 명령어

make tools
make update_deps
make proto
make wire

## 서비스 실행 명령어

micro run --env_vars CONFIGOR_FILE_PATH=./config/config.yaml service/board


## 이미지 빌드 명령어

make build TYPE=service TARGET=board VERSION=v0.1.1


## 패키지 빌드 명령어

이 명령어를 통하여 config 폴더 아래에 있는 모든 yaml파일들의 설정을 pkged.go파일에 생성 그래서 yaml 파일에 수정이 있다면 이 명령어를 통하여 pkged 파일을 업데이트 해야 한다.

make pkger TYPE=service TARGET=board VERSION=v0.1.1

## 에러

GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev
github.com/uber/prototool/cmd/prototool imports
        github.com/uber/prototool/internal/cmd imports
        github.com/uber/prototool/internal/exec imports
        github.com/uber/prototool/internal/grpc imports
        github.com/fullstorydev/grpcurl imports
        google.golang.org/grpc/xds/experimental: cannot find module providing package google.golang.org/grpc/xds/experimental

prototool 설치시 다음과 같은 에러가 발생할 경우 grpc를 다운 그레이드 해야함
go get google.golang.org/grpc@v1.30.0

최신 버전에서는 xds 모듈이 사라졌음
