package spg

import (
  "github.com/otaleghani/spg/internal/dicts"
  // "github.com/otaleghani/spg/internal/formatter"
  "fmt"
)

type Place struct {
  Gen Generator
}

func (g Generator) Place() (Place) {
  return Place{Gen: g}
}

func (p Place) City(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"cities")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Place) CityLocale(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"locale_cities")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Place) Street(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"streets")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Place) StreetLocale(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"locale_streets")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Place) State(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"states")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Place) StateLocale(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"locale_states")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Place) Country(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"countries")
  return p.Gen.randomString(dict, opt.Format)
}
func (p Place) CountryLocale(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"locale_countries")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Place) AddressEn(opt Options) string {
  result := fmt.Sprintf(
    "%v %v\n%v, %v %v\n%v",
    p.Gen.StringNumber(250),
    p.Street(opt),
    p.City(opt),
    p.State(opt),
    p.Gen.StringNumberFixed(5),
    p.Country(opt),
  )
  return result
}

func (p Place) AddressEnLocale(opt Options) string {
  result := fmt.Sprintf(
    "%v %v\n%v, %v %v\n%v",
    p.Gen.StringNumber(250),
    p.StreetLocale(opt),
    p.CityLocale(opt),
    p.StateLocale(opt),
    p.Gen.StringNumberFixed(5),
    p.CountryLocale(opt),
  )
  return result
}
