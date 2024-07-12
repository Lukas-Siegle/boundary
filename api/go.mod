module github.com/hashicorp/boundary/api

go 1.21

toolchain go1.21.5

replace github.com/hashicorp/boundary/sdk => ../sdk

require (
	github.com/hashicorp/boundary/sdk v0.0.47
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-kms-wrapping/v2 v2.0.14
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-secure-stdlib/base62 v0.1.2
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.8
	github.com/hashicorp/go-secure-stdlib/temperror v0.1.1
	github.com/hashicorp/go-uuid v1.0.3
	github.com/mitchellh/copystructure v1.2.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mr-tron/base58 v1.2.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/atomic v1.11.0
	golang.org/x/time v0.3.0
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.33.0
	nhooyr.io/websocket v1.8.11
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/eventlogger v0.2.6-0.20231025104552-802587e608f0 // indirect
	github.com/hashicorp/eventlogger/filters/encrypt v0.1.8-0.20231025104552-802587e608f0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.5 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/pointerstructure v1.2.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
