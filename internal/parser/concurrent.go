/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-14

Description
Part of the parser package, used to parse csv in a concurrent manner.

Usage

Dependency

Todo

Changelog
  [0.0.1]   2024-05-14
  Added     Initial release

********************************************************************/

package parser

import (
  "os"
  "encoding/csv"
) 

func GetColumnData(dict string, column int) ([]string, error) {
  _, err := os.Stat(cachePath() + dict)
  if os.IsNotExist(err) {
    err = initDict(dict)
    if err != nil {
      return []string{}, err
    }
  }
  file, err := os.Open(cachePath() + dict)
  if err != nil {
    return []string{}, err
  }
  defer file.Close()
  reader := csv.NewReader(file)
  records, err := reader.ReadAll()
  if err != nil {
    return []string{}, err
  }
  var words []string
  for _, record := range records {
    words = append(words, record[column])
  }
  return words, nil
}
