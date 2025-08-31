package main

import (
	"fmt"
)

var d int
var e int = 2
var f = 3

// g := 4
func main() {
	println("----- Variablendeklaration mit Anfangswert -----")
	var student1 string = "John"
	var student2 = "Jane"
	x := 2
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	// Note: The variable types of student2 and x is inferred (abgeleitet) from their values
	fmt.Println("")
	fmt.Println("----- Variablendeklaration ohne Anfangswert -----")
	var a string
	var b int
	var c bool

	fmt.Println(a) // a is ""
	fmt.Println(b) // b is 0
	fmt.Println(c) // c is false

	fmt.Println("")
	fmt.Println("----- Wertzuweisung nach Deklaration -----")
	var student3 string
	student3 = "John"
	fmt.Println(student3)

	fmt.Println("")
	fmt.Println("----- Unterschied zwischen keyword var und := -----")
	d = 1
	println(d)
	println(e)
	println(f)
	//println(g)
	/*
		g muss innerhalb einer Funktion sein
		./prog.go:7:1: syntax error: non-declaration statement outside function body
	*/
	multiple_var_declaration()
}
func multiple_var_declaration() {
		println("")
		println("----- Mehrere Variablen deklarieren -----")
		var a,b,c,d int = 1,3,5,7
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		
		fmt.Println("")
		fmt.Println("----- Mehrere Variablen mit verschiedenen Typen -----")
		var e,f = 1,"Hello"
		g,h := 7, "World!"
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println(g)
		fmt.Println(h)
		fmt.Println("")
		fmt.Println("----- Go Variablendeklaration in einem Block -----")
		var (
			ab int
			bb int = 1
			cc string = "Hello"
		)
		fmt.Println(ab)
		fmt.Println(bb)
		fmt.Println(cc)


}

