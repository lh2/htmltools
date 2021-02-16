package cmd

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func readerFromFile(file string) (f io.Reader, name string, err error) {
	if file == "-" {
		name = "[stdin]"
		f = os.Stdin
	} else {
		name = file
		f, err = os.Open(file)
		if err != nil {
			return
		}
	}
	return
}

func logErr(fileName string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", fileName, err)
}

func Fatal(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

func Main(args []string, handleFunc func(*html.Node)) {
	if len(args) == 0 {
		args = append(args, "-")
	}
	for _, a := range args {
		f, fn, err := readerFromFile(a)
		if err != nil {
			logErr(fn, err)
			continue
		}
		doc, err := html.Parse(f)
		if err != nil {
			logErr(fn, err)
			return
		}
		handleFunc(doc)
	}
}
