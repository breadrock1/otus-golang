package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	//Check if input sentence length is empty -> returns input string back.
	if len(inputString) < 1 {
		return inputString, nil
	}

	runeArray := []rune(inputString)
	prevSymbol := runeArray[0]

	//Check first symbol of input sentence -> mustn't be a digit.
	if unicode.IsDigit(prevSymbol) {
		return "", ErrInvalidString
	}

	var deserializedString strings.Builder
	for _, currSymbol := range runeArray[1:] {
		if unicode.IsDigit(currSymbol) {
			//Catches case when near symbols are digits => bigger then 9.
			if unicode.IsDigit(prevSymbol) {
				return "", ErrInvalidString
			}
			if unicode.IsLetter(prevSymbol) {
				str := string(currSymbol)
				number, _ := strconv.Atoi(str)
				str = strings.Repeat(string(prevSymbol), number)
				deserializedString.WriteString(str)
			}
		} else if unicode.IsLetter(prevSymbol) {
			str := string(prevSymbol)
			deserializedString.WriteString(str)
		}

		prevSymbol = currSymbol
	}

	//Processing last sentence symbol.
	str := string(prevSymbol)
	if unicode.IsLetter(prevSymbol) {
		deserializedString.WriteString(str)
	}

	return deserializedString.String(), nil
}
