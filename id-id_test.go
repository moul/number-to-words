package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToIDID() {
	fmt.Println(IntegerToIDID(42))
	// Output: empat puluh dua
}

func TestIntegerToIDID(t *testing.T) {
	Convey("Testing IntegerToIDID()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "nol",
			1:             "satu",
			9:             "sembilan",
			10:            "sepuluh",
			11:            "sebelas",
			19:            "sembilanbelas",
			20:            "dua puluh",
			21:            "dua puluh satu",
			80:            "delapan puluh",
			90:            "sembilan puluh",
			99:            "sembilan puluh sembilan",
			100:           "seratus",
			101:           "seratus satu",
			111:           "seratus sebelas",
			120:           "seratus dua puluh",
			121:           "seratus dua puluh satu",
			900:           "sembilan ratus",
			909:           "sembilan ratus sembilan",
			919:           "sembilan ratus sembilanbelas",
			990:           "sembilan ratus sembilan puluh",
			999:           "sembilan ratus sembilan puluh sembilan",
			1000:          "seribu",
			2000:          "dua ribu",
			4000:          "empat ribu",
			5000:          "lima ribu",
			11000:         "sebelas ribu",
			21000:         "dua puluh satu ribu",
			100000:        "seratus ribu",
			101000:        "seratus satu ribu",
			999000:        "sembilan ratus sembilan puluh sembilan ribu",
			999999:        "sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan",
			1000000:       "satu juta",
			1001000:       "satu juta seribu",
			2000000:       "dua juta",
			4000000:       "empat juta",
			5000000:       "lima juta",
			100100100:     "seratus juta seratus ribu seratus",
			500500500:     "lima ratus juta lima ratus ribu lima ratus",
			606606606:     "enam ratus enam juta enam ratus enam ribu enam ratus enam",
			999000000:     "sembilan ratus sembilan puluh sembilan juta",
			999000999:     "sembilan ratus sembilan puluh sembilan juta sembilan ratus sembilan puluh sembilan",
			999999000:     "sembilan ratus sembilan puluh sembilan juta sembilan ratus sembilan puluh sembilan ribu",
			999999999:     "sembilan ratus sembilan puluh sembilan juta sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan",
			1174315110:    "satu milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sepuluh",
			1174315119:    "satu milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sembilanbelas",
			15174315119:   "limabelas milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sembilanbelas",
			35174315119:   "tiga puluh lima milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sembilanbelas",
			935174315119:  "sembilan ratus tiga puluh lima milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sembilanbelas",
			2935174315119: "dua triliun sembilan ratus tiga puluh lima milyar seratus tujuh puluh empat juta tiga ratus limabelas ribu seratus sembilanbelas",
		}
		for input, expectedOutput := range testing {
			So(IntegerToIDID(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToIDID(-1), ShouldEqual, "minus satu")
	})
}
