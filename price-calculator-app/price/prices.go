package price

import "example.com/price-calculator/helper"

const pricesFile = "price-calculator-app/sources/prices.txt"

func CalculateTaxAddedPrices(priceList, taxRateList []float64) map[float64][]float64 {

	prices, err := helper.ReadFile(pricesFile)
	if err != nil {
		return nil
	}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRateList {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range priceList {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	return result
}
