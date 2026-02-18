package main

import (
	"strconv"
)

// spiegelzahl gibt die Spiegelzahl von zahl zurück.
// zahl muss eine natürliche Zahl sein.
func spiegelzahl(zahl int) int {
	s := strconv.Itoa(zahl)
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	result, _ := strconv.Atoi(reversed)
	return result
}
