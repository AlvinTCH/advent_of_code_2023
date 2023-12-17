package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "slices"
    "strings"
    "strconv"
)

type textMatchStruct struct {
    matchCount int
    cardId string
}

// string to get card id, and to count the number
// of matching winning numbers
func textToMatchCount (txt* string) textMatchStruct {
    txtSplit := strings.Split(*txt, ": ")
    numsData := strings.Split(txtSplit[1], " | ")
    
    winningNumbers := strings.Split(numsData[0], " ")
    cardNumbers := strings.Split(numsData[1], " ")

    numRe := regexp.MustCompile("[0-9]+")
    cardId := numRe.FindString(txtSplit[0])

    matchCount := 0
    for _, winNum := range winningNumbers {
        if winNum == "" {
            continue
        }
        if slices.Contains(cardNumbers, winNum) {
            matchCount += 1
        }
    }

    return textMatchStruct{matchCount, cardId}
}

// part 1, find the total points
func part1(scanner bufio.Scanner) {
    // max of 10 winning numbers per card. just save this as such to reduct calc
    multiplier := []int{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

    totalPoints := 0

    for scanner.Scan() {
        txt := scanner.Text()

        totalPoints += multiplier[textToMatchCount(&txt).matchCount]
    }

    fmt.Println(totalPoints)
}

// part 2, find the total number of cards
func part2(scanner bufio.Scanner) {
    total := 0
    matchAdditionalCards := []int{}

    for scanner.Scan() {
        txt := scanner.Text()
        matchData := textToMatchCount(&txt)

        cardId := matchData.cardId
        matchCount := matchData.matchCount

        cardIdNum, err := strconv.Atoi(cardId);
        if err == nil {
            cardIdIndex := cardIdNum - 1

            // init current index count to 0 if current index does not exist in the slice
            if len(matchAdditionalCards) - 1 < cardIdIndex {
                matchAdditionalCards = append(matchAdditionalCards, 0)
            }

            // add the number of cards that are alrdy in the current index
            totalCardsToAdd := matchAdditionalCards[cardIdIndex] + 1

            // for each current matches, add 1 + number of cards in the current index to the next
            // x indexes based on the match count
            for i := 1; i <= matchCount; i++ {
                additionalIndex := cardIdIndex + i

                if len(matchAdditionalCards) - 1 < additionalIndex {
                    matchAdditionalCards = append(matchAdditionalCards, totalCardsToAdd)
                } else {
                    matchAdditionalCards[additionalIndex] += totalCardsToAdd
                }
            }

            matchAdditionalCards[cardIdIndex] += 1
            // add to total for answer
            total += matchAdditionalCards[cardIdIndex]
        }
    }

    fmt.Println(total)
}


func main() {
    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    // part1(*scanner)
    part2(*scanner)
}
