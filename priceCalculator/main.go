package main

import (
	"fmt"

	"example.com/priceCalc/prices"
)

func main() {
	taxes := []float64{0, 0.7, 0.1, 0.5}

	for _, tax := range taxes {
		priceJob := prices.NewTaxIncludedPricejob(tax)
		priceJob.Process()
		fmt.Println()
	}
}
