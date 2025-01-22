package main

import (
	"PriceCalculator/filemanager"
	"PriceCalculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, rate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fileM := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		//consoleM := consolemanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(rate, *fileM)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Printf("Error processing tax rate %s\n", err.Error())
			}
		case <-doneChans[index]:
			fmt.Printf("Processed tax rate %.2f\n", taxRates[index])
		}
	}
}
