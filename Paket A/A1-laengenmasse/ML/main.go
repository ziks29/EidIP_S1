package main

// Für die formatierten Ein- und Ausgaben muss das Paket fmt importiert werden.
import "fmt"

// Das Programm rechnet eine Länge, die in Metern eingegeben wird,
// in mm, km, Zoll, Seemeilen und Lichtjahre um.
func main() {
	var l float64
	fmt.Println("Geben Sie eine Länge (in Metern) ein:")

	// fmt.Scan liest die Tastatureingabe in die Variable l ein.
	// Dazu wird der Funktion als Argument ein sogenannter Zeiger auf die
	// Variable l übergeben, um der Funktion den direkten Zugriff auf diese
	// Variable zu ermöglichen (siehe Exkurs 5.8 im Lerntexts).
	//
	// Eine Besonderheit von Go ist, dass Funktionen mehr als einen
	// Rückgabewert haben können (mehr dazu in Kapitel 7.1 des Lerntexts).
	// Der zweite Rückgabewert der Funktion fmt.Scan gibt an, ob beim
	// Einlesen ein Fehler aufgetreten ist; so könnten z.B. Buchstaben anstatt
	// einer Gleitkommazahl eingeben worden sein.
	_, err := fmt.Scan(&l)

	// Im fehlerfreien Fall ist die Variable err mit dem speziellen Null-Wert
	// nil belegt; im Fehlerfall hingegen enthält sie eine entsprechende
	// Fehlermeldung. Diese wird dann ggf. hier ausgegeben und das Programm
	// anschließend beendet.
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}

	// Ausgabe der Ergebnisse
	// Anmerkung: Das Verb %g ist bei der Formatierung von Floats flexibler als %f:
	//            %g verwendet bei Bedarf die Exponentialdarstellung
	//            und gibt Nachkommastellen nur dann aus, wenn sie nicht alle 0 sind
	//            (siehe Doku von Printf).
	fmt.Println()
	fmt.Println(l, "Meter entsprechen: ")
	fmt.Printf("%g mm\n", l*1000)                    // 1m = 1000mm
	fmt.Printf("%g km\n", l/1000)                    // 1m = 1/1000 km
	fmt.Printf("%.3f Zoll\n", l*100/2.54)            // 1m = 100cm, 1cm = 1/2.54 Zoll
	fmt.Printf("%.3g sm\n", l/1852)                  // 1m = 1/1852 sm
	fmt.Printf("%.3g Lj\n", l/9_460_730_472_580_800) // 1m = 1/9460730472580800 Lj
}
