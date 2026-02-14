package main

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func encrypt(klartext string, scytale int) string {
	result := ""
	// Wenn die letzte Zeile auf dem Stab am Ende Leerstellen enthält,
	// die nicht mit angegeben wurden, werden diese ergänzt.
	rest := len(klartext) % scytale
	if rest > 0 {
		for i := rest; i < scytale; i++ {
			klartext = klartext + " "
		}
	}

	zlen := len(klartext) / scytale
	// Der folgende im Lerntext nicht eingeführte Befehl erzeugt ein Slice (quasi ein dynamisches Array)
	// mit der Länge scytale.
	raster := make([]string, scytale)
	// In das Slice wird klartext zeilenweise geschrieben...
	pos := 0
	for zeile := 0; zeile < scytale; zeile++ {
		for spalte := 0; spalte < zlen; spalte++ {
			raster[zeile] = raster[zeile] + klartext[pos:pos+1]
			pos++
		}
	}
	// ... und dann spaltenweise wieder ausgelesen, um so die
	// verschlüsselte Version zu erhalten.
	for spalte := 0; spalte < zlen; spalte++ {
		for zeile := 0; zeile < scytale; zeile++ {
			result = result + raster[zeile][spalte:spalte+1]
		}
	}
	// Leerzeichen bzw. Unterstriche am Ende entfernen, da sie auf dem Band nicht erkennbar sind.
	for result[len(result)-1:] == " " {
		result = result[:len(result)-1]
	}
	return result
}

// decrypt gibt die entschlüsselte Version von geheimtext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func decrypt(geheimtext string, scytale int) string {
	// Funktionsweise analog der encrypt-Funktion.
	result := ""
	rest := len(geheimtext) % scytale
	if rest > 0 {
		for i := rest; i < scytale; i++ {
			geheimtext = geheimtext + " "
		}
	}
	zlen := len(geheimtext) / scytale
	raster := make([]string, zlen)
	pos := 0
	for spalte := 0; spalte < zlen; spalte++ {
		for zeile := 0; zeile < scytale; zeile++ {
			raster[spalte] = raster[spalte] + geheimtext[pos:pos+1]
			pos++
		}
	}
	for zeile := 0; zeile < scytale; zeile++ {
		for spalte := 0; spalte < zlen; spalte++ {
			result = result + raster[spalte][zeile:zeile+1]
		}
	}
	for result[len(result)-1:] == " " {
		result = result[:len(result)-1]
	}
	return result
}
