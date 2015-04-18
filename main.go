package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/httpdoc/httpdoc"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n    %s [<options>] <code>\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Where <options> are:")
	flag.PrintDefaults()
}

func main() {
	var doc httpdoc.Doc

	flag.StringVar(&doc.RootDir, "root-dir", "",
		"Documentation root directory")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		printUsage()
		os.Exit(1)
	}

	if doc.RootDir == "" {
		doc = httpdoc.DefaultDoc
	}

	if res, err := doc.GetResourceFor(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", res.PrettyString())
	}
}
