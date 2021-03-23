.SUFFIXES:
MAKEFLAGS+=-r

.PHONY: all clean test build benchmark package
.DEFAULT: all

# GO shit
ROOT_DIR:=$(realpath $(CURDIR))
export GOPATH:=$(ROOT_DIR)

# version & build time
VERSION:=$(shell git describe --dirty --tags)
ifeq (,$(VERSION))
VERSION:="UNKNOWN"
endif
BUILD:=$(shell date -u +%FT%TZ)
HOSTNAME:=$(shell hostname)
ifeq (,$(HOSTNAME))
HOSTNAME:="hostname.unknown"
endif
TARGET:="pef"


all: clean build

clean:
	@echo cleaning...
	@rm -rf $(TARGET)
	@rm -rf *.rpm
	@rm -rf __pycache__/

test:
	@echo unit testing...
	cd src/pef; go test -cover ./...

benchmark:
	@echo benchmark...
	cd src/pef; go test -benchmem -bench=. ./... -run=none

build:
	@echo building...
	cd src/pef; go build -o ../../$(TARGET) -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD) -X main.Hostname=$(HOSTNAME)"

package:
	@echo generating rpm...
	fpm -s dir -t rpm --prefix /usr/local/bin/ -n pef -v $(VERSION) pef
