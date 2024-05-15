package spg

import (
	"testing"
  "fmt"
)

// func Test_NewGenerator(t *testing.T) {
//   g := New("en")
//   opt := Options{Format: "camel", Separator: "-"}
//   for i := 0; i < 5000; i++ {
//     fmt.Printf(
//       "Result: %v | %v | %v | %v | %v\n", 
//       g.Person().MaleFirstName(opt), 
//       g.Person().FemaleFirstName(opt),
//       g.Person().LastName(opt),
//       g.Person().MaleFullName(opt),
//       g.Person().FemaleFullName(opt),
//     )
//   }
// }

func Test_FirstName(t *testing.T) {
  g := New("en")
  opt := Options{Format: "camel", Separator: "-"}
  for i := 0; i < 50; i++ {
    fmt.Printf(
      "Result: %v | %v \n",
      g.Person().FirstName(opt), 
      g.Person().FullName(opt), 
    )
  }
}
