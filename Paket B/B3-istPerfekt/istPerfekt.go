package main

func istPerfekt(zahl int) bool {

	var summe int
	for i := 1; i < zahl; i++ {
		if zahl%i == 0 {
			summe += i // addiere alle teiler von zahl
		}
	}
	return summe == zahl // wenn die summe der teiler gleich der zahl ist, ist sie perfekt

}
