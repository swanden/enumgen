# Enumgen

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/swanden/enumgen/LICENSE)

Enumgen is a tool to automate the creation of methods of the struct that encapsulates
enums-like functionality. This structure can only contain a value form a predefined list of constants.
It enhances type checking in your code.

Given the names of constants defined and its type, enumgen will create a new self-contained Go source file with
all the necessary methods.

Installation
================
~~~
go install github.com/swanden/enumgen
~~~

Run
================
Create file colors.go with batch of constants and main struct.
~~~
package colors

//go:generate enumgen

type Color struct {
	value int
}

const (
	RED   = 1
	GREEN = 2
	BLUE  = 3
)
~~~

Running this command
~~~
go generate ./...
~~~
in the root directory of your project will create the file colors_gen.go, in package colors,
containing all the necessary methods.

Usage
================
~~~
package main

import (
	"myproject/colors"
	"fmt"
)

func initLogDefaultColor(defaultColor colors.Color) {

}

func main() {
	green := colors.NewGREEN() // Creating enum
	blue := colors.NewBLUE()

	fmt.Println(green.IsEqual(blue)) // Comparing two enums
	fmt.Println(green.IsGREEN())     // Comparing enum with const

	initLogDefaultColor(green) // Passing enum to function. Type checking ensures that only a value from the list of constants will be passed to the function

	fmt.Println(colors.Values()) // Get all const values

	fmt.Println(blue) // Printing enum value

	// Deserializing enum from const value (for example from DB)
	valueFromDB := 2
	if color, err := colors.New(valueFromDB); err == nil {
		fmt.Println(color.Value())
	}
}
~~~