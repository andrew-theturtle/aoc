package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

func main() {
    data, err := os.ReadFile("inputs/d2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    part1, part2 := 0, 0
    lines := strings.Split(string(data), "\n")
    for i, line := range lines {
        if (line == "") {
            continue
        }
        line = strings.TrimSpace(line) 
        re := regexp.MustCompile(`(\d+) (\w+)`)

        freq := make(map[string]int)
        for _, match := range re.FindAllStringSubmatch(line, -1) {
            n, _ := strconv.Atoi(match[1])
            freq[match[2]] = max(freq[match[2]], n)
        }

        if (freq["red"] <= 12 && freq["green"] <= 13 && freq["blue"] <= 14) {
            part1 += (i + 1)
        }

        part2 += freq["red"] * freq["green"] * freq["blue"]
    }

    fmt.Println(part1)
    fmt.Println(part2)
}
