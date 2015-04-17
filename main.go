package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bfontaine/httpdoc/Godeps/_workspace/src/github.com/kr/text"
	"github.com/bfontaine/httpdoc/httpdoc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <code>\n", os.Args[0])
		os.Exit(1)
	}

	doc := httpdoc.Doc{RootDir: "./_docs"}

	if code, err := doc.GetStatusCode(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s %s\n\n%s\n\nReferences:\n * %s\n",
			code.Code, code.Text,
			text.Wrap(code.Desc, 75),
			strings.Join(code.Refs, "\n * "),
		)
	}
}
