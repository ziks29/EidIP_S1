package main

import (
	"fmt"
)

// Verbund-Typ für einen binären Knoten
type knoten struct {
	wert          int
	links, rechts *knoten
}

// ----- bekannte Funktionen aus Lerntext -----

// Die rekursive Funktion fügt einen neuen Knoten mit dem Wert
// neuerWert so in einen binären Suchbaum ein,
// dass die Suchbaum-Eigenschaft danach weiterhin erfüllt ist.
// wurzel muss die Wurzel eines Suchbaums referenzieren,
// die bereits einen Wert enthält (d.h. wurzel darf nicht nil sein).
// neuerWert darf noch nicht im Suchbaum enthalten sein.
func einfügen(wurzel *knoten, neuerWert int) {
	if neuerWert < wurzel.wert {
		if wurzel.links == nil {
			// neuen Knoten erzeugen
			wurzel.links = &knoten{wert: neuerWert}
		} else {
			// rekursiv im linken Teilbaum nach
			// passender Einfüge-Position suchen
			einfügen(wurzel.links, neuerWert)
		}
	} else {
		if wurzel.rechts == nil {
			// neuen Knoten erzeugen
			wurzel.rechts = &knoten{wert: neuerWert}
		} else {
			// rekursiv im rechten Teilbaum nach
			// passender Einfüge-Position suchen
			einfügen(wurzel.rechts, neuerWert)
		}
	}
}

// Die rekursive Funktion sucht in einem Suchbaum
// nach dem Wert suchwert.
// wurzel muss die Wurzel eines Suchbaums referenzieren.
// Wenn suchwert im Suchbaum enthalten ist, wird ein Zeiger
// auf den entsprechenden Knoten zurückgeliefert, ansonsten nil.
func suchen(wurzel *knoten, suchwert int) *knoten {
	if wurzel == nil || wurzel.wert == suchwert {
		return wurzel
	}

	if suchwert < wurzel.wert {
		return suchen(wurzel.links, suchwert)
	} else {
		return suchen(wurzel.rechts, suchwert)
	}
}

// Die Funktion liefert true zurück, wenn der Wert suchwert
// in einem Suchbaum enthalten ist, ansonsten false.
// wurzel muss die Wurzel eines Suchbaums referenzieren.
func istWertInBaum(wurzel *knoten, suchwert int) bool {
	return suchen(wurzel, suchwert) != nil
}

// gibt die Werte des Suchbaums in aufsteigender Reihenfolge aus
func baumAusgeben(wurzel *knoten) {
	if wurzel != nil {
		// Traversierung des Baums in sog. In-Order (L-Knoten-R)
		baumAusgeben(wurzel.links)
		fmt.Printf("%d ", wurzel.wert)
		baumAusgeben(wurzel.rechts)
	}
}

// liefert die Summe aller Werte des Baums zurück
func summeDesBaums(wurzel *knoten) int {
	summe := 0
	if wurzel != nil {
		// Traversierung des Baums in sog. Pre-Order (Knoten-L-R)
		summe += wurzel.wert
		summe += summeDesBaums(wurzel.links)
		summe += summeDesBaums(wurzel.rechts)
	}
	return summe
}

// ----- ab hier neu zu implementierende Funktionen -----

// Liefert die Höhe des Baums zurück.
func höheDesBaums(wurzel *knoten) int {
	if wurzel == nil {
		return 0
	} else {
		höheLinks := höheDesBaums(wurzel.links)
		höheRechts := höheDesBaums(wurzel.rechts)
		if höheLinks > höheRechts {
			return höheLinks + 1
		} else {
			return höheRechts + 1
		}
	}
}

// Liefert einen Zeiger auf den Knoten zurück,
// der den kleinsten Wert im Suchbaum enthält.
// wurzel muss die Wurzel eines Suchbaums referenzieren,
// die bereits einen Wert enthält (d.h. wurzel darf nicht nil sein).
func kleinsterKnoten(wurzel *knoten) *knoten {
	if wurzel.links != nil {
		return kleinsterKnoten(wurzel.links)
	} else {
		return wurzel
	}
}

// Sucht in einem Suchbaum nach dem Wert suchwert
// und liefert dessen Höhe zurück.
// wurzel muss die Wurzel eines Suchbaums referenzieren.
// höhe muss beim initialen Aufruf 1 sein.
// Wenn suchwert im Suchbaum enthalten ist, wird die Höhe
// des entsprechenden Knotens zurückgeliefert, ansonsten -1.
func höheDesWerts(wurzel *knoten, suchwert int, höhe int) int {
	if wurzel == nil {
		return -1
	}

	if wurzel.wert == suchwert {
		return höhe
	}

	if suchwert < wurzel.wert {
		return höheDesWerts(wurzel.links, suchwert, höhe+1)
	} else {
		return höheDesWerts(wurzel.rechts, suchwert, höhe+1)
	}
}

// Erhöht jeden Wert im Baum um den Wert 1.
func inkrementiereAlleWerte(wurzel *knoten) {
	if wurzel != nil {
		// Traversierung des Baums in sog. Pre-Order (Knoten-L-R)
		wurzel.wert = wurzel.wert + 1
		inkrementiereAlleWerte(wurzel.links)
		inkrementiereAlleWerte(wurzel.rechts)
	}
}
