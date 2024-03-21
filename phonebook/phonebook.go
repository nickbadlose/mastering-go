package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct{ Name, Surname, Tel, LastAccess string }

const (
	filepath = "/Users/nick/projects/mastering-go/phonebook/records.data"
)

var (
	data  = []Entry{}
	index = map[string]int{}
)

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func list() {
	for _, entry := range data {
		fmt.Println(entry)
	}
}

func readCSVFile() error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, row := range d {
		data = append(data, Entry{
			Name:       row[0],
			Surname:    row[1],
			Tel:        row[2],
			LastAccess: row[3],
		})
	}

	return nil
}

func saveCSVFile() error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, entry := range data {
		err := w.Write([]string{entry.Name, entry.Surname, entry.Tel, entry.LastAccess})
		if err != nil {
			return err
		}
	}
	w.Flush()

	return w.Error()
}

func createIndex() {
	index = make(map[string]int, len(data))
	for i, entry := range data {
		index[entry.Tel] = i
	}
}

func matchTel(t string) bool {
	rgx := regexp.MustCompile(`^\d{10}$`)
	return rgx.Match([]byte(t))
}

func insert(name, surname, tel string) error {
	// If it already exists, do not add it
	_, ok := index[tel]
	if ok {
		return fmt.Errorf("%s already exists", tel)
	}

	data = append(data, Entry{
		Name:       name,
		Surname:    surname,
		Tel:        tel,
		LastAccess: strconv.FormatInt(time.Now().Unix(), 10),
	})

	// Update the index
	createIndex()

	return saveCSVFile()
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}

	data = append(data[:i], data[i+1:]...)

	// Update the index - key does not exist anymore
	// This is pointless since the index is created at the start of each run,
	// also you would need to update the keys of all the entries after the deleted one in the slice. But hey nvm for now.
	delete(index, key)

	return saveCSVFile()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	_, err := os.Stat(filepath)
	// If error is not nil, it means that the file does not exist
	if err != nil {
		fmt.Println("Creating", filepath)
		f, err := os.Create(filepath)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(filepath)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(filepath, "not a regular file!")
		return
	}

	err = readCSVFile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data has %d entries.\n", len(data))

	createIndex()

	// Differentiate between the commands
	switch arguments[1] {
	// The search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		result := search(t)
		if result == nil {
			fmt.Println("Entry not found:", t)
			return
		}
		fmt.Println(*result)
	// The list command
	case "list":
		list()
	// Response to anything that is not a match
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
		}

		err := insert(arguments[2], arguments[3], t)
		if err != nil {
			fmt.Println("could not insert record:", arguments[2:])
			fmt.Println(err)
			return
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("USage: delete Number")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Not a valid option")
	}
}
