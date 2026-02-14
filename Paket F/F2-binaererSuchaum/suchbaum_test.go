package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomTrees(t *testing.T) {
	const numberOfRuns = 10
	const maxValue = 100
	const numberOfNodes = 70

	if numberOfNodes > maxValue {
		t.Error("Falscher Test-Aufbau: Mehr Knoten als Zufallszahlen")
	}

	for run := 1; run <= numberOfRuns; run++ {
		rndNums := rand.Perm(maxValue)
		tree := &knoten{wert: rndNums[0]}
		sum := rndNums[0]
		minVal := rndNums[0]
		maxVal := rndNums[0]

		for i := 1; i < numberOfNodes; i++ {
			einfügen(tree, rndNums[i])
			sum += rndNums[i]
			if rndNums[i] < minVal {
				minVal = rndNums[i]
			}
			if rndNums[i] > maxVal {
				maxVal = rndNums[i]
			}
		}
		baumAusgeben(tree)
		fmt.Println()

		var checkStart *int = new(int)
		*checkStart = -1
		if !checkOrder(t, tree, checkStart) {
			t.Error("Fehler in Run", run)
			return
		}

		if summeDesBaums(tree) != sum {
			t.Errorf("Fehler in Run %d: Summe ist %d, sollte aber %d sein.", run, summeDesBaums(tree), sum)
			return
		}

		if got, want := kleinsterKnoten(tree).wert, minVal; got != want {
			t.Errorf("Fehler in Run %d: kleinsterKnoten ist %d, sollte aber %d sein.", run, got, want)
			return
		}

		maxHöhe := 0
		for i := range numberOfNodes {
			v := rndNums[i]
			result := suchen(tree, v)
			if result == nil {
				t.Errorf("Fehler in Run %d: Knoten mit Wert %d nicht gefunden.", run, v)
				return
			}
			if result.wert != v {
				t.Errorf("Fehler in Run %d: Knoten mit Wert %d gefunden, aber Wert %d war gesucht.", run, result.wert, v)
				return
			}

			maxHöhe = max(maxHöhe, höheDesWerts(tree, v, 1))
		}
		if got, want := maxHöhe, höheDesBaums(tree); got != want {
			t.Errorf("Fehler in Run %d: Maximum aller höheDesWert ist %d, muss aber mit höheDesBaums %d übereinstimmen.", run, got, want)
			return
		}

		// Achtung: ab hier wird der Baum verändert
		inkrementiereAlleWerte(tree)
		if got, want := summeDesBaums(tree), sum+numberOfNodes; got != want {
			t.Errorf("Fehler in Run %d: Nach inkrementiereAlleWerte ist die Summe aller Knoten %d, muss aber %d sein.", run, got, want)
			return
		}

	}
}

func TestFixedTrees(t *testing.T) {
	var tree *knoten
	var checkStart *int = new(int)

	tree = &knoten{wert: 50}
	einfügen(tree, 30)
	einfügen(tree, 20)
	einfügen(tree, 40)
	einfügen(tree, 70)
	einfügen(tree, 60)
	einfügen(tree, 80)
	*checkStart = -1
	checkOrder(t, tree, checkStart)
	fmt.Println()

	tree = &knoten{wert: 57}
	einfügen(tree, 36)
	einfügen(tree, 17)
	einfügen(tree, 42)
	einfügen(tree, 88)
	einfügen(tree, 93)
	einfügen(tree, 71)
	*checkStart = -1
	checkOrder(t, tree, checkStart)
	fmt.Println()

	tree = &knoten{wert: 36}
	einfügen(tree, 88)
	einfügen(tree, 17)
	einfügen(tree, 93)
	einfügen(tree, 57)
	einfügen(tree, 42)
	einfügen(tree, 71)
	*checkStart = -1
	checkOrder(t, tree, checkStart)
	fmt.Println()

	tree = &knoten{wert: 17}
	einfügen(tree, 36)
	einfügen(tree, 42)
	einfügen(tree, 57)
	einfügen(tree, 71)
	einfügen(tree, 88)
	einfügen(tree, 93)
	*checkStart = -1
	checkOrder(t, tree, checkStart)
	fmt.Println()
}

// Hilfsfunktion für Test
func checkOrder(t *testing.T, n *knoten, prevValue *int) bool {
	if n != nil {
		if !checkOrder(t, n.links, prevValue) {
			return false
		}

		//fmt.Printf("%d (prev: %d) ", n.wert, *prevValue)
		if *prevValue >= n.wert {
			t.Errorf("Wert %d ist vor %d im Baum einsortiert.", *prevValue, n.wert)
			return false
		}
		*prevValue = n.wert // zuletzt ermittelten Wert merken (für nächsten Vergleich)

		if !checkOrder(t, n.rechts, prevValue) {
			return false
		}
	}

	return true
}
