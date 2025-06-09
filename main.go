package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {

	var weapon = flag.StringP("weapon", "w", "ak47", "ak47, m4a1, m4a1s")
	var target = flag.StringP("target", "t", "chest", "head, chest")
	var goal = flag.IntP("goal", "g", 95, "Avg% goal")
	var reset = flag.IntP("reset", "r", 30, "Reset Target value")
	var distance = flag.IntP("distance", "d", 0, "Target distance")
	var sets = flag.IntP("sets", "s", 30, "Number of sets (sprays) completed")

	flag.Parse()

	flag.PrintDefaults()

	if !validateWeapon(*weapon) {
		os.Exit(1)
	}

	var results []float64

	for i := 0; i < *sets; i++ {
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

	// date weapon target goal #-sprays avg% distance
	fmt.Printf("%s\n%s, %s, %s, %d, %d, %d, %.2f, %d\n%s\n",
		strings.Repeat("-", 60),
		time.Now().Format(time.DateOnly),
		*weapon,
		*target,
		*goal,
		*sets,
		*reset,
		sum/float64(len(results)),
		*distance,
		strings.Repeat("-", 60),
	)
}

func validateWeapon(weapon string) bool {

	validWeapon := map[string]bool{
		"m4a1":  true,
		"m4a1s": true,
		"ak47":  true,
	}

	if !validWeapon[weapon] {
		return false
	}

	return true
}

func inputToFloat(pctFired string) (float64, error) {
	return strconv.ParseFloat(pctFired, 64)
}

func prompt(count int) string {
	var pctFired string
	fmt.Printf("Spray #%02d | %% of fired: ", count+1)
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
