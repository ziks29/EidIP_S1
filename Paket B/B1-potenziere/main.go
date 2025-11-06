package main

import "fmt"

// Das Programm liest Basis und Exponent ein und gibt den Potenzwert aus.
func main() {
	var basis, exponent int
	fmt.Println("Bitte geben Sie die Basis und den Exponenten ein: ")
	_, err := fmt.Scan(&basis, &exponent)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if exponent < 0 {
		fmt.Println("Sie haben als Exponent", exponent, "eingegeben.")
		fmt.Println("Der Exponent darf aber nicht negativ sein!")
		return
	}
	fmt.Println(basis, "hoch", exponent, "ist", potenziere(basis, exponent))
}
