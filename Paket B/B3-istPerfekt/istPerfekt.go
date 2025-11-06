package main

// istPerfekt bestimmt, ob zahl eine perfekte Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istPerfekt(zahl int) bool {
	sum := 0
	for i := 1; i < zahl; i++ {
		if zahl%i == 0 {
			sum += i
		}
	}
	if sum == zahl {
		return true
	}

	return false
}
