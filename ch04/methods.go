package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

// Traditional Add function
func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}

// Type method Add
func (a *ar2x2) Add(b ar2x2) {
	for i := range a {
		for j := range a[i] {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

// Type method Subtract
func (a *ar2x2) Subtract(b ar2x2) {
	for i := range a {
		for j := range a[i] {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

// Type method Multiply
func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		panic("Need 8 integers")
	}

	k := [8]int{}
	for i, arg := range os.Args[1:] {
		v, err := strconv.Atoi(arg)
		if err != nil {
			panic(fmt.Sprintf("%s is not a valid integer", arg))
		}
		k[i] = v
	}

	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}

	fmt.Println("Traditional a+b", Add(a, b))
	a.Add(b)
	fmt.Println("a+b", a)
	a.Subtract(a)
	fmt.Println("a-a", a)

	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}

	a.Multiply(b)
	fmt.Println("a*b", a)

	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b.Multiply(a)
	fmt.Println("b*a", b)
}
