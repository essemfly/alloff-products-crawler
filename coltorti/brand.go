package coltorti

import "log"

func GetKoreanBrandName(brandName string) string {
	brandMappers := map[string]string{
		"A.P.C.":               "아페쎄",
		"ALEXANDER MCQUEEN":    "알렉산더 맥퀸",
		"AMIRI":                "아미리",
		"BALMAIN":              "발망",
		"BURBERRY":             "버버리",
		"VETEMENTS":            "베트멍",
		"MIU MIU":              "미우미우",
		"MM6 MAISON MARGIELA":  "mm6 마르지엘라",
		"FENDI":                "펜디",
		"GANNI":                "가니",
		"GCDS X WOLFORD":       "GCDS X WOLFORD",
		"GIVENCHY":             "지방시",
		"ISABEL MARANT":        "이자벨 마랑",
		"ISABEL MARANT ETOILE": "이자벨 마랑 에뚜왈",
		"JIMMY CHOO":           "지미추",
		"KENZO":                "켄조",
		"LOEWE":                "로에베",
		"MASON KITUNE":         "메종 키츠네",
		"MARC JACOBS":          "마크 제이콥스",
		"MARNI":                "마르니",
		"MAX MARA":             "막스마라",
		"MAX MARA STUDIO":      "막스마라 스튜디오",
		"MAX MARA THE CUBE":    "막스마라 더 큐브",
		"DOLCE & GABBANA":      "돌체 앤 가바나",
		"MSGM":                 "엠에스지엠",
		"N. 21":                "넘버투에니원",
		"OFF WHITE":            "오프화이트",
		"PALM ANGELS":          "팜 엔젤스",
		"STELLA MCCCARTNEY":    "스텔라 매카트니",
		"VALENTINO":            "발렌티노",
		"LANVIN":               "랑방",
		"BALENCIAGA":           "발렌시아가",
		"ACNE STUDIOS":         "아크네스튜디오",
		"N.21":                 "넘버투에니원",
		"MULBERRY":             "멀버리",
		"POLO RALPH LAUREN":    "폴로랄프로렌",
		"RICK OWENS":           "릭오웬스",
		"RAF SIMONS":           "라프시몬스",
		"SWAROVSKI":            "스와로브스키",
		"TORY BURCH":           "토리버치",
		"THOM BROWNE":          "톰브라운",
		"VANS":                 "반스",
		"MCM":                  "엠씨엠",
		"ALEXANDER WANG":       "알렉산더왕",
		"CHOLE":                "끌로에",
		"DR. MARTENS":          "닥터마틴",
		"GOLDEN GOOSE":         "골든구스",
		"MONCLER":              "몽클레어",
		"MOOSE KNUCKLES":       "무스너클",
		"PRADA":                "프라다",
	}

	if val, ok := brandMappers[brandName]; ok {
		return val
	}

	log.Println("no classified cats", brandName)
	return brandName
}
