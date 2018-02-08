GO ?= go

PACKAGE = goeuler
LIBPACKAGE = euler
SAINT = github.com/bxcodec/saint

_PACKAGE = $(GOPATH)/bin/$(PACKAGE)
_LIBPACKAGE = $(GOPATH)/pkg/$(addsuffix .a, $(addprefix lib, $(LIBPACKAGE)))
_SAINT = $(GOPATH)/src/$(SAINT)

.PHONY: all clean

all: $(_PACKAGE)

$(_PACKAGE): main.go $(_LIBPACKAGE)
	$(GO) build -o $(_PACKAGE) main.go

$(_LIBPACKAGE): $(wildcard $(LIBPACKAGE)/*.go) $(_SAINT)
	cd $(LIBPACKAGE) && $(GO) build -o $(_LIBPACKAGE)

$(_SAINT):
	go get -u $(SAINT)

clean:
	rm -f $(_PACKAGE) $(_LIBPACKAGE)

