package coltorti

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
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

	return []string{
		"신상품",
		"50000805",
		pd.Name,
		strconv.Itoa(pd.Quantity),
		"-",
		"-",
		"mainimagefile",
		"imagefiles",
		"detailimage",
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
		strconv.Itoa(CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType)),
		strings.Join(optionQuantityStrings, ","),
		"",
		"",
		"",
		"",
		pd.Name,
		pd.ProductID,
		"",
		pd.Brand,
		"N",
	}

}

// Product Url
// Image
// Image1
// Image2
// Image3
// Brand
// Sku Styleisnow
// Season
// Year
// Sku Supplier
// Variant
// Color detail
// Color Supplier
// Made in
// Material
// Name
// Description
// Categories
// Qty
// Retail Price
// Discount
// Size Info
// Size
// Qty Detail
// Bag length
// Bag height
// Bag weight
// Handle height
// Shoulder bag length
// Belt length
// Belt height
// Accessory length
// Accessory height
// Accessory weight
// Heel height
// Plateau height
// Insole length
// Color Styleisnow ITA
// FTA
// EAN
// Nome ITA
// Descrizione ITA
// Star

func ReadFile(filePath string) []ColtortiProductInput {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panicln("file not found", err)
	}
	defer file.Close()

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.Comma = ';'
	rdr.LazyQuotes = true

	// csv 내용 모두 읽기
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Panicln("err on read all", err)
	}

	products := []ColtortiProductInput{}

	// 행,열 읽기
	for i, row := range rows {
		if i == 0 {
			continue
		}

		images := FilterEmptyImageUrl([]string{
			row[1], row[2], row[3], row[4],
		})
		if len(images) == 0 {
			continue
		}

		yearInString, _ := strconv.Atoi(row[8])
		quantityInString, _ := strconv.Atoi(row[18])
		discountrateInString, _ := strconv.Atoi(row[20])
		newProduct := ColtortiProductInput{
			ProductURL:    row[0],
			Images:        images,
			Brand:         row[5],
			ProductID:     row[6],
			Season:        row[7],
			Year:          yearInString,
			Color:         row[11],
			MadeIn:        row[13],
			Material:      row[14],
			Name:          row[15],
			Description:   row[16],
			Category:      row[17],
			Quantity:      quantityInString,
			OriginalPrice: priceParser(row[19]),
			CurrencyType:  "EUR",
			DiscountRate:  discountrateInString,
			SizeOptions:   optionParser(row[21], row[22], row[23]),
			FTA:           row[38] == "TRUE",
		}
		products = append(products, newProduct)
	}

	return products
}

func priceParser(priceInCsv string) float64 {
	tmp := strings.Split(priceInCsv, " ")[1]
	tmp = strings.ReplaceAll(tmp, ",", "")
	if ret, err := strconv.ParseFloat(tmp, 32); err == nil {
		return ret
	}
	return 0.0
}

func optionParser(sizeInfo, Size, Qty string) []ProductOption {
	options := []ProductOption{}

	optionSizes := strings.Split(Size, ",")
	optionQuantities := strings.Split(Qty, ",")

	for i, option := range optionSizes {
		optionQuantityInt, err := strconv.Atoi(optionQuantities[i])
		if err != nil {
			log.Panicln("option quantity parsing error", err)
		}
		if optionQuantityInt > 0 {
			options = append(options, ProductOption{
				SizeInfo: sizeInfo,
				SizeName: option,
				Quantity: optionQuantityInt,
			})
		}
	}
	return options
}
