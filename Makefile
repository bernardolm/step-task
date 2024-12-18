ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

APP_ROOT_NAMESPACE="github.com/bernardolm/step-task/pkg/infrastructure"
COMMIT_HASH=$(shell git rev-parse --short HEAD)
IS_DIRTY=$(shell git diff --quiet && echo no || echo yes)
NOW=$(shell date +%y%m%d-%H%M%S)
PWD=$(shell pwd)
SHELL:=/bin/bash

config:
	@ln -sf "${PWD}/.githooks/pre-commit" "${PWD}/.git/hooks/pre-commit"

air:
	@command -v air &>/dev/null || go install github.com/cosmtrek/air@latest
	@expect_unbuffer air -c=dev/.air.toml -d

debug:
	@dlv debug --listen=:2345 --headless=true --api-version=2 cmd/cli/main.go buildGoogleAdsReport

format:
	GOLINES=yes source .githooks/pre-commit

format-all:
	@FILES=all make format

run: format
	@go run cmd/cli/main.go

build: export LDFLAGS_X="-X ${APP_ROOT_NAMESPACE}/config.BuildAt=${NOW} \
-X ${APP_ROOT_NAMESPACE}/config.CommitHash=${COMMIT_HASH} \
-X ${APP_ROOT_NAMESPACE}/config.IsDirty=${IS_DIRTY}"
build: format
	@eval go build -ldflags \"-w -s ${LDFLAGS_X}\" -o bin/step-task cmd/cli/main.go

.PHONY: config air debug format format-all run ldflags build
