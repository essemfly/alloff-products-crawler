package coltorti

import (
	"strconv"
	"strings"

	"github.com/essemfly/alloff-products/utils"
)

type ProductOption struct {
	SizeInfo string
	SizeName string
	Quantity int
}

type ColtortiProductInput struct {
	ProductURL     string
	Images         []string
	ImageFilenames []string
	Brand          string
	ProductID      string
	Season         string
	Year           int
	Color          string
	MadeIn         string
	Material       string
	Name           string
	Description    string
	Category       string
	Quantity       int
	OriginalPrice  float64
	CurrencyType   string
	DiscountRate   int
	SizeOptions    []ProductOption
	FTA            bool
}

func (pd *ColtortiProductInput) ToProductTemplate() []string {
	optionStrings := []string{}
	optionQuantityStrings := []string{}
	for _, optionName := range pd.SizeOptions {
		optionStrings = append(optionStrings, optionName.SizeInfo+optionName.SizeName)
		optionQuantityStrings = append(optionQuantityStrings, strconv.Itoa(optionName.Quantity))
	}

	names := []string{
		GetKoreanBrandName(pd.Brand),
		utils.StringTruncater(pd.Name, 25),
		pd.ProductID,
	}
	nameTranslated := strings.Join(names, " ")
	// nameTranslated, err := utils.TranslateText(language.Korean.String(), pd.Name)
	// if err != nil {
	// 	log.Println("info translate key err", err)
	// }
	// descTranslated, err := utils.TranslateText(language.Korean.String(), pd.Description)
	// if err != nil {
	// 	log.Println("info translate key err", err)
	// }

	originalPrice := CalculatePrice(pd.OriginalPrice, 0, pd.CurrencyType, IsClothing(*pd), pd.FTA)
	ourPrice := CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType, IsClothing(*pd), pd.FTA)
	discountPrice := originalPrice - ourPrice
	originalPriceStr := strconv.Itoa(originalPrice)
	discountPriceStr := strconv.Itoa(discountPrice)

	descImages := pd.Images
	descImages = append(descImages, "https://d3vx04mz0cr7rc.cloudfront.net/alloff-products-detail.jpeg")

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
		strings.Join(descImages, ","),
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
		"개",
		"",
		"원",
		"",
		"원",
		"0",
		"0",
		"0",
		"0",
		"0",
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
	}

}
