package main

import (
	"fmt"

	"example.com/practice-project/cmdmanager"
	"example.com/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, rate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
