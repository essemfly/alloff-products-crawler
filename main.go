package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
	"github.com/essemfly/alloff-products/domain"
	"github.com/essemfly/alloff-products/intrend"
)

func main() {

	sources := SelectSources()
	products := LoadProducts(sources)

	SpawnWorkers(products)
}

func SelectSources() []*domain.Source {
	excelFilename := ""
	coltortiBrands := []string{}
	intrendTabs := []string{}

	return []*domain.Source{
		{Code: "INTREND", Tabs: intrendTabs, ExcelFilename: excelFilename},
		{Code: "COLTORTI", Brands: coltortiBrands},
	}
}

func LoadProducts(sources []*domain.Source) []*domain.Product {
	products := []*domain.Product{}
	for _, source := range sources {
		products = append(products, CrawlSource(source)...)
	}
	return products
}

func SpawnWorkers(products []*domain.Product) {
	const numWorkers = 12

	folders := domain.MakeFolders(len(products))
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

		go domain.WriteFile(workers, done, folder, products[idx*100:lastIndex])
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}

func CrawlSource(source *domain.Source) []*domain.Product {
	if source.Code == "INTRNED" {
		pds := []*domain.Product{}
		for _, url := range source.Tabs {
			intrend.CrawlIntrend(url)
		}
		return pds
	} else if source.Code == "COLTORTI" {
		return coltorti.ReadFile(source.ExcelFilename)
	}
	return nil
}
