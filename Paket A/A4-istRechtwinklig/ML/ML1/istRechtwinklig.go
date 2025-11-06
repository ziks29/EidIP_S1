package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	// Da die Seitenlängen nicht nach Größe sortiert sind, wird der Satz des
	// Pythagoras für alle drei Seiten als mögliche Hypotenuse angewandt.
	// Sollte er für eine der Möglichkeiten gelten, ist das Dreieck rechtwinklig.
	// Daher verwenden wir hier eine oder-Verknüpfung.
	return x*x+y*y == z*z || x*x+z*z == y*y || z*z+y*y == x*x
}
