package ntw

import (
	"fmt"
	"strings"
)

var dutchMegas = []string{"", "duizend", "miljoen", "miljard", "biljoen", "biljard", "triljoen", "triljard", "septiljoen", "octillion", "noniljoen", "decillion"}
var dutchUnits = []string{"nul", "één", "twee", "drie", "vier", "vijf", "zes", "zeven", "acht", "negen"}
var dutchTens = []string{"nul", "tien", "twintig", "dertig", "veertig", "vijftig", "zestig", "zeventig", "tachtig", "negentig"}
var dutchTeens = []string{"tien", "elf", "twaalf", "dertien", "veertien", "vijftien", "zestien", "zeventien", "achttien", "negentien"}

// IntegerToDutch converts an integer to Dutch words
func IntegerToDutch(input int) string {
	words := []string{}

	if input < 0 {
		words = append(words, "TODO")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "nul"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		tripletWords := []string{}
		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		switch hundreds {
		case 0:
			break
		case 1:
			tripletWords = append(tripletWords, "honderd")
			break
		default:
			tripletWords = append(tripletWords, fmt.Sprintf("%shonderd", dutchUnits[hundreds]))
			break
		}

		if tens == 0 && units == 0 {
			words = append(words, strings.Join(tripletWords, ""))
			continue
		}

		switch tens {
		case 0:
			tripletWords = append(tripletWords, dutchUnits[units])
		case 1:
			tripletWords = append(tripletWords, dutchTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%sen%s", dutchUnits[units], dutchTens[tens])
				tripletWords = append(tripletWords, word)
			} else {
				tripletWords = append(tripletWords, dutchTens[tens])
			}
			break
		}

		// megas
		switch idx {
		case 0:
			words = append(words, strings.Join(tripletWords, ""))
		case 1:
			tripletWords = append(tripletWords, dutchMegas[idx])
			words = append(words, strings.Join(tripletWords, ""))
		default:
			words = append(words, strings.Join(tripletWords, ""))
			words = append(words, dutchMegas[idx])
		}

	}

	return strings.Join(words, " ")
}
