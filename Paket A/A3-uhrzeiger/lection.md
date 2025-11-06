# Vorlesungsnotiz: Uhrzeiger, Formatierung und Modulo in Go

Diese Notiz fasst die Unterschiede zwischen zwei Implementierungen im Projekt zusammen, erklärt die Verwendung des Prozentzeichens (`%`) in `fmt.Printf` und den Modulo-Operator (`%`) in Go. Zielgruppe sind Studierende, die die Programme in `A3-uhrzeiger` verstehen und sicher anpassen wollen.

## 1. Ziel
- Verstehen, wie Stunde- und Minutenzeigerwinkel berechnet werden.
- Lernen, wie Format-Verb-Spezifikationen in `fmt.Printf` funktionieren.
- Verstehen, wie der Modulo-Operator in Go funktioniert (für Integer) und wie man Restwert bei Fließkommazahlen mit `math.Mod` erhält.

## 2. Vergleich: `ML/main.go` vs. `main.go`
Kurzfassung der Hauptunterschiede:

- Eingabetypen
  - `ML/main.go`: benutzt `int` für Stunde und Minute; klare Eingabegrenzen (0..23, 0..59).
  - `main.go`: benutzt `float64` für Stunde und Minute; akzeptiert auch nicht-ganzzahlige Werte (kein Range-Check).

- Validierung
  - `ML/main.go`: prüft die Gültigkeit der Uhrzeit und bricht bei ungültigen Eingaben ab.
  - `main.go`: führt keine Validierung durch — kann falsche oder ungültige Eingaben akzeptieren.

- Winkelberechnung
  - Minutenwinkel (grob): Minuten × 6°
    - `ML/main.go` rechnet ganzzahlig: `minute * 360 / 60` (ergibt `minute * 6` als `int`).
    - `main.go` rechnet mit Float: `m * 6` (ergibt genauere Ergebnisse bei nicht-ganzzahligen Minuten).
  - Stundenwinkel: 30° pro Stunde plus 0.5° pro Minute
    - `ML/main.go`: rechnet in float64 und berücksichtigt Minutenanteil korrekt.
    - `main.go`: verwendet `h*30 + m*0.5` — funktioniert, solange `h` im Bereich 0..11 oder man vorher `h % 12` anwendet.

- Ausgabeformatierung
  - `ML/main.go`: verwendet `%02d`, `%3.1f`, `%3d` — passt zu den übergebenen Typen.
  - `main.go`: verwendet `%02.0f` für Stunden/Minuten (OK für float->integer-ähnliche Ausgabe) aber hat einen Bug: `minutenwinkel` ist `float64` und wird mit `%3d` (Integer-Verb) ausgegeben — das verursacht einen Laufzeitfehler.

## 3. Verwendung von `%` in `fmt.Printf`

In Go sind Format-Strings (z. B. für `fmt.Printf`, `fmt.Sprintf`, `fmt.Fprintf`) sehr mächtig. Hier erkläre ich die wichtigsten Verbs, Flags, Breiten/Präzisionen und viele Beispiele, die im Alltag nützlich sind.

Kurzüberblick über gebräuchliche Verbs

- `%v`  : Standardformat (für beliebige Werte)
- `%#v` : Go-syntax-Repräsentation (nützlich beim Debuggen)
- `%+v` : ähnlich wie `%v`, aber mit Feldnamen für Strukturen
- `%T`  : Typ des Werts
- `%d`  : Dezimalzahl (Integer)
- `%b`  : Binär (Integer)
- `%o`  : Oktal (Integer)
- `%x`, `%X` : Hexadezimal (Integer/Bytes)
- `%f`  : Dezimaldarstellung für Floats (Standard 6 Nachkommastellen)
- `%e`, `%E` : Exponentielle Darstellung (Floats)
- `%g`  : Kürzestes von `%e` oder `%f`
- `%c`  : Ein einzelnes Unicode-Zeichen (rune)
- `%s`  : String
- `%q`  : Quoted string oder quoted rune literal
- `%p`  : Adresse/Pointer
- `%t`  : Boolean

Width / Precision / Flags

Syntax: `%[flags][width][.precision]verb`

- width: Minimale Feldbreite (ganzzahlig).
- .precision: Anzahl Nachkommastellen bei Floats oder maximale Länge bei Strings.
- Flags:
  - `0` : mit Nullen auffüllen (bei numerischen Werten, z. B. `%02d`)
  - `-` : linksbündig
  - `+` : immer Vorzeichen anzeigen (`+` oder `-`)
  - ` ` (space) : Vorzeichen nur für positive mit führendem Leerzeichen
  - `#` : alternative Form (z. B. `#x` führt 0x-Präfix ein)

Dynamische Breite/Präzision

Sie können Breite und Präzision auch zur Laufzeit angeben, indem Sie `*` verwenden; die Werte kommen dann als zusätzliche Argumente vor dem zu formatierenden Wert:

```go
fmt.Printf("%*.*f\n", width, precision, value)
// Beispiel: width=6, precision=2 -> "%6.2f"
```

Praktische Beispiele

- Null-gepadete Zeit: `fmt.Printf("%02d:%02d\n", 9, 5)` -> `09:05`
- Float mit einer Nachkommastelle: `fmt.Printf("%3.1f°\n", 7.25)` -> ` 7.2°` (Breite 3)
- Float ohne Nachkommastellen, als Null-padded Ganzzahl: `fmt.Printf("%02.0f\n", 3.7)` -> `04` (gerundet)
- Hex-Ausgabe einer Zahl: `fmt.Printf("0x%X\n", 255)` -> `0xFF`
- Byte-Slice hex dump: `fmt.Printf("%x\n", []byte{0x0a,0x0b})` -> `0a0b`
- Struktur debug: `fmt.Printf("%+v\n", myStruct)` -> zeigt Feldnamen
- Pointer: `fmt.Printf("%p\n", &x)` -> `0xc0000120b0`
- Quoted string: `fmt.Printf("%q\n", "Hi\n")` -> `"Hi\n"`

