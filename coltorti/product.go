package coltorti

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/essemfly/alloff-products/utils"
)

type ProductOption struct {
	SizeInfo string
	SizeName string
	Quantity int
}

type ColtortiProductInput struct {
	ProductURL        string
	Images            []string
	ImageFilenames    []string
	Brand             string
	ProductID         string
	ProductStyleisNow string
	Season            string
	Year              int
	Color             string
	MadeIn            string
	Material          string
	Name              string
	Description       string
	Category          string
	Quantity          int
	OriginalPrice     float64
	CurrencyType      string
	DiscountRate      int
	SizeOptions       []ProductOption
	FTA               bool
}

func (pd *ColtortiProductInput) ToProductTemplate() []string {
	optionStrings := []string{}
	optionQuantityStrings := []string{}
	for _, optionName := range pd.SizeOptions {
		optionStrings = append(optionStrings, optionName.SizeInfo+"-"+optionName.SizeName)
		optionQuantityStrings = append(optionQuantityStrings, strconv.Itoa(optionName.Quantity))
	}

	upperedName := strings.ToUpper(pd.Name)
	brandCuttedName := strings.TrimLeft(upperedName, pd.Brand)
	productIDCutted := strings.Split(pd.ProductID, "-")[0]
	seasonBrief := ""
	if pd.Season == "Spring - Summer" {
		seasonBrief = "SS"
	} else if pd.Season == "Fall - winter" {
		seasonBrief = "FW"
	}
	seasonBrief = seasonBrief + "/" + strconv.Itoa(pd.Year%100)

	mustNameStr := GetKoreanBrandName(pd.Brand) + seasonBrief + productIDCutted
	nameCutLength := 47 - utf8.RuneCountInString(mustNameStr)

	names := []string{
		GetKoreanBrandName(pd.Brand),
		utils.StringTruncater(brandCuttedName, nameCutLength),
		seasonBrief,
		productIDCutted,
	}
	nameTranslated := strings.Join(names, " ")

	ourPrice := CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType, IsClothing(*pd), pd.FTA)
	originalPrice := CalculateOriginalPrice(ourPrice, pd.DiscountRate)
	discountPrice := originalPrice - ourPrice
	originalPriceStr := strconv.Itoa(originalPrice)
	discountPriceStr := strconv.Itoa(discountPrice)

	descImages := pd.Images
	descImages = append(descImages, "https://d3vx04mz0cr7rc.cloudfront.net/alloff-products-detail.jpeg")
	descImageHtml := fmt.Sprintf("<p>시즌: %s</p><p>색상: %s </p><p>소재: %s </p><p>제조국: %s </p>", pd.Season+" "+strconv.Itoa(pd.Year), pd.Color, pd.Material, pd.MadeIn)
	for _, descImageUrl := range descImages {
		descImageHtml = descImageHtml + "<img src='" + descImageUrl + "'>"
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
		"40000",
		"40000", //교환배송비
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
