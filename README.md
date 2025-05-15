# markdown-ast

A simple command-line tool to display the Abstract Syntax Tree (AST) of Markdown documents.

This is a thin wrapper around the [goldmark](https://github.com/yuin/goldmark) Markdown parser, providing easy access to its AST representation capabilities.

## Installation

```bash
go install github.com/micheam/markdown-ast@latest
```

or if you prefer to use `go run`:

```bash
go run github.com/micheam/markdown-ast@latest
```

After installation, the `markdown-ast` command will be available in your `$GOPATH/bin` directory.

## Usage

```
Usage: markdown-ast [options] <markdown_file>

markdown_file:
  The path to the markdown file to be processed.
  If no file is specified, stdin will be used.

Example:
  $ markdown-ast example.md
  $ markdown-ast < example.md
  $ cat example.md | markdown-ast
```

