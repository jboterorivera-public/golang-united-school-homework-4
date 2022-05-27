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

type CustomError struct {
	Message string
	Err     error
}

func (c CustomError) Error() string {
	return c.Message
}

func (c CustomError) Unwrap() error {
	return c.Err
}

func StringSum(input string) (output string, err error) {
	inputTrimmed, prefix := processInput(input)

	if prefix+inputTrimmed == "" {
		return "", fmt.Errorf(errorEmptyInput.Error())
	}

	operand1, operand2, operation, errorOperand := getOperands(inputTrimmed)

	if errorOperand != nil {
		return "", fmt.Errorf(errorOperand.Error())
	}

	if operand1 == "" || operand2 == "" {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	value1, errConv := strconv.Atoi(prefix + operand1)
	if errConv != nil {
		return "", customStrConvError(errConv)
	}

	value2, errConv := strconv.Atoi(operand2)
	if errConv != nil {
		return "", customStrConvError(errConv)
	}

	if operation == "+" {
		return strconv.FormatInt(int64(value1)+int64(value2), 10), nil
	}

	return strconv.FormatInt(int64(value1)-int64(value2), 10), nil
}

func processInput(input string) (result string, prefix string) {
	result = strings.TrimSpace(input)

	if strings.HasPrefix(result, "-") {
		prefix = "-"
		result = result[1:]
	}

	return result, prefix
}

func getOperands(input string) (operand1 string, operand2 string, operation string, err error) {
	runes := []rune(input)
	totalLength := len(runes)
	fillOperand2 := false
	totalOperands := 0

	for i := 0; i < totalLength; i++ {
		//32 = whitespace
		if runes[i] == 32 {
			continue
		}

		//43 = +, 45 = -
		if runes[i] == 43 || runes[i] == 45 {
			totalOperands++

			if totalOperands > 1 {
				return "", "", "", errorNotTwoOperands
			}

			operation = string(runes[i])
			fillOperand2 = true
			continue
		}

		if fillOperand2 {
			operand2 += string(runes[i])
			continue
		}

		operand1 += string(runes[i])
	}

	return operand1, operand2, operation, nil
}

func customStrConvError(e error) (c CustomError) {
	c = CustomError{Message: e.Error()}
	c.Err = fmt.Errorf("Error [%w]", e.(*strconv.NumError))

	return c
}
