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
	danishMegas := []string{"", "tusind", "million", "milliard", "billion", "billiard", "trillion", "trilliard", "kvadrillion", "kvadrilliard", "kvintillion", "kvintilliard", "sekstillion", "sekstilliard", "septillion", "septiliard"}
	danishUnits := []string{"", "en", "to", "tre", "fire", "fem", "seks", "syv", "otte", "ni"}
	danishTens := []string{"", "ti", "tyve", "tredive", "fyrre", "halvtreds", "tres", "halvfjerds", "firs", "halvfems"}
	danishTeens := []string{"ti", "elleve", "tolv", "tretten", "fjorten", "femten", "seksten", "sytten", "atten", "nitten"}

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
		mega := danishMegas[idx]
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		hundredsInWords := hundredsToDaDk(hundreds, danishUnits)
		if hundredsInWords != "" {
			words = append(words, hundredsInWords)
		}

		if tens != 0 || units != 0 {
			if 0 < hundreds {
				words = append(words, "og")
			}

			tensAndUnits := tensAndUnitsToDaDk(mega, tens, units, danishUnits, danishTeens, danishTens)
			words = append(words, tensAndUnits)
		}

		isSingular := triplet == 1
		megas := megasToDaDk(isSingular, mega)
		if megas != "" {
			words = append(words, megas)
		}
	}

	return strings.Join(words, " ")
}

func hundredsToDaDk(hundreds int, danishUnits []string) string {
	switch hundreds {
	case 0:
		return ""
	case 1:
		return "et hundrede"
	default:
		return danishUnits[hundreds] + " hundrede"
	}
}

func tensAndUnitsToDaDk(mega string, tens int, units int, danishUnits []string, danishTeens []string, danishTens []string) string {
	switch tens {
	case 0:
		if mega == "tusind" && units == 1 {
			return "et"
		}
		return danishUnits[units]
	case 1:
		return danishTeens[units]
	default:
		if 0 < units {
			return fmt.Sprintf("%sog%s", danishUnits[units], danishTens[tens])
		}
		return danishTens[tens]
	}
}

func megasToDaDk(isSingular bool, mega string) string {
	if mega != "" {
		if isSingular || mega == "tusind" {
			return mega
		}
		return mega + "er"
	}
	return ""
}
