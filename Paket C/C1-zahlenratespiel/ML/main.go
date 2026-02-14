package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Das Programm implementiert ein Zahlenratespiel gemäß
// der Spielregeln aus der Aufgabenstellung.
func main() {
	const zbmax = 100
	const diffheiß = 6
	var versuch int
	var differenz int
	geheim := rand.Intn(zbmax + 1)
	fmt.Println("Versuchen Sie eine natürliche Zahl zu erraten. Sie liegt zwischen 0 und 100.")
	for nr := 1; ; nr++ {
		fmt.Printf("Ihr %v. Versuch:\n", nr)
		_, err := fmt.Scan(&versuch)
		if err != nil {
			fmt.Println("Fehler: ", err)
			return
		}
		vorigeDifferenz := differenz
		differenz = int(math.Abs(float64(versuch - geheim)))
		switch {
		case differenz == 0:
			fmt.Println("Herzlichen Glückwunsch! Sie haben die Zahl in", nr, "Versuchen erraten.")
			return
		case differenz < diffheiß:
			fmt.Println("heiß")
		default:
			if nr == 1 {
				fmt.Println("kalt")
			} else {
				switch {
				case vorigeDifferenz == differenz:
					fmt.Println("gleicher Abstand")
				case vorigeDifferenz > differenz:
					fmt.Println("wärmer")
				default:
					fmt.Println("kälter")
				}
			}
		}
	}
}
