package main

import "fmt"

func a() {
	var laenge float64
	fmt.Println("Bitte geben Sie die Länge in Metern ein:")
	fmt.Scan(&laenge)
	fmt.Printf("%.2f Metern entsprechen:\n", laenge)

	var mm, km, Zoll, sm, Lj float64

	mm = laenge * 1000
	km = laenge / 1000
	Zoll = laenge * 39.37
	sm = laenge * 3.28
	// Lichtjahr = ca. 9.461.000.000.000 km
	Lj = laenge * 9.461e15

	fmt.Printf("%f Meter sind %f Millimeter\n", laenge, mm)
	fmt.Printf("%f Meter sind %f Kilometer\n", laenge, km)
	fmt.Printf("%f Meter sind %f Zoll\n", laenge, Zoll)
	fmt.Printf("%f Meter sind %f Fuß\n", laenge, sm)
	fmt.Printf("%f Meter sind %f Lichtjahr\n", laenge, Lj)

}
