package main

import (
	"log"

	"github.com/essemfly/alloff-products/coltorti"
)

const numWorkers = 12

func main() {
	log.Println("Running read csvs")
	filePath := "./coltorti/csvs/allProducts_22080317.csv"
	pds := coltorti.ReadFile(filePath)

	log.Println("# of products: ", len(pds))

	folders := coltorti.MakeFolders(len(pds))

	workers := make(chan bool, numWorkers)
	done := make(chan bool, numWorkers)

	for c := 0; c < numWorkers; c++ {
		done <- true
	}

	for idx, folder := range folders {
		log.Println("Folder idx #", idx)
		lastIndex := (idx + 1) * 100
		if lastIndex > len(pds) {
			lastIndex = len(pds)
		}

		workers <- true
		<-done

		go coltorti.WriteFile(workers, done, folder, pds[idx*100:lastIndex])
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}
