package main

// Um die Wurzelfunktion verwenden zu können, muss auch das Paket math
// importiert werden.
import (
	"fmt"
	"math"
)

// Das Programm berechnet einige geometrische Größen eines Quaders aus seinen eingegebenen Seitenlängen.
func main() {
	var a, b, c float64
	fmt.Println("Bitte geben Sie die drei Seitenlängen des Quaders ein:")
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}

	// Berechnung der geometrischen Größen
	volumen := a * b * c
	kanten := 4 * (a + b + c)
	oberfläche := 2.0 * (a*b + a*c + b*c)
	diagonale := math.Sqrt(a*a + b*b + c*c)
	radius := diagonale / 2.0

	// Format-String, der in allen Printf Ausgaben genutzt wird
	format := "%.3g"

	// Ausgabe der Ergebnisse
	fmt.Println()
	fmt.Printf("Ein %v × %v × %v Quader hat die geometrischen Größen:\n", a, b, c)
	fmt.Printf("Volumen: "+format+"\n", volumen)
	fmt.Printf("Kantensumme: "+format+"\n", kanten)
	fmt.Printf("Oberfläche: "+format+"\n", oberfläche)
	fmt.Printf("Umkugelradius: "+format+"\n", radius)
	fmt.Printf("Raumdiagonale: "+format+"\n", diagonale)
}
