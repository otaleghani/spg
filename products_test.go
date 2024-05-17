package spg

import (
	"fmt"
	"testing"
)

func Test_Products(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v, %v, %v, %v, %v\n",
			i,
			g.Product().Thing(opt),
			g.Product().Brand(opt),
			g.Product().Technology(opt),
			g.Product().ProductName(opt),
			g.Product().CompleteProductName(opt),
			g.Product().ProductDescription(opt),
			g.Product().UUID(opt),
		)
	}
}
