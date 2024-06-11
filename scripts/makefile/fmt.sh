#!/bin/bash
# file: makefile.fmt.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/scripts/makefile.fmt.sh
# title: Formatting Go Files
# description: This script formats the Go files using gofmt and golines.
#
# Usage: make fmt

# if gum is not installed just echo the messages

dbs=(
	"cse"
	"logs"
)

gofmt -w .

golines -w --max-len=79 .

templ fmt .

for db in "${dbs[@]}"; do
	cd "./data/""$db" || exit
	sqlc vet
	cd ../..
done
