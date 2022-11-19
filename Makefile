NAME      := ssht
PACKAGE   := github.com/root913/$(NAME)

ifndef VERSION
	VERSION ?= $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
endif

LDFLAGS += -s -w -extldflags "-static" -X "main.version=$(VERSION)"

# 
define print_build_info
	@printf "%-15s\e[38;5;69m%s/\e[38;5;38m%-8s\e[0m\n" "Building for" $(1) $(2)
endef

define build
	$(call print_build_info,$(1),$(2))
	@GOARCH=$(2) GOOS=$(1) GO111MODULE=on go build -ldflags '$(LDFLAGS)' -a -tags netgo -o bin/${NAME}_$(1)_$(2)$(3)
	@printf "\e[38;5;49m%-5s    \e[38;5;29m%s%s\e[0m\n\n" "Finished" bin/${NAME}_$(1)_$(2) $(3)
endef

default: help

build_linux:	## Build for linux
	$(call build,linux,amd64)

build_windows:	## Build for Windows
	$(call build,windows,amd64,.exe)

build_darwin:	## Build for darwin
	$(call build,darwin,amd64)

build: build_darwin build_linux build_windows ## Build for darwin/linux/windows

test:	## Run tests
	@go clean --testcache && go test ./...

clean:	## Clean after build
	@go clean
	@find bin -type f -name "${NAME}_*" -delete -print

docker:	## Enter docker container
	@docker run --rm --workdir=/app -v $(pwd):/app -it golang:1.19-alpine sh

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\e[38;5;69m%-30s\e[38;5;38m %s\e[0m\n", $$1, $$2}'