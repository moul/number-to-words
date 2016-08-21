package ntw

import "strings"

func reduceIfSuperiorLoop(words *[]string, input *int, quantity int, word string) {
	for *input >= quantity {
		*input -= quantity
		*words = append(*words, word)
	}
}

// IntegerToRoman converts an integer to Roman words
func IntegerToRoman(input int) string {
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "-")
		input *= -1
	}

	if input == 0 {
		return "nulla"
	}

	// I  V   X   L    C    D     M
	// 1  4  10  50  100  500  1000

	reduceIfSuperiorLoop(&words, &input, 1000, "M")
	reduceIfSuperiorLoop(&words, &input, 999, "IM")
	reduceIfSuperiorLoop(&words, &input, 995, "VM")
	reduceIfSuperiorLoop(&words, &input, 990, "XM")
	reduceIfSuperiorLoop(&words, &input, 950, "LM")
	reduceIfSuperiorLoop(&words, &input, 900, "CM")

	reduceIfSuperiorLoop(&words, &input, 500, "D")
	reduceIfSuperiorLoop(&words, &input, 499, "ID")
	reduceIfSuperiorLoop(&words, &input, 495, "VD")
	reduceIfSuperiorLoop(&words, &input, 490, "XD")
	reduceIfSuperiorLoop(&words, &input, 450, "LD")
	reduceIfSuperiorLoop(&words, &input, 400, "CD")

	reduceIfSuperiorLoop(&words, &input, 100, "C")
	reduceIfSuperiorLoop(&words, &input, 99, "IC")
	reduceIfSuperiorLoop(&words, &input, 95, "VC")
	reduceIfSuperiorLoop(&words, &input, 90, "XC")

	reduceIfSuperiorLoop(&words, &input, 50, "L")
	reduceIfSuperiorLoop(&words, &input, 49, "IL")
	reduceIfSuperiorLoop(&words, &input, 45, "VL")
	reduceIfSuperiorLoop(&words, &input, 40, "XL")

	reduceIfSuperiorLoop(&words, &input, 10, "X")
	reduceIfSuperiorLoop(&words, &input, 9, "IX")

	reduceIfSuperiorLoop(&words, &input, 5, "V")
	reduceIfSuperiorLoop(&words, &input, 4, "IV")

	reduceIfSuperiorLoop(&words, &input, 1, "I")

	return strings.Join(words, "")
}
