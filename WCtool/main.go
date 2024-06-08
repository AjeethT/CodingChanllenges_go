package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("WC Tool in go")

	linesFlag := flag.Bool("l", false, "Count of Lines")
	wordsFlag := flag.Bool("w", false, "Count of Words")
	charsFlag := flag.Bool("c", false, "Count of Character")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Use correct command")
		return
	}

	filepath := args[0]
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening file", err.Error())
		return
	}
	defer file.Close()

	isLines, isWords, isChars := true, true, true

	if *linesFlag || *wordsFlag || *charsFlag {
		isLines = *linesFlag
		isWords = *wordsFlag
		isChars = *charsFlag
	}

	lineCount, wordCount, charCount := 0, 0, 0
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error in reading file", err.Error())
			return
		}

		if isLines {
			lineCount++
		}

		if isWords {
			words := strings.Fields(line)
			wordCount += len(words)
		}

		if isChars {
			charCount += len(line)
		}

		if err == io.EOF {
			break
		}
	}

	if isLines {
		fmt.Println("Line Counts: ", lineCount)
	}
	if isWords {
		fmt.Println("Word Counts: ", wordCount)
	}
	if isChars {
		fmt.Println("Character Counts: ", charCount)
	}

}
