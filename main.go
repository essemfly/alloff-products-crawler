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
	f, err := os.Create("./outputs/newCsvs.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, pd := range pds {
		if err := w.Write(pd.ToProductTemplate()); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
