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

// IntegerToEnUs converts an integer to American English words
func IntegerToDeDe(input int) string {
	var deMegas = []string{"", "tausend", "million", "milliarde", "billion", "billiarde", "trillion", "sextillion", "siebentausend", "oktillion", "nichtmillion", "dezillion", "undezillion", "duodezillion", "tredemillion", "quattuordeillion"}

	var deUnits = []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	var deTens = []string{"", "zehn", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}
	var deTeens = []string{"zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechszehn", "siebzehn", "achtzehn", "neunzehn"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "null"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		//log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds > 0 {
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
			if units > 0 {
				word := fmt.Sprintf("%s-%s", deTens[tens], deUnits[units])
				words = append(words, word)
			} else {
				words = append(words, deTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := deMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
