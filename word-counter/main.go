package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// calling the count function to count the number of words
	// received from the Standard Input and printing it out

	fmt.Println(count(os.Stdin))
}

// function to perform the actual counting

func count(r io.Reader) int {
	// declare a scanner that will read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wordCounter := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wordCounter++
	}

	// return the total
	return wordCounter
}
