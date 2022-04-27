GOCMD := go
GOTEST := $(GOCMD) test
GOBUILD := $(GOCMD) build

.PHONY: build test cover vet staticcheck gofmt gofumpt clean check lint ci

build:
	$(GOBUILD)

test:
	$(GOTEST) ./...

cover:
	$(GOTEST) -coverprofile fsutil.out ./...

vet:
	$(GOCMD) vet ./...

staticcheck:
	staticcheck ./...

gofmt:
	@echo "gofmt -l ./"
	@test -z "$$(gofmt -l ./ | tee /dev/stderr)"

gofumpt:
	@echo "gofumpt -l ./"
	@test -z "$$(gofumpt -l ./ | tee /dev/stderr)"

lint:
	golangci-lint run

clean:
	$(GOCMD) clean

check: vet staticcheck gofmt #gofumpt

ci: build test check lint
