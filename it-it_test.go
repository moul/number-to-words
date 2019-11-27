package ntw

import (
	"fmt"
	"testing"
)

func ExampleIntegerToItIt() {
	fmt.Println(IntegerToItIt(42))
	// Output: quarantadue
}

func TestIntegerToItIt(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "meno uno",
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
		100100100:     "cento milione cento mille cento",
		500500500:     "cinquecento milione cinquecento mille cinquecento",
		606606606:     "seicento sei milione seicento sei mille seicento sei",
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

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			output := IntegerToItIt(input)
			if expectedOutput != output {
				t.Fatalf("Expected %q, got %q.", expectedOutput, output)
			}
		})
	}
}
