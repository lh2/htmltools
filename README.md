# htmltools

This is a collection of utilities to manipulate HTML documents.

- htmltotext: removes all tags from an HTML document leaving only the text nodes
- htmlremove: removes elements matching a selector from an HTML document
- htmlunwrap: removes elements matching a selector from an HTML document and
  replacing them with their child nodes

You can find more info on how to use these tools in their
[scdoc](https://git.sr.ht/~sircmpwn/scdoc) documents found in their respective
subfolders.

## Installation

Either `go get` a specific tool like this:

```
go get entf.net/htmltools/htmltotext
```

or clone the repository and run `make install`. The latter will install all the
tools and the manpages.
