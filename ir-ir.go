package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["ir-ir"] = Language{
		Name:    "Persian",
		Aliases: []string{"ir", "ir-ir", "ir_IR", "persian"},
		Flag:    "🇮🇷",

		IntegerToWords: IntegerToIrIr,
	}
}

func IntegerToEnUs(input int) string {
	var persianMegas = []string{"", "هزار", "میلیون", "میلیارد", "بیلیون", "بیلیارد", "تریلیون", "تریلیارد"}
	var persianUnits = []string{"", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"}
	var persianTens = []string{"", "ده", "بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"}
	var persianTeens = []string{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هجده", "نوزده"}
	var persianHundreds = []string{"", "صد", "دویست", "سیصد", "چهارصد", "پانصد", "ششصد", "هفتصد", "هشتصد", "نهصد"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "منفی")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "صفر"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds > 0 {
			words = append(words, persianHundreds[hundreds])
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, persianUnits[units])
		case 1:
			words = append(words, persianTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s و %s", persianTens[tens], persianUnits[units])
				words = append(words, word)
			} else {
				words = append(words, persianTens[tens])
			}
			break
		}

	tripletEnd:
		if mega := persianMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
