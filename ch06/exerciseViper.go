package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"os"
	"regexp"
)

var (
	countWords      bool
	countBytes      bool
	countCharacters bool
	countLines      bool
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

func getFiles(args []string) []string {
	files := make([]string, 0, 1)
	for _, arg := range args {
		if arg == "--w" || arg == "--c" || arg == "--m" || arg == "--l" {
			continue
		}

		files = append(files, arg)
	}

	return files
}

func printCounts(lineCount, wordCount, characterCount, byteCount int, descriptor string) {
	if countLines {
		fmt.Printf("\t%d", lineCount)
	}

	if countWords {
		fmt.Printf("\t%d", wordCount)
	}

	if countCharacters {
		fmt.Printf("\t%d", characterCount)
	}

	if countBytes {
		fmt.Printf("\t%d", byteCount)
	}

	if !countLines && !countWords && !countCharacters && !countBytes {
		fmt.Printf("\t%d\t%d\t%d\t%d\t%s\n", lineCount, wordCount, characterCount, byteCount, descriptor)
	} else {
		fmt.Printf("\t%s\n", descriptor)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Usage: file")
	}

	pflag.Bool("w", false, "Count words")
	pflag.Bool("c", false, "Count bytes")
	pflag.Bool("m", false, "Count characters")
	pflag.Bool("l", false, "Count lines")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	countWords = viper.GetBool("w")
	countBytes = viper.GetBool("c")
	countCharacters = viper.GetBool("m")
	countLines = viper.GetBool("l")

	filePaths := getFiles(os.Args[1:])

	var totalLineCount, totalWordCount, totalCharacterCount, totalByteCount int
	for _, filepath := range filePaths {
		fileInfo, err := os.Stat(filepath)
		if err != nil {
			panic(err)
		}

		if !fileInfo.Mode().IsRegular() {
			panic("not a regular file")
		}

		lineCount, wordCount, characterCount, byteCount, err := count(filepath)
		if err != nil {
			panic(err)
		}
		totalLineCount += lineCount
		totalWordCount += wordCount
		totalCharacterCount += characterCount
		totalByteCount += byteCount
		printCounts(lineCount, wordCount, characterCount, byteCount, filepath)
	}

	if len(filePaths) > 1 {
		printCounts(totalLineCount, totalWordCount, totalCharacterCount, totalByteCount, "total")
	}
}
