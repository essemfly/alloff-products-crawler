package coltorti

func ScreenBrands(brandName string) bool {
	brandMappers := map[string]string{
		"ACNE STUDIOS":      "아크네스튜디오",
		"N. 21":             "넘버투에니원",
		"N.21":              "넘버투에니원",
		"MULBERRY":          "멀버리",
		"POLO RALPH LAUREN": "폴로랄프로렌",
		"RICK OWENS":        "릭오웬스",
		"RAF SIMONS":        "라프시몬스",
		"SWAROVSKI":         "스와로브스키",
		"TORY BURCH":        "토리버치",
		"THOM BROWNE":       "톰브라운",
		"VANS":              "반스",
	}
	if _, ok := brandMappers[brandName]; ok {
		return true
	}
	return false
}
