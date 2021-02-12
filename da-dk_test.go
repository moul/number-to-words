package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToDaDk() {
	fmt.Println(IntegerToDaDk(42))
	// Output: toogfyrre
}

func TestIntegerToDaDk(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "minus en",
		0:             "nul",
		1:             "en",
		9:             "ni",
		10:            "ti",
		11:            "elleve",
		19:            "nitten",
		20:            "tyve",
		21:            "enogtyve",
		80:            "firs",
		90:            "halvfems",
		99:            "nioghalvfems",
		100:           "et hundrede",
		101:           "et hundrede og en",
		111:           "et hundrede og elleve",
		120:           "et hundrede og tyve",
		121:           "et hundrede og enogtyve",
		900:           "ni hundrede",
		909:           "ni hundrede og ni",
		919:           "ni hundrede og nitten",
		990:           "ni hundrede og halvfems",
		999:           "ni hundrede og nioghalvfems",
		1000:          "et tusind",
		2000:          "to tusind",
		4000:          "fire tusind",
		5000:          "fem tusind",
		11000:         "elleve tusind",
		21000:         "enogtyve tusind",
		999000:        "ni hundrede og nioghalvfems tusind",
		999999:        "ni hundrede og nioghalvfems tusind ni hundrede og nioghalvfems",
		1000000:       "en million",
		2000000:       "to millioner",
		4000000:       "fire millioner",
		5000000:       "fem millioner",
		100100100:     "et hundrede millioner et hundrede tusind et hundrede",
		500500500:     "fem hundrede millioner fem hundrede tusind fem hundrede",
		606606606:     "seks hundrede og seks millioner seks hundrede og seks tusind seks hundrede og seks",
		999000000:     "ni hundrede og nioghalvfems millioner",
		999000999:     "ni hundrede og nioghalvfems millioner ni hundrede og nioghalvfems",
		999999000:     "ni hundrede og nioghalvfems millioner ni hundrede og nioghalvfems tusind",
		999999999:     "ni hundrede og nioghalvfems millioner ni hundrede og nioghalvfems tusind ni hundrede og nioghalvfems",
		1174315110:    "en milliard et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og ti",
		1174315119:    "en milliard et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og nitten",
		15174315119:   "femten milliarder et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og nitten",
		35174315119:   "femogtredive milliarder et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og nitten",
		935174315119:  "ni hundrede og femogtredive milliarder et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og nitten",
		2935174315119: "to billioner ni hundrede og femogtredive milliarder et hundrede og fireoghalvfjerds millioner tre hundrede og femten tusind et hundrede og nitten",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToDaDk(input))
		})
	}
}
