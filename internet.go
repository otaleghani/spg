package spg

import (
	"github.com/otaleghani/spg/internal/dicts"
	"github.com/otaleghani/spg/internal/formatter"
	"strings"
)

type Internet struct {
	Gen Generator
}

func (g Generator) Internet() Internet {
	return Internet{Gen: g}
}

func (i Internet) Username(opt Options) string {
	dice := i.Gen.randomNumber(2)
	if dice == 1 {
		return i.CasualUsername(opt)
	}
	return i.BusinessUsername(opt)
}

func (i Internet) CasualUsername(opt Options) string {
	dice := i.Gen.randomNumber(9)
	result := ""

	switch dice {
	case 0:
		result = i.Gen.Person().FirstName(opt) + i.Gen.StringNumber(100)
	case 1:
		result = i.Gen.Person().FirstName(opt) + "." + i.Gen.Animal().PetName(opt)
	case 2:
		result = i.Gen.Person().FirstName(opt) + "_" + i.Gen.Animal().PetName(opt)
	case 3:
		result = i.Gen.Person().FirstName(opt) + "-" + i.Gen.Animal().PetName(opt)
	case 4:
		result = i.Gen.Person().LastName(opt) + i.Gen.StringNumber(100)
	case 5:
		result = i.Gen.Animal().PetName(opt)
	case 6:
		result = i.Gen.Animal().PetName(opt) + i.Gen.StringNumber(100)
	case 7:
		result = "xx_" + i.Gen.Animal().PetName(opt) + "_xx"
	}
	return result
}

func (i Internet) BusinessUsername(opt Options) string {
	first_name := i.Gen.Person().FirstName(opt)
	last_name := i.Gen.Person().LastName(opt)
	dice := i.Gen.randomNumber(6)
	result := ""

	switch dice {
	case 0:
		result = string(first_name[0]) + "." + last_name
	case 1:
		result = string(first_name[0]) + last_name
	case 2:
		result = first_name + last_name
	case 3:
		result = first_name + "." + last_name
	case 4:
		result = first_name + "-" + last_name
	case 5:
		result = first_name + "_" + last_name
	}
	return result
}

func (i Internet) TLD(opt Options) string {
	dice := i.Gen.randomNumber(2)
	if dice == 0 {
		return i.CCTLD(opt)
	}
	return i.GTLD(opt)
}

func (i Internet) TLDLocale(opt Options) string {
	dict, _ := dicts.GetDict(i.Gen.Lang, "tld")
	return i.Gen.randomString(dict, opt.Format)
}

func (i Internet) CCTLD(opt Options) string {
	dict, _ := dicts.GetDict("general", "ccTLD")
	return i.Gen.randomString(dict, opt.Format)
}

func (i Internet) GTLD(opt Options) string {
	dict, _ := dicts.GetDict("general", "gTLD")
	return i.Gen.randomString(dict, opt.Format)
}

func (i Internet) DomainName(opt Options) string {
	name := i.Gen.rollDomainName()
	tld := i.TLD(opt)
	return name + tld
}

func (i Internet) EmailProvider(opt Options) string {
	dict, _ := dicts.GetDict(i.Gen.Lang, "email_providers")
	return i.Gen.randomString(dict, opt.Format)
}

func (i Internet) DomainNameLocale(opt Options) string {
	name := i.Gen.rollDomainName()
	tld := i.TLDLocale(opt)
	return name + tld
}

func (g Generator) rollDomainName() string {
	dice := g.randomNumber(9)
	result := ""
	switch dice {
	case 0:
		result = g.Person().FirstName(Options{})
	case 1:
		result = g.Person().LastName(Options{})
	case 2:
		result = g.Place().City(Options{})
	case 3:
		result = g.Animal().PetName(Options{})
	case 4:
		result = g.Animal().Pet(Options{})
	case 5:
		result = g.Product().Thing(Options{})
	case 6:
		result = g.Product().Brand(Options{})
	case 7:
		result = g.Product().Technology(Options{})
	case 8:
		result = g.Product().ProductName(Options{})
	}
	formatter.FormatString(&result, "lower")
	return strings.Replace(result, " ", "", -1)
}
