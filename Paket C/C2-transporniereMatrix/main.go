package main

import (
	"fmt"
)

// transponiere transponiert eine 5x5-Matrix in place
func transponiere(matrix *[5][5]int) {
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// gibAus gibt eine 5x5-Matrix aus
func gibAus(matrix [5][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	for nr := 1; nr <= 4; nr++ {
		matrix := gibMatrix(nr)
		gibAus(matrix)
		fmt.Println()
		fmt.Println("wird transponiert zu")
		fmt.Println()
		transponiere(&matrix)
		gibAus(matrix)
		fmt.Println()
		if nr < 4 {
			fmt.Println()
		}
	}
	fmt.Println()
}
