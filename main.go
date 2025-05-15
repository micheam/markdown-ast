package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func init() {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s [options] <markdown_file>\n", flag.CommandLine.Name())
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "markdown_file:\n")
		fmt.Fprintf(out, "  The path to the markdown file to be processed.\n")
		fmt.Fprintf(out, "  If no file is specified, stdin will be used.\n")
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "Example:\n")
		fmt.Fprintf(out, "  $ %s example.md\n", flag.CommandLine.Name())
		fmt.Fprintf(out, "  $ %s < example.md\n", flag.CommandLine.Name())
		fmt.Fprintf(out, "  $ cat example.md | %s\n", flag.CommandLine.Name())
	}
}

func main() {
	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		return
	}

	var input *os.File
	if flag.NArg() == 0 {
		input = os.Stdin
	} else {
		var err error
		file := flag.Arg(0)
		if input, err = os.Open(file); err != nil {
			fmt.Fprintf(os.Stderr, "Can't open file %s: %v\n", file, err)
			return
		}
		defer input.Close()
	}
	source, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %v\n", err)
		return
	}

	node := goldmark.New().Parser().Parse(text.NewReader(source))
	node.Dump(source, 0)
}
