package main

import "fmt"

// Verbund-Typ für ein Listen-Element
type listenElement struct {
	wert                 int
	vorheriges, nächstes *listenElement
}

// Verbund-Typ für eine Liste
// Hinweis:
// Die Elemente der Liste sind (wie ein Array) null-indiziert,
// d.h. das Element am Anfang hat Position 0 und das am Ende hat Position listenlänge-1.
type doppeltVerketteteListe struct {
	anfang, ende *listenElement
}
	
// Gibt alle Werte der Liste von Anfang bis Ende aus.
func listeAusgeben(liste *doppeltVerketteteListe) {
	element := liste.anfang
	fmt.Printf("-> ")
	for element != nil {
		fmt.Printf("%d -> ", element.wert)
		element = element.nächstes
	}
	fmt.Printf("nil\n")
}

// Gibt alle Werte der Liste von Ende bis Anfang aus.
func listeAusgebenRückwärts(liste *doppeltVerketteteListe) {
	element := liste.ende
	fmt.Printf("<- ")
	for element != nil {
		fmt.Printf("%d <- ", element.wert)
		element = element.vorheriges
	}
	fmt.Printf("nil\n")
}

// Liefert die Länge der Liste (d.h. die Anzahl ihrer Elemente) zurück.
func längeDerListe(liste *doppeltVerketteteListe) int {
	länge := 0
	element := liste.anfang
	for element != nil {
		länge++
		element = element.nächstes
	}
	return länge
}

// Fügt ein Element mit Wert neuerWert an den Anfang der Liste ein,
// d.h. vor das bisherige Anfangs-Element.
func einfügenAmAnfang(liste *doppeltVerketteteListe, neuerWert int) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Fügt ein Element mit Wert neuerWert ans Ende der Liste an,
// d.h. hinter das bisherige End-Element.
func einfügenAmEnde(liste *doppeltVerketteteListe, neuerWert int) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Fügt ein neues Element mit Wert neuerWert an der Position pos eines bereits existierenden Elements
// in die Liste ein, wobei das existierende Element nach rechts verschoben wird.
func einfügenAnPosition(liste *doppeltVerketteteListe, neuerWert int, pos int) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Liefert das Element an Position pos der Liste zurück.
func gibElementAnPosition(liste *doppeltVerketteteListe, pos int) *listenElement {
	// todo: Implementieren Sie den Rumpf der Funktion
	return nil
}

// Liefert den Wert des Elements an Position pos der Liste zurück.
func gibWertAnPosition(liste *doppeltVerketteteListe, pos int) int {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Position %v ist unzulässig in einer Liste der Länge %v.\n", pos, längeDerListe(liste))
		return -1
	}

	return gibElementAnPosition(liste, pos).wert
}

// Entfernt das Element am Anfang der Liste.
func entferneAnfang(liste *doppeltVerketteteListe) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Entfernt das Element am Ende der Liste.
func entferneEnde(liste *doppeltVerketteteListe) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Entfernt das Element an der Position pos aus der Liste.
func entferneAnPosition(liste *doppeltVerketteteListe, pos int) {
	// todo: Implementieren Sie den Rumpf der Funktion
}

// Liefert die Position des ersten Auftretens des Werts gesuchterWert in der Liste zurück,
// oder -1, wenn der Wert nicht in der Liste enthalten ist.
// Wenn gesuchterWert mehrfach in der Liste enthalten ist,
// wird also immer die Position des Wertes zurückgeliefert, der am weitesten links in der Liste steht.
func gibPositionVonWert(liste *doppeltVerketteteListe, gesuchterWert int) int {
	// todo: Implementieren Sie den Rumpf der Funktion
	return -1
}

// Liefert das (am weitesten links stehende) Element der Liste zurück,
// das den Wert gesuchterWert hat,
// oder nil, wenn kein Element mit Wert gesuchterWert in der Liste enthalten ist.
// Wenn mehrere Elemente mit gesuchterWert in der Liste enthalten sind,
// wird also immer dasjenige Element davon zurückgeliefert, das am weitesten links in der Liste steht.
func gibElementMitWert(liste *doppeltVerketteteListe, gesuchterWert int) *listenElement {
	i := gibPositionVonWert(liste, gesuchterWert)
	if i == -1 {
		return nil
	} else {
		return gibElementAnPosition(liste, i)
	}
}

// true, wenn der Wert gesuchterWert in der Liste enthalten ist,
// ansonsten false.
func istWertInListe(liste *doppeltVerketteteListe, gesuchterWert int) bool {
	return gibPositionVonWert(liste, gesuchterWert) != -1
}
