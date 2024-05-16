package spg

import (
	"testing"
  "fmt"
)

func Test_FirstName(t *testing.T) {
  g := New("en-usa")
  opt := Options{Format: "camel", Separator: "-"}
  for i := 0; i < 50; i++ {
    fmt.Printf(
      "Result: %v | %v | %v | %v \n",
      g.Person().FirstName(opt), 
      g.Person().FullName(opt), 
      g.Place().AddressEn(opt), 
      g.Place().AddressEnLocale(opt), 
    )
  }
}
