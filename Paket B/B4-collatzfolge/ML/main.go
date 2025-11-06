package main

import "fmt"

// collatz gibt die Collatz-Folge für den Startwert start aus,
// bis die Glieder 4,2,1 (zum ersten Mal) erreicht sind.
// start muss eine positive ganze Zahl sein.
func collatz(start int) {
	fmt.Println("Die Collatz-Folge mit dem Startwert", start, "lautet:")
	n := start
	fmt.Print(n, ",")
	// solange Folgenglieder berechnen und ausgeben,
	// bis das Folgenglied 1 erreicht ist
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(n, ",")
	}
	fmt.Println("...")
	fmt.Println()
}

// Das Programm gibt die Collatz-Folgen für die Startwerte 19, 23, 42 und 122 aus,
// bis die Glieder 4,2,1 (zum ersten Mal) erreicht sind.
func main() {
	collatz(19)
	collatz(23)
	collatz(42)
	collatz(122)
}
