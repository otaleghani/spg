package spg

import (
	"fmt"
	"testing"
)

func Test_Internet(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n",
			i,
			g.Internet().TLD(opt),
			g.Internet().TLDLocale(opt),
			g.Internet().CCTLD(opt),
			g.Internet().GTLD(opt),
			g.Internet().EmailProvider(Options{}),
			g.Internet().DomainName(Options{}),
			g.Internet().DomainNameLocale(Options{}),
			g.Internet().BusinessUsername(Options{Format: "lower"}),
			g.Internet().CasualUsername(Options{Format: "lower"}),
			g.Internet().Username(Options{Format: "lower"}),
			g.Internet().Email(Options{Format: "lower"}),
			g.Internet().CasualEmail(Options{Format: "lower"}),
			g.Internet().BusinessEmail(Options{Format: "lower"}),
		)
	}
}
