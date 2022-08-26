package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
	"github.com/essemfly/alloff-products/domain"
	"github.com/essemfly/alloff-products/intrend"
	"github.com/essemfly/alloff-products/worker"
)

func main() {
	// hoit, err := utils.TranslateText(language.Korean.String(), "My name is socks")
	// log.Println("HOIT", hoit, err)
	sources := SelectSources()
	products := LoadProducts(sources)

	SpawnWorkers(products)
}

func SelectSources() []*domain.Source {
	excelFilename := ""
	coltortiBrands := []string{}
	intrendTabs := []string{
		"https://it.intrend.it/special-price/abbigliamento-special",      // 티셔츠느낌
		"https://it.intrend.it/special-price/cappotti-e-giacche-special", // 코트느낌
		"https://it.intrend.it/special-price/borse-e-accessori-special",  // 가방
		"https://it.intrend.it/special-price/scarpe-special",             // 신발 가죽

	}

	return []*domain.Source{
		{Code: "INTREND", Tabs: intrendTabs},
		{Code: "COLTORTI", Brands: coltortiBrands, ExcelFilename: excelFilename},
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
		pds := []*domain.Product{}
		for _, url := range source.Tabs {
			intrend.CrawlIntrend(url)
		}
		return pds
	} else if source.Code == "COLTORTI" {
		if source.ExcelFilename != "" {
			return coltorti.ReadFile(source.ExcelFilename)
		}
	}
	return nil
}

func SpawnWorkers(products []*domain.Product) {
	const numWorkers = 12

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

		go worker.WriteFile(workers, done, folder, products[idx*100:lastIndex])
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}
