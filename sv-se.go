package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["sv-se"] = Language{
		Name:    "Swedish",
		Aliases: []string{"sv-se", "sv_SE", "swedish"},
		Flag:    "🇸🇪",

		IntegerToWords: IntegerToSvSe,
	}
}

// IntegerToSvSe converts an integer to Swedish words
func IntegerToSvSe(input int) string {
	var swedishMegas = []string{"", "tusen", "miljoner", "miljarder", "biljoner", "kvadriljon", "kvintiljon", "sextillion", "septillion", "octillion", "nonillion", "decillion"}
	var swedishUnits = []string{"noll", "en", "två", "tre", "fyra", "fem", "sex", "sju", "åtta", "nio"}
	var swedishTens = []string{"noll", "tio", "tjugo", "trettio", "fyrtio", "femtio", "sextio", "sjuttio", "åttio", "nittio"}
	var swedishTeens = []string{"tio", "elva", "tolv", "tretton", "fjorton", "femton", "sexton", "sjutton", "arton", "nitton"}

	words := []string{}

	if input < 0 {
		words = append(words, "mindre")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "noll"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
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
		default:
			words = append(words, swedishUnits[hundreds], "hundra")
			break
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, swedishUnits[units])
		case 1:
			words = append(words, swedishTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", swedishTens[tens], swedishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, swedishTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := swedishMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	return strings.Join(words, " ")
}
