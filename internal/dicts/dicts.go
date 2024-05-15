package dicts

import (
  "errors"
  "fmt"
)

// Animals
// 

var dicts = map[string]map[string][]string{
  "en-usa": {
    "titles_male": {"Mr.", "Doc."},
    "titles_female": {"Mrs.", "Doc."},
    "gender": {"Male", "Female", "Other"},
    "first_names_female": en_females_first_names,
    "first_names_male": en_males_first_names,
    "last_names": en_last_names,
    "places": {},
  },
  "it": {},
  "general": {}
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
