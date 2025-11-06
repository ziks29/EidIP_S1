package main

import "strings"

// buchstabenSortieren gibt einen String zurück, in dem alle Buchstaben
// von wort alphabetisch sortiert sind. In wort mehrfach vorkommende
// Buchstaben kommen im resultierendem String genauso oft vor. Großbuchstaben
// in wort kommen im resultierendem String als Kleinbuchstaben vor.
// wort darf nur die Zeichen a-z und A-Z enthalten.
func buchstabenSortieren(wort string) string {
	// Es soll keine Rolle spielen, ob die Buchstaben groß oder klein geschrieben werden.
	// Daher werden alle Buchstaben in Kleinbuchstaben umgewandelt.
	klein := strings.ToLower(wort)
	result := ""
	// für die Buchstaben a-z
	for buchstabe := "a"[0]; buchstabe <= "z"[0]; buchstabe++ {
		// für jedes Vorkommen des Buchstabens in der kleingeschriebenen Version
		// von wort, diesen in den zurückzugebenden String übertragen.
		for pos := 0; pos < len(klein); pos++ {
			if buchstabe == klein[pos] {
				result = result + klein[pos:pos+1]
			}
		}
	}
	return result
}
