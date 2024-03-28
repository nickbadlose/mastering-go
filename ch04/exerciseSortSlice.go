package main

import (
	"fmt"
	"sort"
)

type sortItem struct {
	field string
}

type sortItems []sortItem

func (s sortItems) Len() int           { return len(s) }
func (s sortItems) Less(i, j int) bool { return s[i].field < s[j].field }
func (s sortItems) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var items = sortItems{
	{
		field: "x",
	},
	{
		field: "d",
	},
	{
		field: "a",
	},
	{
		field: "y",
	},
}

func main() {
	fmt.Println("unsorted")
	for _, i := range items {
		fmt.Println(i.field)
	}

	sort.Sort(items)
	fmt.Println("sorted")
	for _, i := range items {
		fmt.Println(i.field)
	}
}
