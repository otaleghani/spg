package spg

import (
	"fmt"
	"testing"
)

func Test_Animals(t *testing.T) {
	for i := 0; i < num_tests; i++ {
		fmt.Printf(
			"%v: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n",
			i,
			g.Animal().Mammal(opt),
			g.Animal().Fish(opt),
			g.Animal().Bird(opt),
			g.Animal().Insect(opt),
			g.Animal().Pet(opt),
			g.Animal().PetName(opt),
			g.Animal().DogBreed(opt),
			g.Animal().CatBreed(opt),
			g.Animal().Sex(opt),
			g.Animal().MicrochipNumber(opt),
		)
	}
}
