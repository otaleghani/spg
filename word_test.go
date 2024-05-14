package spg

import (
	"testing"
  "fmt"
  "github.com/otaleghani/spg/internal/parser"
)

func Test_General(t *testing.T) {
  firstName, err := FirstName("en", "camel")
  lastName, err := LastName("en", "camel")
  fullName, err := FullName("en", "camel", " ")
  domainName, err := DomainName("camel")

  fmt.Println(firstName)
  fmt.Println(lastName)
  fmt.Println(fullName)
  fmt.Println(domainName)

  err = parser.DeleteCache()

  if err != nil {
    t.Fatal(err)
  }
}

func Test_5000(t *testing.T) {
  for i := 0; i < 50000; i++ {
    _, err := FirstName("en", "camel")
    if err != nil {
      t.Fatal(err)
    }
  }
}

func Test_5000c(t *testing.T) {
  names, err := NamesGenerator()
  if err != nil {
    t.Fatal(err)
  }
  for i := 0; i < 50000; i++ {
    _, err := names.GetRand()
    if err != nil {
      t.Fatal(err)
    }

  }
}
