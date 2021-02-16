VPATH = doc
PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
MANDIR ?= $(PREFIX)/share/man
GO ?= go
GOFLAGS ?=

TOOLS := \
	htmlattr \
	htmlindentheadings \
	htmlremove \
	htmlselect \
	htmltotext \
	htmlunwrap
DOCS := $(addsuffix .1, $(TOOLS))

SRC := $(shell find . -name "*.go")

all: $(TOOLS) $(DOCS)

$(TOOLS): $(SRC)
	$(GO) build $(GOFLAGS) entf.net/htmltools/cmd/$@

%.1: %.1.scd
	scdoc < $< > $@

install: all
	install -Dm755 $(TOOLS) -t "$(DESTDIR)$(BINDIR)"
	install -Dm644 $(DOCS) -t "$(DESTDIR)$(MANDIR)/man1/"

uninstall:
	-rm -- $(addprefix $(DESTDIR)$(BINDIR)/, $(TOOLS))
	-rm -- $(addprefix $(DESTDIR)$(MANDIR)/man1/, $(DOCS))

clean:
	-rm -- $(TOOLS)
	-rm -- $(DOCS)

.PHONY: all install uninstall clean