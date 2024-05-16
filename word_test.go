package spg

import (
	"testing"
  "fmt"
)

var num_tests = 50
var g = New("en-usa")
var opt = Options{Format: "camel", Separator: "-"}

func Test_Names(t *testing.T) {
  for i := 0; i < num_tests; i++ {
    fmt.Printf(
      "Result: %v | %v | %v | %v \n",
      g.Person().FirstName(opt), 
      g.Person().FullName(opt), 
      g.Place().AddressEn(opt), 
      g.Place().AddressEnLocale(opt), 
    )
  }
}

func Test_Things(t *testing.T) {
  for i := 0; i < num_tests; i++ {
    fmt.Printf(
      "Result: %v\n%v\n\n",
      g.Product().CompleteProductName(opt), 
      g.Product().ProductDescription(opt), 
    )
  }
}
