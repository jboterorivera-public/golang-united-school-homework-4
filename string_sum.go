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

	errorNotFoundValidCharacters    = errors.New("Not found valid characters for input")
	errorExpressionWithInvalidChars = errors.New("The input expression contains invalid chars")
	errorInputCanNotStartWithPlus   = errors.New("Input can't start with +")
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
	s, err := validateInput(input)
	s2 := s
	firstNegative := false
	isSum := false
	var tokens []string

	if err != nil {
		return "", err
	}

	if strings.HasPrefix(s, "-") {
		s2 = s[1:]
		firstNegative = true
	}

	tokens = strings.Split(s2, "-")

	if len(tokens) != 2 {
		tokens = strings.Split(s2, "+")
		isSum = true
	}

	if len(tokens) != 2 {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	firstValue := tokens[0]
	if firstNegative {
		firstValue = "-" + tokens[0]
	}

	value1, err := strconv.ParseInt(firstValue, 10, 32)
	if err != nil {
		return "", fmt.Errorf("Operand1 error: %v", err.Error())
	}

	value2, err := strconv.ParseInt(tokens[1], 10, 32)
	if err != nil {
		return "", fmt.Errorf("Operand2 error: %v", err.Error())
	}

	if isSum {
		return strconv.FormatInt(value1+value2, 10), nil
	}

	return strconv.FormatInt(value1-value2, 10), nil
}

func validateInput(input string) (r string, err error) {
	r = ""
	s := strings.TrimSpace(input)
	runes := []rune(s)
	totalLength := len(runes)
	totalOperations := 0

	if totalLength == 0 {
		return "", fmt.Errorf(errorEmptyInput.Error())
	}

	if strings.HasPrefix(s, "+") {
		return "", fmt.Errorf(errorInputCanNotStartWithPlus.Error())
	}

	for i := 0; i < totalLength; i++ {
		if !validChar(runes[i]) {
			return "", fmt.Errorf(errorExpressionWithInvalidChars.Error())
		}

		if runes[i] == 32 {
			continue
		}

		if runes[i] == 43 || runes[i] == 45 {
			if i > 0 {
				totalOperations++
			}
		}

		if totalOperations > 1 {
			return "", fmt.Errorf(errorNotTwoOperands.Error())
		}

		r += string(runes[i])
	}

	if totalOperations == 0 {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	if len(r) == 0 {
		return "", fmt.Errorf(errorNotFoundValidCharacters.Error())
	}

	return r, nil
}

func validChar(input rune) bool {
	//32 = whitespace, 43 = +, 45 = -
	if input == 32 || input == 43 || input == 45 {
		return true
	}

	// numbers from 0 to 9
	if input >= 48 && input <= 57 {
		return true
	}

	return false
}
