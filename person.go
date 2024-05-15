// TO DO:
// Prefix (Sir, Mr., Mrs...), Suffix (the third, VI...), Middlename

package spg

import (
  "github.com/otaleghani/spg/internal/dicts"
  "github.com/otaleghani/spg/internal/formatter"
)

type Person struct {
  Gen Generator
}

func (g Generator) Person() (Person) {
  return Person{Gen: g}
}

func (p Person) FirstName(opt Options) string {
  var dict []string
  rn := p.Gen.randomNumber(2)
  if rn == 1 {
    dict, _ = dicts.GetDict(p.Gen.Lang,"first_names_male")
  } else {
    dict, _ = dicts.GetDict(p.Gen.Lang,"first_names_female")
  }
  return p.Gen.randomString(dict, opt.Format)
}

func (p Person) MaleFirstName(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"first_names_male")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Person) FemaleFirstName(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"first_names_female")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Person) LastName(opt Options) string {
  dict, _ := dicts.GetDict(p.Gen.Lang,"last_names")
  return p.Gen.randomString(dict, opt.Format)
}

func (p Person) FullName(opt Options) string {
  first_name, last_name := "", ""
  rn := p.Gen.randomNumber(2)
  if rn == 1 {
    first_name = p.MaleFirstName(opt)
    last_name = p.LastName(opt)
  } else {
    first_name = p.FemaleFirstName(opt)
    last_name = p.LastName(opt)
  }

  if opt.Separator == "" {
    return first_name + " " + last_name
  }
  return formatter.Separator([]string{first_name, last_name}, opt.Separator)
}

func (p Person) MaleFullName(opt Options) string {
  first_name := p.MaleFirstName(opt)
  last_name := p.LastName(opt)

  if opt.Separator == "" {
    return first_name + " " + last_name
  }
  return formatter.Separator([]string{first_name, last_name}, opt.Separator)
}

func (p Person) FemaleFullName(opt Options) string {
  first_name := p.FemaleFirstName(opt)
  last_name := p.LastName(opt)

  if opt.Separator == "" {
    return first_name + " " + last_name
  }
  return formatter.Separator([]string{first_name, last_name}, opt.Separator)
}

