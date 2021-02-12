package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["da-dk"] = Language{
		Name:    "Dansk",
		Aliases: []string{"da", "da-dk", "da-DK", "dansk"},
		Flag:    "🇩🇰",

		IntegerToWords: IntegerToDaDk,
	}
}

// IntegerToDaDk converts an integer to Danish words
func IntegerToDaDk(input int) string {
	var danishMegas = []string{"", "tusind", "million", "milliard", "billion", "billiard", "trillion", "trilliard", "kvadrillion", "kvadrilliard", "kvintillion", "kvintilliard", "sekstillion", "sekstilliard", "septillion", "septiliard"}
	var danishUnits = []string{"", "en", "to", "tre", "fire", "fem", "seks", "syv", "otte", "ni"}
	var danishTens = []string{"", "ti", "tyve", "tredive", "fyrre", "halvtreds", "tres", "halvfjerds", "firs", "halvfems"}
	var danishTeens = []string{"ti", "elleve", "tolv", "tretten", "fjorten", "femten", "seksten", "sytten", "atten", "nitten"}

	words := []string{}

	// zero is a special case
	if input == 0 {
		return "nul"
	}

	// add minus if the number is negative
	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer into triplets
	triplets := integerToTriplets(input)

	// iterate over triplets
	for idx := len(triplets) - 1; 0 <= idx; idx-- {
		triplet := triplets[idx]

		// nothing to do with an empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		switch hundreds {
		case 0:
			// nothing
			break
		case 1:
			words = append(words, "et hundrede")
			break
		default:
			words = append(words, danishUnits[hundreds], "hundrede")
			break
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		if 0 < hundreds {
			words = append(words, "og")
		}

		switch tens {
		case 0:
			mega := danishMegas[idx]
			if mega == "tusind" && units == 1 {
				words = append(words, "et")
			} else {
				words = append(words, danishUnits[units])
			}
			break
		case 1:
			words = append(words, danishTeens[units])
			break
		default:
			if 0 < units {
				word := fmt.Sprintf("%sog%s", danishUnits[units], danishTens[tens])
				words = append(words, word)
			} else {
				words = append(words, danishTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		isPlural := 1 < triplet
		if mega := danishMegas[idx]; mega != "" {
			if isPlural || mega == "tusind" {
				words = append(words, mega)
			} else {
				words = append(words, mega+"er")
			}
		}
	}

	return strings.Join(words, " ")
}
