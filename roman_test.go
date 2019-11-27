package ntw

import (
	"fmt"
	"testing"
)

func ExampleIntegerToRoman() {
	fmt.Println(IntegerToRoman(42))
	// Output: XLII
}

func TestIntegerToRoman(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:    "-I",
		0:     "nulla",
		1:     "I",
		2:     "II",
		3:     "III",
		4:     "IV",
		5:     "V",
		6:     "VI",
		7:     "VII",
		8:     "VIII",
		9:     "IX",
		10:    "X",
		11:    "XI",
		12:    "XII",
		13:    "XIII",
		14:    "XIV",
		15:    "XV",
		16:    "XVI",
		17:    "XVII",
		18:    "XVIII",
		19:    "XIX",
		20:    "XX",
		21:    "XXI",
		40:    "XL",
		49:    "IL",
		50:    "L",
		60:    "LX",
		70:    "LXX",
		80:    "LXXX",
		90:    "XC",
		99:    "IC",
		100:   "C",
		101:   "CI",
		111:   "CXI",
		120:   "CXX",
		121:   "CXXI",
		499:   "ID",
		500:   "D",
		501:   "DI",
		900:   "CM",
		909:   "CMIX",
		919:   "CMXIX",
		990:   "XM",
		999:   "IM",
		1000:  "M",
		2000:  "MM",
		3000:  "MMM",
		3888:  "MMMDCCCLXXXVIII",
		5000:  "MMMMM",
		9999:  "MMMMMMMMMIM",
		10000: "too big number",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			output := IntegerToRoman(input)
			if expectedOutput != output {
				t.Fatalf("Expected %q, got %q.", expectedOutput, output)
			}
		})
	}
}

func TestIntegerToRomanUnicode(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:    "-Ⅰ",
		0:     "nulla",
		1:     "Ⅰ",
		2:     "Ⅱ",
		3:     "Ⅲ",
		4:     "Ⅳ",
		5:     "Ⅴ",
		6:     "Ⅵ",
		7:     "Ⅶ",
		8:     "Ⅷ",
		9:     "Ⅸ",
		10:    "Ⅹ",
		11:    "Ⅺ",
		12:    "Ⅻ",
		13:    "ⅩⅢ",
		14:    "ⅩⅣ",
		15:    "ⅩⅤ",
		16:    "ⅩⅥ",
		17:    "ⅩⅦ",
		18:    "ⅩⅧ",
		19:    "ⅩⅨ",
		20:    "ⅩⅩ",
		21:    "ⅩⅪ",
		40:    "ⅩⅬ",
		49:    "ⅠⅬ",
		50:    "Ⅼ",
		60:    "ⅬⅩ",
		70:    "ⅬⅩⅩ",
		80:    "ⅬⅩⅩⅩ",
		90:    "ⅩⅭ",
		99:    "ⅠⅭ",
		100:   "Ⅽ",
		101:   "ⅭⅠ",
		111:   "ⅭⅪ",
		120:   "ⅭⅩⅩ",
		121:   "ⅭⅩⅪ",
		499:   "ⅠⅮ",
		500:   "Ⅾ",
		501:   "ⅮⅠ",
		900:   "ⅭⅯ",
		909:   "ⅭⅯⅨ",
		919:   "ⅭⅯⅩⅨ",
		990:   "ⅩⅯ",
		999:   "ⅠⅯ",
		1000:  "Ⅿ",
		2000:  "ⅯⅯ",
		3000:  "ⅯⅯⅯ",
		3888:  "ⅯⅯⅯⅮⅭⅭⅭⅬⅩⅩⅩⅧ",
		10000: "too big number",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			output := IntegerToRomanUnicode(input)
			if expectedOutput != output {
				t.Fatalf("Expected %q, got %q.", expectedOutput, output)
			}
		})
	}
}
