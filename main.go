package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/essemfly/alloff-products/coltorti"
	"github.com/essemfly/alloff-products/domain"
	"github.com/essemfly/alloff-products/intrend"
	"github.com/essemfly/alloff-products/worker"
)

func main() {
	// Sources to be selected
	sources := SelectSources()

	products := []*domain.Product{}
	for _, source := range sources {
		products = append(products, CrawlSource(source)...)
	}

	var productInfosMap map[string][]string
	// Intrned의 경우에만 작동
	// 기존에 만들어진 엑셀에서 파일에 맞는 상품이 있으면 그것을 그대로 가져와서 쓰게끔 되어있음
	// 근데 이것도 모든 outputs 폴더에 모든 걸 가져오게끔 되어있네.
	// outputFileNames := worker.LoadCsvFiles()
	// productInfosMap = intrend.GetCurrentTranslatedInfo(outputFileNames)

	log.Println("Length of products: ", len(products))
	SpawnWorkers(products, productInfosMap)
}

func SelectSources() []*domain.Source {
	// return getColotortiSources()
	return getIntrendSources("220929.json")
}

func getIntrendSources(prevTextFile string) []*domain.Source {
	intrendTabs := []string{
		"https://it.intrend.it/cappotti-e-giacche/giacconi-e-blazer/styleFilter/primavera-estate", // 코트 세일품목
		"https://it.intrend.it/special-price/abbigliamento-special",                               // 티셔츠느낌
		"https://it.intrend.it/special-price/cappotti-e-giacche-special",                          // 코트느낌
		"https://it.intrend.it/special-price/borse-e-accessori-special",                           // 가방
		"https://it.intrend.it/special-price/scarpe-special",                                      // 신발 가죽

	}
	return []*domain.Source{
		{Code: "INTREND", Tabs: intrendTabs, TextFilename: prevTextFile},
	}
}

func getColotortiSources() []*domain.Source {

	brands := []string{
		"ALEXANDER MCQUEEN",
		"AMIRI",
		"BALMAIN",
		"BURBERRY",
		"VETEMENTS",
		"MIU MIU",
		"MM6 MAISON MARGIELA",
		"FENDI ",
		"GANNI",
		"GIVENCHY",
		"ISABEL MARANT",
		"ISABEL MARANT ETOILE",
		"JIMMY CHOO",
		"BALENCIAGA",
		"LOEWE",
		"MAX MARA",
		"MAX MARA STUDIO",
		"MAX MARA THE CUBE",
		"DOLCE & GABBANA",
		"PALM ANGELS",
		"VALENTINO",
		"LANVIN",
		"MCM",
		"ALEXANDER WANG",
		"CHOLE",
		"DR. MARTENS",
		"GOLDEN GOOSE",
		"MONCLER",
		"MOOSE KNUCKLES",
		"MULBERRY",
		"PRADA",
	}
	return []*domain.Source{
		{Code: "COLTORTI", ExcelFilename: "./coltorti/csvs/allProducts_220904.csv", Brands: brands},
	}
}

func CrawlSource(source *domain.Source) []*domain.Product {
	if source.Code == "INTREND" {
		pds := []*domain.Product{}

		pds, err := CrawlIntrendFromPreviousTextFile(source.TextFilename)
		if err == nil {
			return pds
		}

		for idx, url := range source.Tabs {
			log.Println("Intrend tab: ", idx, url)
			pds = append(pds, intrend.CrawlIntrend(url)...)
		}

		file, _ := json.MarshalIndent(pds, "", " ")
		_ = ioutil.WriteFile(source.TextFilename, file, 0644)

		return pds
	} else if source.Code == "COLTORTI" {
		if source.ExcelFilename != "" {
			return coltorti.ReadFile(source.ExcelFilename, source.Brands)
		}
	}
	return nil
}

func CrawlIntrendFromPreviousTextFile(filename string) ([]*domain.Product, error) {
	// 기존에 긁은 자료가 있는경우, 그냥 사용
	file, err := ioutil.ReadFile(filename)
	if err == nil {
		var ret []*domain.Product
		if err := json.Unmarshal(file, &ret); err != nil {
			log.Panicln("file unmarshal failed", filename)
		}
		return ret, nil
	}
	return nil, err
}

func SpawnWorkers(products []*domain.Product, prevProducts map[string][]string) {
	const numWorkers = 21

	folders := worker.MakeFolders(len(products))
	workers := make(chan bool, numWorkers)
	done := make(chan bool, numWorkers)

	for c := 0; c < numWorkers; c++ {
		done <- true
	}

	for idx, folder := range folders {
		log.Println("Folder idx #", idx)
		lastIndex := (idx + 1) * 100
		if lastIndex > len(products) {
			lastIndex = len(products)
		}

		workers <- true
		<-done

		// Include Image downloading + Writing excel files + Translating
		go worker.WriteFile(workers, done, folder, products[idx*100:lastIndex], prevProducts)
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}
