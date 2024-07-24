package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

type TextSentiment string

var (
	positiveSentiment TextSentiment = "positive"
	neutralSentiment  TextSentiment = "neutral"
	negativeSentiment TextSentiment = "negative"
)

type FileAnalyzer struct {
	// WordCount represents the number of words that the file contains
	WordCount int

	// LineCount represents the number of line which the file contains
	LineCount int

	// CharacterCount represents the total number of characters in the file
	CharacterCount int

	// Sentiment analysis determines the overall sentiment of the text (positive, negative, neutral).
	Sentiment TextSentiment

	// FrequentWord represents words which is frequently used in file
	FrequentWord []string

	// LongestWord represents the word which is the larger in the file
	LongestWord []string

	// ShortestWord represents the word which is the shorter in the file
	ShortestWord []string

	// Language represents the language in which the file is written
	Language string
}

func (fa *FileAnalyzer) ReadFile(filePath string) error {
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

		fa.WordCount += len(words)
	}

	return nil
}
