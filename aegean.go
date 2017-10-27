package ntw

func init() {
	// register the language
	Languages["aegean"] = Language{
		Name:    "Aegean",
		Aliases: []string{"aegean"},
		Flag:    "",

		IntegerToWords: IntegerToAegean,
	}
}

// IntegerToAegean converts an integer to Aegean words
func IntegerToAegean(input int) string {
	if input < 0 || input == 0 || input >= 100000 {
		return ""
	}

	output := ""
	if i := input / 10000 % 10; i > 0 {
		output += string([]rune(" 𐄫𐄬𐄭𐄮𐄯𐄰𐄱𐄲𐄳")[i])
	}
	if i := input / 1000 % 10; i > 0 {
		output += string([]rune(" 𐄢𐄣𐄤𐄥𐄦𐄧𐄨𐄩𐄪")[i])
	}
	if i := input / 100 % 10; i > 0 {
		output += string([]rune(" 𐄙𐄚𐄛𐄜𐄝𐄞𐄟𐄠𐄡")[i])
	}
	if i := input / 10 % 10; i > 0 {
		output += string([]rune(" 𐄐𐄑𐄒𐄓𐄔𐄕𐄖𐄗𐄘")[i])
	}
	if i := input % 10; i > 0 {
		output += string([]rune(" 𐄇𐄈𐄉𐄊𐄋𐄌𐄍𐄎𐄏")[i])
	}

	return output
}
