package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

func main() {
    input, _ := os.ReadFile("inputs/d6.txt")
    split := strings.Split(strings.TrimSpace(string(input)), "\n")
    times := strings.Fields(split[0])[1:]
    distances := strings.Fields(split[1])[1:]

    fmt.Println(solve(times, distances))

    combineTime := strings.Join(times, "")
    combineDist := strings.Join(distances, "")
    fmt.Println(solve(
        []string{combineTime},
        []string{combineDist},
    ))
}

func solve(times, distances []string) int {
    res := 1
    for i := range times {
        time, _ := strconv.ParseFloat(times[i], 64)
        dist, _ := strconv.ParseFloat(distances[i], 64)
        dist++
        b := math.Sqrt(math.Pow(time, 2) - 4 * dist)

        count := int(math.Floor((time + b) / 2) - math.Ceil((time - b) / 2) + 1)
        res *= count
    }

    return res
}
