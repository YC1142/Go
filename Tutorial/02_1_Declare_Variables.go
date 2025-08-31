package main

import (
	"fmt"
)

var d int
var e int = 2
var f = 3

// g := 4
func main() {
	var student1 string = "John"
	var student2 = "Jane"
	x := 2
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	// Note: The variable types of student2 and x is inferred (abgeleitet) from their values
	fmt.Println("")
	fmt.Println("Variablendeklaration ohne Anfangswert")
	var a string
	var b int
	var c bool

	fmt.Println(a) // a is ""
	fmt.Println(b) // b is 0
	fmt.Println(c) // c is false

	fmt.Println("")
	fmt.Println("Wertzuweisung nach Deklaration")
	var student3 string
	student3 = "John"
	fmt.Println(student3)

	fmt.Println("")
	fmt.Println("Unterschied zwischen keyword var und :=")
	d = 1
	println(d)
	println(e)
	println(f)
	//println(g)
	/*
		g muss innerhalb einer Funktion sein
		./prog.go:7:1: syntax error: non-declaration statement outside function body
	*/
}
