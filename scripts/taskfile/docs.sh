# name: makefile/docs.sh
# description: A script to generate the go docs for the cse-ncaa project.
# 
# Usage: make docs

golds -s -gen -wdpkgs-listing=promoted -dir=./docs -footer=verbose+qrcode
xdg-open ./docs/index.html
