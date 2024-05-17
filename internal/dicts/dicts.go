package dicts

import (
	"fmt"
)

var dicts = map[string]map[string][]string{
	"en-usa": {

		// People
		"titles_male":        {"Mr.", "Doc."},
		"titles_female":      {"Mrs.", "Doc."},
		"gender":             {"Male", "Female"},
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

		// Animals
		"mammals":    en_animals_mammals,
		"birds":      en_animals_birds,
		"fishes":     en_animals_fishes,
		"insects":    en_animals_insects,
		"pets":       en_pets,
		"pets_names": en_pets_names,
		"dog_breeds": en_dog_breeds,
		"cat_breeds": en_cat_breeds,

    // Internet
    "tld": en_tld,
    "email_providers": en_email_providers,
	},
  "general": {
    "ccTLD": cctld,
    "gTLD": gtld,
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
