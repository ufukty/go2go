package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"

	"github.com/ufukty/go2go/pkg/go2go"
)

type Args struct {
	In  string
	Out string
}

func Main() error {
	args := Args{}

	flag.StringVar(&args.In, "in", "", "input file. omit to pipe.")
	flag.StringVar(&args.Out, "out", "", "output file. omit to pipe.")
	flag.Parse()

	var in io.Reader
	if args.In != "" {
		i, err := os.Open(args.In)
		if err != nil {
			return fmt.Errorf("opening input file: %w", err)
		}
		defer i.Close()
		in = i
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("checking stdin: %w", err)
		}
		if stat.Size() == 0 {
			return fmt.Errorf("stdin is empty")
		}
		in = os.Stdin
	}

	var out io.WriteCloser
	if args.Out != "" {
		o, err := os.Create(args.Out)
		if err != nil {
			return fmt.Errorf("creating output file: %w", err)
		}
		defer o.Close()
		out = o
	} else {
		out = os.Stdout
	}

	f, err := parser.ParseFile(token.NewFileSet(), "", in, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	go2go.Fprint(out, f)

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
