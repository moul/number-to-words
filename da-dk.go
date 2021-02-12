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
		return "nul"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; 0 <= idx; idx-- {
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
		// log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		switch hundreds {
		case 0:
			// Nothing
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
		isPlural := hundreds == 0 && tens == 0 && units == 1
		if mega := danishMegas[idx]; mega != "" {
			if isPlural || mega == "tusind" {
				words = append(words, mega)
			} else {
				words = append(words, mega+"er")
			}
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
