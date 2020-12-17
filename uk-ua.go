package ntw

import (
	"strings"
)

func init() {
	// register the language
	Languages["uk-ua"] = Language{
		Name:    "Ukrainian",
		Aliases: []string{"uk", "uk-ua", "uk_UA", "ukrainian"},
		Flag:    "🇺🇦",

		IntegerToWords: IntegerToUkUa,
	}
}

func ukPlural(n int, words []string) string {
	return plural(n, words)
}

// IntegerToUkUa converts an integer to UK words
func IntegerToUkUa(input int) string {
	var ukUnits = []string{
		"",
		"один",
		"два",
		"три",
		"чотири",
		"п'ять",
		"шість",
		"сім",
		"вісім",
		"дев'ять",
	}
	var ukTeens = []string{
		"десять",
		"одинадцять",
		"дванадцять",
		"тринадцять",
		"чотирнадцять",
		"п'ятнадцять",
		"шістнадцять",
		"сімнадцять",
		"вісімнадцять",
		"дев'ятнадцять",
	}
	var ukTens = []string{
		"",
		"десять",
		"двадцять",
		"тридцять",
		"сорок",
		"п'ятдесят",
		"шістдесят",
		"сімдесят",
		"вісімдесят",
		"дев'яносто",
	}
	var ukHundreds = []string{
		"",
		"сто",
		"двісті",
		"триста",
		"чотириста",
		"п'ятсот",
		"шістсот",
		"сімсот",
		"вісімсот",
		"дев'ятсот",
	}
	var ukMegas = [][]string{
		{"", "", ""},
		{"тисяча", "тисячі", "тисяч"},                    // 10^3
		{"мільйон", "мільйона", "мільйонів"},             // 10^6
		{"мільярд", "мільярда", "мільярдів"},             // 10^9
		{"трильйон", "трильйона", "трильйонів"},          // 10^12
		{"квадрильйон", "квадрильйона", "квадрильйонів"}, // 10^15
		{"квінтильйон", "квінтильйона", "квінтильйонів"}, // 10^18
		{"секстильйон", "секстильйона", "секстильйонів"}, // 10^21
		{"септильйон", "септильйона", "септильйонів"},    // 10^34
		{"октильйон", "октильйона", "октильйонів"},       // 10^27
	}

	var words []string

	if input < 0 {
		words = append(words, "мінус")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "нуль"
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
			words = append(words, ukHundreds[hundreds])
		}

		if tens > 0 || units > 0 {
			switch tens {
			case 0:
				words = append(words, ukOneTwoUnitFix(units, idx, ukUnits))
			case 1:
				words = append(words, ukTeens[units])
				break
			default:
				words = append(words, ukTens[tens])
				if units > 0 {
					words = append(words, ukOneTwoUnitFix(units, idx, ukUnits))
				}
				break
			}
		}

		// mega
		if idx >= 1 && idx < len(ukMegas) {
			mega := ukMegas[idx]
			tens = tens*10 + units
			if len(mega) > 0 {
				words = append(words, ukPlural(tens, mega))
			}
		}
	}

	return strings.Join(words, " ")
}

func ukOneTwoUnitFix(unit, idx int, arr []string) string {
	if idx == 1 {
		switch unit {
		case 1:
			return "одна"
		case 2:
			return "дві"
		}
	}

	return arr[unit]
}
