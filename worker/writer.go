package worker

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/essemfly/alloff-products/coltorti"
	"github.com/essemfly/alloff-products/domain"
	"github.com/essemfly/alloff-products/intrend"
)

func LoadCsvFiles() []string {
	csvFiles := []string{}
	files, err := ioutil.ReadDir("./outputs")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			err := filepath.Walk("./outputs/"+f.Name(),
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					r, err := regexp.MatchString(".csv", path)
					if err == nil && r {
						csvFiles = append(csvFiles, path)
					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
		}
	}

	return csvFiles
}

func WriteFile(worker chan bool, done chan bool, foldername string, pds []*domain.Product, prevProducts map[string][]string) {
	filepath := foldername + "/" + "output.csv"
	// f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, pd := range pds {
		filenames := CacheProductImages(foldername, pd)
		pd.ImageFilenames = filenames

		row := GetProductTemplate(pd)
		// for intrend
		// row := intrend.CheckAlreadyHaveProductRow(prevProducts, pd)
		// if row == nil {
		// 	row = GetProductTemplate(pd)
		// }

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	<-worker
	done <- true
}

func GetProductTemplate(pd *domain.Product) []string {
	if pd.Source.Code == "INTREND" {
		return intrend.GetIntrendTemplate(pd)
	} else if pd.Source.Code == "COLTORTI" {
		return coltorti.GetColtortiTemplate(pd)
	}
	return []string{}
}

func MakeFolders(numPds int) []string {
	t := time.Now()
	foldernames := []string{}
	for i := 0; i <= numPds/100; i++ {
		foldername := "./outputs/" + t.Format("2006-01-02") + "-" + strconv.Itoa(i)
		os.Mkdir(foldername, 0755)
		foldernames = append(foldernames, foldername)
	}
	return foldernames
}
