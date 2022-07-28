package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
)

func main() {
	log.Println("Running read csvs")
	filePath := "./coltorti/csvs/allProducts_220728_0100.csv"
	pds := coltorti.ReadFile(filePath)

	log.Println("# of products: ", len(pds))

	folders := coltorti.MakeFolders(len(pds))
	for idx, folder := range folders {
		lastIndex := (idx + 1) * 100
		if lastIndex > len(pds) {
			lastIndex = len(pds)
		}
		coltorti.WriteFile(folder, pds[idx*100:lastIndex])
	}
}
