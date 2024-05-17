package spg

import (
	"fmt"
	"testing"
)

func Test_Numbers(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v\n",
			i,
			g.randomNumber(10),
			g.StringNumber(10),
			g.StringNumberFixed(10),
		)
	}
}
