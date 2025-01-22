package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate      float64
	Prices       []float64
	PriceWithTax map[string]float64
}

func New(tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:  []float64{10, 20, 30},
		TaxRate: tax,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	//job.PriceWithTax = result

	fmt.Println(result)

}

func (job *TaxIncludedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not opne prices.txt")
		fmt.Println(err)
		return
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for index, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price failed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[index] = floatPrice
	}
	job.Prices = prices

}
