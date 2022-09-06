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
	sources := SelectSources()
	products := LoadProducts(sources)

	// Intrned의 경우에만 필요한듯?
	// outputFileNames := worker.LoadCsvFiles()
	// productInfosMap := intrend.GetCurrentTranslatedInfo(outputFileNames)

	log.Println("Length of products: ", len(products))
	SpawnWorkers(products, nil)
}

func SelectSources() []*domain.Source {
	// intrendTabs := []string{
	// 	"https://it.intrend.it/special-price/abbigliamento-special",      // 티셔츠느낌
	// 	"https://it.intrend.it/special-price/cappotti-e-giacche-special", // 코트느낌
	// 	"https://it.intrend.it/special-price/borse-e-accessori-special",  // 가방
	// 	"https://it.intrend.it/special-price/scarpe-special",             // 신발 가죽

	// }

	// return []*domain.Source{
	// 	{Code: "INTREND", Tabs: intrendTabs, TextFilename: "intrend.json"},
	// }

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

func LoadProducts(sources []*domain.Source) []*domain.Product {
	products := []*domain.Product{}
	for _, source := range sources {
		products = append(products, CrawlSource(source)...)
	}
	return products
}

func CrawlSource(source *domain.Source) []*domain.Product {
	if source.Code == "INTREND" {
		if source.TextFilename == "intrend.json" {
			return intrend.ReadFromFile(source.TextFilename)
		}

		pds := []*domain.Product{}
		for idx, url := range source.Tabs {
			log.Println("Intrend tab: ", idx, url)
			pds = append(pds, intrend.CrawlIntrend(url)...)
		}

		file, _ := json.MarshalIndent(pds, "", " ")
		_ = ioutil.WriteFile("intrend.json", file, 0644)

		return pds
	} else if source.Code == "COLTORTI" {
		if source.ExcelFilename != "" {
			return coltorti.ReadFile(source.ExcelFilename, source.Brands)
		}
	}
	return nil
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

		go worker.WriteFile(workers, done, folder, products[idx*100:lastIndex], prevProducts)
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}
