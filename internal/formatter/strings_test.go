package formatter

import (
	"fmt"
	"testing"
)

func Test_Formatter(t *testing.T) {
	words := []string{"some ", "words"}
	lower, upper, title, camel := "", "", "", ""

	for i := 0; i < len(words); i++ {
		lower, upper, title, camel = words[i], words[i], words[i], words[i]
		FormatString(&lower, "lower")
		FormatString(&upper, "upper")
		FormatString(&title, "title")
		FormatString(&camel, "camel")
		fmt.Printf("%v: %v, %v, %v, %v", i, lower, upper, title, camel)
	}
	fmt.Printf("%v", Separator(words, "-"))
}
