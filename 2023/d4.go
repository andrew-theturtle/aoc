package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    data, err := os.ReadFile("inputs/d4.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    part1, part2 := 0, 0
    counts := []int{}
    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        if len(line) > 0 {
            score, count := solvePart1(line)
            part1 += score
            counts = append(counts, count)
        } 
    }


    counter := []int{}
    for i := 0; i < len(counts) + 1; i++ {
        counter = append(counter, 1)
    }
    for i, count := range counts {
        for j := 1; j <= count; j++ {
            counter[i + 1 + j] += counter[i + 1]
        }
    }

    for i := 1; i < len(counter); i++ {
        part2 += counter[i]
    }

    fmt.Println(part1)
    fmt.Println(part2)
}

func solvePart1(line string) (int, int) {
    winningNumbers := map[int]struct{}{}

    re := regexp.MustCompile(`: ([\d ]+) \| ([\d ]+)`) 
    matches := re.FindStringSubmatch(line)

    for _, num := range strings.Fields(matches[1]) {
        number, _ := strconv.Atoi(num)
        winningNumbers[number] = struct{}{}
    }

    count := 0
    for _, num := range strings.Fields(matches[2]) {
        number, _ := strconv.Atoi(num)
        if _, ok := winningNumbers[number]; ok {
            count++
        }
    }

    if count == 0 {
        return 0, 0
    }

    res := math.Pow(2, float64(count - 1))
    return int(res), count
} 
