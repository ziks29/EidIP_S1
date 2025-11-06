package main

import (
	"fmt"  // "fmt" package is imported to handle input and output
	"math" // "math" package is imported to use mathematical functions
)

func main() {
	// Declare a variable to hold 3 length in meters
	var a, b, c float64
	fmt.Print("Bitte geben Sie die drei Seitenlängen des Quaders ein:\n")
	_, err := fmt.Scan(&a, &b, &c) // Read input values
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}

	// Calculate volume, surface area, and total edge length
	volumen := a * b * c
	oberfläche := 2 * (a*b + a*c + b*c)
	kanten := 4 * (a + b + c)
	diagonale := math.Sqrt(a*a + b*b + c*c)
	radius := diagonale / 2.0

	// Output the results
	fmt.Printf("\nEin %v × %v × %v Quader hat die geometrischen Größen:\n", a, b, c)
	fmt.Printf("Volumen: %.3g\n", volumen)
	fmt.Printf("Oberfläche: %.3g\n", oberfläche)
	fmt.Printf("Kantensumme: %.3g\n", kanten)
	fmt.Printf("Umkugelradius: %.3g\n", radius)
	fmt.Printf("Raumdiagonale: %.3g\n", diagonale)
}
