package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/sync/semaphore"
	"io"
	"os"
	"regexp"
	"runtime"
	"sync/atomic"
)

var (
	countWords      bool
	countBytes      bool
	countCharacters bool
	countLines      bool

	maxProcs = runtime.GOMAXPROCS(0)

	workers = int64(limitWorkers(5))

	sem = semaphore.NewWeighted(workers)
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

type counter struct{ lineCount, wordCount, characterCount, byteCount int }

func limitWorkers(n int) int {
	if n < maxProcs {
		return n
	}
	return maxProcs
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

	ctx := context.Background()
	var totalLineCount, totalWordCount, totalCharacterCount, totalByteCount atomic.Int64
	for _, filepath := range filePaths {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			panic(err)
		}

		go func(fp string) {
			defer sem.Release(1)
			fileInfo, err := os.Stat(fp)
			if err != nil {
				panic(err)
			}

			if !fileInfo.Mode().IsRegular() {
				panic("not a regular file")
			}

			lineCount, wordCount, characterCount, byteCount, err := count(fp)
			if err != nil {
				panic(err)
			}

			printCounts(lineCount, wordCount, characterCount, byteCount, fp)
			totalLineCount.Add(int64(lineCount))
			totalWordCount.Add(int64(wordCount))
			totalCharacterCount.Add(int64(characterCount))
			totalByteCount.Add(int64(byteCount))
		}(filepath)
	}

	err := sem.Acquire(ctx, workers)
	if err != nil {
		panic(err)
	}

	if len(filePaths) > 1 {
		printCounts(
			int(totalLineCount.Load()),
			int(totalWordCount.Load()),
			int(totalCharacterCount.Load()),
			int(totalByteCount.Load()),
			"total",
		)
	}

}
