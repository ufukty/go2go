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

## License

Check LICENSE file.
