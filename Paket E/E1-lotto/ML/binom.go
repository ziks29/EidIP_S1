package main

// gibt den Binomialkoeffizienten n über k zurück.
func binom(n, k int) int {
	if k == 0 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return binom(n-1, k-1) * n / k
	}
}
