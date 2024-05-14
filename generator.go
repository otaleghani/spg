/******************************************************************

Author      Oliviero Taleghani
Date        2024-05-14

Description
Helper functions for creating and quering generators.

Usage

Dependency
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"
  "math/rand"
  "time"

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

package spg

import (
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"
  "math/rand"
  "time"
)

type GeneratorMetadata struct {
  Lang string
  Format string
  Separator string
}

type Generator struct {
  Options GeneratorMetadata
  Data []string
}

func NewGenerator(
  lang string, 
  format string, 
  source string,
  column int,
) (Generator, error) {
  data, err := parser.ParseCsvColumnData(
    source + "/" + lang + ".csv",
    column,
  )
  if err != nil {
    return Generator{}, err
  }
  lowerFormat, err := formatter.FormatterCheck(format)
  if err != nil {
    return Generator{}, err
  }
  result := Generator{
    Options: GeneratorMetadata{
      Lang: lang,
      Format: lowerFormat,
    },
    Data: data,
  }
  return result, nil
}

func (generator Generator) Get() string {
  rand.Seed(time.Now().UnixNano())
  if len(generator.Data) > 0 {
    randomWord := generator.Data[rand.Intn(len(generator.Data))]
    formatter.Switch(&randomWord, generator.Options.Format)
    return randomWord
  } else {
    return ""
  }
}
