package go2go

import (
	"fmt"
	"go/ast"
	"io"
)

func Fprint(w io.Writer, n ast.Node) {
	p := printer{
		indent:           0,
		writeBuffer:      []byte{},
		indentBeforeType: false,
		target:           w,
	}
	p.Print(n)
	fmt.Fprintln(w, "")
}
