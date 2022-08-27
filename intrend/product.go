package intrend

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/essemfly/alloff-products/domain"
	"github.com/essemfly/alloff-products/utils"
	"golang.org/x/text/language"
)

func GetIntrendTemplate(pd *domain.Product) []string {
	optionStrings := []string{}
	optionQuantityStrings := []string{}
	for _, optionName := range pd.SizeOptions {
		optionStrings = append(optionStrings, optionName.SizeName)
		optionQuantityStrings = append(optionQuantityStrings, strconv.Itoa(optionName.Quantity))
	}

	translatedTitle, err := utils.TranslateText(language.Korean.String(), pd.Name)
	if err != nil {
		log.Println("err occured on translating text")
	}
	translatedDescription, err := utils.TranslateText(language.Korean.String(), pd.Description)
	if err != nil {
		log.Println("err occured on translating text")
	}
	translatedMaterial, err := utils.TranslateText(language.Korean.String(), pd.Material)
	if err != nil {
		log.Println("err occured on translating text")
	}

	names := []string{
		pd.Brand,
		translatedTitle,
		pd.ProductStyleisNow,
	}
	nameTranslated := strings.Join(names, " ")

	ourPrice := CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType, IsClothing(pd.Category), pd.FTA)
	originalPrice := CalculateOriginalPrice(ourPrice, pd.DiscountRate)
	discountPrice := originalPrice - ourPrice
	originalPriceStr := strconv.Itoa(originalPrice)
	discountPriceStr := strconv.Itoa(discountPrice)

	descImages := pd.Images
	descImageHtml := ""
	for _, descImageUrl := range descImages {
		descImageHtml = descImageHtml + "<img src='" + descImageUrl + "'>"
	}
	descImageHtml += fmt.Sprintf("<p></p><p>색상: %s </p><p>소재: %s </p><p>상품 설명: %s</p><p></p>", pd.Color, translatedMaterial, translatedDescription)

	infoImages := []string{}
	infoImages = append(infoImages, "https://d3vx04mz0cr7rc.cloudfront.net/detail_220820_1.jpeg")
	infoImages = append(infoImages, "https://d3vx04mz0cr7rc.cloudfront.net/detail_220820_2.jpeg")
	infoImages = append(infoImages, "https://d3vx04mz0cr7rc.cloudfront.net/detail_220820_3.jpeg")

	for _, infoImgUrl := range infoImages {
		descImageHtml = descImageHtml + "<img src='" + infoImgUrl + "'>"
	}

	return []string{
		"신상품",
		GetNaverCategoryCode(pd.Category),
		nameTranslated,
		originalPriceStr,
		strconv.Itoa(pd.Quantity),
		"-",
		"010-4118-1406",
		pd.ImageFilenames[0],
		strings.Join(pd.ImageFilenames, ","),
		descImageHtml,
		"",
		"",
		"",
		pd.Brand,
		"",
		"",
		"과세상품",
		"Y",
		"Y",
		"0201038",
		"올오프",
		"N",
		"",
		"택배‚ 소포‚ 등기",
		"무료",
		"0",
		"선결제",
		"",
		"",
		"30000",
		"30000", //교환배송비
		"",
		"",
		"",
		discountPriceStr,
		"원",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"6", // 무이자 할부개월
		"",  // 사은품
		"단독형",
		"사이즈",
		strings.Join(optionStrings, ","),
		"",
		strings.Join(optionQuantityStrings, ","),
		"",
		"",
		"",
		"",
		nameTranslated,
		pd.ProductID,
		"",
		pd.Brand,
		"N",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
}
