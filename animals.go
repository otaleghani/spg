package spg

import (
	"github.com/otaleghani/spg/internal/dicts"
	// "github.com/otaleghani/spg/internal/formatter"
)

type Animal struct {
	Gen Generator
}

func (g Generator) Animal() Animal {
	return Animal{Gen: g}
}

func (a Animal) Mammal(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "mammals")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) Bird(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "birds")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) Fish(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "fishes")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) Insect(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "insects")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) Pet(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "pets")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) PetName(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "pets_names")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) DogBreed(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "dog_breeds")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) CatBreed(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "cat_breeds")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) Sex(opt Options) string {
	dict, _ := dicts.GetDict(a.Gen.Lang, "gender")
	return a.Gen.randomString(dict, opt.Format)
}

func (a Animal) MicrochipNumber(opt Options) string {
	return a.Gen.StringNumberFixed(15)
}
