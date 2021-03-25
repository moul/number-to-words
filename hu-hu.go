package ntw

import (
	"strings"
)

func init() {
	// register the language
	Languages["hu-hu"] = Language{
		Name:    "Hungarian Language",
		Aliases: []string{"hu", "magyar", "hungarian"},
		Flag:    "🇭🇺",

		IntegerToWords: IntegerToHuHu,
	}
}

// IntegerToHuHu converts an integer to Hungarian words
func IntegerToHuHu(input int) string {
	var huMegas = []string{"száz", "ezer", "millió", "milliárd", "trillió", "kvadtrillió", "kvintrillió", "szextrillió", "szeptrillió", "oktrillió", "nonillió", "decillió", "undecillió", "duodecillió", "tridecillió", "kvattuordecillió"}
	var huUnits = []string{"", "egy", "kettő", "három", "négy", "öt", "hat", "hét", "nyolc", "kilenc"}
	var huTens = []string{"", "tíz", "húsz", "harminc", "negyven", "ötven", "hatvan", "hetven", "nyolcvan", "kilencven"}

	var buf strings.Builder
	if input < 0 {
		buf.WriteString("mínusz ")
		input *= -1
	}

	//log.Printf("Input: %d\n", input)
	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "zéró"
	}

	groupSep := "-"
	if input < 2000 {
		groupSep = ""
	}
	empty := true
	A := func(ss ...string) {
		empty = false
		for _, s := range ss {
			buf.WriteString(s)
		}
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}
		if !empty {
			buf.WriteString(groupSep)
		}

		// three-digits
		thousands := triplet / 1000 % 10
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		//log.Printf("%d. thousands:%d, hundreds:%d, Tens:%d, Units:%d\n", idx, thousands, hundreds, tens, units)
		if thousands != 0 {
			A(huUnits[thousands], huMegas[1])
		}
		if hundreds != 0 {
			A(huUnits[hundreds], huMegas[0])
		}

		if units == 0 {
			if tens != 0 {
				A(huTens[tens])
			}
		} else {
			if tens != 0 {
				switch tens {
				case 1:
					A("tizen")
				case 2:
					A("huszon")
				default:
					A(huTens[tens])
				}
			}
			A(huUnits[units])
		}

		if idx > 0 {
			if mega := huMegas[idx]; mega != "" {
				A(mega)
			}
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return buf.String()
}
