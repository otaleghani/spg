package spg

import (
  "github.com/otaleghani/spg/internal/dicts"
  "fmt"
  "strings"
)

type Product struct {
  Gen Generator
}

func (g Generator) Product() (Product) {
  return Product{Gen: g}
}

func (p Product) Thing(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"products")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Product) Brand(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"brands")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Product) Technology(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"technologies")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Product) ProductName(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"product_names")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Product) CompleteProductName(opt Options) string {
  return fmt.Sprintf(
    "%v %v %v with %v",
    p.Thing(opt),
    p.Brand(opt),
    p.ProductName(opt),
    p.Technology(opt),
  )
}

func (p Product) ProductDescription(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"product_descriptions")
  description := p.Gen.randomString(dict, opt.Format)
  return strings.Replace(
    description,
    "{{technology}}",
    p.Technology(opt),
    -1,
  )
}
