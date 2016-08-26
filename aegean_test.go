package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToAegean() {
	fmt.Println(IntegerToAegean(42))
	// Output: 𐄓𐄈
}

func TestIntegerToAegean(t *testing.T) {
	Convey("Testing IntegerToAegean()", t, FailureContinues, func() {
		testing := map[int]string{
			0:      "zero not supported",
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
			100000: "too big number",
			100001: "too big number",
		}
		for input, expectedOutput := range testing {
			So(IntegerToAegean(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToAegean(-1), ShouldEqual, "negative values not supported")

		// FIXME: large numbers using ansi code
	})
}
