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
//
package spg
//
//import (
//  "github.com/otaleghani/spg/internal/parser"
//  "github.com/otaleghani/spg/internal/formatter"
//  "math/rand"
//  "time"
//)
//
// type GeneratorMetadata struct {
//   Lang string
//   Format string
//   Separator string
// }
// 
// type Generator struct {
//   Options GeneratorMetadata
//   Data []string
// }
// 
// type Options struct {
//   Lang string
//   Format string
//   Separator string
// }
// 
// type Generator struct {
//   Data map[string][]string
// }
// 
// // gen := spg.GetGenerator()
// // options := spg.Options{Lang: "en", Format: "random", Separator: "-"}
// // gen.Email(options)
// 
// // 1. Download all of the dict in based on the given language
// 
// func NewGenerator(
//   lang string, 
//   format string, 
//   source string,
//   column int,
// ) (Generator, error) {
//   data, err := parser.ParseCsvColumnData(
//     source + "/" + lang + ".csv",
//     column,
//   )
//   if err != nil {
//     return Generator{}, err
//   }
//   lowerFormat, err := formatter.FormatterCheck(format)
//   if err != nil {
//     return Generator{}, err
//   }
//   result := Generator{
//     Options: GeneratorMetadata{
//       Lang: lang,
//       Format: lowerFormat,
//     },
//     Data: data,
//   }
//   return result, nil
// }
// 
// func (gen Generator) get(query string) string {
//   rand.Seed(time.Now().UnixNano())
//   if len(generator.Data) > 0 {
//     randomWord := generator.Data[rand.Intn(len(generator.Data))]
//     formatter.Switch(&randomWord, generator.Options.Format)
//     return randomWord
//   } else {
//     return ""
//   }
// }
// 
// func (gen Generator) Email() string {
// // You know that you will need to get the email! Simple.
//   firstName := gen.FirstName()
//   domain := gen.Domain()
// }
// 
// func (gen Generator) FirstName() string {
//   gen.get("FirstName")
// 
//   gen.Data[FirstName] --> // This is the one that you'll query. 
// }
// 
// USED TO GET THE DATA 
// THOSE FUNCTIONS WILL BE USED TO GET FROM THE GENERATOR THE DATA.
//
// func FirstNames(lang string, format string) (Generator, error) {
//   gen, err := NewGenerator(lang, format, "names", 0)
//   if err != nil {
//     return Generator{}, err
//   }
//   return gen, nil
// }
// func LastNames(lang string, format string) (Generator, error) {
//   gen, err := NewGenerator(lang, format, "surnames", 0)
//   if err != nil {
//     return Generator{}, err
//   }
//   return gen, nil
// }
// func TopLevelDomain() (Generator, error) {
//   gen, err := NewGenerator("country", "lower", "domains", 0)
//   if err != nil {
//     return Generator{}, err
//   }
//   return gen, nil
// }
// func Words(lang string, format string) (Generator, error) {
//   gen, err := NewGenerator(lang, format, "words", 0)
//   if err != nil {
//     return Generator{}, err
//   }
//   return gen, nil
// }
