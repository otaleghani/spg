package parser

import (
  "os"
  "encoding/csv"
  "math/rand"
  "time"
)


func ParseCvs(dict string) (string, error) {
  file, err := os.Open(dict)
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
