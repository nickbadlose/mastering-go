package main

import "fmt"

func concatArrayToSlice(a1, a2 [2]string) []string {
	s := make([]string, 0, 4)
	for i := range a1 {
		s = append(s, a1[i])
	}
	for i := range a2 {
		s = append(s, a2[i])
	}

	return s
}

func concatArray(a1, a2 [2]string) [4]string {
	a := [4]string{}
	for i := range a1 {
		a[i] = a1[i]
	}
	for i := range a2 {
		a[i+2] = a2[i]
	}

	return a
}

func concatSlicesToArray(s1, s2 []string) [4]string {
	a := [4]string{}
	for i := range s1[:2] {
		a[i] = s1[i]
	}
	for i := range s2[:2] {
		a[i+2] = s2[i]
	}

	return a
}

func main() {
	a1 := [2]string{"A", "B"}
	a2 := [2]string{"C", "D"}

	s := concatArrayToSlice(a1, a2)
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	a := concatArray(a1, a2)
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	s1 := []string{"A", "B"}
	s2 := []string{"C", "D"}

	a = concatSlicesToArray(s1, s2)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
