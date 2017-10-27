package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["fr-be"] = Language{
		Name:    "Belgian French",
		Aliases: []string{"fr-be", "fr_BE", "belgian"},
		Flag:    "🇧🇪",

		IntegerToWords: IntegerToFrBe,
	}
}

// IntegerToFrBe converts an integer to French (belgium) words
func IntegerToFrBe(input int) string {
	var frBeMegas = []string{"", "mille", "million", "milliard", "billion", "billiard", "trillion", "trilliard", "quadrillion", "quadrilliard", "quintillion", "quintilliard"}
	var frBeUnits = []string{"", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf"}
	var frBeTens = []string{"", "dix", "vingt", "trente", "quarante", "cinquante", "soixante", "septante", "quatre-vingt", "nonante"}
	var frBeTeens = []string{"dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"}

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
					words = append(words, frBeUnits[hundreds], "cents")
					goto tripletEnd
				} else {
					words = append(words, frBeUnits[hundreds], "cent")
				}
			}
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, frBeUnits[units])
		case 1:
			words = append(words, frBeTeens[units])
			break
		case 8:
			switch units {
			case 0:
				words = append(words, frBeTens[tens]+"s")
				break
			default:
				words = append(words, frBeTens[tens], frBeUnits[units])
				break
			}
			break
		default:
			switch units {
			case 0:
				words = append(words, frBeTens[tens])
				break
			case 1:
				words = append(words, frBeTens[tens], "et", frBeUnits[units])
				break
			default:
				word := fmt.Sprintf("%s-%s", frBeTens[tens], frBeUnits[units])
				words = append(words, word)
				break
			}
			break
		}

	tripletEnd:
		// mega
		mega := frBeMegas[idx]
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
