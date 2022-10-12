package ntw

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	// register the language
	Languages["en-in"] = Language{
		Name:    "Indian English",
		Aliases: []string{"en", "en-in", "indian", "english"},
		Flag:    "🇮🇳",

		IntegerToWords: IntegerToEnIn,
	}
}

// IntegerToEnUs converts an integer to American English words
func IntegerToEnIn(input int) string {
	var indianMegas = []string{"", "thousand", "lakh", "crore", "arab", "kharab", "neel", "padma", "shankh", "mahashankh"}
	var indianUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var indianTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	var indianTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	var hybrids []int
	hybrids = integerToDHybrid(input)

	
	log.Printf("Hybrids: %v\n", hybrids)

	// zero is a special case
	if len(hybrids) == 0 {
		return "zero"
	}

	// iterate over triplets
	for idx := len(hybrids) - 1; idx >= 0; idx-- {
		hybrid := hybrids[idx]
		//log.Printf("hybrid: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if hybrid == 0 {
			continue
		}

		// three-digits
		hundreds := hybrid / 100 % 10
		tens := hybrid / 10 % 10
		units := hybrid % 10
		//log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds > 0 {
			words = append(words, indianUnits[hundreds], "hundred")
		}

		if tens == 0 && units == 0 {
			goto hybridEnd
		}

		switch tens {
		case 0:
			words = append(words, indianUnits[units])
		case 1:
			words = append(words, indianTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", indianTens[tens], indianUnits[units])
				words = append(words, word)
			} else {
				words = append(words, indianTens[tens])
			}
			break
		}

		hybridEnd:
		// mega
		if mega := indianMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
