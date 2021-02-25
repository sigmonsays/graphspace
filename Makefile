.PHONY: install

TOPDIR = $(shell pwd)

export GOWORKSPACE := $(shell pwd)
export GOBIN := $(GOWORKSPACE)/bin
export GO111MODULE := on

GO_BINS =
GO_BINS += graphspace

all:
	$(MAKE) compile

compile:
	mkdir -p tmp
	mkdir -p $(GOBIN)
	go install github.com/...

install:
	mkdir -p $(DESTDIR)/$(INSTALL_PREFIX)/bin/ 
	$(MAKE) install-bins

install-bins: $(addprefix installbin-, $(GO_BINS))

$(addprefix installbin-, $(GO_BINS)):
	$(eval BIN=$(@:installbin-%=%))
	cp -v $(GOBIN)/$(BIN) $(DESTDIR)/$(INSTALL_PREFIX)/bin/
