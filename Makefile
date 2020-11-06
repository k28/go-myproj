# メタ情報
NAME := myproj
VERSION := $(godump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO11MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# 開発に必要な依存をインストールする
## Setup
.PHONY: deps
devel-deps: deps
	GO11MODULE=off go get	\
	golang.org/x/lint/golint	\
	github.com/motemen/gobump/cmd/gobump	\
	github.com/Songmu/make2help/cmd/make2help

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries ex. make bin/myproj
bin%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o %@ $<

## build binary
.PHONY: build
build: bin/myproj

## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)

