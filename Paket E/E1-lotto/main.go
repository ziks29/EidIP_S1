package main

import (
	"fmt"
	"math"
)

func main() {
	var r, n int
	println("Zur Berechnung der Wahrscheinlichkeit von r Richtigen bei der")
	print("Ziehung von 6 aus N Zahlen geben Sie bitte r und N ein: ")
	_, err := fmt.Scan(&r, &n)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if r <= 0 || r > 6 || n <= 6 {
		fmt.Printf("r (Ihre Eingabe: %v) muss im Bereich [1..6] liegen und N (Ihre Eingabe: %v) größer als 6.", r, n)
		return
	}
	fmt.Println()
	Pr := int(math.Round(float64(binom(n, 6)) / float64(binom(n-6, 6-r)) / float64(binom(6, r))))
	fmt.Printf("Die Wahrscheinlichkeit für %v Richtige im Lotto 6 aus %v ist 1 : %v.\n", r, n, Pr)
}
