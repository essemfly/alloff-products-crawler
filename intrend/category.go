package intrend

import (
	"log"
	"strings"
)

func GetNaverCategoryCode(category string) string {
	categoryMappers := map[string]string{
		"/abbigliamento/abiti":                    "50000805",
		"/abbigliamento/tute":                     "50000811",
		"/abbigliamento/maglie-e-cardigan":        "50000806",
		"/abbigliamento/top-e-t-shirt":            "50000803",
		"/abbigliamento/camicie-e-bluse":          "50000804",
		"/abbigliamento/pantaloni":                "50000810",
		"/abbigliamento/gonne":                    "50000808",
		"/abbigliamento/felpe":                    "50000805",
		"/abbigliamento/jeans":                    "50000809",
		"/abbigliamento/tailleur":                 "50000813",
		"/abbigliamento/leggings":                 "50000812",
		"/abbigliamento/costumi-da-bagno":         "50005462",
		"/cappotti-e-giacche/cappotti":            "50000813",
		"/cappotti-e-giacche/giacconi-e-blazer":   "50000815",
		"/cappotti-e-giacche/impermeabili":        "50000813",
		"/cappotti-e-giacche/giacche-in-pelle":    "50000815",
		"/cappotti-e-giacche/pellice":             "50000813",
		"/cappotti-e-giacche/piumini-e-imbottiti": "50000814",
		"/borse-e-accessori/sciarpe-e-colli":      "50004010",
		"/borse-e-accessori/borse":                "50000639",
		"/borse-e-accessori/guanti-e-cappelli":    "50000551",
		"/borse-e-accessori/bigiotteria":          "50000434",
		"/borse-e-accessori/accessori":            "50000434",
		"/borse-e-accessori/calze":                "50003995",
		"/borse-e-accessori/cinture":              "50000539",
		"/borse-e-accessori/occhiali":             "50000556",
		"/scarpe/tutte":                           "50003840",
		"/scarpe/décolleté":                       "50003827",
		"/scarpe/sneakers":                        "50003822",
		"/scarpe/ballerine-e-mocassini":           "50003820",
		"/scarpe/sandali":                         "50003842",
		"/scarpe/stivali-e-tronchetti":            "50004191",
		"/scarpe/stringate":                       "50003842",
	}

	if val, ok := categoryMappers[category]; ok {
		return val
	}

	log.Println("no classified cats for intrend", category)
	return "50000846"
}

func IsClothing(category string) bool {
	return !strings.Contains(category, "accessori")
}
