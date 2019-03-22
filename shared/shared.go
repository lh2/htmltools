package shared

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var currentFile string

func readerFromFile(file string) (f io.Reader, err error) {
	if file == "-" {
		currentFile = "[stdin]"
		f = os.Stdin
	} else {
		currentFile = file
		f, err = os.Open(file)
		if err != nil {
			return
		}
	}
	return
}

func LogErr(err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", currentFile, err)
}

func Main(args []string, handleFunc func(*html.Node)) {
	if len(args) == 0 {
		args = append(args, "-")
	}
	for _, a := range args {
		f, err := readerFromFile(a)
		if err != nil {
			LogErr(err)
			continue
		}
		doc, err := html.Parse(f)
		if err != nil {
			LogErr(err)
			return
		}
		handleFunc(doc)
	}
}
