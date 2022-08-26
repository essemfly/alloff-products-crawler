package intrend

import (
	"log"

	"github.com/essemfly/alloff-products/utils"
)

const (
	AlloffDiscount      = 0.08
	ForeignDevlieryFee  = 35000
	DomesticDeliveryFee = 3000
	VAT                 = 0.1
	Margin              = 0.05
	ClothingTaxRate     = 0.13
	NonClothingTaxRate  = 0.08
	VATCriterion        = 150
)

func CalculatePrice(originalPrice float64, discountRate int, currencyType string, isClothing, FTA bool) int {
	ourPrice := originalPrice * float64(100-discountRate) / 100.0
	ourPrice = ourPrice * utils.EURO_EXCHANGE_RATE
	ourPrice = ourPrice + ForeignDevlieryFee
	taxPrice := 0.0

	if FTA {
		log.Println("FTA True cases are here", FTA)
	}

	if !FTA && ourPrice < VATCriterion*utils.DOLLOR_EXCHANGE_RATE {
		if isClothing {
			taxPrice = ourPrice * ClothingTaxRate
		} else {
			taxPrice = ourPrice * NonClothingTaxRate
		}
	}
	vatPrice := 0.0
	if ourPrice < VATCriterion*utils.DOLLOR_EXCHANGE_RATE {
		vatPrice = (ourPrice + taxPrice) * VAT
	}

	totalPrice := ourPrice + taxPrice + vatPrice
	totalPrice = totalPrice * (1 + Margin)

	intPrice := int(totalPrice)
	intPrice = intPrice / 1000
	intPrice = intPrice * 1000
	return intPrice
}

func CalculateOriginalPrice(ourPrice int, discountRate int) int {
	newOriginalPrice := float64(ourPrice) * 10000.0 / (float64(100-discountRate) * float64(1-AlloffDiscount) * 100)
	intPrice := int(newOriginalPrice)
	intPrice = intPrice / 1000
	intPrice = intPrice * 1000

	return intPrice
}
