package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["es-es"] = Language{
		Name:    "European Spanish",
		Aliases: []string{"es", "es-es", "es_ES", "spanish"},
		Flag:    "🇪🇸",

		IntegerToWords: IntegerToEsEs,
	}
}

// IntegerToEsEs converts an integer to spanish words
func IntegerToEsEs(input int) string {
	var spanishMegasSingular = []string{"", "mil", "millón", "mil millones", "billón"}
	var spanishMegasPlural = []string{"", "mil", "millones", "mil millones", "billones"}
	var spanishUnitsAdjectives = []string{"", "un", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nove"}
	var spanishUnits = []string{"", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
	var spanishHundreds = []string{"", "ciento", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos"}
	var spanishTens = []string{"", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
	var spanishTeens = []string{"diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve"}
	var spanishTwenties = []string{"veinte", "veintiuno", "veintidós", "veintitrés", "veinticuatro", "veinticinco", "veintiséis", "veintisiete", "veintiocho", "veintinueve"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// special cases
	switch {
	case len(triplets) == 0:
		return "cero"
	case input == 1000:
		words = append(words, "mil")
		goto end
	case input == 1:
		words = append(words, "uno")
		goto end
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
			words = append(words, spanishHundreds[hundreds])
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			if idx > 0 {
				words = append(words, spanishUnitsAdjectives[units])
			} else {
				words = append(words, spanishUnits[units])
			}
		case 1:
			words = append(words, spanishTeens[units])
			break
		case 2:
			if idx > 0 && units == 1 {
				words = append(words, "veintiún")
			} else {
				words = append(words, spanishTwenties[units])
			}
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s y %s", spanishTens[tens], spanishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, spanishTens[tens])
			}
			break
		}

	tripletEnd:
		switch triplet {
		case 0:
			break
		case 1:
			if mega := spanishMegasSingular[idx]; mega != "" {
				words = append(words, mega)
			}
			break
		default:
			if mega := spanishMegasPlural[idx]; mega != "" {
				words = append(words, mega)
			}
			break
		}
	}

end:
	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
