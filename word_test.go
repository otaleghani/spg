package spg

import (
	"testing"
  "fmt"
)

func Test_Robe(t *testing.T) {
  firstName, err := FirstName("en", "camel")
  lastName, err := LastName("en", "camel")
  fullName, err := FullName("en", "camel", " ")

  fmt.Println(firstName)
  fmt.Println(lastName)
  fmt.Println(fullName)
  if err != nil {
    t.Fatal(err)
  }
}
