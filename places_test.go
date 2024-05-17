package spg

import (
	"fmt"
	"testing"
)

func Test_Places(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v, %v, %v, %v, %v, %v,\n\n%v,\n\n%v\n\n\n",
			i,
			g.Place().City(opt),
			g.Place().CityLocale(opt),
			g.Place().Street(opt),
			g.Place().StreetLocale(opt),
			g.Place().State(opt),
			g.Place().StateLocale(opt),
			g.Place().Country(opt),
			g.Place().CountryLocale(opt),
			g.Place().AddressEn(opt),
			g.Place().AddressEnLocale(opt),
		)
	}
}
