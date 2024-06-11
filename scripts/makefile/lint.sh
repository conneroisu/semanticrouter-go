#!/bin/bash
# file: makefile.lint.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/makefile.lint.sh
# description: Runs all linters
#
# Usage: make lint

# if gum is not installed just echo the messages
if ! command -v gum &>/dev/null; then
	CMD="staticcheck ./..." && echo "Running $CMD" && $CMD
	CMD="golangci-lint run" && echo "Running $CMD" && $CMD
	CMD="sqlc vet" && echo "Running $CMD" && $CMD
	exit 0
fi

staticcheck ./...

golangci-lint run

sqlc vet

revive -config .revive.toml -formatter friendly ./...
