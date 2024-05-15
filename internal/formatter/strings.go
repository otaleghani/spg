/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Part of the formatter package, contains helper functions to format
either strings or []strings.

Dependency
  "strings"
  "errors"
  "unicode"

Todo

Changelog
  [0.0.2]   2024-05-15
  Add       Refactor 

  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

package formatter

import (
  "strings"
  "unicode"
)

func Separator(words []string, div string) string {
  return strings.Join(words, div)
}

func FormatString(word *string, format string) {
  switch format {
  case "lower":
    *word = strings.ToLower(*word) 
    return 
  case "upper":
    *word = strings.ToUpper(*word) 
    return
  case "title":
    *word = strings.ToTitle(*word) 
    return
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
    return 
  default: 
    return 
  }
}
