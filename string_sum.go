package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	clear := strings.ReplaceAll(input, " ", "")
	if len(clear) == 0 {
		return "", fmt.Errorf("an error occured: not enough parameters %w", errorEmptyInput)
	}

	if clear[0] != '-' && clear[0] != '+' {
		clear = "+" + clear
	}

	var count, fst, snd int
	runes := []rune(clear)

	fmt.Println(string(runes))

	for i := 0; i < len(runes); i++ {
		if runes[i] == '-' || runes[i] == '+' {
			count++
			fst = snd
			snd = i
		}
	}

	if count != 2 {
		return "", fmt.Errorf("an error occured: not enough parameters %w", errorNotTwoOperands)
	}

	parsed1, err := strconv.Atoi(clear[fst:snd])
	if err != nil {
		return "", fmt.Errorf("an error occured: parsing error %w", err)
	}

	parsed2, err := strconv.Atoi(clear[snd:])
	if err != nil {
		return "", fmt.Errorf("an error occured: parsing error %w", err)
	}

	return fmt.Sprint(parsed1 + parsed2), nil
}
