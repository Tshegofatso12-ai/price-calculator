package prices

import (
	"fmt"

	"expample.com/price-calculator/conversion"
	"expample.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate      float64                 `json:"tax_rate"`
	Prices       []float64               `json:"prices"`
	PriceWithTax map[string]string       `json:"price_with_tax"`
	IOManager    filemanager.FileManager `json:"-"`
}

func New(fm filemanager.FileManager, tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:    []float64{10, 20, 30},
		TaxRate:   tax,
		IOManager: fm,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadPrices()
	if err != nil {
		errChan <- err
		return
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.PriceWithTax = result

	fmt.Println(result)
	job.IOManager.WriteResult(job)
	doneChan <- true

}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.Prices = prices
	return nil
}
