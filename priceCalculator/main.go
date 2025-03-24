package main

import (
	"fmt"

	"example.com/priceCalc/fileManager"
	"example.com/priceCalc/prices"
)

func main() {
	taxes := []float64{0, 0.7, 0.1, 0.5}

	// Channels for Goroutines
	doneChans := make([]chan bool, len(taxes))
	errorChans := make([]chan error, len(taxes))

	for i, tax := range taxes {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.NewTaxIncludedPricejob(fm, tax)

		// Goroutine
		go priceJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i := range taxes {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
