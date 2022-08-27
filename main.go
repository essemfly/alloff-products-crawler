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
	outputFileNames := intrend.LoadCsvFiles()
	productInfosMap := intrend.GetCurrentTranslatedInfo(outputFileNames)

	sources := SelectSources()
	products := LoadProducts(sources)
	SpawnWorkers(products, productInfosMap)
}

func SelectSources() []*domain.Source {
	intrendTabs := []string{
		"https://it.intrend.it/special-price/abbigliamento-special",      // 티셔츠느낌
		"https://it.intrend.it/special-price/cappotti-e-giacche-special", // 코트느낌
		"https://it.intrend.it/special-price/borse-e-accessori-special",  // 가방
		"https://it.intrend.it/special-price/scarpe-special",             // 신발 가죽

	}

	return []*domain.Source{
		{Code: "INTREND", Tabs: intrendTabs, TextFilename: "intrend.json"},
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
			return coltorti.ReadFile(source.ExcelFilename)
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
