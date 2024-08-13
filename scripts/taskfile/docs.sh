# name: makefile/docs.sh
# description: A script to generate the go docs for the cse-ncaa project.
# 
# Usage: make docs

gomarkdoc --exclude-dirs ./testData/... ./... > docs/Home.md
