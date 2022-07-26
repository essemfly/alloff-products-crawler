package coltorti

import "strings"

func IsClothing(pd ColtortiProductInput) bool {
	if strings.Contains(pd.Category, "Clothing") {
		return true
	}
	return false
}

func IsFemale(pd ColtortiProductInput) bool {
	if strings.Contains(pd.Category, "Men") {
		return false
	}
	return true
}

func Classifier(pd ColtortiProductInput) string {
	subCats := strings.Split(pd.Category, ">")
	lastCat := subCats[len(subCats)-1]
	return lastCat
}
