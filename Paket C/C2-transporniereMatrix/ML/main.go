package main

import "fmt"

// matrixAusgeben gibt ihr Argument in Tabellenform aus.
func matrixAusgeben(matrix [5][5]int) {
	for _, zeile := range matrix {
		for _, wert := range zeile {
			fmt.Printf("%3d", wert)
		}
		fmt.Println()
	}
}

// Das Programm gibt Matrizen aus, transponiert sie und gibt sie transponiert wieder aus.
func main() {
	const anzMatrizen = 4
	for mnr := 1; mnr <= anzMatrizen; mnr++ {
		matrix := gibMatrix(mnr)
		matrixAusgeben(matrix)
		fmt.Println()
		fmt.Println("wird transponiert zu")
		fmt.Println()
		// Vertauschung mit einer doppelten for-Schleife:
		// Nur die obere Dreiecksmatrix wird durchlaufen und ihre Elemente werden mit den Elementen
		// der untern Dreiecksmatrix getauscht.
		// (denn ein Durchlaufen und Vertauschen der gesamten Matrix hätte zur Folge,
		// dass danach wieder aller Elemente an ihrer ursprünglichen Position stehen würden.)
		// Die Elemente der Diagonalen müssen nicht getaucht werden,
		// sodass schließlich in der letzten Zeile der Matrix  gar nichts getauscht werden muss.
		for zeile := 0; zeile < len(matrix)-1; zeile++ {
			for spalte := zeile + 1; spalte < len(matrix[zeile]); spalte++ {
				matrix[zeile][spalte], matrix[spalte][zeile] = matrix[spalte][zeile], matrix[zeile][spalte]
			}
		}
		matrixAusgeben(matrix)
		fmt.Println()
		fmt.Println()
	}
}
