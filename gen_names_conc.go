/******************************************************************

Author      Oliviero Taleghani
Date        2024-05-14

Description
Entry point for the package. Here you will have all of the functions
needed to call concurrently the generator. 

Usage

Dependency
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"

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


type FullNameGenerator struct {
  Options GeneratorMetadata
  Data_FirstName []string
  Data_LastName []string
}

func FirstNames(lang string, format string) (Generator, error) {
  if gen, err := NewGenerator(lang, format, "names", 0); err != nil {
    return Generator{}, err
  }
  return gen, nil
}

func LastNames(lang string, format string) (Generator, error) {
  if gen, err := NewGenerator(lang, format, "surnames", 0); err != nil {
    return Generator{}, err
  }
  return gen, nil
}

func FullNames(lang string, format string, separator string) (FullNameGenerator, error) {
  firstNames, err := FirstNames(lang, format)
  if err != nil {
    return FullNameGenerator{}, err
  }
  lastNames, err := LastNames(lang, format)
  if err != nil {
    return FullNameGenerator{}, err
  }
  lowerFormat, err := formatter.FormatterCheck(format)
  if err != nil {
    return Generator{}, err
  }
  result := FullNameGenerator{
    Options: GeneratorMetadata{
      Lang: lang,
      Format: lowerFormat,
      Separator: separator,
    },
    Data_FirstName: firstNames.Data,
    Data_LastName: lastNames.Data,
  }
  return result, nil
}

func (gen FullNameGenerator) Get() string {
  var array []string
  rand.Seed(time.Now().UnixNano())
  if len(gen.FirstNames) > 0 {
    randomFN := gen.FirstNames[rand.Intn(len(gen.FirstNames))]
    array = append(array, randomFN)
  } else {
    return ""
  }
  if len(gen.LastNames) > 0 {
    randomLN := gen.LastNames[rand.Intn(len(gen.LastNames))]
    array = append(array, randomLN)
  } else {
    return ""
  }
  randomName := formatter.Divider(array, gen.Options.Divider)
  formatter.Switch(&randomName, generator.Options.format)
  return randomName
}
