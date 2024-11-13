MAKE_PATH=$(GOPATH)/bin:/bin:/usr/bin:/usr/local/bin:$PATH

.PHONY: all
all: gen

BUF_VERSION=1.46.0

.PHONY: gen
gen:
	sudo chmod +x gen.sh && ./gen.sh

.PHONY: buf-install
buf-install:
ifeq ($(shell uname -s), Darwin)
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Darwin-arm64 && \
	mv buf $(GOPATH)/bin/buf && \
	chmod +x $(GOPATH)/bin/buf
else
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64 && \
	sudo mv buf $(GOPATH)/bin/buf && \
	sudo chmod +x $(GOPATH)/bin/buf
endif