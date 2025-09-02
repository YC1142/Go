package main

import (
	"fmt"
)

func main() {
	output_functions()
	println("")
	formatting_verbs()
}
func output_functions() {
	fmt.Println("----- The Print() Function -----")
	/*
		Die Funktion Print() gibt ihre Argumente im Standardformat aus.
	*/
	var i, j string = "Hello", "World"
	fmt.Print(i)
	fmt.Print(j)
	/*
		new lines with \n
	*/
	fmt.Println("")
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	/*
		Ein Print() für mehrere Variablen
	*/
	fmt.Println("")
	fmt.Print(i, "\n", j)
	/*
		Space between string arguments
	*/
	fmt.Println("")
	fmt.Print(i, " ", j)
	/*
		Print() für ein leerzeichen zwischen Argumente, wenn keines dieser ein String ist
	*/
	fmt.Println("")
	a := 5
	b := 3
	fmt.Print(a, b)
	fmt.Println("")
	fmt.Println("----- The Println() Function -----")
	/*
		Ähnlich wie Print(), mit dem Unterschied, dass ein Leerzeichen zwischen den Argumenten hinzugefügt wird, sowie auch ein Zeilenumbruch
		am Ende
	*/
	var c, d string = "Hello", "World"
	fmt.Println(c, d)
	fmt.Println("")
	fmt.Println("----- The Printf() Function -----")
	/*
			Die Funktion Printf() formatiert zunächst ihr Argument basierend auf dem angegebenen Formatierungsverb und gibt es dann aus.

			Hier verwenden wir zwei Formatierungsverben:

		    %v wird verwendet, um den Wert der Argumente auszugeben.
		    %T wird verwendet, um den Typ der Argumente auszugeben.

	*/
	var va string = "Hello"
	var ji int = 15
	fmt.Printf("i has value: %v and type %T\n", va, va)
	fmt.Printf("i has value: %v and type %T\n", ji, ji)
}

func formatting_verbs() {

}
