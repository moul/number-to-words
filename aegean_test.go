package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToAegean() {
	fmt.Println(IntegerToAegean(42))
	// Output: 𐄓𐄈
}

func TestIntegerToAegean(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:     "", // unsupported
		0:      "", // not supported
		1:      "𐄇",
		2:      "𐄈",
		3:      "𐄉",
		9:      "𐄏",
		10:     "𐄐",
		11:     "𐄐𐄇",
		12:     "𐄐𐄈",
		19:     "𐄐𐄏",
		20:     "𐄑",
		21:     "𐄑𐄇",
		40:     "𐄓",
		49:     "𐄓𐄏",
		50:     "𐄔",
		60:     "𐄕",
		90:     "𐄘",
		99:     "𐄘𐄏",
		100:    "𐄙",
		101:    "𐄙𐄇",
		111:    "𐄙𐄐𐄇",
		12345:  "𐄫𐄣𐄛𐄓𐄋",
		99999:  "𐄳𐄪𐄡𐄘𐄏",
		100000: "", // too big
		100001: "", // too big
	}
	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToAegean(input))
		})
	}
}
