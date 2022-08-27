package intrend

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/essemfly/alloff-products/domain"
)

func TranslateMapper(originalName string) string {
	translatedName := ""

	// mappers := map[string]string{}

	return translatedName

}

func LoadCsvFiles() []string {
	csvFiles := []string{}
	files, err := ioutil.ReadDir("./outputs")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			err := filepath.Walk("./outputs/"+f.Name(),
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					r, err := regexp.MatchString(".csv", path)
					if err == nil && r {
						csvFiles = append(csvFiles, path)
					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
		}
	}

	return csvFiles
}

func GetCurrentTranslatedInfo(outputFileNames []string) map[string][]string {
	productPrevInfo := map[string][]string{}
	for _, fileName := range outputFileNames {
		file, err := os.Open(fileName)
		if err != nil {
			log.Println("Err in file opening", err)
		}
		defer file.Close()
		rdr := csv.NewReader(bufio.NewReader(file))
		rows, err := rdr.ReadAll()
		if err != nil {
			log.Panicln("err on read all", err, fileName)
		}

		for _, row := range rows {
			splitsOfName := strings.Split(row[2], " ")
			key := splitsOfName[len(splitsOfName)-1]
			productPrevInfo[key] = row
		}
	}
	return productPrevInfo
}

func CheckAlreadyHaveProductRow(prevProducts map[string][]string, pd *domain.Product) []string {
	if val, ok := prevProducts[pd.ProductID]; ok {
		log.Println("Yes it has!", pd.ProductID, pd.ProductStyleisNow)
		splitsOfName := strings.Split(val[2], " ")
		key := splitsOfName[len(splitsOfName)-1]
		newName := strings.Replace(val[2], key, pd.ProductStyleisNow, 1)
		val[2] = newName

		val[29] = "30000"
		val[30] = "30000"

		ourPrice := CalculatePrice(pd.OriginalPrice, pd.DiscountRate, pd.CurrencyType, IsClothing(pd.Category), pd.FTA)
		originalPrice := CalculateOriginalPrice(ourPrice, pd.DiscountRate)
		discountPrice := originalPrice - ourPrice
		originalPriceStr := strconv.Itoa(originalPrice)
		discountPriceStr := strconv.Itoa(discountPrice)

		val[3] = originalPriceStr
		val[34] = discountPriceStr
		return val
	}

	return nil
}
