status ?= onetime
version ?= 1.14.2
logging ?= false
kubefed ?= false
deploytool ?= helm

TARGETS := $(shell ls scripts)

.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-`uname -s`-`uname -m` > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

clean:
	rm -rf bin dist output vendor package/strongswan-*

$(TARGETS): .dapper
	./.dapper -m bind $@ $(status) $(version) $(logging) $(kubefed) $(deploytool)

.DEFAULT_GOAL := ci

.PHONY: $(TARGETS)

