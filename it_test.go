package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToItalian() {
	fmt.Println(IntegerToItalian(42))
	// Output: quarantadue
}

func TestIntegerToItalian(t *testing.T) {
	Convey("Testing IntegerToItalian()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "zero",
			1:             "uno",
			9:             "nove",
			10:            "dieci",
			11:            "undici",
			19:            "diciannove",
			20:            "venti",
			21:            "ventiuno",
			80:            "ottanta",
			90:            "novanta",
			99:            "novantanove",
			100:           "cento",
			101:           "cento uno",
			111:           "cento undici",
			120:           "cento venti",
			121:           "cento ventiuno",
			900:           "novecento",
			909:           "novecento nove",
			919:           "novecento diciannove",
			990:           "novecento novanta",
			999:           "novecento novantanove",
			1000:          "uno mille",
			2000:          "due mille",
			4000:          "quattro mille",
			5000:          "cinque mille",
			11000:         "undici mille",
			21000:         "ventiuno mille",
			999000:        "novecento novantanove mille",
			999999:        "novecento novantanove mille novecento novantanove",
			1000000:       "uno milione",
			2000000:       "due milione",
			4000000:       "quattro milione",
			5000000:       "cinque milione",
			999000000:     "novecento novantanove milione",
			999000999:     "novecento novantanove milione novecento novantanove",
			999999000:     "novecento novantanove milione novecento novantanove mille",
			999999999:     "novecento novantanove milione novecento novantanove mille novecento novantanove",
			1174315110:    "uno miliardo cento settantaquattro milione trecento quindici mille cento dieci",
			1174315119:    "uno miliardo cento settantaquattro milione trecento quindici mille cento diciannove",
			15174315119:   "quindici miliardo cento settantaquattro milione trecento quindici mille cento diciannove",
			35174315119:   "trentacinque miliardo cento settantaquattro milione trecento quindici mille cento diciannove",
			935174315119:  "novecento trentacinque miliardo cento settantaquattro milione trecento quindici mille cento diciannove",
			2935174315119: "due triliardo novecento trentacinque miliardo cento settantaquattro milione trecento quindici mille cento diciannove",
		}
		for input, expectedOutput := range testing {
			So(IntegerToItalian(input), ShouldEqual, expectedOutput)
		}

		// FIXME: test negative values
	})
}
