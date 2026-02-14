package main

// gibtMatrix gibt eine von vier quadratischen 5x5-Integer-Matrizen zurück.
// nr (1 bis 4) gibt an, welche Beispiel-Matrizen zurückgegeben wird.
func gibMatrix(nr int) [5][5]int {
	var matrix [5][5]int
	switch nr {
	case 1:
		matrix = [5][5]int{
			{15, 8, 1, 24, 17},
			{16, 14, 7, 5, 23},
			{22, 20, 13, 6, 4},
			{3, 21, 19, 12, 10},
			{9, 2, 25, 18, 11},
		}
	case 2:
		matrix = [5][5]int{
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}
	case 3:
		matrix = [5][5]int{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{5, 6, 7, 8, 9},
			{5, 6, 7, 8, 9},
			{10, 11, 12, 13, 14},
		}
	case 4:
		matrix = [5][5]int{
			{1, 6, 11, 16, 21},
			{2, 7, 12, 17, 22},
			{3, 8, 13, 18, 23},
			{4, 9, 14, 19, 24},
			{5, 10, 15, 20, 25},
		}
	}
	return matrix
}
