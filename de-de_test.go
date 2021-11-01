package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToDeDe() {
	fmt.Println(IntegerToDeDe(42))
	// Output: zweiundvierzig
}

func TestIntegerToDeDe(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "minus eins",
		0:             "null",
		1:             "eins",
		9:             "neun",
		10:            "zehn",
		11:            "elf",
		19:            "neunzehn",
		20:            "zwanzig",
		21:            "einundzwanzig",
		80:            "achtzig",
		90:            "neunzig",
		99:            "neunundneunzig",
		100:           "einhundert",
		101:           "einhunderteins",
		111:           "einhundertelf",
		120:           "einhundertzwanzig",
		121:           "einhunderteinundzwanzig",
		900:           "neunhundert",
		909:           "neunhundertneun",
		919:           "neunhundertneunzehn",
		990:           "neunhundertneunzig",
		999:           "neunhundertneunundneunzig",
		1000:          "eintausend",
		2000:          "zweitausend",
		4000:          "viertausend",
		5000:          "fünftausend",
		11000:         "elftausend",
		21000:         "einundzwanzigtausend",
		999000:        "neunhundertneunundneunzigtausend",
		999999:        "neunhundertneunundneunzigtausendneunhundertneunundneunzig",
		1000000:       "eine Million",
		2000000:       "zwei Millionen",
		4000000:       "vier Millionen",
		5000000:       "fünf Millionen",
		100100100:     "einhundert Millionen einhunderttausendeinhundert",
		500500500:     "fünfhundert Millionen fünfhunderttausendfünfhundert",
		606606606:     "sechshundertsechs Millionen sechshundertsechstausendsechshundertsechs",
		999000000:     "neunhundertneunundneunzig Millionen",
		999000999:     "neunhundertneunundneunzig Millionen neunhundertneunundneunzig",
		999999000:     "neunhundertneunundneunzig Millionen neunhundertneunundneunzigtausend",
		999999999:     "neunhundertneunundneunzig Millionen neunhundertneunundneunzigtausendneunhundertneunundneunzig",
		1174315110:    "eine Milliarde einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertzehn",
		1174315119:    "eine Milliarde einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertneunzehn",
		15174315119:   "fünfzehn Milliarden einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertneunzehn",
		35174315119:   "fünfunddreißig Milliarden einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertneunzehn",
		935174315119:  "neunhundertfünfunddreißig Milliarden einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertneunzehn",
		2935174315119: "zwei Billionen neunhundertfünfunddreißig Milliarden einhundertvierundsiebzig Millionen dreihundertfünfzehntausendeinhundertneunzehn",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToDeDe(input))
		})
	}
}
