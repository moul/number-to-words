package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["en-gb"] = Language{
		Name:    "British English",
		Aliases: []string{"en", "en-gb", "es_GB", "british", "english"},
		Flag:    "🇬🇧",

		IntegerToWords: IntegerToEnGb,
	}
}

// IntegerToEnGb converts an integer to British English words
func IntegerToEnGb(input int) string {
	var englishMegas = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
	var englishUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var englishTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	var englishTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "zero"
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

		if hundreds > 0 {
			words = append(words, englishUnits[hundreds], "hundred")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		if hundreds > 0 && (tens > 0 || units > 0) {
			words = append(words, "and")
		}

		switch tens {
		case 0:
			words = append(words, englishUnits[units])
		case 1:
			words = append(words, englishTeens[units])
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", englishTens[tens], englishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, englishTens[tens])
			}
		}

	tripletEnd:
		// mega
		if mega := englishMegas[idx]; mega != "" {
			words = append(words, mega)
		}

		if idx > 0 {
			words = append(words, ",")
		}
	}

	var result string
	result = strings.Join(words, " ")
	result = strings.Replace(result, " ,", ",", -1)
	result = strings.TrimSuffix(result, ",")

	return result
}
