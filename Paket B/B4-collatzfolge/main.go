package main

// collatzFolge gibt die Collatz-Folge für die gegebene Startzahl zurück.
import "fmt"

func collatzFolge(startzahl int) []int {
	folge := []int{startzahl} // Initialisiere die Folge mit der Startzahl
	zahl := startzahl
	for zahl != 1 {
		if zahl%2 == 0 {
			zahl = zahl / 2 // Wenn die Zahl gerade ist, teile sie durch 2
		} else {
			zahl = 3*zahl + 1 // Wenn die Zahl ungerade ist, multipliziere sie mit 3 und addiere 1
		}
		folge = append(folge, zahl) // Füge die neue Zahl zur Folge hinzu
	}
	return folge
}

func main() {
	// Ask user and scan
	var startzahl int
	fmt.Print("Geben Sie eine Startzahl für die Collatz-Folge ein: ")
	fmt.Scanln(&startzahl)
	folge := collatzFolge(startzahl)
	fmt.Println("Die Collatz-Folge ist:", folge)

}
