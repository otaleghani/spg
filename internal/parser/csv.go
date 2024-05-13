/********************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Part of the parser package, used to parse csv files. 

Usage
  func ParseCvs(dict string) (string, error)
  | Parses a result given a dict. If the dict is not cached, it
  | downloads it

Dependency
  "os"
  "encoding/csv"
  "math/rand"
  "time"

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

package parser

import (
  "os"
  "encoding/csv"
  "math/rand"
  "time"
)

func ParseCvs(dict string) (string, error) {
  _, err := os.Stat(cachePath() + dict)
  if os.IsNotExist(err) {
    err = initDict(dict)
    if err != nil {
      return "", err
    }
  }
  file, err := os.Open(cachePath() + dict)
  if err != nil {
    return "", err
  }
  defer file.Close()
  reader := csv.NewReader(file)
  records, err := reader.ReadAll()
  if err != nil {
    return "", err
  }
  var words []string
  for _, record := range records {
    words = append(words, record[0])
  }
  rand.Seed(time.Now().UnixNano())
  if len(words) > 0 {
    randomWord := words[rand.Intn(len(words))]
    return randomWord, nil
  } else {
    return "", nil
  }
}
