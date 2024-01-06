package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const MAX_TOP_SIZE = 10

type UniqWord struct {
	Key   string
	Value int
}

func Top10(inputText string) []string {
	wordsFreq := computeWordsFrequency(&inputText)
	uniqWords := extractUniqWords(wordsFreq)
	if len(uniqWords) < 1 {
		return []string{}
	}
	topWords := extractTopArray(uniqWords)
	return topWords
}

func computeWordsFrequency(inputText *string) *map[string]int {
	var topWordsMap = make(map[string]int)
	for _, word := range splitToWords(inputText) {
		if strings.EqualFold("", word) {
			continue
		}

		_, isKeyAlreadyExists := topWordsMap[word]
		if isKeyAlreadyExists {
			topWordsMap[word]++
		} else {
			topWordsMap[word] = 1
		}
		fmt.Println("")
	}
	return &topWordsMap
}

func splitToWords(inputText *string) []string {
	var reSlice = regexp.MustCompile("[\\s!?.,:;`'\"(){}\\[\\]]+")
	sentenceWords := reSlice.Split(*inputText, -1)
	sort.Strings(sentenceWords)
	return sentenceWords
}

func extractTopArray(uniqWords []UniqWord) []string {
	sortWordsByFreq(uniqWords)
	test := groupByValue(uniqWords)
	fmt.Println(test)
	topWords := make([]string, 0, MAX_TOP_SIZE)
	for _, uWord := range uniqWords[:MAX_TOP_SIZE] {
		//fmt.Printf("%s ->  %d\n", uWord.Key, uWord.Value)
		topWords = append(topWords, uWord.Key)
	}
	return topWords
}

func groupByValue(uniqWords []UniqWord) map[int][]string {
	var groupedWords = make(map[int][]string)
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

func sortWordsByFreq(uniqWords []UniqWord) {
	sort.Slice(uniqWords, func(srcCmp, dstCmp int) bool {
		return uniqWords[srcCmp].Value > uniqWords[dstCmp].Value
	})
}

func extractUniqWords(wordsFreq *map[string]int) []UniqWord {
	var uniqWords = make([]UniqWord, 0, len(*wordsFreq))
	for key, value := range *wordsFreq {
		currWord := UniqWord{key, value}
		uniqWords = append(uniqWords, currWord)
	}
	return uniqWords
}
