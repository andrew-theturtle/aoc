package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("inputs/d3.txt")

	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				grid[image.Point{x, y}] = r
			}
		}
	}

	part1, part2 := 0, 0
	parts := map[image.Point][]int{}
	for y, s := range strings.Fields(string(input)) {
		for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1) {
			bounds := []image.Point{}
			for x := m[0]; x < m[1]; x++ {
				for _, d := range []image.Point{
                    {0, -1}, {0, 1}, 
				} {
					bounds = append(bounds, image.Point{x, y}.Add(d))
				}
			}

            bounds = append(bounds, 
                image.Point{m[0]-1, y},
                image.Point{m[1], y},
                image.Point{m[0]-1, y - 1},
                image.Point{m[1], y - 1},
                image.Point{m[0]-1, y + 1},
                image.Point{m[1], y + 1},
            )

			n, _ := strconv.Atoi(s[m[0]:m[1]])
			for _, p := range bounds {
				if _, ok := grid[p]; ok {
					parts[p] = append(parts[p], n)
					part1 += n
                    break
				}
			}
		}
	}

	for p, ns := range parts {
		if grid[p] == '*' && len(ns) == 2 {
			part2 += ns[0] * ns[1]
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
