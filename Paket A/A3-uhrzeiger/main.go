package main

import (
	"fmt" // "fmt" package is imported to handle input and output
)

func main() {
	var h, m int
	fmt.Print("Geben Sie bitte eine Uhrzeit an, indem Sie zunächst\ndie Stunde (von 0 bis 23)und dann die Minute (von 0 bis 59)eingeben:\n")
	_, err := fmt.Scan(&h, &m) // Read input values
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}

	// Calculate the angles of the hour and minute hands
	minutenwinkel := m * 6
	stundenwinkel := float64(h%12)*30 + float64(m)*0.5

	// Output the results
	fmt.Println()

	fmt.Printf("Zeigerstellung um %02d:%02d: \n", h, m)
	fmt.Printf("Stundenwinkel: %3.1f° \n", stundenwinkel)
	fmt.Printf("Minutenwinkel: %3d° \n", minutenwinkel)

}
