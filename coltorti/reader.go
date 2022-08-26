package coltorti

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/essemfly/alloff-products/domain"
)

func ReadFile(filePath string) []*domain.Product {
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

	products := []*domain.Product{}

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
		discountrateRawString := strings.Split(row[20], ".")[0]
		discountrateInString, _ := strconv.Atoi(discountrateRawString)
		newProduct := domain.Product{
			Source:            domain.Source{Code: "COLTORTI"},
			ProductURL:        row[0],
			Images:            images,
			Brand:             row[5],
			ProductID:         row[9],
			ProductStyleisNow: row[6],
			Season:            row[7],
			Year:              yearInString,
			Color:             row[11],
			MadeIn:            row[13],
			Material:          row[14],
			Name:              row[15],
			Description:       row[16],
			Category:          row[17],
			Quantity:          quantityInString,
			OriginalPrice:     priceParser(row[19]),
			CurrencyType:      "EUR",
			DiscountRate:      discountrateInString,
			SizeOptions:       optionParser(row[21], row[22], row[23]),
			FTA:               row[38] == "true",
		}

		products = append(products, &newProduct)
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

func optionParser(sizeInfo, Size, Qty string) []domain.ProductOption {
	options := []domain.ProductOption{}

	optionSizes := strings.Split(Size, ",")
	optionQuantities := strings.Split(Qty, ",")

	for i, option := range optionSizes {
		optionQuantityInt, err := strconv.Atoi(optionQuantities[i])
		if err != nil {
			log.Panicln("option quantity parsing error", err)
		}
		if optionQuantityInt > 0 {
			options = append(options, domain.ProductOption{
				SizeInfo: sizeInfo,
				SizeName: option,
				Quantity: optionQuantityInt,
			})
		}
	}
	return options
}
