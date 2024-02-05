package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    input, _ := os.ReadFile("inputs/d5.txt")
    splits := strings.Split(strings.TrimSpace(string(input)), "\n\n")

    seedString := strings.Fields(strings.Split(splits[0], ": ")[1])
    seeds := make([]int, len(seedString))
    for i := range seedString {
        val, _ := strconv.Atoi(seedString[i])
        seeds[i] = val
    }

    maps := [][][3]int{}
    for i, s1 := range splits[1:] {
        maps = append(maps, [][3]int{})

        for j, s2 := range strings.Split(strings.Split(s1, ":\n")[1], "\n") {
            maps[i] = append(maps[i], [3]int{})
            fmt.Sscanf(s2, "%d %d %d", &maps[i][j][0], &maps[i][j][1], &maps[i][j][2])
        }
    }

    part1, part2 := math.MaxInt, math.MaxInt
    for i, seed := range seeds {
        part1 = min(part1, calDist(&maps, seed))

        if i % 2 == 0 {
            for s := seeds[i]; s < seeds[i] + seeds[i + 1]; s++ {
                part2 = min(part2, calDist(&maps, s))
            }
        }
    }

    fmt.Println(part1)
    fmt.Println(part2)
}

func calDist(maps *[][][3]int, seed int) int {
    for _, m := range *maps {
        for _, r := range m {
            if seed >= r[1] && seed < r[1] + r[2] {
                seed = seed + r[0] - r[1]
                break
            }
        }
    }
    return seed
}

func calDistWithChannel(maps *[][][3]int, seed int, c chan int) {
    for _, m := range *maps {
        for _, r := range m {
            if seed >= r[1] && seed < r[1] + r[2] {
                seed = seed + r[0] - r[1]
                break
            }
        }
    }
    c <- seed
}
