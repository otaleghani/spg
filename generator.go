package spg

import (
	"github.com/otaleghani/spg/internal/formatter"
	"math/rand"
  "crypto/rand"
	"time"
)

type Generator struct {
	Rng  *rand.Rand // #nosec G404: Intentianally using math/rand
	Lang string
}

type Options struct {
	Format    string
	Separator string
}

func New(lang string) (g Generator) {
	g = Generator{
		Rng:  rand.New(rand.NewSource(time.Now().Unix())),
		Lang: lang,
	}
	return
}

func (g Generator) randomString(data []string, format string) string {
	if len(data) > 0 {
    randomWord := data[rand.Intn(len(data))] // #nosec G404: Intentianally using math/rand
		formatter.FormatString(&randomWord, format)
		return randomWord
	}
	return ""
}
