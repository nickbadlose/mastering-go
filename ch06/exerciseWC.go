package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func count(fp string) (lineCount, wordCount, characterCount, byteCount int, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	r := bufio.NewReader(f)
	for {
		line, rErr := r.ReadString('\n')
		if rErr != nil && rErr != io.EOF {
			return 0, 0, 0, 0, rErr
		}

		rgx := regexp.MustCompile("[^\\s]+")
		//rgx := regexp.MustCompile("\\S+") TODO try this
		words := rgx.FindAllString(line, -1)
		wordCount += len(words)
		byteCount += len(line) // can do this since line is essentially a []byte

		// looping through a line gets each individual rune (character), not byte
		for _ = range line {
			characterCount++
		}

		if rErr == io.EOF {
			break
		}
		lineCount++
	}

	return
}

func main() {
	if len(os.Args) != 2 {
		panic("Usage: file")
	}

	filepath := os.Args[1]

	fileInfo, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}

	if fileInfo.Mode() == os.ModeIrregular {
		panic("not a regular file")
	}

	lineCount, wordCount, characterCount, byteCount, err := count(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\t%d\t%d\t%d\t%d\t%s\n", lineCount, wordCount, characterCount, byteCount, filepath)
}
