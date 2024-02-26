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
	sort.Slice(uniqWords, func(i, j int) bool {
		a, b := uniqWords[i], uniqWords[j]
		if a.Value == b.Value {
			return a.Key < b.Key
		}
		return a.Value > b.Value
	})

	topWords := make([]string, 0)
	for i := 0; i < len(uniqWords) && i < MaxTopSize; i++ {
		topWords = append(topWords, uniqWords[i].Key)
	}
	return topWords
}
