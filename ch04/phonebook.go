package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type nameNumber interface {
	name() string
	number() string
}

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

func (e Entry) name() string   { return e.Surname + e.Name }
func (e Entry) number() string { return e.Tel }

type EntryWithAreaCode struct {
	Name       string
	Surname    string
	AreaCode   string
	Tel        string
	LastAccess string
}

func (e EntryWithAreaCode) name() string   { return e.Surname + e.Name }
func (e EntryWithAreaCode) number() string { return e.AreaCode + e.Tel }

// CSVFILE resides in the home directory of the current user
var CSVFILE = "/Users/nick/projects/mastering-go/phonebook/recordsWithAreaCode.data"

type PhoneBook []nameNumber

var data = PhoneBook{}
var index map[string]int

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		var temp nameNumber
		switch len(line) {
		case 4:
			temp = Entry{
				Name:       line[0],
				Surname:    line[1],
				Tel:        line[2],
				LastAccess: line[3],
			}
		case 5:
			temp = EntryWithAreaCode{
				Name:       line[0],
				Surname:    line[1],
				AreaCode:   line[2],
				Tel:        line[3],
				LastAccess: line[4],
			}
		default:
			return errors.New("Unknown File Format!")
		}

		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		var temp []string
		switch v := row.(type) {
		case Entry:
			temp = []string{v.Name, v.Surname, v.Tel, v.LastAccess}
		case EntryWithAreaCode:
			temp = []string{v.Name, v.Surname, v.AreaCode, v.Tel, v.LastAccess}
		default:
			return fmt.Errorf("unsupported data type, %T", v)

		}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, d := range data {
		t, ok := d.(nameNumber)
		if !ok {
			return errors.New("unsupported data type for indexing")
		}
		key := t.number()
		index[key] = i
	}
	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initS(N, S, T string) *Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	// Update the index
	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) interface{} {
	i, ok := index[key]
	if !ok {
		return nil
	}

	switch v := data[i].(type) {
	case Entry:
		v.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
		return &v
	case EntryWithAreaCode:
		v.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
		return &v
	default:
		return fmt.Errorf("unsupported data type, %T", v)
	}
}

func list(reverse bool) {
	if reverse {
		sort.Sort(sort.Reverse(data))
	} else {
		sort.Sort(data)
	}
	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func setCSVFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		CSVFILE = filepath
	}

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", CSVFILE)
	}
	return nil
}

// Implement sort.Interface
func (a PhoneBook) Len() int {
	return len(a)
}

// First based on surname. If they have the same
// surname take into account the name.
func (a PhoneBook) Less(i, j int) bool { return a[i].name() < a[j].name() }

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	err := setCSVFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := initS(arguments[2], arguments[3], t)
		// If it was nil, there was an error
		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
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
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		rv := reflect.ValueOf(temp)
		if rv.Kind() == reflect.Ptr {
			fmt.Println(rv.Elem().Interface())
		} else {
			fmt.Println(temp)
		}
	case "list":
		list(false)
	case "reverse":
		list(true)
	default:
		fmt.Println("Not a valid option")
	}
}
