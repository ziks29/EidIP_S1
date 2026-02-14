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
	neuesElement := &listenElement{wert: neuerWert} // Element erzeugen
	neuesElement.vorheriges = nil
	neuesElement.nächstes = liste.anfang

	// falls die Liste leer ist, muss das neue Element zugleich auch Ende der Liste werden.
	if liste.anfang == nil {
		liste.ende = neuesElement
	} else {
		// ansonsten wird das neue Element vorne vorgehangen
		liste.anfang.vorheriges = neuesElement
	}

	// neuesElement Element wird Anfang der Liste
	liste.anfang = neuesElement

	return
}

// Fügt ein Element mit Wert neuerWert ans Ende der Liste an,
// d.h. hinter das bisherige End-Element.
func einfügenAmEnde(liste *doppeltVerketteteListe, neuerWert int) {
	neuesElement := &listenElement{wert: neuerWert} // Element erzeugen
	neuesElement.vorheriges = liste.ende
	neuesElement.nächstes = nil

	// falls die Liste leer ist, muss das neue Element zugleich auch Anfang der Liste werden
	if liste.ende == nil {
		liste.anfang = neuesElement
	} else {
		// ansonsten wird das neue Element hinten angehangen
		liste.ende.nächstes = neuesElement
	}

	// neuesElement Element wird Ende der Liste
	liste.ende = neuesElement

	return
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

	// beide Zeiger des neuen Elements so setzen,
	// dass es links vom Element, das bisher an Position pos war, platziert wird
	neuesElement.vorheriges = element.vorheriges
	neuesElement.nächstes = element

	// Die jeweiligen Zeiger der Elemente links und rechts vom neuen Element so "umbiegen",
	// dass sie auf das neue Element zeigen:

	// wenn links vom neuen Element ist kein Element existiert, ...
	if neuesElement.vorheriges == nil {
		// ... dann bildet das neue Element nun den Anfang dr Liste
		liste.anfang = neuesElement
	} else {
		// ... ansonsten den Zeiger "nächstes" des Elements links vom neuen Element anpassen
		neuesElement.vorheriges.nächstes = neuesElement
	}
	// den Zeiger "vorheriges" des Elements rechts vom neuen Element anpassen
	neuesElement.nächstes.vorheriges = neuesElement
}

// Liefert das Element an Position pos der Liste zurück.
func gibElementAnPosition(liste *doppeltVerketteteListe, pos int) *listenElement {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Position %v ist unzulässig in einer Liste der Länge %v.\n", pos, längeDerListe(liste))
		return nil
	}

	// abhängig von der gesuchten Position wird die Liste von vorne oder von hinten durchlaufen,
	// um möglichst schnell zum entsprechenden Element zu gelangen
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
		// Das Element rechts vom Anfangs-Element soll neuer Anfang werden,
		// daher muss sein "vorheriges" Zeiger auf nil gesetzt werden.
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
		// Das Element links vom End-Element soll neues Ende werden,
		// daher muss sein "nächstes" Zeiger auf nil gesetzt werden.
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
		// die jeweiligen Zeiger der Elemente links und rechts vom zu entfernenden Element entsprechend "umbiegen"
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
