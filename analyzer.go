package main

type TextSentiment string

var (
	positiveSentiment 	TextSentiment = "positive"
	neutralSentiment 	TextSentiment = "neutral"
	negativeSentiment 	TextSentiment = "negative"
)

type FileAnalyzer struct {
	// WordCount represents the number of words that the file contains
	WordCount 		int

	// LineCount represents the number of line which the file contains
	LineCount 		int

	// CharacterCount represents the total number of characters in the file
	CharacterCount 	int

	// Sentiment analysis determines the overall sentiment of the text (positive, negative, neutral).
	Sentiment		TextSentiment

	// FrequentWord represents words which is frequently used in file
	FrequentWord 	[]string

	// LongestWord represents the word which is the larger in the file
	LongestWord 	[]string

	// ShortestWord represents the word which is the shorter in the file
	ShortestWord 	[]string

	// Language represents the language in which the file is written
	Language 		string
}
