package ntw

import (
	"math"
	"strings"
)

func init() {
	// register the language
	Languages["sv-se"] = Language{
		Name:    "Swedish",
		Aliases: []string{"sv-se", "sv_SE", "swedish"},
		Flag:    "🇸🇪",

		IntegerToWords: IntegerToSvSe,
	}
}

var (
	swedishMegas = []string{"tusen", "miljon", "miljard", "biljon", "kvadriljon", "kvintiljon"} //"sextillion", "septillion", "octillion", "nonillion", "decillion"
	swedishUnits = []string{"noll", "en", "två", "tre", "fyra", "fem", "sex", "sju", "åtta", "nio", "tio", "elva", "tolv", "tretton", "fjorton", "femton", "sexton", "sjutton", "arton", "nitton"}
	swedishTens  = []string{"", "", "tjugo", "trettio", "fyrtio", "femtio", "sextio", "sjuttio", "åttio", "nittio"}
)

// IntegerToSvSe converts an integer to Swedish words
func IntegerToSvSe(n int) string {
	result := integerToSvSe(n, true, true)
	result = strings.Replace(result, "etttusen", "ettusen", -1)

	return strings.TrimSpace(result)
}

func integerToSvSe(input int, firstCall bool, isSuffix bool) string {

	if input < 0 {
		return "mindre " + integerToSvSe(-input, true, isSuffix)
	}

	if input == 0 && firstCall {
		return "noll"
	} else if input == 0 {
		return ""
	}

	if input < 20 {
		return addGroupSvSe(input, isSuffix)
	}

	if input < 100 {
		return swedishTens[input/10] + addGroupSvSe(input%10, isSuffix)
	}

	if input < 1000 {
		if input/100 == 1 {
			return "etthundra" + addGroupSvSe(input%100, isSuffix)
		}
		return addGroupSvSe(input/100, false) + "hundra" + addGroupSvSe(input%100, isSuffix)
	}

	for i := len(swedishMegas) - 1; i >= 0; i-- {
		if power := int(math.Pow(1000, float64(i+1))); input > power-1 {
			mega := swedishMegas[i]

			if power > 1000 {
				if input/power > 1 {
					mega += "er"
				}
				mega = " " + mega + " "
				return integerToSvSe(input/power, false, isSuffix) + mega + integerToSvSe(input%power, false, isSuffix)
			}

			return integerToSvSe(input/power, false, false) + mega + integerToSvSe(input%power, false, isSuffix)
		}
	}
	return ""
}

func addGroupSvSe(n int, isSuffix bool) string {
	if n == 0 {
		return ""
	}
	if !isSuffix && n == 1 {
		return "ett"
	}
	if n < 20 {
		return swedishUnits[n]
	}
	if n%10 == 0 {
		return swedishTens[n/10]
	}

	if !isSuffix && n%10 == 1 {
		return swedishTens[n/10] + "ett"
	}

	return swedishTens[n/10] + swedishUnits[n%10]
}
