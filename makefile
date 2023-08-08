ITCOMMIT := $(shell git rev-parse --short HEAD)
BUILD_FLAGS := -ldflags="-w -s -X github.com/luoruofeng/naval-cli/pkg/version.GITCOMMIT=$(GITCOMMIT)"

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

default: bin

.PHONY: all
all: bin

.PHONY: bin
bin:
	CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -o naval-cli main.go

.PHONY: install
install:
	go install ${BUILD_FLAGS}

# kompile naval-cli for multiple platforms
.PHONY: cross
cross:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-linux-amd64" main.go
	GOOS=linux GOARCH=arm CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-linux-arm" main.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-linux-arm64" main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-windows-amd64.exe" main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-darwin-amd64" main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 GO111MODULE=on go build  ${BUILD_FLAGS} -installsuffix cgo  -o "bin/naval-cli-darwin-arm64" main.go

.PHONY: clean
clean:
	rm -f naval-cli
	rm -r -f bundles

# build docker image that is used for running all test locally
.PHONY: test-image
test-image:
	docker build -t $(TEST_IMAGE) -f script/test_in_container/Dockerfile script/test_in_container/