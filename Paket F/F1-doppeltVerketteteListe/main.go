package main

import "fmt"

func main() {

	liste := &doppeltVerketteteListe{}

	listeAusgeben(liste)
	einfügenAmEnde(liste, 8)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 17)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 12)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 2)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 23)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 20)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 17)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 5)
	listeAusgeben(liste)
	einfügenAmEnde(liste, 14)
	listeAusgeben(liste)

	fmt.Println()
	listeAusgeben(liste)
	listeAusgebenRückwärts(liste)
	fmt.Println("Länge der Liste:", längeDerListe(liste))

	fmt.Println()
	listeAusgeben(liste)
	einfügenAnPosition(liste, 333, 3)
	listeAusgeben(liste)
	einfügenAnPosition(liste, 44, 0)
	listeAusgeben(liste)
	einfügenAnPosition(liste, 111, 1)
	listeAusgeben(liste)
	einfügenAnPosition(liste, 999, 11)
	listeAusgeben(liste)
	fmt.Println("Länge der Liste:", längeDerListe(liste))
	fmt.Println()

	fmt.Println(gibWertAnPosition(liste, 0))
	fmt.Println(gibWertAnPosition(liste, 3))
	fmt.Println(gibWertAnPosition(liste, 5))
	fmt.Println(gibWertAnPosition(liste, 6))
	fmt.Println(gibWertAnPosition(liste, 7))
	fmt.Println(gibWertAnPosition(liste, 12))
	listeAusgeben(liste)
	fmt.Println()

	entferneAnPosition(liste, 12)
	entferneAnPosition(liste, 0)
	listeAusgeben(liste)
	entferneAnPosition(liste, 3)
	listeAusgeben(liste)
	entferneAnPosition(liste, 1)
	listeAusgeben(liste)
	entferneAnPosition(liste, 7)
	listeAusgeben(liste)
	fmt.Println()

	fmt.Println(istWertInListe(liste, 333))
	fmt.Println(istWertInListe(liste, 66))
	fmt.Println()

	fmt.Println(gibPositionVonWert(liste, 333))
	fmt.Println(gibPositionVonWert(liste, 66))
	fmt.Println()

	fmt.Println(gibElementMitWert(liste, 333))
	fmt.Println(gibElementMitWert(liste, 66))
	fmt.Println()
}
