#!/bin/bash
# file: makefile/lint.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/makefile/lint.sh
# description: Runs all linters
#
# Usage: make lint

staticcheck ./...

golangci-lint run

sqlc vet

revive -config .revive.toml -formatter friendly ./...
