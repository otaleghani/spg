/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Part of the formatter package, contains helper functions to format
either strings or []strings.

Usage
  func Divider(words []string, div string) string
  | Joins a []string into a string suing the given divider

  func Switch(word *string, format string) error
  | Formats word in either lower, upper, camel or title case

Dependency
  "strings"
  "errors"
  "unicode"

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

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
