package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToHuHu() {
	fmt.Println(IntegerToHuHu(42))
	// Output: negyvenkettő
}

func TestIntegerToHuHu(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		2001:    "kettőezer-egy",
		3016:    "háromezer-tizenhat",
		47563:   "negyvenhétezer-ötszázhatvanhárom",
		4600:    "négyezer-hatszáz",
		2014:    "kettőezer-tizennégy",
		7490530: "hétmillió-négyszázkilencvenezer-ötszázharminc",

		-1:            "mínusz egy",
		0:             "zéró",
		1:             "egy",
		9:             "kilenc",
		10:            "tíz",
		11:            "tizenegy",
		19:            "tizenkilenc",
		20:            "húsz",
		21:            "huszonegy",
		80:            "nyolcvan",
		90:            "kilencven",
		99:            "kilencvenkilenc",
		100:           "egyszáz",
		101:           "egyszázegy",
		111:           "egyszáztizenegy",
		120:           "egyszázhúsz",
		121:           "egyszázhuszonegy",
		900:           "kilencszáz",
		909:           "kilencszázkilenc",
		919:           "kilencszáztizenkilenc",
		990:           "kilencszázkilencven",
		999:           "kilencszázkilencvenkilenc",
		1000:          "egyezer",
		2000:          "kettőezer",
		4000:          "négyezer",
		5000:          "ötezer",
		11000:         "tizenegyezer",
		21000:         "huszonegyezer",
		999000:        "kilencszázkilencvenkilencezer",
		999999:        "kilencszázkilencvenkilencezer-kilencszázkilencvenkilenc",
		1000000:       "egymillió",
		2000000:       "kettőmillió",
		4000000:       "négymillió",
		5000000:       "ötmillió",
		100100100:     "egyszázmillió-egyszázezer-egyszáz",
		500500500:     "ötszázmillió-ötszázezer-ötszáz",
		606606606:     "hatszázhatmillió-hatszázhatezer-hatszázhat",
		999000000:     "kilencszázkilencvenkilencmillió",
		999000999:     "kilencszázkilencvenkilencmillió-kilencszázkilencvenkilenc",
		999999000:     "kilencszázkilencvenkilencmillió-kilencszázkilencvenkilencezer",
		999999999:     "kilencszázkilencvenkilencmillió-kilencszázkilencvenkilencezer-kilencszázkilencvenkilenc",
		1174315110:    "egymilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztíz",
		1174315119:    "egymilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztizenkilenc",
		15174315119:   "tizenötmilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztizenkilenc",
		35174315119:   "harmincötmilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztizenkilenc",
		935174315119:  "kilencszázharmincötmilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztizenkilenc",
		2935174315119: "kettőtrillió-kilencszázharmincötmilliárd-egyszázhetvennégymillió-háromszáztizenötezer-egyszáztizenkilenc",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToHuHu(input))
		})
	}
}
