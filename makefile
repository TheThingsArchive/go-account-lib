SHELL = bash

go = go
golint = golint
pkgs = $(go) list ./... | grep -vE go-account-lib/vendor

cover_file=coverage.out
tmp_cover_dir ?= .cover

.PHONY: test tools cover watch

tools:
	@echo fething dev deps...
	@command -v cover $(go) get -u github.com/kardianos/govendor
	@command -v cover $(go) get -u github.com/golang/lint/golint

test:
	@for pkg in $$($(pkgs)); do   \
		profile=$$([ "$$COVER" = "1" ] && echo "-coverprofile=$(tmp_cover_dir)/$$(echo $$pkg | tr -d '/').cover"); \
		$(go) test -cover $$profile $$pkg; \
	done

cover:
	@mkdir -p $(tmp_cover_dir)
	@touch $(tmp_cover_dir)/empty.cover
	@export COVER=1; make test
	@echo "mode: set" > $(cover_file)
	@cat $(tmp_cover_dir)/*.cover | grep -v mode | sort -r >> $(cover_file)
	@rm -rf $(tmp_cover_dir)

vet:
	$(pkgs) | xargs $(go) vet

watch:
	@ginkgo watch -coverprofile=/dev/null $$($(pkgs) | sed "s/.*$$(basename $$PWD)\//.\//")

deps:
	govendor sync -v

fmt:
	[[ -z "`$(pkgs) | xargs go fmt | tee -a /dev/stderr`" ]]

lint:
	for pkg in `$(pkgs)`; do $(golint) $$pkg; done

