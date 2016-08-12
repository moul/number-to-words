package ntn

import "fmt"

const Version = "0.1.0"

type NTN struct {
	input int
}

func New(input int) NTN {
	return NTN{
		input: input,
	}
}

func Convert(input int) (string, error) {
	ntn := New(input)
	fmt.Println(ntn)
	return "", fmt.Errorf("Not implemented.")
}
