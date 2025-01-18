package main

import (
	"PriceCalculator/consolemanager"
	"PriceCalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range taxRates {
		//fileM := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		consoleM := consolemanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(rate, *consoleM)
		priceJob.Process()
	}
}
