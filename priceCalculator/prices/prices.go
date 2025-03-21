package prices

import (
	"fmt"

	"example.com/priceCalc/fileManager"
)

type TaxIncludedPriceJob struct {
	Tax                      float64
	Prices                   []float64
	TaxIncludedPrices        map[string]string
	TaxIncludedPricesWithTax map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	pricesFromFile, err := fileManager.ReadLinesFromFile("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices = pricesFromFile
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	result := make(map[string]string)
	fmt.Println("Tax rate: ", job.Tax*100, "%")

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.Tax)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	fmt.Println(result)
	fileManager.WriteJsonToFile(fmt.Sprintf("result_%.0f.json", job.Tax*100), job)
}

func NewTaxIncludedPricejob(tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices: []float64{10, 20, 30},
		Tax:    tax,
	}
}
