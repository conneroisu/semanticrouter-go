#!/bin/sh
# file: makefile.release.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/scripts/makefile.release.sh
# title: makefile.release.sh
# description: A script to release the cse-ncaa project.
#
# Usage: make release

# if gum is not installed just echo the messages
if ! command -v gum &>/dev/null; then
	echo "releasing cse-ncaa"
	goreleaser release
	exit 0
fi
