package main

import "fmt"

// Das Programm liest eine Zeichenfolge ein und gibt aus, ob diese
// ein Palindrom darstellt.
func main() {
	var zeichenfolge string
	fmt.Println("Bitte geben Sie eine Zeichenfolge ein.")
	_, err := fmt.Scan(&zeichenfolge)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	negation := "k"
	if istPalindrom(zeichenfolge) {
		negation = ""
	}
	fmt.Printf("%v ist %vein Palindrom.\n", zeichenfolge, negation)
}
