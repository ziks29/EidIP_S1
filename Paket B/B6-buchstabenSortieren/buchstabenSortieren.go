package main

// buchstabenSortieren gibt einen String zurück, in dem alle Buchstaben
// von wort alphabetisch sortiert sind. In wort mehrfach vorkommende
// Buchstaben kommen im resultierendem String genauso oft vor. Großbuchstaben
// in wort kommen im resultierendem String als Kleinbuchstaben vor.
func buchstabenSortieren(wort string) string {
	// First to lowercase
	wortLower := ""
	for i := 0; i < len(wort); i++ {
		char := wort[i]
		// If uppercase letter, convert to lowercase
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}
		wortLower += string(char)
	}

	// Then sort the letters using a simple counting sort
	counts := make([]int, 26) // Assuming only a-z letters
	for i := 0; i < len(wortLower); i++ {
		char := wortLower[i]
		if char >= 'a' && char <= 'z' {
			counts[char-'a']++
		}
	}

	sortedWort := ""
	for i := 0; i < 26; i++ {
		for j := 0; j < counts[i]; j++ {
			sortedWort += string('a' + byte(i))
		}
	}

	return sortedWort
}
