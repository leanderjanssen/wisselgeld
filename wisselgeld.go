package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func main() {

	var invoer []float64
	var input float64

	munten := map[float64]string{
		200: "2 euro",
		100: "1 euro",
		50:  "50 cent",
		20:  "20 cent",
		10:  "10 cent",
		5:   "5 cent",
	}

	aantalMunt := [6]float64{0, 0, 0, 0, 0, 0}

	//Vraag om wisselgeld indien niet opgegeven als argument
	if len(os.Args[1:]) < 1 {
		fmt.Printf("\nVoer wisselgeld bedrag in (gebruik een punt als decimaal teken): ")
		_, err := fmt.Scanf("%f", &input)
		invoer = append(invoer, input)

		if err != nil {
			fmt.Println("\nOngeldige invoer!")
			os.Exit(1)
		}
		//Verwerk anders de argumenten
	} else {
		for _, arg := range os.Args[1:] {
			if n, err := strconv.ParseFloat(arg, 64); err == nil {
				invoer = append(invoer, n)
			}
		}
	}

	for _, value := range invoer {
		fmt.Printf("\n%-30s %-1s %0.2f\n", "Ingevoerd wisselgeld bedrag: ", "€", value)
		fmt.Printf("%-30s %-1s %0.2f\n\n", "Afgerond wisselgeld bedrag: ", "€", Round(value, 0.05))

		// Rond het invoer bedrag af en converteer naar centen
		wisselgeld := Round(value, 0.05) * 100

		// Creeer array met de keys van munten
		var keys []float64
		for k := range munten {
			keys = append(keys, k)
		}
		// Sorteer de array, grootste munt eerst
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

		fmt.Println("Het minimum aantal munten is:\n")

		// Bereken per munt het max aantal
		for key, value := range keys {
			munt := keys[key]
			aantal := math.Floor(wisselgeld / munt)

			if aantal >= 1 {
				aantalMunt[key] = aantal
				wisselgeld -= (aantal * keys[key])

				fmt.Printf("%4.f%s %7s\n", aantal, "x", munten[value])
			}
		}
	}
}
