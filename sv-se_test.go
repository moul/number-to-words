package ntw

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToSvSe() {
	fmt.Println(IntegerToSvSe(42))
	// Output: fyrtiotvå
}

func TestIntegerToSvSe(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1241521:      "mindre en miljon tvåhundrafyrtioettusenfemhundratjugoen",
		-200:          "mindre tvåhundra",
		-1:            "mindre en",
		0:             "noll",
		1:             "en",
		9:             "nio",
		10:            "tio",
		11:            "elva",
		19:            "nitton",
		20:            "tjugo",
		21:            "tjugoen",
		80:            "åttio",
		90:            "nittio",
		99:            "nittionio",
		100:           "etthundra",
		101:           "etthundraen",
		111:           "etthundraelva",
		120:           "etthundratjugo",
		121:           "etthundratjugoen",
		900:           "niohundra",
		909:           "niohundranio",
		919:           "niohundranitton",
		990:           "niohundranittio",
		999:           "niohundranittionio",
		1000:          "ettusen",
		1500:          "ettusenfemhundra",
		2000:          "tvåtusen",
		4000:          "fyratusen",
		5000:          "femtusen",
		11000:         "elvatusen",
		21000:         "tjugoettusen",
		999000:        "niohundranittioniotusen",
		999999:        "niohundranittioniotusenniohundranittionio",
		1000000:       "en miljon",
		2000000:       "två miljoner",
		4000000:       "fyra miljoner",
		5000000:       "fem miljoner",
		100100100:     "etthundra miljoner etthundratusenetthundra",
		500500500:     "femhundra miljoner femhundratusenfemhundra",
		606606606:     "sexhundrasex miljoner sexhundrasextusensexhundrasex",
		999000000:     "niohundranittionio miljoner",
		999000999:     "niohundranittionio miljoner niohundranittionio",
		999999000:     "niohundranittionio miljoner niohundranittioniotusen",
		999999999:     "niohundranittionio miljoner niohundranittioniotusenniohundranittionio",
		1001000000:    "en miljard en miljon",
		1002000000:    "en miljard två miljoner",
		1111111111:    "en miljard etthundraelva miljoner etthundraelvatusenetthundraelva",
		1010101010:    "en miljard tio miljoner etthundraettusentio",
		1174315110:    "en miljard etthundrasjuttiofyra miljoner trehundrafemtontusenetthundratio",
		1174315119:    "en miljard etthundrasjuttiofyra miljoner trehundrafemtontusenetthundranitton",
		10101010101:   "tio miljarder etthundraen miljoner tiotusenetthundraen",
		15174315119:   "femton miljarder etthundrasjuttiofyra miljoner trehundrafemtontusenetthundranitton",
		35174315119:   "trettiofem miljarder etthundrasjuttiofyra miljoner trehundrafemtontusenetthundranitton",
		935174315119:  "niohundratrettiofem miljarder etthundrasjuttiofyra miljoner trehundrafemtontusenetthundranitton",
		2935174315119: "två biljoner niohundratrettiofem miljarder etthundrasjuttiofyra miljoner trehundrafemtontusenetthundranitton",
		math.MaxInt64: "nio kvintiljoner tvåhundratjugotre kvadriljoner trehundrasjuttiotvå biljoner trettiosex miljarder åttahundrafemtiofyra miljoner sjuhundrasjuttiofemtusenåttahundrasju",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToSvSe(input))
		})
	}
}
