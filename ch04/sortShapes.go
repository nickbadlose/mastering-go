package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

const (
	min = 1
	max = 5
)

func rF64() float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type cube struct {
	x float64
}

func (c cube) Vol() float64 { return math.Pow(c.x, 3) }

type cuboid struct {
	x, y, z float64
}

func (c cuboid) Vol() float64 { return c.x * c.y * c.z }

type sphere struct {
	r float64
}

func (s sphere) Vol() float64 { return 4 / 3 * math.Pi * math.Pow(s.r, 3) }

type shapes []Shape3D

func (s shapes) Len() int           { return len(s) }
func (s shapes) Less(i, j int) bool { return s[i].Vol() < s[j].Vol() }
func (s shapes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func PrintShapes(s shapes) {
	for _, v := range s {
		switch v.(type) {
		case cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}
	for i := 0; i < 3; i++ {
		cb := cube{rF64()}
		cbo := cuboid{
			x: rF64(),
			y: rF64(),
			z: rF64(),
		}
		s := sphere{rF64()}

		data = append(data, cb)
		data = append(data, cbo)
		data = append(data, s)
	}

	PrintShapes(data)

	sort.Sort(data)
	PrintShapes(data)

	sort.Sort(sort.Reverse(data))
	PrintShapes(data)
}
