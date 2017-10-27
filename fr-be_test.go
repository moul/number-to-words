package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToFrBe() {
	fmt.Println(IntegerToFrBe(42))
	fmt.Println(IntegerToFrBe(92))
	// Output:
	// quarante-deux
	// nonante-deux
}

func TestIntegerToFrBe(t *testing.T) {
	Convey("Testing IntegerToFrBe()", t, FailureContinues, func() {
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
			70:            "septante",
			71:            "septante et un",
			72:            "septante-deux",
			73:            "septante-trois",
			79:            "septante-neuf",
			80:            "quatre-vingts",
			81:            "quatre-vingt un",
			82:            "quatre-vingt deux",
			90:            "nonante",
			91:            "nonante et un",
			92:            "nonante-deux",
			99:            "nonante-neuf",
			100:           "cent",
			101:           "cent un",
			111:           "cent onze",
			120:           "cent vingt",
			121:           "cent vingt et un",
			900:           "neuf cents",
			909:           "neuf cent neuf",
			919:           "neuf cent dix-neuf",
			990:           "neuf cent nonante",
			999:           "neuf cent nonante-neuf",
			1000:          "mille",
			2000:          "deux mille",
			4000:          "quatre mille",
			5000:          "cinq mille",
			11000:         "onze mille",
			21000:         "vingt et un mille",
			999000:        "neuf cent nonante-neuf mille",
			999999:        "neuf cent nonante-neuf mille neuf cent nonante-neuf",
			1000000:       "un million",
			2000000:       "deux millions",
			4000000:       "quatre millions",
			5000000:       "cinq millions",
			100100100:     "cent millions cent mille cent",
			123456789:     "cent vingt-trois millions quatre cent cinquante-six mille sept cent quatre-vingt neuf",
			500500500:     "cinq cents millions cinq cents mille cinq cents",
			606606606:     "six cent six millions six cent six mille six cent six",
			999000000:     "neuf cent nonante-neuf millions",
			999000999:     "neuf cent nonante-neuf millions neuf cent nonante-neuf",
			999999000:     "neuf cent nonante-neuf millions neuf cent nonante-neuf mille",
			999999999:     "neuf cent nonante-neuf millions neuf cent nonante-neuf mille neuf cent nonante-neuf",
			1174315110:    "un milliard cent septante-quatre millions trois cent quinze mille cent dix",
			1174315119:    "un milliard cent septante-quatre millions trois cent quinze mille cent dix-neuf",
			15174315119:   "quinze milliards cent septante-quatre millions trois cent quinze mille cent dix-neuf",
			35174315119:   "trente-cinq milliards cent septante-quatre millions trois cent quinze mille cent dix-neuf",
			935174315119:  "neuf cent trente-cinq milliards cent septante-quatre millions trois cent quinze mille cent dix-neuf",
			2935174315119: "deux billions neuf cent trente-cinq milliards cent septante-quatre millions trois cent quinze mille cent dix-neuf",
		}
		for input, expectedOutput := range testing {
			So(IntegerToFrBe(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToFrBe(-1), ShouldEqual, "moins un")
	})
}
