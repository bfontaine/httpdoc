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

	flag.StringVar(&doc.RootDir, "root-dir", "./_docs",
		"Documentation root directory")
	flag.Parse()

	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	if code, err := doc.GetStatusCode(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", code.PrettyString())
	}
}
