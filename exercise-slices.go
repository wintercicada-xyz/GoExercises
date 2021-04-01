package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var p = make([][]uint8, dy)
	for y, _ := range p {
		p[y] = make([]uint8, dx)
		for x, _ := range p[y] {
			p[y][x] = uint8((x+y)/2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
