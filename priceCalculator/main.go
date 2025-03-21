package main

import (
	"fmt"

	"example.com/priceCalc/fileManager"
	"example.com/priceCalc/prices"
)

func main() {
	taxes := []float64{0, 0.7, 0.1, 0.5}

	for _, tax := range taxes {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.NewTaxIncludedPricejob(fm, tax)
		priceJob.Process()
		fmt.Println()
	}
}
