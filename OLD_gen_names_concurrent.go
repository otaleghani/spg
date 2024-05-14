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

// import (
//   // "github.com/otaleghani/spg/internal/parser"
//   "github.com/otaleghani/spg/internal/formatter"
//   "math/rand"
//   "time"
// )

// type FullNameGenerator struct {
//   Options GeneratorMetadata
//   Data_FirstNames []string
//   Data_LastNames []string
// }
// func FullNames(lang string, format string, separator string) (FullNameGenerator, error) {
//   firstNames, err := FirstNames(lang, format)
//   if err != nil {
//     return FullNameGenerator{}, err
//   }
//   lastNames, err := LastNames(lang, format)
//   if err != nil {
//     return FullNameGenerator{}, err
//   }
//   lowerFormat, err := formatter.FormatterCheck(format)
//   if err != nil {
//     return FullNameGenerator{}, err
//   }
//   result := FullNameGenerator{
//     Options: GeneratorMetadata{
//       Lang: lang,
//       Format: lowerFormat,
//       Separator: separator,
//     },
//     Data_FirstNames: firstNames.Data,
//     Data_LastNames: lastNames.Data,
//   }
//   return result, nil
// }
// 
// func (gen FullNameGenerator) Get() string {
//   var array []string
//   rand.Seed(time.Now().UnixNano())
//   if len(gen.Data_FirstNames) > 0 {
//     randomFN := gen.Data_FirstNames[rand.Intn(len(gen.Data_FirstNames))]
//     formatter.Switch(&randomFN, gen.Options.Format)
//     array = append(array, randomFN)
//   } else {
//     return ""
//   }
//   if len(gen.Data_LastNames) > 0 {
//     randomLN := gen.Data_LastNames[rand.Intn(len(gen.Data_LastNames))]
//     formatter.Switch(&randomLN, gen.Options.Format)
//     array = append(array, randomLN)
//   } else {
//     return ""
//   }
//   randomName := formatter.Divider(array, gen.Options.Separator)
//   return randomName
// }
