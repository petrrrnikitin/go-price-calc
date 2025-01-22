package prices

import (
	"PriceCalculator/IOManager"
	"PriceCalculator/conversion"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         IOManager.IOManager `json:"-"`
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
}

func NewTaxIncludedPriceJob(taxRate float64, ioManager IOManager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: ioManager,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxInclPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}

	job.TaxIncludedPrices = result

	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}
	doneChan <- true
}
