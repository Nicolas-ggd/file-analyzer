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
}
