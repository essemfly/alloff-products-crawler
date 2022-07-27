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
		"Men > clothing > suits ":                      "50000840",
		"Men > clothing > Trousers > Trousers":         "50000836",
		"Men > clothing > suits > Polo shirts":         "50000833",
		"Women > accessories > Jewellery":              "50004174",
		"Men > Clothing > Trousers > Trousers":         "50000836",
		"Women > Accessories > Jewellery":              "50004174",
		"Men > Shoes > Lace-ups":                       "50000787",
		"Women > Clothing > Jackets/blazers":           "50000815",
		"Women > Bags >Bags":                           "50000639",
		"Women > Shoes >Boots and booties":             "50004191",
		"Women > Shoes > Pumps":                        "50003825",
		"Men > Accessories > Small leather goods":      "50003988",
		"Womn > Shoes >Boots and booties":              "50004191",
		"Women > Bags > Bags":                          "50000639",
		"Men > Accessories > Scarves, hats and gloves": "50000547",
		"Men > Bags > Crossbody bags":                  "50000648",
		"Women > Shoes > Boots and booties":            "50004191",
		"Women > Shoes >Mules":                         "50003847",
		"Women > Shoes > Boots and boties":             "50004191",
		"Women > Clothing > Dresses":                   "50000807",
		"Women > Shoes > Sandals":                      "50003842",
		"Men > Shoes > Moccasins":                      "50000784",
		"Men > Shoes > Boots":                          "50000792",
	}

	if val, ok := categoryMappers[category]; ok {
		return val
	}

	log.Println("no classified cats", category)
	return "50000846"
}
