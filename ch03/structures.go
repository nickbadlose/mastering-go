package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

// initialised by Go
func zeroS() Entry {
	return Entry{}
}

// initialised by Go - returns pointer
func zeroPtoS() *Entry { return &Entry{} }

// initialised by the user
func initS(n, s string, y int) Entry {
	if y < 2000 {
		return Entry{Name: n, Surname: s, Year: 2000}
	}

	return Entry{Name: n, Surname: s, Year: y}
}

// initialised by the user - returns pointer
func initPtoS(n, s string, y int) *Entry {
	if y < 2000 {
		return &Entry{Name: n, Surname: s, Year: 2000}
	}

	return &Entry{Name: n, Surname: s, Year: y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)
	s2 := initS("Mihalis", "Tsoukalos", 2020)
	p2 := initPtoS("Mihalis", "Tsoukalos", 2020)
	fmt.Println("s2:", s2, "p2:", *p2)
	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)
	pS := new(Entry)
	fmt.Println("pS:", pS)
}
