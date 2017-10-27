package ntw

import (
	"fmt"
	"strings"
)

var portugueseMegasSingular = []string{"", "mil", "milhão", "mil milhões", "bilião"}
var portugueseMegasPlural = []string{"", "mil", "milhões", "mil milhões", "bilhões"}
var portugueseUnits = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseHundreds = []string{"", "cem", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos", "cento"}
var portugueseTens = []string{"", "dez", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
var portugueseTeens = []string{"dez", "onze", "doze", "treze", "catorze", "quinze", "dezasseis", "dezasete", "dezoito", "dezanove"}

func IntegerToPortuguesePT(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}

	triplets := integerToTriplets(input)
	switch {
	case len(triplets) == 0:
		return "zero"
	}
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
		if hundreds > 0 && units == 0 && tens == 0 {
			var word string
			if idx == 0 && len(words) != 0 {
				word = fmt.Sprintf("e %s", portugueseHundreds[hundreds])
			} else {
				word = fmt.Sprintf("%s", portugueseHundreds[hundreds])
			}
			words = append(words, word)
		} else if hundreds > 0 {
			if hundreds == 1 {
				hundreds = 10
			}
			word := fmt.Sprintf("%s e", portugueseHundreds[hundreds])
			words = append(words, word)

		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, portugueseUnits[units])
		case 1:
			word := fmt.Sprintf("%s", portugueseTeens[units])
			words = append(words, word)
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s e %s", portugueseTens[tens], portugueseUnits[units])
				words = append(words, word)
			} else {

				word := fmt.Sprintf("%s", portugueseTens[tens])
				words = append(words, word)
			}
			break
		}
	tripletEnd:
		switch triplet {
		case 1:
			if mega := portugueseMegasSingular[idx]; mega != "" {
				if idx == 4 {
					sum := 0
					for i := 0; i < len(triplets)-1; i++ {
						sum += triplets[i]
					}
					if sum == 0 {

					} else {
						words = append(words, "um")
					}

				} else if idx == 1 && portugueseUnits[idx] == words[0] {
					words = words[1:]
				}
				word := fmt.Sprintf("%s", mega)
				words = append(words, word)
			}
			break
		default:
			if mega := portugueseMegasPlural[idx]; mega != "" {
				if idx == 1 && portugueseUnits[idx] == words[0] {
					words = words[1:]
				}
				word := fmt.Sprintf("%s", mega)
				words = append(words, word)
			}
			break
		}
	}

	return strings.Join(words, " ")

}
