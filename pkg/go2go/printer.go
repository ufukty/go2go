package go2go

import (
	"fmt"
	"go/token"
	"io"
	"reflect"
	"slices"
)

type printer struct {
	indent           int
	writeBuffer      []byte
	indentBeforeType bool
	target           io.Writer
}

func (p *printer) bufferIndentation() {
	for j := 0; j < p.indent; j++ {
		p.writeBuffer = append(p.writeBuffer, indentation...)
	}
	p.indentBeforeType = false
}

func (p *printer) write(data []byte) {
	var bufferedSoFar = 0
	for i, b := range data {
		if b == '\n' {
			// add the passed line into buffer
			p.writeBuffer = append(p.writeBuffer, data[bufferedSoFar:i+1]...)
			bufferedSoFar += (i + 1) - bufferedSoFar

			// don't print the indentation, but leave a mark to print before next write
			p.indentBeforeType = true
		} else {
			if p.indentBeforeType {
				p.bufferIndentation()
			}
		}
	}
	if len(data) > bufferedSoFar {
		p.writeBuffer = append(p.writeBuffer, data[bufferedSoFar:]...)
	}
	p.target.Write(p.writeBuffer)
	p.writeBuffer = []byte{}
}

func (p *printer) lineEnd() {
	p.write([]byte{'\n'})
}

func (p *printer) printf(format string, value ...any) {
	p.write([]byte(fmt.Sprintf(format, value...)))
}

func (p *printer) recur(value reflect.Value) {
	if isNil(value) {
		p.printf("nil")
		return
	}

	kind := value.Kind()

	// open curly braces
	switch kind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		p.printf("%s{", value.Type())
		p.indent++
	}

	switch kind {

	case reflect.Interface:
		p.recur(value.Elem())

	case reflect.Map:
		if value.Len() > 0 {
			p.lineEnd()
			for _, key := range value.MapKeys() {
				p.recur(key)
				p.printf(": ")
				p.recur(value.MapIndex(key))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Pointer:
		p.printf("&")
		p.recur(value.Elem())

	case reflect.Array:
		if value.Len() > 0 {
			p.lineEnd()
			for i, n := 0, value.Len(); i < n; i++ {
				p.recur(value.Index(i))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Slice:
		if s, ok := value.Interface().([]byte); ok {
			p.printf("%#q", s)
			return
		}
		if value.Len() > 0 {
			p.lineEnd()
			for i, n := 0, value.Len(); i < n; i++ {
				p.recur(value.Index(i))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Struct:
		typ := value.Type()
		first := true
		for i, n := 0, typ.NumField(); i < n; i++ {
			if name := typ.Field(i).Name; token.IsExported(name) {
				if slices.Index(skipFields, name) == -1 {
					value := value.Field(i)
					if first {
						p.lineEnd()
						first = false
					}
					p.printf("%s: ", name)
					p.recur(value)
					p.printf(",")
					p.lineEnd()
				}
			}
		}

	default:
		v := value.Interface()
		switch v := v.(type) {
		case string:
			p.printf("%q", v)
		case token.Token:
			p.printf("%s", "token."+tokenMap[v.String()])

		default:
			p.printf("%v", v)
		}
	}

	// close curly braces
	switch kind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		p.indent--
		p.printf("}")
	}
}

func (p *printer) Print(x any) {
	p.recur(reflect.ValueOf(x))
}
