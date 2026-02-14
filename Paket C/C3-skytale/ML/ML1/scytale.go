package main

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func encrypt(klartext string, scytale int) string {
	result := ""
	// wenn die letzte Zeile auf dem Stab am Ende Leerstellen enthält
	// die nicht mit angegeben wurden, werden diese ergänzt.
	rest := len(klartext) % scytale
	if rest > 0 {
		for i := rest; i < scytale; i++ {
			klartext = klartext + " "
		}
	}
	// Wie lang sind die Textzeilen auf dem Stab?
	zlen := len(klartext) / scytale

	// Der Text befindet sich in einem zlen x scytale großen Raster auf dem Stab.
	// Dieses Raster spaltenweise durchlaufen ...
	for spalte := 0; spalte < zlen; spalte++ {
		for zeile := 0; zeile < scytale; zeile++ {
			/*
				 ... und die Position des nächsten Buchstabens in klartext bestimmen,
				 der in result übertragen werden muss.
				Pro Buchstaben geht es mit der Schrittweite zlen durch den String.
				Immer wenn die letzte Zeile erreicht wurde, muss für die nächste Spalte
				um eine Position versetzt wieder vorne begonnen werden.

				 Das lässt sich am besten anhand eines Beispiels nachvollziehen:
				 012345 soll mit dem Stabdurchmesser 2 verschlüsselt werden,
				 wird also folgendermaßen auf den Stab geschrieben:
				 0 1 2
				 3 4 5
				 Zur Verschlüsselung müssen die Buchstaben an den Positionen
				 0,3,1,4,2,5 in result übertragen werden.
				 \ / ^- nach der letzten Zeile um 1 versetzt von vorn
				  Schrittweite zlen=3
			*/
			pos := zeile*zlen + spalte
			result = result + klartext[pos:pos+1]
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
	for zeile := 0; zeile < scytale; zeile++ {
		for spalte := 0; spalte < zlen; spalte++ {
			pos := zeile + spalte*scytale
			result = result + geheimtext[pos:pos+1]
		}
	}
	for result[len(result)-1:] == " " {
		result = result[:len(result)-1]
	}
	return result
}
