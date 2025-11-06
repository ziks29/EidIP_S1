package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	// längste Seite (d.h. mögliche Hypotenuse) bestimmen
	// und dann mittels Satz des Pythagoras prüfen
	if z > x && z > y {
		// Pythagoras bzgl. z als Hypotenuse
		return x*x+y*y == z*z
	} else if y > x && y > z {
		// Pythagoras bzgl. y als Hypotenuse
		return x*x+z*z == y*y
	} else {
		// Pythagoras bzgl. x als Hypotenuse
		return z*z+y*y == x*x
	}
}
