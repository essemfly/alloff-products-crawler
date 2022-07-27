package coltorti

import (
	"strconv"
	"strings"
)

type ProductOption struct {
	SizeInfo string
	SizeName string
	Quantity int
}

type ColtortiProductInput struct {
	ProductURL    string
	Images        []string
	Brand         string
	ProductID     string
	Season        string
	Year          int
	Color         string
	MadeIn        string
	Material      string
	Name          string
	Description   string
	Category      string
	Quantity      int
	OriginalPrice float64
	CurrencyType  string
	DiscountRate  int
	SizeOptions   []ProductOption
	FTA           bool
}

func (pd *ColtortiProductInput) ToProductTemplate() []string {
	optionStrings := []string{}
	optionQuantityStrings := []string{}
	for _, optionName := range pd.SizeOptions {
		optionStrings = append(optionStrings, optionName.SizeInfo+optionName.SizeName)
		optionQuantityStrings = append(optionQuantityStrings, strconv.Itoa(optionName.Quantity))
	}

	nameTranslated := pd.Name
	// nameTranslated, err := utils.TranslateText(language.Korean.String(), pd.Name)
	// if err != nil {
	// 	log.Println("info translate key err", err)
	// }
	// descTranslated, err := utils.TranslateText(language.Korean.String(), pd.Description)
	// if err != nil {
	// 	log.Println("info translate key err", err)
	// }

	ourPrice := strconv.Itoa(CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType, IsClothing(*pd), pd.FTA))

	return []string{
		"신상품",
		"50000805",
		nameTranslated,
		ourPrice,
		strconv.Itoa(pd.Quantity),
		"-",
		"010-4118-1406",
		pd.Images[0],
		strings.Join(pd.Images, ","),
		"alloff-products-detail.jpeg",
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
		"택배",
		"무료",
		"0",
		"선결제",
		"",
		"",
		"80000",
		"80000", //교환배송비
		"",
		"",
		"0",
		"%",
		"0",
		"%",
		"0",
		"0",
		"0",
		"0",
		"%",
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
		ourPrice,
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
