package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/essemfly/alloff-products/coltorti"
)

func main() {
	log.Println("Running read csvs")
	filePath := "./coltorti/csvs/allProducts.csv"
	pds := coltorti.ReadFile(filePath)

	log.Println("# of products: ", len(pds))
	f, err := os.OpenFile("./outputs/outputTemplate.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, pd := range pds {
		filenames := coltorti.CacheProductImages(pd)
		log.Println("Product", filenames)
		pd.Images = filenames
		if err := w.Write(pd.ToProductTemplate()); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
