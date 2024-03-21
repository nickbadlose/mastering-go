package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}

	m := make(map[int]string, 2)
	for i := range a1 {
		m[i] = a1[i]
	}

	fmt.Println(a1)
	fmt.Println(m)

	m4 := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}

	s1, s2 := make([]int, 0, 4), make([]string, 0, 4)
	for k, v := range m4 {
		s1 = append(s1, k)
		s2 = append(s2, v)
	}

	fmt.Println(m4)
	fmt.Println(s1)
	fmt.Println(s2)
}
