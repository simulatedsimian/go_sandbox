package elitecore

import (
	"fmt"
	"testing"
)

func TestEliteCore(t *testing.T) {

	seed := NewRNGSeed(1, 2, 3, 4)

	for n := 0; n < 100; n++ {
		fmt.Printf("rnd = %d\n", seed.GenRnd())
	}

}
