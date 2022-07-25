package coltorti

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func CacheProductImages(pd ColtortiProductInput) []string {
	newImageUrls, err := cacheImages(pd.ProductID, pd.Images)
	if err != nil {
		log.Println("cache image error", err)
		return nil
	}

	return newImageUrls
}

func cacheImages(pdInfoID string, images []string) ([]string, error) {
	newImageUrls := []string{}
	for idx, imgURL := range images {
		imgResp, err := http.Get(imgURL)
		if err != nil {
			log.Println("failed to get image from url: "+imgURL, err)
			return nil, err
		}
		defer imgResp.Body.Close()

		if imgResp.StatusCode != 200 {
			return nil, errors.New("status code: " + strconv.Itoa(imgResp.StatusCode))
		}

		extension, err := getFileExtensionFromUrl(imgURL)
		if err != nil {
			log.Println("failed to get extension from url: "+imgURL, err)
			return nil, err
		}

		filename := pdInfoID + "-" + strconv.Itoa(idx)

		file, err := os.Create("./coltorti/images/" + filename + "." + extension)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, imgResp.Body)
		if err != nil {
			log.Fatal(err)
		}

		newImageUrls = append(newImageUrls, filename+"."+extension)
	}

	return newImageUrls, nil
}

func getFileExtensionFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		return "", errors.New("couldn't find a period to indicate a file extension")
	}
	return u.Path[pos+1 : len(u.Path)], nil
}
