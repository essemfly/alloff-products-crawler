package coltorti

import (
	"log"
	"strings"
)

func IsClothing(pd ColtortiProductInput) bool {
	return strings.Contains(pd.Category, "Clothing")
}

func IsFemale(pd ColtortiProductInput) bool {
	return !strings.Contains(pd.Category, "Men")
}

func Classifier(pd ColtortiProductInput) string {
	subCats := strings.Split(pd.Category, ">")
	lastCat := subCats[len(subCats)-1]
	return lastCat
}

func GetNaverCategoryCode(category string) string {
	categoryMappers := map[string]string{
		"Men > clothing > suits ":                             "50000840",
		"Men > clothing > Trousers > Trousers":                "50000836",
		"Men > clothing > suits > Polo shirts":                "50000833",
		"Women > accessories > Jewellery":                     "50004174",
		"Men > Clothing > Trousers > Trousers":                "50000836",
		"Women > Accessories > Jewellery":                     "50004174",
		"Men > Shoes > Lace-ups":                              "50000787",
		"Women > Clothing > Jackets/blazers":                  "50000815",
		"Women > Bags >Bags":                                  "50000639",
		"Women > Shoes >Boots and booties":                    "50004191",
		"Women > Shoes > Pumps":                               "50003825",
		"Men > Accessories > Small leather goods":             "50003988",
		"Womn > Shoes >Boots and booties":                     "50004191",
		"Women > Bags > Bags":                                 "50000639",
		"Men > Accessories > Scarves, hats and gloves":        "50000547",
		"Men > Bags > Crossbody bags":                         "50000648",
		"Women > Shoes > Boots and booties":                   "50004191",
		"Women > Shoes >Mules":                                "50003847",
		"Women > Shoes > Boots and boties":                    "50004191",
		"Women > Clothing > Dresses":                          "50000807",
		"Women > Shoes > Sandals":                             "50003842",
		"Men > Shoes > Moccasins":                             "50000784",
		"Men > Shoes > Boots":                                 "50000792",
		"Women > Clothing > Outerwear > Parkas":               "50000814",
		"Women > Shoes > Sneakers":                            "50003822",
		"Men > Clothing > Beachwear/Underwear > Beach towels": "50006380",
		"Men > Clothing > Trousers > Bermuda shorts":          "50000810",
		"Women > Bags > Clutches":                             "50000642",
		"Beauty & Lifestyle > Beauty > Face":                  "50004395",
		"Men > Clothing > Knitwear > Pullovers":               "50000831",
		"Men > Shoes > Sneakers":                              "50000788",
		"Men > Clothing > Beachwear/Underwear > Swimsuits":    "50006379",
		"Women > Shoes > Mules":                               "50003847",
		"Men > Clothing > Topwear > T-shirts":                 "50000830",
		"Men > Clothing > Outerwear > Jackets":                "50000838",
		"Men > Clothing > Outerwear > Bomber":                 "50000837",
		"Women > Bags > Tote bags":                            "50000647",
		"Men > Shoes > Sandals":                               "50000789",
		"Men > Accessories > Wallets":                         "50003985",
		"Men > Clothing > Topwear > Sweaters":                 "50000831",
		"Women > Accessories > Scarves, hats and gloves":      "50004010",
		"Men > Bags > Backpacks":                              "50000651",
		"Women > Clothing > Outerwear > Jackets":              "50000838",
		"Beauty & Lifestyle > Lifestyle > Books":              "50010561",
		"Women > Clothing > Trousers > Trousers":              "50000810",
		"Women > Shoes > Lace-ups":                            "50003819",
		"Men > Clothing > Shirts":                             "50000830",
		"Women > Clothing > Beachwear > Swimsuits":            "50006381",
		"Men > Bags > Business and travel bags":               "50000658",
		"Men > Accessories > Belts":                           "50003988",
		"Men > Accessories > Small leather goods > Beauty & Lifestyle > Lifestyle > Umbrella": "50004018",
		"Men > Clothing > Jeans":                                                                             "50000836",
		"Women > Accessories > Belt":                                                                         "50000539",
		"Women > Accessories > Small leather goods":                                                          "50004194",
		"Women > Clothing > Tops > T-shirts and polo shirts":                                                 "50000804",
		"Women > Shoes > Flats":                                                                              "50003818",
		"Women > Clothing > Knitwear > Pullovers":                                                            "50000805",
		"Beauty & Lifestyle > Lifestyle > Fragrances & Candles":                                              "50003353",
		"Women > Clothing > Jeans > Jeans":                                                                   "50000810",
		"Women > Clothing > Skirts":                                                                          "50000808",
		"Women > Clothing > Tops > Sweaters":                                                                 "50000805",
		"Women > Clothing > Shirts > Shirts":                                                                 "50000804",
		"Women > Clothing > Jeans > Trousers > Denim shorts > Shorts":                                        "50000810",
		"Women > Clothing > Shirts > Blouses":                                                                "50000804",
		"Women > Accessories > Wallets":                                                                      "50003982",
		"Women > Clothing > Jeans > Shorts":                                                                  "50000810",
		"Women > Clothing > Tops > Tops":                                                                     "50000803",
		"Women > Clothing > Trousers > Denim shorts":                                                         "50000810",
		"Men > Clothing > Outerwear > Coats":                                                                 "50000839",
		"Men > Clothing > Outerwear > Pea coats":                                                             "50000839",
		"Men > Clothing > Jackets/blazers":                                                                   "50000838",
		"Men > Clothing > Outerwear > Parkas":                                                                "50000837",
		"Women > Clothing > Trousers > Leggings":                                                             "50000812",
		"Men > Clothing > Outerwear > Trench coats and rain coats":                                           "50000839",
		"Women > Clothing > Knitwear > Cardigans":                                                            "50000806",
		"Men > Clothing > Outerwear > Puffer jackets":                                                        "50000838",
		"Men > Clothing > Outerwear > Coats > Bomber":                                                        "50000839",
		"Women > Clothing > Outerwear > Trench coats and rain coats":                                         "50000822",
		"Men > Clothing > Outerwear > Jackets > Bomber":                                                      "50000837",
		"Men > Accessories > Hi-tech accessories > Beauty & Lifestyle > Lifestyle":                           "50000570",
		"Men > Bags > Belt bags":                                                                             "50000648",
		"Women > Shoes > Moccasins":                                                                          "50003820",
		"Women > Clothing > Outerwear > Leather clothing":                                                    "50000815",
		"Women > Clothing > Jumpsuits":                                                                       "50000811",
		"Men > Clothing > Outerwear > Leather jackets":                                                       "50000838",
		"Women > Clothing > Outerwear > Bomber > Bomber jackets":                                             "50000814",
		"Women > Clothing > Outerwear > Bomber":                                                              "50000814",
		"Women > Clothing > Outerwear > Pea coats":                                                           "50000813",
		"Women > Clothing > Outerwear > Puffer jackets":                                                      "50000815",
		"Men > Clothing > Knitwear > Cardigans":                                                              "50000832",
		"Men > Clothing > Knitwear > Pullovers > Cardigans":                                                  "50000832",
		"Women > Accessories > Socks":                                                                        "50003995",
		"Women > Clothing > Outerwear > Coats":                                                               "50000813",
		"Men > Shoes > Women > Shoes > Sneakers > Sneakers":                                                  "50000788",
		"Men > Clothing > Topwear > Polo shirts":                                                             "50000830",
		"Women > Bags > Backpacks":                                                                           "50000644",
		"Women > Clothing > Outerwear > Jackets > Bomber":                                                    "50000814",
		"Beauty & Lifestyle > Beauty > Body":                                                                 "50000408",
		"Men > Clothing > Suits":                                                                             "50000840",
		"Men > Clothing > Outerwear > Vests":                                                                 "50000834",
		"Women > Clothing > Outerwear > Vests":                                                               "50000817",
		"Women > Clothing > Outerwear > Knitwear > Pullovers > Vests":                                        "50000817",
		"Men > Accessories > Beauty & Lifestyle > Lifestyle > Home & Living":                                 "50001311",
		"Women > Accessories > Hi-tech accessories > Beauty & Lifestyle > Lifestyle":                         "50001311",
		"Women > Accessories > Beauty & Lifestyle > Lifestyle > Beauty > Home & Living > Beauty Accessories": "50001311",
		"Beauty & Lifestyle > Lifestyle > Home & Living":                                                     "50001311",
		"Women > Accessories > Small leather goods > Beauty & Lifestyle > Lifestyle > Umbrella":              "50004018",
		"Women > Clothing > Outerwear > Capes":                                                               "50000815",
		"Beauty & Lifestyle > Beauty > Beauty Accessories":                                                   "50001311",
		"Women > Beauty & Lifestyle > Beauty > Face":                                                         "50004392",
		"Women > Clothing > Lingerie":                                                                        "50003626",
		"Men > Accessories > Women > Accessories > Belt > Belts":                                             "50000539",
		"Beauty & Lifestyle > Lifestyle > Hi-Tech":                                                           "50001311",
		"Men > Clothing > Jackets/blazers > Outerwear > Coats":                                               "50000838",
		"Men > Clothing > Jeans > Trousers > Bermuda shorts":                                                 "50000836",
		"Women > Shoes > Sandals > Mules":                                                                    "50003847",
		"Men > Bags > Tote bags":                                                                             "50000647",
		"Beauty & Lifestyle > Beauty > Body > Beauty Accessories":                                            "50001311",
		"Beauty & Lifestyle > Lifestyle > Fragrances & Candles > Home & Living":                              "50001311",
		"Women > Accessories > Beauty & Lifestyle > Lifestyle > Home & Living":                               "50001311",
		"Men > Accessories > Jewellery":                                                                      "50004174",
		"Men > Clothing > Outerwear > Jackets > Puffer jackets":                                              "50000838",
		"Men > Clothing > Knitwear > Vests":                                                                  "50000834",
		"Beauty & Lifestyle > Beauty > Face > Body":                                                          "50004392",
		"Men > Accessories > Socks":                                                                          "50004002",
		"Women > Bags > Belt bags":                                                                           "50000539",
		"Women > Accessories > Beauty & Lifestyle > Lifestyle > Fragrances & Candles":                        "50003353",
		"Women > Bags > Tote bags > Bags":                                                                    "50000640",
		"Women > Clothing > Jeans > Skirts > Shorts":                                                         "50000808",
		"Women > Clothing > Trousers > Tops > Denim shorts > Tops":                                           "50000810",
		"Women > Accessories > Beauty & Lifestyle > Lifestyle > Umbrella":                                    "50004018",
		"Men > Clothing > Topwear > T-shirts > Sweaters":                                                     "50000830",
		"Women > Clothing > Beachwear > Cover-ups":                                                           "50003055",
		"Women > Accessories > Beauty & Lifestyle > Lifestyle > Textiles":                                    "50001311",
		"Women > Clothing > Jeans > Jeans > Trousers > Trousers":                                             "50000810",
		"Women > Clothing > Outerwear > Trench coats and rain coats > Jackets":                               "50000813",
		"Men > Shoes > Espadrilles":                                                                          "50000785",
		"Women > Clothing > Knitwear > Cardigans > Skirts":                                                   "50000806",
		"Women > Clothing > Outerwear > Parkas > Jackets":                                                    "50000815",
		"Women > Clothing > Bomber jackets":                                                                  "50000814",
		"Women > Accessories > Sunglasses":                                                                   "50000554",
		"Women > Clothing > Knitwear > Pullovers > Tops > Sweaters":                                          "50000805",
		"Men > Accessories > Women > Accessories > Wallets > Wallets":                                        "50003982",
		"Women > Clothing > Knitwear > Pullovers > Trousers > Trousers":                                      "50000805",
		"Men > Accessories > Beauty & Lifestyle > Lifestyle > Textiles":                                      "50001311",
		"Men > Clothing > Trousers > Bermuda shorts > Trousers":                                              "50000836",
		"Men > Clothing > Knitwear > Cardigans > Vests":                                                      "50000834",
		"Women > Shoes > Flats > Mules":                                                                      "50003847",
		"Men > Clothing > Outerwear > Jackets > Coats":                                                       "50000839",
		"Women > Clothing > Outerwear > Trousers > Trousers > Jackets":                                       "50000815",
		"Women > Clothing > Shirts > Tops > Shirts > Tops":                                                   "50000803",
		"Women > Clothing > Shirts > Blouses > Shirts":                                                       "50000804",
		"Women > Accessories > Small leather goods > Hi-tech accessories":                                    "50001311",
		"Men > Clothing > Knitwear > Pullovers > Vests":                                                      "50000834",
		"Men > Accessories > Hi-tech accessories":                                                            "50004194",
		"Men > Accessories > Lifestyle":                                                                      "50004194",
		"Men > Clothing > Topwear > T-shirts > Women > Clothing > Tops > T-shirts and polo shirts":           "50000830",
		"Men > Accessories":                                                                                  "50004194",
		"Men > Accessories > Small leather goods > Women > Accessories > Small leather goods":                "50004194",
		"Women > Accessories > Hi-tech accessories":                                                          "50004194",
		"Women > Clothing > Dresses > Beachwear > Cover-ups":                                                 "50003055",
		"Women > Beauty > Beauty > Beauty & Lifestyle > Beauty > Body":                                       "50000281",
		"Women > Accessories > Lifestyle":                                                                    "50004194",
		"Women > Bags > Clutches > Bags":                                                                     "50000642",
		"Women > Accessories > Wallets > Small leather goods":                                                "50003982",
		"Women > Clothing > Shirts > Shirts > Beachwear > Swimsuits":                                         "50005462",
		"Men > Clothing > Jumpsuits":                                                                         "50008960",
		"Women > Clothing > Shirts > Dresses > Shirts":                                                       "50000807",
		"Women > Beauty > Beauty > Beauty & Lifestyle > Beauty > Face":                                       "50004392",
		"Beauty & Lifestyle > Lifestyle > Umbrella":                                                          "50004018",
		"Women > Bags": "50000639",
		"Men > Clothing > Women > Clothing > Outerwear > Outerwear > Jackets > Jackets": "50000838",
		"Women > Clothing > Jackets/blazers > Skirts":                                   "50000808",
		"Men > Clothing > Jackets/blazers > Outerwear > Jackets":                        "50000838",
		"Women > Clothing > Outerwear > Fur jackets":                                    "50000815",
		"Men > Accessories > Sunglasses":                                                "50000554",
		"Beauty & Lifestyle > Beauty > Hair":                                            "50004009",
		"Beauty & Lifestyle > Lifestyle > Pets":                                         "50006730",
	}

	if val, ok := categoryMappers[category]; ok {
		return val
	}

	log.Println("no classified cats", category)
	return "50000846"
}

