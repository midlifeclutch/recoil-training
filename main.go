package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var results []float64

	for i := 0; i < 30; i++ {
		pctFired := prompt(i)

		if pct, err := inputToFloat(pctFired); err == nil {
			if pct < 10 {
				fmt.Println("%% of fired seems low. Confirm:")
				pctFired = prompt(i)
				pct, _ = inputToFloat(pctFired)

				results = append(results, pct)
			} else {
				results = append(results, pct)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var sum float64
	for _, i := range results {
		sum += i
	}

	fmt.Println(sum / float64(len(results)))
}

func inputToFloat(pctFired string) (float64, error) {
	return strconv.ParseFloat(pctFired, 64)
}

func prompt(count int) string {
	var pctFired string
	fmt.Printf("Spray #%d | %% of fired: ", count+1)
	_, err := fmt.Scanf("%s", &pctFired)

	if err != nil {
		fmt.Println("Invalid input.")

		fmt.Printf("Spray #%d | %% of fired: ", count+1)
		_, err := fmt.Scanf("%s", &pctFired)

		if err != nil {
			os.Exit(1)
		}
	}

	return pctFired
}
