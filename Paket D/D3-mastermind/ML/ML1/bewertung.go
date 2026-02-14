package main

// häufigkeit gibt die Häufigkeit von farbe in farbfolge zurück.
func häufigkeit(farbfolge [anzStellen]string, farbe string) int {
	result := 0
	for i := 0; i < len(farbfolge); i++ {
		if farbfolge[i] == farbe {
			result = result + 1
		}
	}
	return result
}

// bewertung gibt die Bewertung von versuch an Hand von code zurück.
// Die zurückgegebene Bewertung ist ein Paar (schwarze, weiße).
func bewertung(versuch, code [anzStellen]string) (int, int) {
	schwarze := 0
	// Schwarze müssen in Position und Farbe übereinstimmen
	for i := 0; i < len(code); i++ {
		if versuch[i] == code[i] {
			schwarze = schwarze + 1
		}
	}
	weiße := 0
	// Bei Weißen ist die Position egal, sodass nur alle farblich
	// Übereinstimmende gezählt werden müssen ...
	for f := 0; f < len(farben); f++ {
		farbe := farben[f : f+1]
		weiße = weiße + min(häufigkeit(versuch, farbe), häufigkeit(code, farbe))
	}
	// ... und die Schwarzen davon abgezogen werden
	weiße = weiße - schwarze
	return schwarze, weiße
}
