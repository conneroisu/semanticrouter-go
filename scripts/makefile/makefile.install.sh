#!/bin/bash
# file: makefile.install.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/scripts/makefile.install.sh
# title: makefile.install.sh
# description: A script to install the cse-ncaa project.
#
# Usage: make install

if [ -e go ]; then
	echo "installing go"
	curl -sL https://go.dev/dl/go1.22.3.linux-amd64.tar.gz | tar -C /usr/local -xz
fi

if [ -e gum ]; then
	echo "installing gum"
	go install github.com/charmbracelet/gum@latest
fi

if [ -e sqlc ]; then
	go install github.com/kyleconroy/sqlc@latest
fi
