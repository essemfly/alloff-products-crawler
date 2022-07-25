package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
)

func main() {
	log.Println("Running read csvs")
	filePath := "./coltorti/csvs/allProducts.csv"
	pds := coltorti.ReadFile(filePath)
	log.Println("length", len(pds))
}
