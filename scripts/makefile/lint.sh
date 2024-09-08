#!/bin/bash
# file: makefile/lint.sh
# description: Runs all linters
#
# Usage: make lint

staticcheck ./...

golangci-lint run

revive -config .revive.toml -formatter friendly ./...
