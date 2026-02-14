package main

// bewertung gibt die Bewertung von versuch an Hand von code zurück.
// Die zurückgegebene Bewertung ist ein Paar (schwarze, weiße).
func bewertung(versuch, code [anzStellen]string) (int, int) {
	const vBewertet = "_"
	const cBewertet = "*"
	schwarze := 0
	// Schwarze müssen in Position und Farbe übereinstimmen.
	for pos := 0; pos < anzStellen; pos++ {
		if versuch[pos] == code[pos] {
			schwarze = schwarze + 1
			// Schon bewertete Stifte „herausstreichen“, damit sie nicht
			// nochmal bewertet werden.
			versuch[pos] = vBewertet
			code[pos] = cBewertet
		}
	}
	weiße := 0
	// bei weissen ist die Position egal, nur die Farbe muss übereinstimmen.
	for posC := 0; posC < anzStellen; posC++ {
		for posV := 0; posV < anzStellen; posV++ {
			if code[posC] == versuch[posV] {
				weiße = weiße + 1
				versuch[posV] = vBewertet
				code[posC] = cBewertet
			}
		}
	}
	return schwarze, weiße
}
