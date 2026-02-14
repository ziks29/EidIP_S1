package main

import (
	"fmt"
)

func m() {
	var meter float64 // Eingabe der Länge in Metern
	fmt.Print("Geben Sie eine Länge in Metern ein: ")
	fmt.Scanln(&meter)
	fuss := meter * 3.28084
	km := meter / 1000
	mm := meter * 1000
	zoll := meter * 39.3701
	seemeile := meter / 1852
	lichtjahr := meter / 9.461e+15

	fmt.Printf("Länge in Fuß: %.2f ft\n", fuss)
	fmt.Printf("Länge in Kilometern: %.2f km\n", km)
	fmt.Printf("Länge in Millimetern: %.2f mm\n", mm)
	fmt.Printf("Länge in Zoll: %.2f in\n", zoll)
	fmt.Printf("Länge in Seemeilen: %.6f sm\n", seemeile)
	fmt.Printf("Länge in Lichtjahren: %.3g Lj\n", lichtjahr)

}
