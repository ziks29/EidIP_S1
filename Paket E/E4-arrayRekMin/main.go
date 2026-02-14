package main

import (
	"fmt"
	"math/rand"
)

// ***Länge des Arrays hier festlegen ***
const arrLen = 225

// das Programm belegt einen Teilbereich eines Arrays der Länge arrLen mit zufälligen Werten,
// gibt diesen Teilbereich aus und ermittelt den kleinsten Wert in diesem Teilbereich.
func main() {
	var arr [arrLen]int
	tatsLen := rand.Intn(arrLen-1) + 1
	for i := 0; i < tatsLen; i++ {
		arr[i] = rand.Intn(150) - rand.Intn(50) + rand.Intn(50)
	}
	fmt.Printf("Von den %v Werten im Teilbereich\n%v\nist %v der kleinste Wert.\n", tatsLen, arr[0:tatsLen], arrayRekMin(arr, 0, tatsLen))
}
