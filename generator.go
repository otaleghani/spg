/******************************************************************

Author      Oliviero Taleghani
Date        2024-05-14

Description
Helper functions for creating and quering the Generator. This will be
the struct used to store all of the parsed csv data and it will be
used to generate the different random results.

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
  "strings"
  "strconv"
)

type Options struct {
  Format string
  Separator string
}

type Spg struct {
  Format string
  Separator string
}

type Generator struct {
  Data map[string][]string
  Lang string
}

// You could just add everything into a []string, so that you don't have to do anything strange with this. No parser, no Formatter. Nothing.

func CreateGenerator(lang string, format string) (Generator, error) {
  dicts := make(map[string]int)
  dicts["first_names"] = 0
  dicts["last_names"] = 0
  dicts["words"] = 0

  gen := Generator{
    Lang:strings.ToLower(lang), 
    Data: make(map[string][]string),
  }
  for name, column := range(dicts) {
    data, err := parser.ParseCsvColumnData(
      name + "/" + strings.ToLower(lang) + ".csv",
      column,
    )
    if err != nil {
      return Generator{}, err
    }
    gen.Data[name] = data
  }
  
  generalDicts := make(map[string]int)
  generalDicts["domains"] = 1

  for name, column := range(generalDicts) {
    data, err := parser.ParseCsvColumnData(
      "general/" + strings.ToLower(name) + ".csv",
      column,
    )
    if err != nil {
      return Generator{}, err
    }
    gen.Data[name] = data
  }

  return gen, nil
}

func (gen Generator) getRand(query string) string {
  rand.Seed(time.Now().UnixNano())
  if len(gen.Data[query]) > 0 {
    randomWord := gen.Data[query][rand.Intn(len(gen.Data[query]))]
    return randomWord
  } else {
    return ""
  }
}

func (gen Generator) FirstName(opt Options) string {
  result := gen.getRand("first_names")
  formatter.Switch(&result, opt.Format)
  return result
}

func (gen Generator) LastName(opt Options) string {
  result := gen.getRand("last_names")
  formatter.Switch(&result, opt.Format)
  return result
}

func (gen Generator) FullName(opt Options) string {
  var array []string
  array = append(array, gen.FirstName(opt), gen.LastName(opt))
  result := formatter.Separator(array, opt.Separator)
  return result
}

func (gen Generator) Word(opt Options) string {
  result := gen.getRand("words")
  formatter.Switch(&result, opt.Format)
  return result
}

func (gen Generator) Domain(opt Options) string {
  result := gen.getRand("domains")
  formatter.Switch(&result, opt.Format)
  return result
}

func (gen Generator) Email(opt Options) string {
  username := gen.FirstName(opt)
  provider := gen.Word(opt)
  domain := gen.Domain(opt)
  email := username + "@" + provider + domain
  formatter.Switch(&email, opt.Format)
  return email
}

func (gen Generator) CreditCard(opt Options) string {
  var array []string
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 4; i++ {
    currentString := ""
    for j := 0; j < 4; j++ {
      currentString = currentString + strconv.Itoa(rand.Intn(9))
    }
    array = append(array, currentString)
  }
  return strings.Join(array, opt.Separator)
}



