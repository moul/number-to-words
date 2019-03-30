package ntw

import (
	"fmt"
	"strings"
)

func init() {
	// register the language
	Languages["pl-pl"] = Language{
		Name:    "Polish",
		Aliases: []string{"pl", "pl-pl", "pl_PL", "polish"},
		Flag:    "🇵🇱",

		IntegerToWords: IntegerToPlPl,
	}
}

func IntegerToPlPl(input int) string {
	var polishMegas = [][]string{
		{"", "", ""},
		{"tysiąc", "tysiące", "tysięcy"},
		{"milion", "miliony", "milionów"},
		{"miliard", "miliardy", "miliardów"},
		{"bilion", "biliony", "bilionów"},
		{"biliard", "biliardy", "biliardów"},
		{"trylion", "tryliony", "trylionów"},
		{"tryliard", "tryliardy", "tryliardów"},
		{"kwadrylion", "kwadryliony", "kwadrylionów"},
		{"kwintylion", "kwintyliony", "kwintylionów"},
		{"sekstylion", "sekstyliony", "sekstylionów"},
		{"septylion", "septyliony", "septylionów"},
		{"oktylion", "oktyliony", "oktylionów"},
		{"nonylion", "nonyliony", "nonylionów"},
		{"decylion", "decyliony", "decylionów"}}
	var polishUnits = []string{"", "jeden", "dwa", "trzy", "cztery", "pięć", "sześć", "siedem", "osiem", "dziewięć"}
	var polishTens = []string{"", "dziesięć", "dwadzieścia", "trzydzieści", "czterdzieści", "pięćdziesiąt", "sześćdziesiąt", "siedemdziesiąt", "osiemdziesiąt", "dziewięćdziesiąt"}
	var polishTeens = []string{"dziesięć", "jedenaście", "dwanaście", "trzynaście", "czternaście", "piętnaście", "szesnaście", "siedemnaście", "osiemnaście", "dziewiętnaście"}
	var polishHundreds = []string{"", "sto", "dwieście", "trzysta", "czterysta", "pięćset", "sześćset", "siedemset", "osiemset", "dziewięćset"}

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "zero"
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
			words = append(words, polishHundreds[hundreds])
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, polishUnits[units])
		case 1:
			words = append(words, polishTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", polishTens[tens], polishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, polishTens[tens])
			}
			break
		}

	tripletEnd:
		if triplet == 1 {
			words = append(words, polishMegas[idx][0])
			continue
		}

		megaIndex := 2
		if units >= 2 && units <= 4 {
			megaIndex = 1
		}

		if mega := polishMegas[idx][megaIndex]; mega != "" {
			words = append(words, mega)
		}
	}

	return strings.TrimSpace(strings.Join(words, " "))
}
