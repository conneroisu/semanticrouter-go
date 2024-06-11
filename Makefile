# -include .env
#  file: Makefile

export MAKEFLAGS += --always-make --print-directory
SHELLFLAGS = -e

.PHONY: fmt
fmt:
	@sh ./scripts/makefile.fmt.sh

.PHONY: test
test:
	@sh ./scripts/makefile.test.sh

.PHONY: lint
lint:
	@sh ./scripts/makefile.lint.sh

.PHONY: tidy
tidy:
	@sh ./scripts/makefile.tidy.sh

.PHONY: docs
docs:
	@sh ./scripts/makefile.docs.sh

.PHONY: testdata
testdata:
	@sh ./scripts/makefile.testdata.sh

.PHONY: release
release:
	@sh ./scripts/makefile.release.sh

.PHONY: build
build:
	@sh ./scripts/makefile.build.sh

.PHONY: prod
prod:
	@sh ./scripts/makefile.prod.sh
