.PHONY: build


PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell git describe --tags)
COMMIT := $(shell [ -z "${COMMIT_ID}" ] && git log -1 --format='%H' || echo ${COMMIT_ID} )
BUILDTIME := $(shell date -u +"%Y%m%d.%H%M%S" )
DOCKER ?= docker
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
GOFLAGS:=""

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=zetacore \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=zetacored \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=zetaclientd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/zeta-chain/zetacore/common.Version=$(VERSION) \
	-X github.com/zeta-chain/zetacore/common.CommitHash=$(COMMIT) \
	-X github.com/zeta-chain/zetacore/common.BuildTime=$(BUILDTIME)


BUILD_FLAGS := -ldflags '$(ldflags)'
TEST_DIR?="./..."
TEST_BUILD_FLAGS :=  -tags mocknet

clean: clean-binaries clean-dir

clean-binaries:
	@rm -rf ${GOBIN}/zeta*

clean-dir:
	@rm -rf ~/.zetacored

all: install

test-coverage-exclude-core:
	@go test {TEST_BUILD_FLAGS} -v -coverprofile coverage.out $(go list ./... | grep -v /x/zetacore/)

test-coverage:
	@go test ${TEST_BUILD_FLAGS} -v -coverprofile coverage.out ${TEST_DIR}

coverage-report: test-coverage
	@go tool cover -html=cover.txt

test:
	@go test ${TEST_BUILD_FLAGS} ${TEST_DIR}

gosec:
	gosec  -exclude-dir=localnet ./...

install: go.sum
		@echo "--> Installing zetacored & zetaclientd"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/zetacored
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/zetaclientd

install-zetaclient: go.sum
		@echo "--> Installing zetaclientd"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/zetaclientd

# running with race detector on will be slow
install-zetaclient-race-test-only-build: go.sum
		@echo "--> Installing zetaclientd"
		@go install -race -mod=readonly $(BUILD_FLAGS) ./cmd/zetaclientd

install-zetacore: go.sum
		@echo "--> Installing zetacored"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/zetacored

install-indexer: go.sum
		@echo "--> Installing indexer"
		@go install -mod=readonly $(BUILD_FLAGS) ./cmd/indexer

install-mockmpi:
	@echo "--> Installing MockMPI"
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/mockmpi


go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

init: clean install-zetacore
	./localnet/zetachain/standalone-network/init.sh

run:
	./localnet/zetachain/standalone-network/run.sh

init-run: init run

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.3
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen.sh
.PHONY: proto-gen

proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace $(protoImageName) \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;
.PHONY: proto-format

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-any:
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen-any.sh
.PHONY: proto-gen-any

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh
.PHONY: proto-swagger-gen

proto-lint:
	$(DOCKER_BUF) lint --error-format=json
.PHONY: proto-lint

proto-check-breaking:
	# we should turn this back on after our first release
	# $(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master
.PHONY: proto-check-breaking