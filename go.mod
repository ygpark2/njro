module github.com/ygpark2/njro

go 1.25.0

require (
	entgo.io/contrib v0.7.0
	entgo.io/ent v0.14.4
	github.com/golang/protobuf v1.5.4
	github.com/google/uuid v1.6.0
	github.com/gosimple/slug v1.15.0
	github.com/lib/pq v1.12.3
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/micro/micro/v3 v3.19.0
	github.com/rs/zerolog v1.35.1
	github.com/stretchr/testify v1.12.1
	golang.org/x/crypto v0.54.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.36.11
	gopkg.in/yaml.v3 v3.0.1
)

require (
	ariga.io/atlas v0.31.1-0.20250212144724-069be8033e83 // indirect
	github.com/DATA-DOG/go-sqlmock v1.5.2 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/go-acme/lego/v3 v3.4.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang-jwt/jwt v0.0.0-20210529014511-0f726ea0e725 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/jhump/protoreflect v1.10.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/nightlyone/lockfile v1.0.0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quic-go/quic-go v0.49.1 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	golang.org/x/tools v0.47.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

// exclude github.com/ygpark2/njro v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/njro/shared => ./shared/

replace google.golang.org/grpc => google.golang.org/grpc v1.65.0
