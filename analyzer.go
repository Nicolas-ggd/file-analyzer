package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

type FileAnalyzer struct {
	// WordCount represents the number of words that the file contains
	WordCount int

	// LineCount represents the number of line which the file contains
	LineCount int

	// CharacterCount represents the total number of characters in the file
	CharacterCount int

	// FrequentWord represents words which is frequently used in file
	FrequentWord map[string]int

	// LongestWord represents the word which is the larger in the file
	LongestWord string

	// ShortestWord represents the word which is the shorter in the file
	ShortestWord string

	// Language represents the language in which the file is written
	Language string
}

func NewFileAnalyzer() *FileAnalyzer {
	return &FileAnalyzer{
		FrequentWord: make(map[string]int),
	}
}

func (fa *FileAnalyzer) ReadFile(filePath string) error {
	if fa.FrequentWord == nil {
		fa.FrequentWord = make(map[string]int)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fa.LineCount++
		fa.CharacterCount += len(line)

		words := strings.FieldsFunc(line, func(c rune) bool {
			return unicode.IsSpace(c) || unicode.IsPunct(c)
		})

		for _, word := range words {
			fa.longestWord(word)
			_, ok := fa.FrequentWord[word]
			if ok {
				fa.FrequentWord[word]++
			} else {
				fa.FrequentWord[word] = 1
			}
		}

		fa.WordCount += len(words)
	}

	return nil
}

func (fa *FileAnalyzer) longestWord(s string) (best string) {
	for _, words := range strings.Split(s, ", ") {
		if len(words) > len(best) {
			best = words
			fa.LongestWord = best
		}
	}

	return
}
