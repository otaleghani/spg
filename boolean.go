package spg

import (
	"github.com/otaleghani/spg/internal/formatter"
	"strconv"
)

func (g Generator) Boolean() bool {
	return g.randomNumber(2) == 0
}

func (g Generator) StringBoolean(opt Options) string {
	result := strconv.FormatBool(g.Boolean())
	formatter.FormatString(&result, opt.Format)
	return result
}
