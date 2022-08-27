package intrend

import (
	"github.com/essemfly/alloff-products/utils"
)

const (
	AlloffDiscount            = 0.0
	ForeignDevlieryFeeInEuro  = 20
	DomesticDeliveryFeeInEuro = 3
	VAT                       = 0.1
	CommissionUnder80Euro     = 0.15
	CommissionOver80Euro      = 0.12
	MarginUnder80Euro         = 0.17
	MarginOver80Euro          = 0.17
	ClothingTaxRate           = 0.13
	NonClothingTaxRate        = 0.08
	VATCriterion              = 150
)

/* 2022.08.27

거래조건
개인 팩킹 수수료  건당 3 유로
배송비 이태리- 한국 개별 주소지
: 실비 정산  ( 대략 17 - 22 유로 건당 예상, 평균 20유로)

80유로 이하 15% -> 구매수수료
80유로 초과 12% -> 구매수수료
건별 +3유로


-의류, 신발-
상품이 80유로 이하일 경우 (구매수수료 15%)
1.상품가 + 해외 배송비 20유로(약 25,000)=a
2.a+(a가 $150 넘을경우 13%(관세)) = b
2.b+(a가 $150이 넘을경우 ‘a+13%’에대한 부가세10%)=c
3.c+구매수수료(배송비 미포함 상품가의 15%)+3유로(개별배송 서비스료) = d
4.d +(d의 17%마진)= 최종가격

-의류, 신발-
상품이 80유로 초과일 경우 (구매수수료 12%)
1.상품가 + 해외 배송비 20유로(약 25,000)=a
2.a+(a가 $150 넘을경우 13%(관세)) = b
2.b+(상품가가 $150이 넘을경우 ‘a+13%’에대한 부가세10%)=c
3.c+구매수수료(배송비 미포함 상품가의 12%)+3유로(개별배송 서비스료) = d
4.d +(d의 15%마진)= 최종가격

-잡화-
상품이 80유로 이하일 경우 (구매수수료 15%)
1.상품가 + 해외 배송비 20유로(약 25,000)+상품가(상품가+해외 배송비)가 $150 넘을경우 8%(관세) = a
2.(a+상품가가 $150이 넘을경우 ‘a’에대한 부가세10%)=b
3.b+구매수수료(배송비 미포함 상품가의 15%)+3유로(개별배송 서비스료) = c
4.c +(c의17%마진)= 최종가격

-잡화-
상품이 80유로 이상일 경우 (구매수수료 12%)
1.상품가 + 해외 배송비 20유로(약 25,000)+상품가(상품가+해외 배송비)가 $150 넘을경우 8%(관세) = a
2.(a+상품가가 $150이 넘을경우 ‘a’에대한 부가세10%)=b
3.b+구매수수료(배송비 미포함 상품가의 12%)+3유로(개별배송 서비스료) = c
4.c +(c의15%마진)= 최종가격

*/

func CalculatePrice(originalPrice float64, discountRate int, currencyType string, isClothing, FTA bool) int {
	originalPriceInEuro := originalPrice * float64(100-discountRate) / 100.0
	ourPrice := originalPriceInEuro + ForeignDevlieryFeeInEuro
	ourPrice = ourPrice * utils.EURO_EXCHANGE_RATE
	taxPrice := 0.0

	// 관세
	// 총 상품액이 $150인경우 관세를 붙이는 여부 -> 관세는 13%, 8%
	if ourPrice > VATCriterion*utils.DOLLOR_EXCHANGE_RATE {
		if isClothing {
			taxPrice = ourPrice * ClothingTaxRate
		} else {
			taxPrice = ourPrice * NonClothingTaxRate
		}
	}

	// 부가세
	vatPrice := 0.0
	if ourPrice > VATCriterion*utils.DOLLOR_EXCHANGE_RATE {
		vatPrice = (ourPrice + taxPrice) * VAT
	}

	// 구매수수료
	marginRate := MarginUnder80Euro
	commissionFee := originalPriceInEuro * CommissionUnder80Euro * utils.EURO_EXCHANGE_RATE
	if originalPriceInEuro >= 80 {
		marginRate = MarginOver80Euro
		commissionFee = originalPriceInEuro * CommissionOver80Euro * utils.EURO_EXCHANGE_RATE
	}

	// 마진
	totalPrice := ourPrice + taxPrice + vatPrice + commissionFee + DomesticDeliveryFeeInEuro*utils.EURO_EXCHANGE_RATE
	totalPrice = totalPrice * (1 + marginRate)

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
