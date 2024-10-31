package go2go

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

// this function tests if the Fprint can produce code in valid Go syntax by putting it into a Go file, compile and run
func Test_ReverseCheck(t *testing.T) {
	content, err := os.ReadFile("testdata/walk.go")
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	f, err := parser.ParseFile(token.NewFileSet(), "testdata/walk.go", nil, parser.AllErrors)
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	module, err := os.MkdirTemp(os.TempDir(), "Test_ReverseCheck")
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	file, err := os.Create(filepath.Join(module, "main.go"))
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	fmt.Fprintf(file, `package main

import (
	"os"
	"go/printer"
	"go/ast"
	"go/token"
)

var f = `)
	Fprint(file, f)
	fmt.Fprintf(file, `


func main() {
	printer.Fprint(os.Stdout, token.NewFileSet(), f)
}`)

	// fmt.Println(module)

	cmd := exec.Command("go", "mod", "init", "na/na")
	cmd.Dir = module
	err = cmd.Run()
	if err != nil {
		t.Error(fmt.Errorf("performing module creation: %w", err))
	}

	stdout := bytes.NewBuffer([]byte{})

	cmd = exec.Command("go", "run", ".")
	cmd.Dir = module
	cmd.Stdout = stdout
	cmd.Run()
	if err != nil {
		t.Error(fmt.Errorf("performing running test program: %w", err))
	}

	fmt.Println(diff.Diff(string(content), stdout.String()))
}
