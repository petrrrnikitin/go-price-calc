package main

import (
	"PriceCalculator/filemanager"
	"PriceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		priceJob := prices.NewTaxIncludedPriceJob(rate, *fm)
		priceJob.Process()
	}
}
