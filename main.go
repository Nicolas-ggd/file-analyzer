package main

import (
	"bufio"
    "fmt"
    "log"
    "os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}

	fa := FileAnalyzer{
		WordCount: 		10,
		LineCount: 		1,
		FrequentWord: 	[]string{"and"},
		LongestWord: 	[]string{"example"},
		ShortestWord: 	[]string{"and"},
		Sentiment: 		positiveSentiment,
	}

	fmt.Printf("%+v\n", fa)
}
