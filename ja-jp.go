package ntw

func init() {
	// register the language
	Languages["ja-jp"] = Language{
		Name:    "Japanese",
		Aliases: []string{"jp", "ja-jp", "ja_JP", "japanese"},
		Flag:    "🇯🇵",

		IntegerToWords: IntegerToJaJp,
	}
}

func integerToQuadruplets(number int) []int {
	quadruplet := []int{}

	for number > 0 {
		quadruplet = append(quadruplet, number%10000)
		number = number / 10000
	}
	return quadruplet
}

// IntegerToJaJp converts an integer to Japanese words
func IntegerToJaJp(input int) string {
	var japaneseUnits = []string{
		"",
		"一",
		"二",
		"三",
		"四",
		"五",
		"六",
		"七",
		"八",
		"九",
	}

	var japaneseDigits = []string{
		"",
		"万",
		"億",
		"兆",
	}

	//log.Printf("Input: %d\n", input)
	words := ""

	if input < 0 {
		words = "マイナス"
		input *= -1
	}

	// split integer in quadruplets
	quadruplets := integerToQuadruplets(input)
	//log.Printf("Quadruplets: %v\n", quadruplets)

	// zero is a special case
	if len(quadruplets) == 0 {
		return "零"
	}

	// iterate over quadruplet
	for idx := len(quadruplets) - 1; idx >= 0; idx-- {
		quadruplet := quadruplets[idx]

		// nothing todo for empty quadruplet
		if quadruplet == 0 {
			continue
		}

		// four-digits
		thousands := quadruplet / 1000 % 10
		hundreds := quadruplet / 100 % 10
		tens := quadruplet / 10 % 10
		units := quadruplet % 10

		if thousands > 0 {
			if thousands == 1 {
				words += "千"
			} else {
				words += japaneseUnits[thousands] + "千"
			}
		}

		if hundreds > 0 {
			if hundreds == 1 {
				words += "百"
			} else {
				words += japaneseUnits[hundreds] + "百"
			}
		}

		if tens > 0 {
			if tens == 1 {
				words += "十"
			} else {
				words += japaneseUnits[tens] + "十"
			}
		}

		if units > 0 {
			words += japaneseUnits[units]
		}

		if idx > 0 {
			words += japaneseDigits[idx]
		}
	}

	return words
}
