TOOLS = htmlremove htmltotext htmlunwrap htmlselect htmlindentheadings
PREFIX = /usr/local
MANS = $(shell find . -name '*.scd' | sed s/\.scd//)

all: $(TOOLS) $(MANS)

$(TOOLS):
	mkdir -p bin
	go build -o bin/$@ entf.net/htmltools/$@

%.1: %.1.scd
	scdoc < $< > $@

install: all
	mkdir -p "$(PREFIX)/bin"
	cp $(addprefix bin/, $(TOOLS)) "$(PREFIX)/bin/"
	mkdir -p "$(PREFIX)/share/man/man1"
	cp $(MANS) "$(PREFIX)/share/man/man1/"

uninstall:
	-rm -- $(addprefix $(PREFIX)/bin/, $(TOOLS))
	-rm -- $(addprefix $(PREFIX)/share/man/man1/, $(notdir $(MANS)))

clean:
	-rm -r bin/
	-rm -- $(MANS)

.PHONY: all $(TOOLS) install uninstall clean