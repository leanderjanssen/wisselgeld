package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func main() {

	var invoer float64
	munten := map[float64]string{
		200: "2 euro",
		100: "1 euro",
		50:  "50 cent",
		20:  "20 cent",
		10:  "10 cent",
		5:   "5 cent",
	}

	aantalMunt := [6]float64{0, 0, 0, 0, 0, 0}

	fmt.Printf("\nVoer wisselgeld bedrag in (gebruik een punt als decimaal teken): ")
	_, err := fmt.Scanf("%f", &invoer)

	if err != nil {
		fmt.Println("\nOngeldige invoer!")
		os.Exit(1)
	}

	fmt.Printf("\n%-30s %-1s %0.2f\n", "Ingevoerd wisselgeld bedrag: ", "€", invoer)
	fmt.Printf("%-30s %-1s %0.2f\n\n", "Afgerond wisselgeld bedrag: ", "€", Round(invoer, 0.05))

	// Rond het invoer bedrag af en converteer naar centen
	wisselgeld := Round(invoer, 0.05) * 100

	var keys []float64
	// Creeer array met de keys van munten
	for k := range munten {
		keys = append(keys, k)
	}
	// Sorteer de array, grootste munt eerst
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	// Bereken per munt het max aantal
	for key, value := range keys {
		var munt = keys[key]
		var aantal = math.Floor(wisselgeld / munt)

		if aantal >= 1 {
			aantalMunt[key] = aantal
			wisselgeld -= (aantal * keys[key])

			fmt.Printf("%4.f %s %7s\n", aantal, "x", munten[value])
			//fmt.Printf("%-13s %7s", "Muntsoort: ", munten[value])
			//fmt.Println(" | Aantal: ", aantal)
		}
	}
}
