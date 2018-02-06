GO ?= go

PACKAGE = goeuler
LIBPACKAGE = euler

_PACKAGE = $(GOPATH)/bin/$(PACKAGE)
_LIBPACKAGE = $(GOPATH)/pkg/$(addsuffix .a, $(addprefix lib, $(LIBPACKAGE))) 

.PHONY: all clean

all: $(_PACKAGE)

$(_PACKAGE): main.go $(_LIBPACKAGE)
	$(GO) build -o $(_PACKAGE) main.go

$(_LIBPACKAGE): $(wildcard $(LIBPACKAGE)/*.go)
	cd $(LIBPACKAGE) && $(GO) build -o $(_LIBPACKAGE)

clean:
	rm -f $(_PACKAGE) $(_LIBPACKAGE)

