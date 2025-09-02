package main

import (
	"fmt"
)

func main() {
	constant_tutorial()
}

func constant_tutorial() {
	fmt.Println("")
	fmt.Println("----- Eine Konstante deklarieren -----")
	const PI = 3.14
	fmt.Println(PI)
	/*
		Konstante Namen werden in der Regel in Großbuchstaben geschrieben (zur leichteren Identifizierung und Unterscheidung von Variablen).
		Konstanten können sowohl innerhalb als auch außerhalb einer Funktion deklariert werden.
	*/
	fmt.Println("")
	fmt.Println("----- Typisierte Konstanten -----")
	const A int = 1
	fmt.Println(A)
	/*
		Typisierte Konstanten werden mit einem definierten Typ deklariert.
	*/
	fmt.Println("")
	fmt.Println("----- Nicht typisierte Konstanten -----")
	const B = 1
	fmt.Println(B)
	/*
		Nicht typisierte Konstanten werden ohne Typ deklariert.
	*/
	/*
		Wenn eine Konstante deklariert wurde, ist es nicht möglich den Wert zu ver
	*/
	fmt.Println("")
	fmt.Println("----- Deklaration mehrerer Konstanten -----")
	const (
		D int = 5
		E     = 3.14
		F     = "Hi!"
	)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	/*
		Mehrere Konstanten können zur besseren Lesbarkeit in einem Block zusammengefasst werden:
	*/
}
