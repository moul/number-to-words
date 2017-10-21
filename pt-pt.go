package ntw

import (
	"fmt"
	"math"
)

var portugueseMegasSingular = []string{"", "mil", "milhão", "mil milhões", "bilião"}
var portugueseMegasPlural = []string{"", "mil", "milhões", "mil milhões", "bilhões"}
var portugueseUnitsAdjectives = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseUnits = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseHundreds = []string{"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos", "cento"}
var portugueseTens = []string{"", "dez", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
var portugueseTeens = []string{"dez", "onze", "doze", "treze", "catorze", "quinze", "dezasseis", "dezasete", "dezoito", "dezanove"}

func IntegerToPortuguese_PT(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}
	fmt.Printf("%d  %f\n", input, math.Floor(math.Log10(math.Abs(float64(input))))+1)

	triplets := integerToTriplets(input)
	fmt.Println(triplets)
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

	}
	// split integer in triplets
	return Join(words, "")

}

// JOIN From strings.Join
// following robs pike advice from golang proverbs "A little copying is better than a little dependency."
func Join(a []string, sep string) string {

	switch len(a) {

	case 0:

		return ""

	case 1:

		return a[0]

	case 2:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1]

	case 3:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1] + sep + a[2]

	}

	n := len(sep) * (len(a) - 1)

	for i := 0; i < len(a); i++ {

		n += len(a[i])

	}

	b := make([]byte, n)

	bp := copy(b, a[0])

	for _, s := range a[1:] {

		bp += copy(b[bp:], sep)

		bp += copy(b[bp:], s)

	}

	return string(b)

}