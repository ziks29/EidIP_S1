package main

import "strings"

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func encrypt(klartext string, scytale int) string {
	// Länge des Textes berechnen
	length := len(klartext)

	// Wenn die Länge ein Vielfaches von scytale ist, ist eine Padding nicht nötig
	// Ansonsten Padding mit Leerzeichen
	if length%scytale != 0 {
		padding := scytale - (length % scytale)
		klartext = klartext + strings.Repeat(" ", padding)
		length = len(klartext)
	}

	// Anzahl der Spalten berechnen
	columns := length / scytale

	// Verschlüsselte Nachricht spaltenweise aus dem Raster lesen
	result := ""
	for col := 0; col < columns; col++ {
		for row := 0; row < scytale; row++ {
			result += string(klartext[row*columns+col])
		}
	}

	// Trailing spaces entfernen
	result = strings.TrimRight(result, " ")

	return result
}

// decrypt gibt die entschlüsselte Version von geheimtext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func decrypt(geheimtext string, scytale int) string {
	// Länge des Texts berechnen
	length := len(geheimtext)

	// Wenn die Länge ein Vielfaches von scytale ist, ist eine Padding nicht nötig
	// Ansonsten Padding mit Leerzeichen
	if length%scytale != 0 {
		padding := scytale - (length % scytale)
		geheimtext = geheimtext + strings.Repeat(" ", padding)
		length = len(geheimtext)
	}

	// Anzahl der Spalten berechnen
	columns := length / scytale

	// Raster mit Leerzeichen initialisieren
	raster := make([][]rune, scytale)
	for i := range raster {
		raster[i] = make([]rune, columns)
	}

	// Geheimtext spaltenweise in das Raster schreiben
	index := 0
	for col := 0; col < columns; col++ {
		for row := 0; row < scytale; row++ {
			raster[row][col] = rune(geheimtext[index])
			index++
		}
	}

	// Klartext zeilenweise aus dem Raster auslesen
	result := ""
	for row := 0; row < scytale; row++ {
		for col := 0; col < columns; col++ {
			result += string(raster[row][col])
		}
	}

	// Trailing spaces entfernen
	result = strings.TrimRight(result, " ")

	return result
}
