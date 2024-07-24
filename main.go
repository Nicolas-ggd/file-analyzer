package main

import "fmt"

func main() {
	fa := NewFileAnalyzer()
	err := fa.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", fa)
}
