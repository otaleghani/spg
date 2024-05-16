package dicts

import (
  "errors"
  "fmt"
)

var dicts = map[string]map[string][]string{
  "en-usa": {

    // People
    "titles_male": {"Mr.", "Doc."},
    "titles_female": {"Mrs.", "Doc."},
    "gender": {"Male", "Female", "Other"},
    "first_names_female": en_females_first_names,
    "first_names_male": en_males_first_names,
    "last_names": en_last_names,

    // Places 
    "locale_cities": en_usa_cities,
    //"": en_usa_
    "locale_streets": en_usa_streets,
    "locale_states": en_usa_states,
    "locale_countries": []string{"United States"},

    "cities": en_cities,
    "streets": en_streets,
    "states": en_states,
    "countries": en_countries,
  },
}

func GetDict(lang, source string) ([]string, error) {
  if dict, ok := dicts[lang]; ok {
    if values, ok := dict[source]; ok {
      return values, nil
    }
    return nil, errors.New(fmt.Sprint("Cannot find source %v\n", source))
  }
  return nil, errors.New(fmt.Sprint("Cannot find language %v\n", lang))
}
