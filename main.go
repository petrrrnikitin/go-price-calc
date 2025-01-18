package main

import (
	"PriceCalculator/filemanager"
	"PriceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range taxRates {
		fileM := filemanager.New("prices1.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		//consoleM := consolemanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(rate, *fileM)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process price job")
			fmt.Println(err)
		}
	}
}
