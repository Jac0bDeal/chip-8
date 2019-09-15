GOLINT := $(shell command -v golint)
GOIMPORTS := $(shell command -v goimports)

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
ifndef VERSION
	VERSION := $(shell git tag -l --points-at HEAD)
	ifndef VERSION
		VERSION = $(COMMIT)
    endif
endif

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean disassembler emulator

clean:
	rm -f bin/*

disassembler: test
	@echo "---build---"
	go build ${LDFLAGS} -o bin/disassemble cmd/disassembler/*

emulator: test
	@echo "---build---"
	go build ${LDFLAGS} -o bin/chip8 cmd/emulator/*

fmt:
	@echo "---goimports---"
ifndef GOIMPORTS
	@go install golang.org/x/tools/cmd/goimports
endif
	@go list -f {{.Dir}} ./... | xargs -n1 -P8 goimports -w

lint:
	@echo "---golint---"
ifndef GOLINT
	@go install golang.org/x/lint/golint
endif
	@golint ./...

test: tidy fmt vet lint
	@echo "---go test---"
	@go test -count=1 --race ./...

tidy:
	@echo "---tidy---"
	@go mod tidy

vet:
	@echo "---go vet---"
	@go vet ./...

.PHONY: disassembler emulator test clean tidy fmt vet lint
