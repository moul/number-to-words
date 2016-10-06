package ntw

import (
	"fmt"
	"strings"
)

var spanishMegas = []string{"", "mil", "millon", "billon", "trillon", "cuatrillon", "quintillon", "sextillon", "septillon", "octillon", "nonillon", "decillon", "oncillon", "docillon", "trecillon", "catorcillon"}
var spanishUnits = []string{"", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
var spanishTens = []string{"", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
var spanishTeens = []string{"diez", "once", "doce", "trece", "catorce", "quince", "dieciseis", "diecisiete", "dieciocho", "diecinueve"}

// IntegerToSpanish converts an integer to spanish words
func IntegerToSpanish(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "zero"
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
			words = append(words, spanishUnits[hundreds], "mil")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, spanishUnits[units])
		case 1:
			words = append(words, spanishTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", spanishTens[tens], spanishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, spanishTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := spanishMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
