package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToEnglish() {
	fmt.Println(IntegerToEnglish(42))
	// Output: forty-two
}

func TestIntegerToEnglish(t *testing.T) {
	Convey("Testing IntegerToEnglish()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "zero",
			1:             "one",
			9:             "nine",
			10:            "ten",
			11:            "eleven",
			19:            "nineteen",
			20:            "twenty",
			21:            "twenty-one",
			80:            "eighty",
			90:            "ninety",
			99:            "ninety-nine",
			100:           "one hundred",
			101:           "one hundred one",
			111:           "one hundred eleven",
			120:           "one hundred twenty",
			121:           "one hundred twenty-one",
			900:           "nine hundred",
			909:           "nine hundred nine",
			919:           "nine hundred nineteen",
			990:           "nine hundred ninety",
			999:           "nine hundred ninety-nine",
			1000:          "one thousand",
			2000:          "two thousand",
			4000:          "four thousand",
			5000:          "five thousand",
			11000:         "eleven thousand",
			21000:         "twenty-one thousand",
			999000:        "nine hundred ninety-nine thousand",
			999999:        "nine hundred ninety-nine thousand nine hundred ninety-nine",
			1000000:       "one million",
			2000000:       "two million",
			4000000:       "four million",
			5000000:       "five million",
			999000000:     "nine hundred ninety-nine million",
			999000999:     "nine hundred ninety-nine million nine hundred ninety-nine",
			999999000:     "nine hundred ninety-nine million nine hundred ninety-nine thousand",
			999999999:     "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
			1174315110:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred ten",
			1174315119:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
			15174315119:   "fifteen billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
			35174315119:   "thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
			935174315119:  "nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
			2935174315119: "two trillion nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
		}
		for input, expectedOutput := range testing {
			So(IntegerToEnglish(input), ShouldEqual, expectedOutput)
		}

		// FIXME: test negative values
	})
}
