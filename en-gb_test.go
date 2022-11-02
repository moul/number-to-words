package ntw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleIntegerToEnGb() {
	fmt.Println(IntegerToEnGb(42))
	// Output: forty-two
}

func TestIntegerToEnGb(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "minus one",
		0:             "zero",
		1:             "one",
		9:             "nine",
		10:            "ten",
		11:            "eleven",
		19:            "nineteen",
		20:            "twenty",
		21:            "twenty-one",
		80:            "eighty",
		90:            "ninety",
		99:            "ninety-nine",
		100:           "one hundred",
		101:           "one hundred and one",
		111:           "one hundred and eleven",
		120:           "one hundred and twenty",
		121:           "one hundred and twenty-one",
		900:           "nine hundred",
		909:           "nine hundred and nine",
		919:           "nine hundred and nineteen",
		990:           "nine hundred and ninety",
		999:           "nine hundred and ninety-nine",
		1000:          "one thousand",
		2000:          "two thousand",
		4000:          "four thousand",
		5000:          "five thousand",
		11000:         "eleven thousand",
		21000:         "twenty-one thousand",
		999000:        "nine hundred and ninety-nine thousand",
		999999:        "nine hundred and ninety-nine thousand, nine hundred and ninety-nine",
		1000000:       "one million",
		2000000:       "two million",
		4000000:       "four million",
		5000000:       "five million",
		100100100:     "one hundred million, one hundred thousand, one hundred",
		500500500:     "five hundred million, five hundred thousand, five hundred",
		606606606:     "six hundred and six million, six hundred and six thousand, six hundred and six",
		999000000:     "nine hundred and ninety-nine million",
		999000999:     "nine hundred and ninety-nine million, nine hundred and ninety-nine",
		999999000:     "nine hundred and ninety-nine million, nine hundred and ninety-nine thousand",
		999999999:     "nine hundred and ninety-nine million, nine hundred and ninety-nine thousand, nine hundred and ninety-nine",
		1174315110:    "one billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and ten",
		1174315119:    "one billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and nineteen",
		15174315119:   "fifteen billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and nineteen",
		35174315119:   "thirty-five billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and nineteen",
		935174315119:  "nine hundred and thirty-five billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and nineteen",
		2935174315119: "two trillion, nine hundred and thirty-five billion, one hundred and seventy-four million, three hundred and fifteen thousand, one hundred and nineteen",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, IntegerToEnGb(input))
		})
	}
}
