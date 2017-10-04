package ntw

import (
	"fmt"
	"strings"
)

var turkishMegs = []string{"", "bin", "milyon", "milyar", "trilyon", "katrilyon", "kentilyon", "sekstilyon", "septilyon", "oktilyon", "nonilyon", "desilyon", "andesilyon", "dodesilyon", "tredesilyon", "katordesilyon"}
var turkishUnits = []string{"", "bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz", "dokuz"}
var turkishTens = []string{"", "on", "yirmi", "otuz", "kırk", "elli", "altmış", "yetmiş", "seksen", "doksan"}
var turkishTeens = []string{"on", "on bir", "on iki", "on üç", "on dört", "on beş", "on altı", "on yedi", "on sekiz", "on dokuz"}

// IntegerToTurkish converts an integer to Turkish words
func IntegerToTurkish(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "eksi")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "sıfır"
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
		//log.Printf("Input: %d, Idx: %d, Hundreds:%d, Tens:%d, Units:%d\n", input, idx, hundreds, tens, units)

		switch hundreds {
		case 0:
			break
		case 1:
			words = append(words, "yüz")
		default:
			words = append(words, turkishUnits[hundreds], "yüz")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			if !(units == 1 && idx == 1) {
				words = append(words, turkishUnits[units])
			}
		case 1:
			words = append(words, turkishTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", turkishTens[tens], turkishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, turkishTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := turkishMegs[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
