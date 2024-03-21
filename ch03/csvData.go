package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// CSV file read all at once
	// lines data type is [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string, delimiter rune) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	// changing the default field delimiter to tab
	csvWriter.Comma = rune(delimiter)
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		err = csvWriter.Write(temp)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	if csvWriter.Error() != nil {
		return csvWriter.Error()
	}

	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("csvData input output delimiter!")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if len(os.Args[3]) != 1 {
		fmt.Println("delimiter should be one character only!")
		return
	}

	delimiter := os.Args[3][0]

	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// CSV data is read in columns - each line is a slice
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Println(temp)
	}

	err = saveCSVFile(output, rune(delimiter))
	if err != nil {
		fmt.Println(err)
		return
	}
}
