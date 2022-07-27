package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
)

func main() {
	log.Println("Running read csvs")
	filePath := "./coltorti/csvs/allProducts.csv"
	pds := coltorti.ReadFile(filePath)

	log.Println("# of products: ", len(pds))

	folders := coltorti.MakeFolders(len(pds))
	for idx, folder := range folders {
		coltorti.WriteFile(folder, pds[idx*100:(idx+1)*100])
	}
}
