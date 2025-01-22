package main

import "expample.com/price-calculator/prices"

func main() {
	taxes := []float64{0.05, 0.1, 0.25}

	for _, tax := range taxes {
		priceJob := prices.New(tax)
		priceJob.Process()
	}

}
