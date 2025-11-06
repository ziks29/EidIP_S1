package main

import "fmt"

// Das Programm gibt alle perfekten Zahlen zwischen 1 und 10000 aus.
func main() {
	const Min = 1
	const Max = 10000
	fmt.Println("Zwischen", Min, "und", Max, "liegen folgende perfekten Zahlen:")
	for kandidat := Min; kandidat <= Max; kandidat++ {
		if istPerfekt(kandidat) {
			fmt.Println(kandidat)
		}
	}
}
