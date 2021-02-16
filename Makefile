VPATH = doc
PREFIX = /usr/local

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
	go build entf.net/htmltools/cmd/$@

%.1: %.1.scd
	scdoc < $< > $@

install: all
	install -Dm755 $(TOOLS) -t "$(PREFIX)/bin/"
	install -Dm644 $(DOCS) -t "$(PREFIX)/share/man/man1/"

uninstall:
	-rm -- $(addprefix $(PREFIX)/bin/, $(TOOLS))
	-rm -- $(addprefix $(PREFIX)/share/man/man1/, $(DOCS))

clean:
	-rm -- $(TOOLS)
	-rm -- $(DOCS)

.PHONY: all install uninstall clean