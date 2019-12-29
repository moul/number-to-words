package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["id-id"] = Language{
		Name:    "Indonesian",
		Aliases: []string{"id", "id-id", "id_ID", "indonesian"},
		Flag:    "🇮🇩",

		IntegerToWords: IntegerToIDID,
	}
}

// IntegerToIDID converts an integer to Indonesian words
func IntegerToIDID(input int) string {
	var indonesianMegas = []string{"", "ribu", "juta", "milyar", "triliun", "kuadriliun", "kuintiliun", "sekstiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "kuatuordesiliun"}
	var indonesianUnits = []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	var indonesianTens = []string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}
	var indonesianTeens = []string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	// log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "nol"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		// log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		// log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds == 1 {
			words = append(words, "seratus")
		} else if hundreds > 0 {
			words = append(words, indonesianUnits[hundreds], "ratus")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, indonesianUnits[units])
		case 1:
			words = append(words, indonesianTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", indonesianTens[tens], indonesianUnits[units])
				words = append(words, word)
			} else {
				words = append(words, indonesianTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := indonesianMegas[idx]; mega != "" {
			// exception for 1000
			if idx == 1 && triplet == 1 {
				words = append(words[0:len(words)-1], "seribu")
			} else {
				words = append(words, mega)
			}
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
