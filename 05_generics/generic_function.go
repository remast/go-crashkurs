package main

import "fmt"

func humanAge[T int32 | float64](dogAge T) T {
	return dogAge * 7
}

func main() {
	// Generische Funktion mit Typ int aufrufen
	var i int32 = humanAge[int32](3)
	fmt.Println("int:", i)

	// Generische Funktion mit Typ float aufrufen
	var f float64 = humanAge[float64](4.9)
	fmt.Println("float", f)
}
