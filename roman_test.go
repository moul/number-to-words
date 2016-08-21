package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToRoman() {
	fmt.Println(IntegerToRoman(42))
	// Output: XLII
}

func TestIntegerToRoman(t *testing.T) {
	Convey("Testing IntegerToRoman()", t, FailureContinues, func() {
		testing := map[int]string{
			0:    "nulla",
			1:    "I",
			2:    "II",
			3:    "III",
			4:    "IV",
			5:    "V",
			6:    "VI",
			7:    "VII",
			8:    "VIII",
			9:    "IX",
			10:   "X",
			11:   "XI",
			12:   "XII",
			13:   "XIII",
			14:   "XIV",
			15:   "XV",
			16:   "XVI",
			17:   "XVII",
			18:   "XVIII",
			19:   "XIX",
			20:   "XX",
			21:   "XXI",
			40:   "XL",
			49:   "IL",
			50:   "L",
			60:   "LX",
			70:   "LXX",
			80:   "LXXX",
			90:   "XC",
			99:   "IC",
			100:  "C",
			101:  "CI",
			111:  "CXI",
			120:  "CXX",
			121:  "CXXI",
			499:  "ID",
			500:  "D",
			501:  "DI",
			900:  "CM",
			909:  "CMIX",
			919:  "CMXIX",
			990:  "XM",
			999:  "IM",
			1000: "M",
			2000: "MM",
			3000: "MMM",
			3888: "MMMDCCCLXXXVIII",
			//5000:          "",
			//11000:         "",
			//21000:         "",
			//999000:        "",
			//999999:        "",
			//1000000:       "",
			//2000000:       "",
			//4000000:       "",
			//5000000:       "",
			//999000000:     "",
			//999000999:     "",
			//999999000:     "",
			//999999999:     "",
			//1174315110:    "",
			//1174315119:    "",
			//15174315119:   "",
			//35174315119:   "",
			//935174315119:  "",
			//2935174315119: "",
		}
		for input, expectedOutput := range testing {
			So(IntegerToRoman(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToRoman(-1), ShouldEqual, "-I")

		// FIXME: large numbers using ansi code
	})
}
