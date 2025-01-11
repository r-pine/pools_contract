.PHONY: all
all: gen

BUF_VERSION=1.49.0

.PHONY: gen
gen:
	sudo chmod +x gen.sh && ./gen.sh

.PHONY: buf-install
buf-install:
	curl -L -o buf	https://github.com/bufbuild/buf/releases/download/v1.49.0/buf-Linux-x86_64 && \
	mv buf /usr/local/bin/buf && \
	chmod +x /usr/local/bin/buf
