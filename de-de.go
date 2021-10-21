package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["de-de"] = Language{
		Name:    "German",
		Aliases: []string{"de", "de-de", "de_DE", "german", "Deutsch", "deutsch"},
		Flag:    "🇩🇪",

		IntegerToWords: IntegerToDeDe,
	}
}

// IntegerToDeDe converts an integer to German words
func IntegerToDeDe(input int) string {
	var deMegasSingular = []string{"", "eintausend", "eine Million", "eine Milliarde", "eine Billion", "eine Billiarde", "eine Trillion", "eine Trilliarde", "eine Quadrillion", "eine Quadrilliarde", "eine Quintillion", "eine Quintilliarde", "eine Sextillion", "eine Sextilliarde", "eine Septillion", "eine Septilliarde"}
	var deMegasPlural = []string{"", "tausend", "Millionen", "Milliarden", "Billionen", "Billiarden", "Trillionen", "Trilliarden", "Quadrillionen", "Quadrilliarden", "Quintillionen", "Quintilliarden", "Sextillionen", "Sextilliarden", "Septillionen", "Septilliarden"}
	var deUnits = []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	var deTens = []string{"", "zehn", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}
	var deTeens = []string{"zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	words := []string{}

	if input < 0 {
		words = append(words, "minus ")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "null"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		// nothing to do for empty triplet
		if triplet == 0 {
			continue
		}

		var mega string
		if triplet == 1 && idx != 0 { // handle singular megas
			mega = deMegasSingular[idx]

			if idx > 1 { // Million and above, megas need separator
				mega = fmt.Sprintf(" %s ", mega)
			}

			words = append(words, mega)
			continue
		}

		mega = deMegasPlural[idx]

		if idx > 1 { // Million and above, megas need separator
			mega = fmt.Sprintf(" %s ", mega)
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		if hundreds == 1 {
			words = append(words, "einhundert")
		} else if hundreds > 0 {
			words = append(words, deUnits[hundreds], "hundert")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, deUnits[units])
		case 1:
			words = append(words, deTeens[units])
			break
		default:
			if units == 1 {
				word := fmt.Sprintf("%sund%s", "ein", deTens[tens])
				words = append(words, word)
			} else if units > 0 {
				word := fmt.Sprintf("%sund%s", deUnits[units], deTens[tens])
				words = append(words, word)
			} else {
				words = append(words, deTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega != "" {
			words = append(words, mega)
		}
	}

	joint := strings.Join(words, "")
	paddingFix := strings.TrimSpace(joint)

	return paddingFix
}