func GetProductKeywords(category string) []string {
	isMale := true
	isClothing := false
	isAccessory := false
	isBag := false
	isShoes := false

	if strings.Contains(category, "Women") {
		isMale = false
	}

	if strings.Contains(category, "Clothing") {
		isClothing = true
	} else if strings.Contains(category, "Accessories") {
		isAccessory = true
	} else if strings.Contains(category, "Bag") {
		isBag = true
	} else if strings.Contains(category, "Shoes") {
		isShoes = true
	}

	if isAccessory {
		return []string{
			"패션액세서리",
			"코디하기좋은",
			"코디하기좋은옷",
			"코디하기좋은액세서리",
			"선물하기좋은",
			"선물용",
			"남자친구선물",
			"여자친구선물",
			"꾸안꾸 ",
			"요즘유행하는패션",
		}
	}

	if isMale {
		if isBag {
			return []string{
				"2030가방",
				"코디하기좋은가방",
				"코디하기좋은가방추천",
				"유행하는가방",
				"동창회코디",
				"기념일선물",
				"남자가방코디",
				"가을에어울리는스타일 ",
				"클러치백대용 ",
			}
		} else if isShoes {
			return []string{
				"남자신발코디",
				"유행하는신발",
				"신발코디",
				"가벼운신발",
				"편하게신는신발",
				"여름에신기좋은",
				"런닝신발",
				"남자신발선물",
				"남자신발브랜드",
				"남자유행하는코디",
			}
		} else if isClothing {
			return []string{
				"남자유행하는룩",
				"코디하기좋은옷",
				"여름에시원한옷",
				"여름에좋은",
				"올여름유행룩",
				"올여름유행옷",
				"올여름유행패션",
				"남자유행하는옷 ",
				"요즘유행하는패션",
				"남자유행코디 ",
			}
		}
	} else {
		if isBag {
			return []string{
				"2030가방",
				"코디하기좋은가방",
				"코디하기좋은가방추천",
				"유행하는가방",
				"유행하는가방",
				"동창회선물",
				"기념일선물",
				"여자가방코디",
				"가을에어울리는스타일",
			}
		} else if isShoes {
			return []string{
				"여자신발코디",
				"유행하는신발",
				"여자신발종류 ",
				"가벼운신발",
				"편하게신는신발",
				"여름에신기좋은",
				"런닝신발",
				"여자신발단화",
				"여자신발브랜드",
				"여자유행하는코디",
			}
		} else if isClothing {
			return []string{
				"이쁜미시룩",
				"코디하기좋은",
				"코디하기좋은옷",
				"여름에시원한옷",
				"여름에좋은",
				"올여름유행룩",
				"올여름유행옷",
				"올여름유행패션",
				"꾸안꾸",
				"요즘유행하는패션",
			}
		}
	}
	return nil
}
