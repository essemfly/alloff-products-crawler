package intrend

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/essemfly/alloff-products/domain"
	"github.com/gocolly/colly"
)

const (
	DEFAULT_STOCK = 2
)

func GetSources() []string {
	return []string{
		/*
			"https://it.intrend.it/promozioni",                               // 이거 안에 엄청 여러개가 있다.
			"https://it.intrend.it/promozioni-curvy",                         // 이거 안에도 엄청 여러가지가 있다.
		*/
		"https://it.intrend.it/special-price/abbigliamento-special",      // 총체적 의류
		"https://it.intrend.it/special-price/cappotti-e-giacche-special", // 코트
		"https://it.intrend.it/special-price/borse-e-accessori-special",  // 악세서리
		"https://it.intrend.it/special-price/scarpe-special",             // 가죽신발?
	}
}

func CrawlIntrend(source string) []*domain.Product {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	products := []*domain.Product{}
	currentPageNum := 0

	c.OnHTML(".js-product-list .js-pagination-container .js-product-card", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".full-price")
		originalPrice := 0.0
		if originalPriceStr != "" {
			originalPriceStr = strings.Split(originalPriceStr, " ")[1]
			originalPriceStr = strings.Replace(originalPriceStr, ".", "", -1)
			originalPriceStr = strings.Replace(originalPriceStr, ",", ".", -1)
			originalPrice, err := strconv.ParseFloat(originalPriceStr, 32)
			if err != nil {
				msg := fmt.Sprintf("err on parse original price %v : ", originalPrice)
				log.Println("err", msg)
				return
			}
		}

		discountedPriceStr := e.ChildText(".price")
		discountedPriceStr = strings.Split(discountedPriceStr, " ")[1]
		discountedPriceStr = strings.Replace(discountedPriceStr, ".", "", -1)
		discountedPriceStr = strings.Replace(discountedPriceStr, ",", ".", -1)
		discountedPrice, err := strconv.ParseFloat(discountedPriceStr, 32)
		if err != nil {
			msg := fmt.Sprintf("err on parse discount price %v : ", originalPrice)
			log.Println("err", msg)
			return
		}

		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0.0 {
			originalPrice = float64(genOriginalPrice(float32(discountedPrice)))
		}

		productUrl := "https://it.intrend.it" + e.ChildAttr(".js-anchor", "href")

		title, composition, productColor, productID, images, sizes, categories, inventories, description := getIntrendDetail(productUrl)
		options := []domain.ProductOption{}
		if len(sizes) == 0 {
			options = append(options, domain.ProductOption{
				SizeInfo: "FREE",
				SizeName: "FREE",
				Quantity: 1,
			})
		} else {
			for _, inv := range inventories {
				options = append(options, domain.ProductOption{
					SizeInfo: inv.SizeInfo,
					SizeName: inv.SizeName,
					Quantity: inv.Quantity,
				})
			}
		}

		// forbidden 403 case
		if title == "" {
			msg := fmt.Sprintf("not allowed access by intrend server on : %s\n", source)
			log.Println("err", msg)
			return
		}

		// Title & Description Translate
		addRequest := &domain.Product{
			Source:            domain.Source{Code: "INTREND"},
			ProductURL:        productUrl,
			Images:            images,
			Brand:             "막스마라 인트렌드",
			ProductID:         productID,
			ProductStyleisNow: productID,
			Color:             productColor,
			MadeIn:            "Italy",
			Name:              title,
			Description:       description["설명"],
			Quantity:          2,
			OriginalPrice:     originalPrice,
			Material:          composition,
			CurrencyType:      "EUR",
			DiscountRate:      int(PercentageChange(int(originalPrice), int(discountedPrice))),
			Season:            "",
			Category:          categories[0],
			Year:              0,
			SizeOptions:       options,
			FTA:               false,
		}

		log.Println("Add Request - Category", addRequest.Category)
		products = append(products, addRequest)
	})

	c.OnHTML(".js-pager .container-fluid ul", func(e *colly.HTMLElement) {
		lastPageStr := e.ChildAttr("li:last-child a", "data-page")
		lastPageNum, _ := strconv.Atoi(lastPageStr)
		if currentPageNum < lastPageNum {
			currentPageNum += 1
			url := source + "?page=" + strconv.Itoa(currentPageNum)
			c.Visit(url)
		}
	})

	err := c.Visit(source)
	if err != nil {
		log.Println("error occurred in crawl intrend ", err)
	}

	return products
}

