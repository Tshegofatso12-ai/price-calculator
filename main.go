package main

import (
	"fmt"

	"expample.com/price-calculator/filemanager"
	"expample.com/price-calculator/prices"
)

func main() {
	taxes := []float64{0.05, 0.1, 0.25}
	doneChans := make([]chan bool, len(taxes))
	errorChans := make([]chan error, len(taxes))
	for index, tax := range taxes {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.New(fm, tax)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxes {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done")
		}
	}
}
