package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	Tax                      float64
	Prices                   []float64
	TaxIncludedPrices        map[string]float64
	TaxIncludedPricesWithTax map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	// Open a file
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file.")
		return
	}

	// Helps in reading contents from a file
	scanner := bufio.NewScanner(file)
	var lines []float64
	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error in parsing price")
			return
		}
		lines = append(lines, floatPrice)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed.")
		file.Close()
		return
	}

	job.Prices = lines
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	result := make(map[string]string)
	fmt.Println("Tax rate: ", job.Tax*100, "%")

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.Tax)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPricejob(tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices: []float64{10, 20, 30},
		Tax:    tax,
	}
}
