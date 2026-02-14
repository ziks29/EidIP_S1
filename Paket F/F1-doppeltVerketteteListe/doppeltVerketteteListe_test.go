package main

import (
	"testing"
)

func Test_leereListe(t *testing.T) {
	liste := &doppeltVerketteteListe{}

	if got, want := längeDerListe(liste), 0; got != want {
		t.Errorf("längeDerListe(liste): %v, want: %v", got, want)
	}

	if liste.anfang != nil || liste.ende != nil {
		t.Error("Anfang und Ende einer leeren Liste müssen beide nil sein.")
	}

}

func Test_einfügenAmAnfang(t *testing.T) {

	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}

	// Liste "von vorne" sukzessive mit Testwerten aus Slice befüllen,
	// wodurch sich die umgekehrte Reihenfolge der Testdaten ergibt
	for _, v := range testwerte {
		einfügenAmAnfang(liste, v)
	}
	listeAusgeben(liste)

	// Slice mit Testwerten in umgekehrter Reihenfolge erzeugen
	var umgekehrteReihenfolge []int
	for i := len(testwerte) - 1; i >= 0; i-- {
		umgekehrteReihenfolge = append(umgekehrteReihenfolge, testwerte[i])
	}

	subtest_listeMitVorgabeVergleichen(t, liste, umgekehrteReihenfolge)

}

func Test_einfügenAmEnde(t *testing.T) {

	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}

	// Liste "am Ende" sukzessive mit Testwerten aus Slice befüllen,
	// wodurch sich die gleiche Reihenfolge wie im Slice ergibt
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	subtest_listeMitVorgabeVergleichen(t, liste, testwerte)

}

func Test_einfügenAnPosition(t *testing.T) {
	liste := &doppeltVerketteteListe{}
	einfügenAmEnde(liste, 0)
	einfügenAmEnde(liste, 10)
	einfügenAmEnde(liste, 20)
	einfügenAmEnde(liste, 30)
	einfügenAmEnde(liste, 40)
	einfügenAmEnde(liste, 50)
	listeAusgeben(liste)

	einfügenAnPosition(liste, 15, 2)
	einfügenAnPosition(liste, 99, 6)
	einfügenAnPosition(liste, -11, 0)
	listeAusgeben(liste)

	subtest_listeMitVorgabeVergleichen(t, liste, []int{-11, 0, 10, 15, 20, 30, 40, 99, 50})

	// weiterer Test mit einer Liste mit nur einem Element
	liste = &doppeltVerketteteListe{}
	einfügenAmAnfang(liste, 22)
	einfügenAnPosition(liste, 5, 0)
	listeAusgeben(liste)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{5, 22})

}

func Test_gibElementAnPosition(t *testing.T) {
	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	for i, v := range testwerte {
		got, want := gibElementAnPosition(liste, i).wert, v
		if got != want {
			t.Errorf("gibElementAnPosition(liste, %v): %v, want: %v", i, got, want)
		}
	}

	// weiterer Test mit einer Liste mit nur einem Element
	liste = &doppeltVerketteteListe{}
	i, v := 0, 22
	einfügenAmAnfang(liste, v)
	got, want := gibElementAnPosition(liste, i).wert, v
	if got != want {
		t.Errorf("gibElementAnPosition(liste, %v): %v, want: %v", i, got, want)
	}

	// weiterer Test mit einer Liste mit nur zwei Elementen
	testwerte = []int{5, 10}
	liste = &doppeltVerketteteListe{}
	for _, v = range testwerte {
		einfügenAmEnde(liste, v)
	}
	for i, v = range testwerte {
		got, want = gibElementAnPosition(liste, i).wert, v
		if got != want {
			t.Errorf("gibElementAnPosition(liste, %v): %v, want: %v", i, got, want)
		}

	}
}

func Test_entferneAnfang(t *testing.T) {
	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	for i := 1; i < len(testwerte); i++ {
		entferneAnfang(liste)
		listeAusgeben(liste)
		got, want := liste.anfang.wert, testwerte[i]
		if got != want {
			t.Errorf("entferneAnfang(liste): %v, want: %v", got, want)
		}
	}

}

func Test_entferneEnde(t *testing.T) {
	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	for i := len(testwerte) - 2; i >= 0; i-- {
		entferneEnde(liste)
		listeAusgeben(liste)
		got, want := liste.ende.wert, testwerte[i]
		if got != want {
			t.Errorf("entferneEnde(liste): %v, want: %v", got, want)
		}
	}
}

func Test_entferneAnPosition(t *testing.T) {
	testwerte := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}

	liste := &doppeltVerketteteListe{}
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	entferneAnPosition(liste, 4)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{10, 11, 12, 13, 15, 16, 17, 18})

	entferneAnPosition(liste, 1)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{10, 12, 13, 15, 16, 17, 18})

	entferneAnPosition(liste, 5)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{10, 12, 13, 15, 16, 18})

	entferneAnPosition(liste, 0)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{12, 13, 15, 16, 18})

	entferneAnPosition(liste, 4)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{12, 13, 15, 16})

	entferneAnPosition(liste, 1)
	subtest_listeMitVorgabeVergleichen(t, liste, []int{12, 15, 16})

}

func Test_gibPositionVonWert(t *testing.T) {
	testwerte := []int{5, 10, 15, 20, 25, 30}

	liste := &doppeltVerketteteListe{}
	for _, v := range testwerte {
		einfügenAmEnde(liste, v)
	}
	listeAusgeben(liste)

	for i, v := range testwerte {
		got, want := gibPositionVonWert(liste, v), i
		if got != want {
			t.Errorf("gibPositionVonWert(liste, %v): %v, want: %v", v, got, want)
		}
	}
}

// Hilfsfunktion für Tests
func subtest_listeMitVorgabeVergleichen(t *testing.T, liste *doppeltVerketteteListe, vorgabe []int) {
	var got, want, i int
	var element *listenElement

	// Länge der Liste mit Vorgabe vergleichen
	got, want = längeDerListe(liste), len(vorgabe)
	if got != want {
		t.Errorf("längeDerListe(liste): %v, want: %v", got, want)
		return
	}

	// Vorwärts-Reihenfolge in Liste mit Vorgabe vergleichen
	element = liste.anfang
	i = 0
	for element != nil {
		got, want = element.wert, vorgabe[i]
		if got != want {
			t.Errorf("Fehler beim Vorwärts-Durchlauf: got=%v, want=%v", got, want)
		}
		element = element.nächstes
		i++
	}

	// Rückwärts-Reihenfolge in Liste mit Vorgabe vergleichen
	element = liste.ende
	i = len(vorgabe) - 1
	for element != nil {
		got, want = element.wert, vorgabe[i]
		if got != want {
			t.Errorf("Fehler beim Rückwärts-Durchlauf: got=%v, want=%v", got, want)
		}
		element = element.vorheriges
		i--
	}

}
