package prices

import (
	"fmt"

	"example.com/priceCalc/fileManager"
)

type TaxIncludedPriceJob struct {
	IOManager                fileManager.FileManager `json:"io_manager"`
	Tax                      float64                 `json:"tax"`
	Prices                   []float64               `json:"input_prices"`
	TaxIncludedPrices        map[string]string       `json:"tax_included_prices"`
	TaxIncludedPricesWithTax map[string]float64      `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	pricesFromFile, err := job.IOManager.ReadLinesFromFile()

	if err != nil {
		return err
	}

	job.Prices = pricesFromFile
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadPrices()

	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)
	fmt.Println("Tax rate: ", job.Tax*100, "%")

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.Tax)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	fmt.Println(result)
	job.IOManager.WriteResultToFile(job)

	// Once write is complete, pass true to the channel
	doneChan <- true
}

func NewTaxIncludedPricejob(fm fileManager.FileManager, tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		Prices:    []float64{10, 20, 30},
		Tax:       tax,
	}
}
