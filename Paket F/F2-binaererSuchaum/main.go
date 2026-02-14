package main

import "fmt"

func main() {

	//      50
	//     /  \
	//    30   70
	//   / \   / \
	//  20 40 60 80
	//baum := &knoten{wert: 50}
	//einfügen(baum, 30)
	//einfügen(baum, 20)
	//einfügen(baum, 40)
	//einfügen(baum, 70)
	//einfügen(baum, 60)
	//einfügen(baum, 80)

	//      57
	//     /  \
	//    36   88
	//   / \   / \
	//  17 42 71 93
	//baum := &knoten{wert: 57}
	//einfügen(baum, 36)
	//einfügen(baum, 17)
	//einfügen(baum, 42)
	//einfügen(baum, 88)
	//einfügen(baum, 93)
	//einfügen(baum, 71)

	//      17
	//       \
	//       36
	//        \
	//        42
	//         \
	//         57
	//          \
	//          57
	//           \
	//           71
	//            \
	//            88
	//             \
	//             93
	//baum := &knoten{wert: 17}
	//einfügen(baum, 36)
	//einfügen(baum, 42)
	//einfügen(baum, 57)
	//einfügen(baum, 71)
	//einfügen(baum, 88)
	//einfügen(baum, 93)

	//      36
	//     /  \
	//    17   88
	//        /  \
	//       57  93
	//      / \
	//    42  71
	// der Wurzel-Knoten muss explizit erzeugt werden
	baum := &knoten{wert: 36}
	// alle weitern Knoten werden implizit beim Einfügen erzeugt
	einfügen(baum, 88)
	einfügen(baum, 17)
	einfügen(baum, 93)
	einfügen(baum, 57)
	einfügen(baum, 42)
	einfügen(baum, 71)

	fmt.Print("Werte im Baum: ")
	baumAusgeben(baum) // 17 36 42 57 71 88 93
	fmt.Println()
	fmt.Println()

	fmt.Println("Summe des Baums:", summeDesBaums(baum)) // 404
	fmt.Println()

	fmt.Println("Knoten mit Wert 57:", suchen(baum, 57))
	fmt.Println("Knoten mit Wert 13:", suchen(baum, 13))

	fmt.Println("42 im Baum?", istWertInBaum(baum, 42)) // 42? true
	fmt.Println("11 im Baum?", istWertInBaum(baum, 11)) // 11? false
	fmt.Println()

	fmt.Println("kleinster Wert:", kleinsterKnoten(baum).wert)
	fmt.Println()

	fmt.Println("Höhe des Baumes:", höheDesBaums(baum))
	fmt.Println()

	fmt.Println("Höhe des Werts 36: ", höheDesWerts(baum, 36, 1))
	fmt.Println("Höhe des Werts 88: ", höheDesWerts(baum, 88, 1))
	fmt.Println("Höhe des Werts 93: ", höheDesWerts(baum, 93, 1))
	fmt.Println("Höhe des Werts 42: ", höheDesWerts(baum, 42, 1))
	fmt.Println()

	fmt.Print("Werte im Baum vorher:  ")
	baumAusgeben(baum)
	fmt.Println()
	fmt.Println("inkrementiere alle Werte")
	inkrementiereAlleWerte(baum)
	fmt.Print("Werte im Baum nachher: ")
	baumAusgeben(baum)
	fmt.Println()
	fmt.Println()

}
