package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := investmentAmount * math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: ", futureValue)
	fmt.Printf("Future real value: %.2f", futureRealValue)
}
