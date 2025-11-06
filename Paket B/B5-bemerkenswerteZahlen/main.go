package main

import "fmt"

// Das Programm liest eine natürliche Zahl ein und gibt die nächste natürliche
// Zahl aus, deren Quadrat die Summe der ganzen Zahlen von 1 bis n (mit
// beliebigem natürlichen n) ist.
func main() {
	var start int
	fmt.Println("Bitte geben Sie als Startwert eine natürliche Zahl ein: ")
	_, err := fmt.Scan(&start)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if start < 0 {
		fmt.Println("Sie haben als Startwert", start, "eingegeben.")
		fmt.Println("Der Startwert muß aber eine natürliche Zahl sein!")
		return
	}

	fmt.Println("Die erste bemerkenswerte Zahl größer als", start, "ist", bemerkenswert(start))
}
