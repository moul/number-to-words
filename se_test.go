package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToSwedish() {
	fmt.Println(IntegerToSwedish(42))
	// Output: fyrtio-två
}

func TestIntegerToSwedish(t *testing.T) {
	Convey("Testing IntegerToSwedish()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "noll",
			1:             "en",
			9:             "nio",
			10:            "tio",
			11:            "elva",
			19:            "nitton",
			20:            "tjugo",
			21:            "tjugo-en",
			80:            "åttio",
			90:            "nittio",
			99:            "nittio-nio",
			100:           "en hundra",
			101:           "en hundra en",
			111:           "en hundra elva",
			120:           "en hundra tjugo",
			121:           "en hundra tjugo-en",
			900:           "nio hundra",
			909:           "nio hundra nio",
			919:           "nio hundra nitton",
			990:           "nio hundra nittio",
			999:           "nio hundra nittio-nio",
			1000:          "en tusen",
			2000:          "två tusen",
			4000:          "fyra tusen",
			5000:          "fem tusen",
			11000:         "elva tusen",
			21000:         "tjugo-en tusen",
			999000:        "nio hundra nittio-nio tusen",
			999999:        "nio hundra nittio-nio tusen nio hundra nittio-nio",
			1000000:       "en miljoner",
			2000000:       "två miljoner",
			4000000:       "fyra miljoner",
			5000000:       "fem miljoner",
			999000000:     "nio hundra nittio-nio miljoner",
			999000999:     "nio hundra nittio-nio miljoner nio hundra nittio-nio",
			999999000:     "nio hundra nittio-nio miljoner nio hundra nittio-nio tusen",
			999999999:     "nio hundra nittio-nio miljoner nio hundra nittio-nio tusen nio hundra nittio-nio",
			1174315110:    "en miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra tio",
			1174315119:    "en miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra nitton",
			15174315119:   "femton miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra nitton",
			35174315119:   "trettio-fem miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra nitton",
			935174315119:  "nio hundra trettio-fem miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra nitton",
			2935174315119: "två biljoner nio hundra trettio-fem miljarder en hundra sjuttio-fyra miljoner tre hundra femton tusen en hundra nitton",
		}
		for input, expectedOutput := range testing {
			So(IntegerToSwedish(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToSwedish(-1), ShouldEqual, "mindre en")
	})
}
