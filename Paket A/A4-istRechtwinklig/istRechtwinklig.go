package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	// c^2 = a^2 + b^2
	if x*x+y*y != z*z && x*x+z*z != y*y && y*y+z*z != x*x {
		return false
	}

	return true
}
