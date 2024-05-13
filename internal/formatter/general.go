package formatter

import (
  "strings"
  "errors"
  "unicode"
)

func Divider(words []string, div string) string {
  return strings.Join(words, div)
}

func Switch(word *string, format string) error {
  switch format {
  case "lower":
    *word = strings.ToLower(*word) 
    return nil
  case "upper":
    *word = strings.ToUpper(*word) 
    return nil
  case "title":
    *word = strings.ToTitle(*word) 
    return nil
  case "camel":
    *word = strings.ToLower(*word) 
    var result strings.Builder
    
    for i, v := range *word {
      if i == 0 {
        result.WriteRune(unicode.ToUpper(v))
      } else {
        result.WriteRune(v)
      }
    }

    *word = result.String()
    return nil
  default: 
    return errors.New("No formatter found")
  }
}
