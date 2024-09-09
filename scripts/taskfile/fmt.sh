#!/bin/bash
# file: makefile.fmt.sh
# title: Formatting Go Files
# description: This script formats the Go files using gofmt and golines.
#
# Usage: make fmt

# if gum is not installed just echo the messages

gofmt -w .

golines -w --max-len=79 .
