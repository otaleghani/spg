package spg

import (
	"testing"
  "fmt"
  "github.com/otaleghani/spg/internal/parser"
)

func Test_ClearCache(t *testing.T) {
  err := parser.DeleteCache()
  if err != nil {
    t.Fatal(err)
  }
}

func Test_Generator(t *testing.T) {
  gen, err := CreateGenerator("en", "camel")
  if err != nil {
    t.Fatal(err)
  }
  
  for i := 0; i < 50000; i++ {
    fmt.Println(gen.FirstName(Options{}) + " " + gen.LastName(Options{}))
  }
}
