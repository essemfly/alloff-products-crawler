package coltorti

import "github.com/essemfly/alloff-products/utils"

func CalculatePrice(originalPrice float64, diescountRate int, currencyType string) int {
	// Assume currency type is euro
	krwOriginalPrice := originalPrice * utils.EURO_EXCHANGE_RATE
	krwPrice := int(krwOriginalPrice * float64(100-diescountRate) / 100)
	krwPrice = krwPrice * (100 - diescountRate) / 100
	return krwPrice
}
