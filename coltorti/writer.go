package coltorti

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteFile(filePath string, pds []ColtortiProductInput) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, pd := range pds {
		filenames := CacheProductImages(pd)
		log.Println("Product", filenames)
		pd.Images = filenames
		if err := w.Write(pd.ToProductTemplate()); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
