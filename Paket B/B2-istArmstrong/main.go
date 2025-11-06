package main

import "fmt"

// Das Programm liest zwei natürliche Zahlen ein und gibt alle
// Armstrong-Zahlen aus, die innerhalb dieser Grenzen liegen
func main() {
	var start, ende int
	fmt.Println("Wir bestimmen die Armstrongzahlen in einem Zahlenbereich.")
	fmt.Println("Geben Sie bitte zwei natürliche Zahlen als Grenzen ein: ")
	_, err := fmt.Scan(&start, &ende)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if start < 0 || ende < 0 || ende < start {
		fmt.Println("Sie haben die Grenzen", start, "und", ende, "eingegeben.")
		fmt.Println("Die beiden Zahlen müssen aber positiv sein und die zweite darf nicht größer als die erste sein.")
		return
	}
	fmt.Println("Im Bereich von", start, "bis", ende, "liegen folgende Armstrongzahlen:")
	for kandidat := start; kandidat <= ende; kandidat++ {
		if istArmstrong(kandidat) {
			fmt.Println(kandidat)
		}
	}
}
