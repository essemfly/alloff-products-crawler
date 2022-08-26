package coltorti

import (
	"log"

	"github.com/essemfly/alloff-products/utils"
)

/*
의류 FTA 불가능
(원가-13%)+20000 해외 배송비 +원가가 $150이상일 경우 13% 관세 + 원가가 $150이상일 경우 10% 부가세 + 10% 마진 +3000 = 최종가격

의류 FTA 가능
(원가-13%)+20000 해외 배송비 + 원가가 $150이상일 경우 10% 부가세 + 10% 마진 +3000 = 최종가격

잡화 FTA 불가능
(원가-13%)+20000 해외 배송비 +원가가 $150이상일 경우 8% 관세+  원가가 $150이상일 경우 10% 부가세 + 10% 마진 +3000 = 최종가격

잡화 FTA 가능
(원가-13%)+20000 해외 배송비 + 원가가 $150이상일 경우 10% 부가세 + 10% 마진 +3000 = 최종가격
*/

// -의류/신발- FTX 불가능
// 상품가격 -13% +23000(배송비)= a
// a+(관세 a=<$150=13%) + (부가세 a=<$150= 10%)= b
// b+5%마진= c

// -의류/신발- FTX 가능
// 상품가격 -13% +23000(배송비)= a
// a + (부가세 a=<$150=10%)= b
// b+5%마진= c

// -잡화- FTX 불가능
// 상품가격 -13% +23000(배송비)= a
// a+(관세 a=<$150=8%) + (부가세 a=<$150=10%)=b
// b+5%마진= c

// -잡화- FTX 가능
// 상품가격 -13% +23000(배송비)= a
// a+=(부가세 a=<$150=10%)= b
// b+5%마진= c

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
