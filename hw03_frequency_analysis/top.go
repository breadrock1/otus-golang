package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const MaxTopSize = 10

type UniqWord struct {
	Key   string
	Value int
}

func Top10(inputText string) []string {
	wordsFreq := computeWordsFrequency(&inputText)
	uniqWords := extractUniqWords(wordsFreq)
	topWords := make([]string, 0)
	if len(uniqWords) > 0 {
		topWords = extractTopArray(uniqWords)
	}

	return topWords
}

func computeWordsFrequency(inputText *string) *map[string]int {
	topWordsMap := make(map[string]int)
	for _, word := range splitToWords(inputText) {
		if strings.EqualFold("", word) {
			continue
		}
		topWordsMap[word]++
	}
	return &topWordsMap
}

func splitToWords(inputText *string) []string {
	reSlice := regexp.MustCompile(`\s+`)
	sentenceWords := reSlice.Split(*inputText, -1)
	return sentenceWords
}

func extractUniqWords(wordsFreq *map[string]int) []UniqWord {
	uniqueWords := make([]UniqWord, 0, len(*wordsFreq))
	for key, value := range *wordsFreq {
		currWord := UniqWord{key, value}
		uniqueWords = append(uniqueWords, currWord)
	}
	return uniqueWords
}

func extractTopArray(uniqWords []UniqWord) []string {
	topWords := make([]string, 0, MaxTopSize)

	groupedByFreq := groupByValue(uniqWords)
	groupKeys := getSortedKeys(groupedByFreq)
	for _, freqValue := range groupKeys {
		if len(topWords) >= 10 {
			break
		}
		words := groupedByFreq[freqValue]
		topWords = append(topWords, words...)
	}

	return topWords[:MaxTopSize]
}

func groupByValue(uniqWords []UniqWord) map[int][]string {
	groupedWords := make(map[int][]string)
	for _, uWord := range uniqWords {
		_, isKeyAlreadyExists := groupedWords[uWord.Value]
		if isKeyAlreadyExists {
			groupedWords[uWord.Value] = append(groupedWords[uWord.Value], uWord.Key)
		} else {
			groupedWords[uWord.Value] = []string{uWord.Key}
		}
	}

	for _, group := range groupedWords {
		sort.Strings(group)
	}

	return groupedWords
}

func getSortedKeys(groupedWords map[int][]string) []int {
	keys := make([]int, 0)
	for key := range groupedWords {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}
