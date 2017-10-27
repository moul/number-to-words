package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToFrench() {
	fmt.Println(IntegerToFrench(42))
	// Output: quarante-deux
}

func TestIntegerToFrench(t *testing.T) {
	Convey("Testing IntegerToFrench()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "zéro",
			1:             "un",
			4:             "quatre",
			9:             "neuf",
			10:            "dix",
			11:            "onze",
			12:            "douze",
			17:            "dix-sept",
			19:            "dix-neuf",
			20:            "vingt",
			21:            "vingt et un",
			42:            "quarante-deux",
			60:            "soixante",
			61:            "soixante et un",
			62:            "soixante-deux",
			70:            "soixante-dix",
			71:            "soixante et onze",
			72:            "soixante-douze",
			73:            "soixante-treize",
			79:            "soixante-dix-neuf",
			80:            "quatre-vingts",
			81:            "quatre-vingt un",
			82:            "quatre-vingt deux",
			90:            "quatre-vingt-dix",
			91:            "quatre-vingt-onze",
			92:            "quatre-vingt-douze",
			99:            "quatre-vingt-dix-neuf",
			100:           "cent",
			101:           "cent un",
			111:           "cent onze",
			120:           "cent vingt",
			121:           "cent vingt et un",
			900:           "neuf cents",
			909:           "neuf cent neuf",
			919:           "neuf cent dix-neuf",
			990:           "neuf cent quatre-vingt-dix",
			999:           "neuf cent quatre-vingt-dix-neuf",
			1000:          "mille",
			2000:          "deux mille",
			4000:          "quatre mille",
			5000:          "cinq mille",
			11000:         "onze mille",
			21000:         "vingt et un mille",
			999000:        "neuf cent quatre-vingt-dix-neuf mille",
			999999:        "neuf cent quatre-vingt-dix-neuf mille neuf cent quatre-vingt-dix-neuf",
			1000000:       "un million",
			2000000:       "deux millions",
			4000000:       "quatre millions",
			5000000:       "cinq millions",
			100100100:     "cent millions cent mille cent",
			123456789:     "cent vingt-trois millions quatre cent cinquante-six mille sept cent quatre-vingt neuf",
			500500500:     "cinq cents millions cinq cents mille cinq cents",
			606606606:     "six cent six millions six cent six mille six cent six",
			999000000:     "neuf cent quatre-vingt-dix-neuf millions",
			999000999:     "neuf cent quatre-vingt-dix-neuf millions neuf cent quatre-vingt-dix-neuf",
			999999000:     "neuf cent quatre-vingt-dix-neuf millions neuf cent quatre-vingt-dix-neuf mille",
			999999999:     "neuf cent quatre-vingt-dix-neuf millions neuf cent quatre-vingt-dix-neuf mille neuf cent quatre-vingt-dix-neuf",
			1174315110:    "un milliard cent soixante-quatorze millions trois cent quinze mille cent dix",
			1174315119:    "un milliard cent soixante-quatorze millions trois cent quinze mille cent dix-neuf",
			15174315119:   "quinze milliards cent soixante-quatorze millions trois cent quinze mille cent dix-neuf",
			35174315119:   "trente-cinq milliards cent soixante-quatorze millions trois cent quinze mille cent dix-neuf",
			935174315119:  "neuf cent trente-cinq milliards cent soixante-quatorze millions trois cent quinze mille cent dix-neuf",
			2935174315119: "deux billions neuf cent trente-cinq milliards cent soixante-quatorze millions trois cent quinze mille cent dix-neuf",
		}
		for input, expectedOutput := range testing {
			So(IntegerToFrench(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToFrench(-1), ShouldEqual, "moins un")
	})
}
