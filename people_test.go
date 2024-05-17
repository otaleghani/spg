package spg

import (
	"fmt"
	"testing"
)

func Test_People(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v, %v, %v, %v, %v\n",
			i,
			g.Person().FirstName(opt),
			g.Person().MaleFirstName(opt),
			g.Person().FemaleFirstName(opt),
			g.Person().LastName(opt),
			g.Person().FullName(opt),
			g.Person().MaleFullName(opt),
			g.Person().FemaleFullName(opt),
		)
	}
}
