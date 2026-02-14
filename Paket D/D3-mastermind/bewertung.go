package main

// bewertung gibt die Bewertung von versuch an Hand von code zurück.
// Die zurückgegebene Bewertung ist ein Paar (schwarze, weiße).
func bewertung(versuch, code [anzStellen]string) (int, int) {
	schwarze := 0
	weiße := 0
	
	// Track which positions and colors have been matched
	codeUsed := [anzStellen]bool{}
	versuchUsed := [anzStellen]bool{}
	
	// First pass: Count black pegs (correct color in correct position)
	for i := 0; i < anzStellen; i++ {
		if versuch[i] == code[i] {
			schwarze++
			codeUsed[i] = true
			versuchUsed[i] = true
		}
	}
	
	// Second pass: Count white pegs (correct color in wrong position)
	for i := 0; i < anzStellen; i++ {
		if versuchUsed[i] {
			continue // Already matched as black
		}
		for j := 0; j < anzStellen; j++ {
			if !codeUsed[j] && versuch[i] == code[j] {
				weiße++
				codeUsed[j] = true
				break // Each code position can only be used once
			}
		}
	}
	
	return schwarze, weiße
}
