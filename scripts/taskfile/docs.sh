# name: makefile/docs.sh
# description: A script to generate the go docs for the cse-ncaa project.
# 
# Usage: make docs


gum spin --spinner dot --title "Generating Docs" --show-output -- \
	gomarkdoc --exclude-dirs ./testData/... ./... > ./Code-Generated.md
