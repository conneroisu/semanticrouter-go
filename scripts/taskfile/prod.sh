#!/bin/bash
# file: makefile.prod.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/scripts/makefile.prod.sh
# title: makefile.prod.sh
# description: A script to build the cse-ncaa project into the folder used as the input for the scrape action.
#
# Usage: make build (SCRIPT RUNS IN GH ACTIONS DO NOT RUN MANUALLY UNLESS YOU KNOW WHAT YOU ARE DOING)
#
# if the BOB env var is not set, ask if your really want to build

VERSION=$(sh git describe --always --long --dirty)
if [ -z "$BOB" ]; then
	read -p "Are you sure you want to build? (y/n) " -n 1 -r
	echo
	if [[ ! $REPLY =~ ^[Yy]$ ]]
	then
		echo "Exiting..."
		exit 1
	else
		echo "Building..."
		go build -tags production -o ./bin/cse-ncaa -v -ldflags="-X main.version=$(git describe --always --long --dirty)" github.com/conneroisu/go-semantic-router ./main.go
		exit 0
	fi
else
	if true; then
		exit 1
	fi
fi
