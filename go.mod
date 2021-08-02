module github.com/ygpark2/mboard

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/asim/go-micro/cmd/protoc-gen-micro/v3 v3.0.0-20210726052521-c3107e6843e2 // indirect
	github.com/asim/go-micro/v3 v3.5.2
	github.com/envoyproxy/protoc-gen-validate v0.3.0
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/google/wire v0.5.0
	github.com/gosimple/slug v1.9.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.1
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.9.0
	github.com/markbates/pkger v0.17.1
	github.com/matryer/is v1.3.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro/v3 v3.3.1-0.20210713153811-bb922eccdbd3
	github.com/micro/services v0.10.0
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.1.0
	github.com/xmlking/micro-starter-kit v0.3.7
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	google.golang.org/genproto v0.0.0-20210726200206-e7812ac95cc0
	google.golang.org/grpc v1.40.0-dev.0.20210708170655-30dfb4b933a5
	google.golang.org/grpc/examples v0.0.0-20210728214646-ad0a2a847cdf // indirect
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
)

// exclude github.com/ygpark2/mboard v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/mboard/shared => ./shared/
