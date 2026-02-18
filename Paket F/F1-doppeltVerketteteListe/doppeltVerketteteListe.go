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
	neuesElement := &listenElement{wert: neuerWert}
	neuesElement.vorheriges = nil
	neuesElement.nächstes = liste.anfang

	if liste.anfang == nil {
		liste.ende = neuesElement
	} else {
		liste.anfang.vorheriges = neuesElement
	}

	liste.anfang = neuesElement
}

// Fügt ein Element mit Wert neuerWert ans Ende der Liste an,
// d.h. hinter das bisherige End-Element.
func einfügenAmEnde(liste *doppeltVerketteteListe, neuerWert int) {
	neuesElement := &listenElement{wert: neuerWert}
	neuesElement.vorheriges = liste.ende
	neuesElement.nächstes = nil

	if liste.ende == nil {
		liste.anfang = neuesElement
	} else {
		liste.ende.nächstes = neuesElement
	}

	liste.ende = neuesElement
}

// Fügt ein neues Element mit Wert neuerWert an der Position pos eines bereits existierenden Elements
// in die Liste ein, wobei das existierende Element nach rechts verschoben wird.
func einfügenAnPosition(liste *doppeltVerketteteListe, neuerWert int, pos int) {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Einfügen an Position %v einer Liste mit Länge %v nicht möglich.\n", pos, längeDerListe(liste))
		return
	}

	element := gibElementAnPosition(liste, pos)
	neuesElement := &listenElement{wert: neuerWert}

	neuesElement.vorheriges = element.vorheriges
	neuesElement.nächstes = element

	if neuesElement.vorheriges == nil {
		liste.anfang = neuesElement
	} else {
		neuesElement.vorheriges.nächstes = neuesElement
	}
	neuesElement.nächstes.vorheriges = neuesElement
}

// Liefert das Element an Position pos der Liste zurück.
func gibElementAnPosition(liste *doppeltVerketteteListe, pos int) *listenElement {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Position %v ist unzulässig in einer Liste der Länge %v.\n", pos, längeDerListe(liste))
		return nil
	}

	var element *listenElement
	if pos < längeDerListe(liste)/2 {
		element = liste.anfang
		for i := 0; i < pos; i++ {
			element = element.nächstes
		}
	} else {
		element = liste.ende
		for i := längeDerListe(liste) - 1; i > pos; i-- {
			element = element.vorheriges
		}
	}

	return element
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
	if liste.anfang == nil {
		fmt.Println("FEHLER: Die Liste ist leer.")
		return
	}

	if längeDerListe(liste) == 1 {
		liste.anfang = nil
		liste.ende = nil
	} else {
		liste.anfang.nächstes.vorheriges = nil
		liste.anfang = liste.anfang.nächstes
	}
}

// Entfernt das Element am Ende der Liste.
func entferneEnde(liste *doppeltVerketteteListe) {
	if liste.anfang == nil {
		fmt.Println("FEHLER: Die Liste ist leer.")
		return
	}

	if längeDerListe(liste) == 1 {
		liste.anfang = nil
		liste.ende = nil
	} else {
		liste.ende.vorheriges.nächstes = nil
		liste.ende = liste.ende.vorheriges
	}
}

// Entfernt das Element an der Position pos aus der Liste.
func entferneAnPosition(liste *doppeltVerketteteListe, pos int) {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Entfernen an Position %v einer Liste mit Länge %v nicht möglich.\n", pos, längeDerListe(liste))
		return
	}

	element := gibElementAnPosition(liste, pos)

	if element == liste.anfang {
		entferneAnfang(liste)
	} else if element == liste.ende {
		entferneEnde(liste)
	} else {
		element.vorheriges.nächstes = element.nächstes
		element.nächstes.vorheriges = element.vorheriges
	}
}

// Liefert die Position des ersten Auftretens des Werts gesuchterWert in der Liste zurück,
// oder -1, wenn der Wert nicht in der Liste enthalten ist.
// Wenn gesuchterWert mehrfach in der Liste enthalten ist,
// wird also immer die Position des Wertes zurückgeliefert, der am weitesten links in der Liste steht.
func gibPositionVonWert(liste *doppeltVerketteteListe, gesuchterWert int) int {
	if liste.anfang == nil {
		fmt.Println("FEHLER: Die Liste ist leer.")
		return -1
	}

	element := liste.anfang
	pos := 0
	for element != nil {
		if element.wert == gesuchterWert {
			return pos
		} else {
			element = element.nächstes
			pos++
		}
	}

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
