#!/bin/bash
# file: makefile.install.sh
# title: makefile.install.sh
# description: A script to install the cse-ncaa project.
#
# Usage: make install

if [ -e gum ]; then
	echo "installing gum"
	go install github.com/charmbracelet/gum@latest
fi

if [ -e sqlc ]; then
	gum spin --spinner dot --title "Installing SQLC" --show-output -- \
		go install github.com/kyleconroy/sqlc@latest
fi
