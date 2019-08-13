package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["fr-fr"] = Language{
		Name:    "French",
		Aliases: []string{"fr", "fr-fr", "fr_FR", "french"},
		Flag:    "🇫🇷",

		IntegerToWords: IntegerToFrFr,
	}
}

// IntegerToFrFr converts an integer to French words
func IntegerToFrFr(input int) string {
	var frenchMegas = []string{"", "mille", "million", "milliard", "billion", "billiard", "trillion", "trilliard", "quadrillion", "quadrilliard", "quintillion", "quintilliard"}
	var frenchUnits = []string{"", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf"}
	var frenchTens = []string{"", "dix", "vingt", "trente", "quarante", "cinquante", "soixante", "soixante", "quatre-vingt", "quatre-vingt"}
	var frenchTeens = []string{"dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "moins")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "zéro"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// special cases
		if triplet == 1 && idx == 1 {
			words = append(words, "mille")
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		//log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds > 0 {
			if hundreds == 1 {
				words = append(words, "cent")
			} else {
				if tens == 0 && units == 0 {
					words = append(words, frenchUnits[hundreds], "cents")
					goto tripletEnd
				} else {
					words = append(words, frenchUnits[hundreds], "cent")
				}
			}
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, frenchUnits[units])
		case 1:
			words = append(words, frenchTeens[units])
			break
		case 7:
			switch units {
			case 1:
				words = append(words, frenchTens[tens], "et", frenchTeens[units])
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchTeens[units])
				words = append(words, word)
				break
			}
			break
		case 8:
			switch units {
			case 0:
				words = append(words, frenchTens[tens]+"s")
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchUnits[units])
				words = append(words, word)
				break
			}
			break
		case 9:
			word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchTeens[units])
			words = append(words, word)
			break
		default:
			switch units {
			case 0:
				words = append(words, frenchTens[tens])
				break
			case 1:
				words = append(words, frenchTens[tens], "et", frenchUnits[units])
				break
			default:
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchUnits[units])
				words = append(words, word)
				break
			}
			break
		}

	tripletEnd:
		// mega
		mega := frenchMegas[idx]
		if mega != "" {
			if mega != "mille" && triplet > 1 {
				mega += "s"
			}
			words = append(words, mega)
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
