// Enumgen is a tool to automate the creation of methods of the struct that encapsulates
// enums-like functionality. This structure can only contain a value form a predefined list of constants.
// It enhances type checking in your code.
//
// For example, given this snippet
//
// package colors
//
// //go:generate enumgen
//
// type Color struct {
//     value int
// }
//
// const (
//     RED   = 1
//     GREEN = 2
//     BLUE  = 3
// )
//
// running this command
//
// go generate ./...
//
// in the root directory of your project will create the file colors_gen.go, in package colors,
// containing all the necessary methods.
package main

import (
	"github.com/swanden/enumgen/template"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	goFile := os.Getenv("GOFILE")
	if goFile == "" {
		log.Fatal("GOFILE is not set")
	}

	astFile, err := parser.ParseFile(
		token.NewFileSet(),
		goFile,
		nil,
		0,
	)
	if err != nil {
		log.Fatalf("parse input file: %v", err)
	}

	data := template.GetData(astFile)
	content, err := template.GenFileContent(data)
	if err != nil {
		log.Fatalf("generate file content: %v", err)
	}

	genFile := strings.TrimSuffix(goFile, ".go") + "_gen.go"
	err = os.WriteFile(genFile, content.Bytes(), 0644)
	if err != nil {
		log.Fatalf("create file: %v", err)
	}
}
