package coltorti

import "github.com/essemfly/alloff-products/utils"

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

func CalculatePrice(originalPrice float64, diescountRate int, currencyType string, productType string, FTA bool) int {
	// Assume currency type is euro
	krwOriginalPrice := originalPrice * utils.EURO_EXCHANGE_RATE
	krwPrice := int(krwOriginalPrice * float64(100-diescountRate) / 100)
	krwPrice = krwPrice * (100 - diescountRate) / 100
	return krwPrice
}