type IntrendStock struct {
	STYCD    string `json:"STYCD"`
	SIZECD   string `json:"SIZECD"`
	COLCD    string `json:"COLCD"`
	SIZECDNM string `json:"SIZECDNM"`
	SALECNT  int    `json:"SALECNT"`
	STOCKQTY int    `json:"STOCKQTY"`
}

func getIntrendDetail(productUrl string) (title, composition, productColor, productID string, imageUrls []string, sizes, categories []string, inventories []*domain.ProductOption, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	isDigit := regexp.MustCompile(`^\d*\.?\d+$`)

	description = map[string]string{}

	c.OnHTML(".pp_mod-prod-desc-head", func(e *colly.HTMLElement) {
		title = e.ChildText(".title")
	})

	c.OnHTML(".product-gallery", func(e *colly.HTMLElement) {
		e.ForEach(".js-item-image .item img", func(_ int, el *colly.HTMLElement) {
			imageUrls = append(imageUrls, strings.Split(el.Attr("src"), "#")[0])
		})
	})

	c.OnHTML(".sizes .sizes-select-wrapper .sizes-select-list", func(e *colly.HTMLElement) {
		e.ForEach(".list-inline li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("span .value")
			if isDigit.MatchString(size) {
				size = "IT" + size
			}
			sizes = append(sizes, size)
			if el.Attr("class") != "li-disabled" {
				inventories = append(inventories, &domain.ProductOption{
					Quantity: DEFAULT_STOCK,
					SizeInfo: size,
					SizeName: size,
				})
			}
		})
	})

	c.OnHTML(".breadcrumb", func(e *colly.HTMLElement) {
		e.ForEach(".cta-underlined", func(i int, h *colly.HTMLElement) {
			if i == 1 {
				categories = append(categories, h.Attr("href"))
			}
		})
	})

	c.OnHTML(".product-name-value", func(h *colly.HTMLElement) {
		productID = h.Text
	})

	c.OnHTML(".swatches .title", func(e *colly.HTMLElement) {
		productColor = e.ChildText(".value")
	})

	c.OnHTML("#description .details-tab-content", func(e *colly.HTMLElement) {
		description["설명"] = e.ChildText("p")
	})

	c.OnHTML("#composition .details-tab-content", func(e *colly.HTMLElement) {
		texts := ""
		e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
			texts += el.Text
		})
		composition = texts
	})

	// c.OnHTML("#fitting .details-tab-content", func(e *colly.HTMLElement) {
	// 	texts := ""
	// 	e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
	// 		texts += el.Text
	// 	})
	// 	description["모델"] = texts
	// })

	c.Visit(productUrl)
	return
}

func genOriginalPrice(discountedPrice float32) float32 {
	originalPrice := discountedPrice
	if discountedPrice <= 20 {
		disRate := genRandRate(78, 80)
		originalPrice = discountedPrice * disRate
	} else if 20 < discountedPrice && discountedPrice <= 50 {
		disRate := genRandRate(70, 78)
		originalPrice = discountedPrice * disRate
	} else if 50 < discountedPrice && discountedPrice <= 70 {
		disRate := genRandRate(65, 75)
		originalPrice = discountedPrice * disRate
	} else if 70 < discountedPrice && discountedPrice <= 100 {
		disRate := genRandRate(55, 72)
		originalPrice = discountedPrice * disRate
	} else if 100 < discountedPrice && discountedPrice <= 300 {
		disRate := genRandRate(45, 60)
		originalPrice = discountedPrice * disRate
	} else if 300 < discountedPrice && discountedPrice <= 400 {
		disRate := genRandRate(40, 55)
		originalPrice = discountedPrice * disRate
	} else if 400 < discountedPrice {
		disRate := genRandRate(30, 45)
		originalPrice = discountedPrice * disRate
	}
	return originalPrice
}

func genRandRate(min, max int) float32 {
	rand.Seed(time.Now().UnixNano())
	rng := max - min + 1
	randFloat := (float32(rand.Intn(rng)) + float32(min) + 100.00) / 100.00
	return randFloat
}

func PercentageChange(old, new int) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}
