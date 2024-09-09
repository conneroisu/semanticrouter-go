#!/bin/bash
# file: makefile.install.sh
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

# install gum (https://charm.sh/gum) binary-name: [gum]
go install github.com/charmbracelet/gum@latest

# install the taskfile cli (https://taskfile.dev/) binary-name: [task]
gum spin --spinner dot --title "Installing Taskfile CLI[task]" --show-output -- \
	go install github.com/go-task/task/v3/cmd/task@latest

# install gomarkdoc (https://github.com/princjef/gomarkdoc) binary-name: [gomarkdoc]
gum spin --spinner dot --title "Installing gomarkdoc[gomarkdoc]" --show-output -- \
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

gh extension install https://github.com/nektos/gh-act

# install sqlc (https://www.sqlc.dev/) binary-name: [sqlc]
gum spin --spinner dot --title "Installing sqlc[sqlc]" --show-output -- \
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# install revive (https://github.com/mgechev/revive) binary-name: [revive]
gum spin --spinner dot --title "Installing revive[revive]" --show-output -- \
	go install github.com/mgechev/revive@latest

# install air (https://github.com/cosmtrek/air) binary-name: [air]
gum spin --spinner dot --title "Installing air[air]" --show-output -- \
	go install github.com/air-verse/air@latest

# install golangci-lint (https://golangci-lint.run/) binary-name: [golangci-lint]
gum spin --spinner dot --title "Installing golangci-lint[golangci-lint]" --show-output -- \
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

# install staticcheck (https://staticcheck.io/) binary-name: [staticcheck]
gum spin --spinner dot --title "Installing staticcheck[staticcheck]" --show-output -- \
	go install honnef.co/go/tools/cmd/staticcheck@latest

# install golines (https://github.com/segmentio/golines) binary-name: [golines]
gum spin --spinner dot --title "Installing golines[golines]" --show-output -- \
	go install github.com/segmentio/golines@latest

# install go-task (https://github.com/go-task/task) binary-name: [go-task]
gum spin --spinner dot --title "Installing go-task[go-task]" --show-output -- \
	go install github.com/go-task/task/v3/cmd/task@latest

# isntall gocovsh (https://github.com/orlangure/gocovsh) binary-name: [gocovsh]
gum spin --spinner dot --title "Installing gocovsh[gocovsh]" --show-output -- \
	go install github.com/orlangure/gocovsh@v0.5.1

