# `go2go`

go2go is a program that prints AST structure of a Go code in Go syntax, to let you construct them later in your runtime partially or fully. go2go enables ad-hoc code generator development.

Example output:

```go
&ast.File{
  Doc: nil,
  Name: &ast.Ident{
    Name: "public",
  },
  Decls: []ast.Decl{
    &ast.GenDecl{
      Doc: nil,
      Tok: token.IMPORT,
      Specs: []ast.Spec{
        &ast.ImportSpec{
          Doc: nil,
          Name: nil,
          Path: &ast.BasicLit{
            Kind: token.STRING,
            Value: "\"fmt\"",
          },
          // ...
```

Since, the tree is in valid syntax, it can be directly copied into a Go file.

This is the fastest way to create templates for code generators I've seen yet.

## Install

```sh
go install github.com/ufukty/go2go@latest
```

## Usage

```sh
$ go2go --help
Usage of go2go:
  -in string
        input file. omit to pipe.
  -out string
        output file. omit to pipe.
```

## Prior art

I am publishing this after not touching to it for over one year of time. So I don't remember.

With quick search, I found [lu4p/astextract](https://github.com/lu4p/astextract). It contains a [file](https://github.com/lu4p/astextract/blob/master/print.go) written in similar logic to go2go with usage of `reflect`.

Although it refers to the built-in `go/ast` package's [printer](https://golang.org/src/go/ast/print.go).

## License

Check LICENSE file.
