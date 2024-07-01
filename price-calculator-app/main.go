package main

import (
	"example.com/price-calculator/helper"
	"example.com/price-calculator/price"
)

const pricesFile = "price-calculator-app/sources/prices.txt"
const taxRatesFile = "price-calculator-app/sources/tax_rates.txt"
const updatedPricesFile = "price-calculator-app/out"

func main() {
	prices, err := helper.ReadFile(pricesFile)
	if err != nil {
		return
	}
	taxRates, err := helper.ReadFile(taxRatesFile)
	if err != nil {
		return
	}

	taxRateList, err := helper.ParseFloats(taxRates)
	if err != nil {
		return
	}

	priceList, err := helper.ParseFloats(prices)
	if err != nil {
		return
	}

	result := price.CalculateTaxAddedPrices(priceList, taxRateList)

	helper.WriteToJSON(updatedPricesFile, result)
}
