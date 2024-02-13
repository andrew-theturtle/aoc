package main

import (
    "fmt"
    "strings"
    "os"
    "slices"
)

type Hand struct {
    Cards string
    Bid int
}

func main() {
    data, _ := os.ReadFile("inputs/d7.txt")

    hands := []Hand{}
    for _, s := range strings.Split(strings.TrimSpace(string(data)), "\n") {
        h := Hand{}
        fmt.Sscanf(s, "%s %d", &h.Cards, &h.Bid)
        hands = append(hands, h)
    }

    winning := func(joker bool) int {
        slices.SortFunc(hands, func(a, b Hand) int {
            return cmp(a.Cards, b.Cards, joker)
        })
        res := 0
        for i, h := range hands {
            res += h.Bid * (i + 1)
        }
        return res
    }

    fmt.Println(winning(false))
    fmt.Println(winning(true))
}

func cmp(a, b string, joker bool) int {
    jCards, replaces := "J", "AEKDQCJBTA"
    if joker {
        jCards, replaces = "23456789TQKA", "AEKDQCJ0TA"
    }

    // return the string representing the order of the type of the cards
    // i.e five of kinds has the highest order, hence having value "6",
    // we are sorting the cards ascendingly
    cardType := func (cards string) string {
        typeToOrder := map[int]string{5: "0", 7: "1", 9: "2", 11: "3", 13: "4", 17: "5", 25: "6"}
        k := 0

        // try to replace the joker with all possible cards
        for _, j := range strings.Split(jCards, "") {
            newCards := strings.ReplaceAll(cards, "J", j)
            t := 0
            for _, s := range newCards {
                t += strings.Count(newCards, string(s))
            }
            k = max(k, t)
        }
        return typeToOrder[k]
    }

    // replace cards with characters representing their order in the deck
    // so that we can compare them lexicographically, such as A -> E > K -> D
    modifiedA := strings.NewReplacer(strings.Split(replaces, "")...).Replace(a)
    modifiedB := strings.NewReplacer(strings.Split(replaces, "")...).Replace(b)

    return strings.Compare(
            cardType(a) + modifiedA,
            cardType(b) + modifiedB,
        )
}
