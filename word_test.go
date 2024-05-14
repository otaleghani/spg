package spg

import (
	"testing"
  "fmt"
  "github.com/otaleghani/spg/internal/parser"
)

func Test_General(t *testing.T) {
  firstName, err := SingleFirstName("en", "camel")
  lastName, err := SingleLastName("en", "camel")
  fullName, err := SingleFullName("en", "camel", " ")
  domainName, err := SingleDomainName("camel")

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
    _, err := SingleFirstName("en", "camel")
    if err != nil {
      t.Fatal(err)
    }
  }
}

func Test_5000c(t *testing.T) {
  array := []string{}
  names, err := FullNames("en", "camel", "-")
  if err != nil {
    t.Fatal(err)
  }
  for i := 0; i < 50000; i++ {
    array = append(array, names.Get())
  }
  fmt.Println(array)
}