Fehlerquellen & Tipps

- Typ-Mismatch: Ein sehr häufiger Laufzeitfehler ist, ein falsches Verb für den Typ zu verwenden (z. B. `%d` für `float64`). Das führt in Go zu einem Laufzeitfehler. Verwende `%v` beim Debuggen, wenn du dir nicht sicher bist.
- Rundung: `%f` rundet das Ergebnis auf die angegebene Anzahl von Nachkommastellen. Wenn du Rundungsverhalten kontrollieren musst, runde vorher mit `math.Round`.
- Kontrolle: `fmt.Sprintf` erzeugt einen String (statt ihn zu drucken). `fmt.Fprintf` schreibt in ein `io.Writer`.

Beispiele im Kontext (kleine Programme)

1) Formatierungen mischen:

```go
package main

import "fmt"

func main() {
    name := "Anna"
    score := 95.3456
    fmt.Printf("%s hat %6.2f Punkte\n", name, score) // Anna hat 95.35 Punkte
}
```

2) Dynamische Breite / Präzision:

```go
package main

import "fmt"

func main() {
    w, p := 8, 3
    fmt.Printf("%*.*f\n", w, p, 12.34567) // Breite 8, Präzision 3 -> '  12.346'
}
```

3) Sicher debuggen (kein Typfehler):

```go
package main

import "fmt"

func main() {
    var f float64 = 7.5
    // statt fmt.Printf("%d", f) => Laufzeitfehler
    fmt.Printf("mit %%v: %v, mit %%T: %T\n", f, f)
}
```

Literal-Prozentzeichen

Um ein `%` im ausgegebenen Text zu haben, benutze `%%`:

```go
fmt.Printf("Prozent: %d%%\n", 50) // Ausgabe: "Prozent: 50%"
```

Weiteres: `fmt`-Familie

- `fmt.Print`, `fmt.Println` : einfache Ausgabe, keine Format-Strings nötig.
- `fmt.Printf` : formatierte Ausgabe (wie oben).
- `fmt.Sprintf` : gibt formatierte Ausgabe als String zurück (nützlich zum Aufbau von Texten).
- `fmt.Fprintf` : schreibt formatierte Ausgabe in ein `io.Writer` (z. B. Datei, Netzwerk-Conn).

Mit diesen Beispielen solltest du `fmt.Printf` sicherer und effektiver nutzen können. Nutze `%v` beim Debuggen, benutze passende Verbs für deinen Typ und verwende dynamische Breiten, wenn du formatierte Spalten brauchst.

## 4. Modulus-Operator (`%`) in Go (arithmetisch)
- Für ganze Zahlen: `a % b` liefert den Rest der ganzzahligen Division `a / b`.
  - Beispiel: `7 % 3 == 1`, `10 % 2 == 0`.
- Vorzeichenregel in Go: das Ergebnis hat das Vorzeichen des Dividenden (linker Operand):
  - `-7 % 3 == -1`
  - `7 % -3 == 1`
- Division durch 0 (`x % 0`) führt zur Laufzeit zu einem Panic — vermeiden!

- Für Fließkommazahlen: `math.Mod` verwenden
  - `import "math"`
  - Beispiel: `math.Mod(7.5, 3.0) == 1.5`

## 5. Typische Fehler und Korrekturen (Empfehlungen)
- Fehler: Benutzen von `%d` für einen `float64`-Wert → Laufzeit-Panic.
  - Korrektur: entweder das Argument in `int(...)` umwandeln oder das Format zu `%f`/`%3.1f` ändern.

- Verbesserungsvorschlag für `main.go` (schnell & sicher):
  - Validierung der Eingaben (Stunde 0..23, Minute 0..59).
  - Reduktion auf 12-Stunden-Format vor der Winkelberechnung: `h12 := int(h) % 12` oder bei floats `math.Mod(h, 12)`.
  - Einheitliche Typwahl: entweder alles `int` (für ganzzahlige Minuten) oder alles `float64` wenn Bruchteile erlaubt sind.

## 6. Kurze Beispiele (Code-Snippets)
- Ganzzahliger Modulo:

```go
// int modulo
a := 10
b := 3
fmt.Println(a % b) // Ausgabe: 1
```

- Float-Rest mit math.Mod:

```go
import "math"

fmt.Println(math.Mod(10.5, 3.0)) // Ausgabe: 1.5
```

- Sicheres printf:

```go
minutenwinkel := 6.0 * m // m ist float64
fmt.Printf("Minutenwinkel: %3.1f°\n", minutenwinkel)
```

## 7. Zusammenfassung
- Verwende in Format-Strings passende Verbs für den tatsächlichen Typ.
- Nutze die Validierung von Eingaben, um unerwartete Laufzeitfehler zu vermeiden.
- Für Restwertrechnung bei Ganzzahlen benutze `%`, für Floats `math.Mod`.

---

Datei: `lection.md` wurde im Ordner `A3-uhrzeiger` erstellt. Wenn du möchtest, kann ich noch eine Version auf Deutsch/Englisch umstellen, Beispielprogramme ergänzen oder die genannten Korrekturen direkt in den Quellcode anwenden und testen.