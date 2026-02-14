package main

import "fmt"

// Das Programm liest eine Zahl ein und gibt ihre Spiegelzahl aus.
func main() {
	var zahl int
	fmt.Println("Bitte geben Sie eine natürliche Zahl ein:")
	_, err := fmt.Scan(&zahl)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if zahl < 0 {
		fmt.Println(zahl, "ist keine natürliche Zahl.")
	} else if zahl != 0 && zahl%10 == 0 {
		fmt.Println("Zu", zahl, "gibt es keine Spiegelzahl.")
	} else {
		fmt.Printf("Die Spiegelzahl von %v ist %v.\n", zahl, spiegelzahl(zahl))
	}
}
