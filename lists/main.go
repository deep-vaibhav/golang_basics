package main

import "fmt"

type Product struct {
	title string
	id    string
	price string
}

func lists() {
	var productNames [4]string = [4]string{"Book"}
	prices := [4]float64{10.99, 20.00, 9.99, 40.45}

	productNames[2] = "Carpet"

	fmt.Println(productNames)
	fmt.Println(prices)

	featuredPrices := prices[:3]
	fmt.Println(featuredPrices)
}
