package spg

import (
  "math/rand"
  "time"
)

type Generator struct {
  Rng *Rand
  Lang string
}

type Options struct {
  Format string
  Separator string
}

type YouTube struct {
  g Generator
}

func New(lang string) (g Generator) {
  g = Generator{
    Rng: rand.New(rand.NewSource(time.Now().Unix())),
    Lang: lang,
  }
  return
}


g := spg.New("it")

options := Options{
  Separator: " ",
  Format: "camel",
}

gender := g.Person().Gender()
name := g.Person().Name()
name := g.Person().NameTaited(gender, options)

// So every single file will have
// Dictionaries?
// 
