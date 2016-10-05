package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToSpanish() {
	fmt.Println(IntegerToSpanish(42))
	// Output: forty-two
}

func TestIntegerToSpanish(t *testing.T) {
	Convey("Testing IntegerToSpanish()", t, FailureContinues, func() {
		testing := map[int]string{
			0:             "cero",
			1:             "uno",
			9:             "nueve",
			10:            "diez",
			11:            "once",
			19:            "diecinueve",
			20:            "veinte",
			21:            "veintiuno",
			80:            "ochenta",
			90:            "noventa",
			99:            "noventa y nueve",
			100:           "cien",
			101:           "ciento uno",
			111:           "ciento once",
			120:           "ciento veinte",
			121:           "ciento veinte y uno",
			900:           "novecientos",
			909:           "novecientos nueve",
			919:           "novecientos diecinueve",
			990:           "novecientos noventa",
			999:           "novecientos noventa y nueve",
			1000:          "mil",
			2000:          "dos mil",
			4000:          "cuatro mil",
			5000:          "cinco mil",
			11000:         "once mil",
			21000:         "veintiún mil",
			999000:        "novecientos noventa y nueve mil",
			999999:        "Novecientos noventa y nueve mil novecientos noventa y nueve",
			1000000:       "un millon",
			2000000:       "dos millon",
			4000000:       "cuatro millones",
			5000000:       "cinco millones",
			100100100:     "cien millones cien mil cien",
			500500500:     "quinientos millones quinientos mil quinientos",
			606606606:     "seiscientos seis millones seiscientos seis mil seiscientos seis",
			999000000:     "novecientos noventa y nueve millones",
			999000999:     "novecientos noventa y nueve millones novecientos noventa y nueve",
			999999000:     "novecientos noventa y nueve millones novecientos noventa y nueve mil",
			999999999:     "novecientos noventa y nueve millones novecientos noventa y nueve mil novecientos noventa y nueve",
			1174315110:    "mil ciento setenta y cuatro millones trescientos quince mil ciento diez",
			1174315119:    "mil ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
			15174315119:   "quince mil ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
			35174315119:   "treinta y cinco mil ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
			935174315119:  "novecientos treinta y cinco mil ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
			2935174315119: "dos billones novecientos treinta y cinco mil ciento setenta y cuatro millones trescientos quince mil ciento diecinueve",
		}
		for input, expectedOutput := range testing {
			So(IntegerToSpanish(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToSpanish(-1), ShouldEqual, "menos uno")
	})
}
