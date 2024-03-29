package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	// Check if input sentence length is empty -> returns input string back.
	if len(inputString) < 1 {
		return "", nil
	}

	var deserializedSlice []string
	runeArray := []rune(inputString)
	for offset := len(runeArray) - 1; offset >= 0; offset-- {
		rCurrSymbol := runeArray[offset]
		if offset == 0 {
			if isArabicDigit(rCurrSymbol) {
				return "", ErrInvalidString
			}
			deserializedSlice = append(deserializedSlice, string(rCurrSymbol))
			break
		}

		rPrevSymbol := runeArray[offset-1]
		if isNonDigitSymbol(rCurrSymbol) {
			deserializedSlice = append(deserializedSlice, string(rCurrSymbol))
		}

		if unicode.IsDigit(rCurrSymbol) {
			if !isArabicDigit(rCurrSymbol) {
				deserializedSlice = append(deserializedSlice, string(rCurrSymbol))
				continue
			}

			if isArabicDigit(rPrevSymbol) {
				return "", ErrInvalidString
			}

			repeated := extractSymbolData(rPrevSymbol, rCurrSymbol)
			deserializedSlice = append(deserializedSlice, repeated)

			offset--
		}
	}

	resultString := revertDeserializedData(deserializedSlice)
	return resultString, nil
}

func isArabicDigit(runeData rune) bool {
	if unicode.IsDigit(runeData) {
		_, isError := strconv.Atoi(string(runeData))
		return isError == nil
	}
	return false
}

func isNonDigitSymbol(runeData rune) bool {
	return unicode.IsSymbol(runeData) || unicode.IsLetter(runeData)
}

func revertDeserializedData(sliceData []string) string {
	reversedSlice := make([]string, len(sliceData))
	for offset := len(sliceData) - 1; offset >= 0; offset-- {
		currSymbol := sliceData[offset]
		newOffset := len(sliceData) - 1 - offset
		reversedSlice[newOffset] = currSymbol
	}
	return strings.Join(reversedSlice, "")
}

func extractSymbolData(runeSymbol rune, decodeSymbol rune) string {
	number, _ := strconv.Atoi(string(decodeSymbol))
	return strings.Repeat(string(runeSymbol), number)
}
