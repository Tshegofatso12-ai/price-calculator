package main

import (
	"fmt"

	"expample.com/price-calculator/filemanager"
	"expample.com/price-calculator/prices"
)

func main() {
	taxes := []float64{0.05, 0.1, 0.25}

	for _, tax := range taxes {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.New(fm, tax)
		priceJob.Process()
	}

}
