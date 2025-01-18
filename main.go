package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, rate := range taxRates {
		taxInclPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxInclPrices[priceIndex] = price * (1 + rate)
		}
		result[rate] = taxInclPrices
	}

	fmt.Println(result)
}
