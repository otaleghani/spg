package spg

import (
	"testing"
  "fmt"
  // "github.com/otaleghani/spg/internal/parser"
)

func Test_Robe(t *testing.T) {
  firstName, err := FirstName("en", "camel")
  lastName, err := LastName("en", "camel")
  fullName, err := FullName("en", "camel", " ")
  domainName, err := DomainName("camel")

  fmt.Println(firstName)
  fmt.Println(lastName)
  fmt.Println(fullName)
  fmt.Println(domainName)

  // err = parser.DeleteCache()

  if err != nil {
    t.Fatal(err)
  }
}
