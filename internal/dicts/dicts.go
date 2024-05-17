package dicts

import (
	"fmt"
)

var dicts = map[string]map[string][]string{
	"en-usa": {

		// People
		"titles_male":        {"Mr.", "Doc."},
		"titles_female":      {"Mrs.", "Doc."},
		"gender":             {"Male", "Female", "Other"},
		"first_names_female": en_females_first_names,
		"first_names_male":   en_males_first_names,
		"last_names":         en_last_names,

		// Places locale
		"locale_cities":    en_usa_cities,
		"locale_streets":   en_usa_streets,
		"locale_states":    en_usa_states,
		"locale_countries": []string{"United States"},

		// Places
		"cities":    en_cities,
		"streets":   en_streets,
		"states":    en_states,
		"countries": en_countries,

		// Products
		"products":             en_products,
		"brands":               en_brands,
		"technologies":         en_technologies,
		"product_names":        en_product_names,
		"product_descriptions": en_product_descriptions,
	},
}

func GetDict(lang, source string) ([]string, error) {
	if dict, ok := dicts[lang]; ok {
		if values, ok := dict[source]; ok {
			return values, nil
		}
		return nil, fmt.Errorf("CANNOT FIND SOURCE %v", source)
	}
	return nil, fmt.Errorf("CANNOT FIND LANGUAGE %v", lang)
}
