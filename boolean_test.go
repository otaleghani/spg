package spg

import (
	"fmt"
	"testing"
)

func Test_Boolean(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v %v\n",
			i,
			g.Boolean(),
			g.StringBoolean(Options{Format: "lower"}),
		)
	}
}
