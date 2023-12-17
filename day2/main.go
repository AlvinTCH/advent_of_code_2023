package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "regexp"
    "fmt"
    "strconv"
    "sort"
)

type cubeStruct struct{
    red int
    green int
    blue int
}

type colorGroup struct {
    red []int
    green []int
    blue []int
}

func findByColor(gameSets []string, numRe *regexp.Regexp) colorGroup {
    colorGroupData := colorGroup{[]int{}, []int{}, []int{}}

    for _, gameSet := range gameSets {
        greenRe := regexp.MustCompile("[0-9]+ green")
        matchedGreen := greenRe.FindAllString(gameSet, -1)

        greenCount := 0
        for _, green := range matchedGreen {
            greenNum, err := strconv.Atoi(numRe.FindString(green));
            if err == nil {
                greenCount += greenNum
            }
        }

        colorGroupData.green = append(colorGroupData.green, greenCount)

        redRe := regexp.MustCompile("[0-9]+ red")
        matchedRed := redRe.FindAllString(gameSet, -1)

        redCount := 0
        for _, red := range matchedRed {
            redNum, err := strconv.Atoi(numRe.FindString(red));
            if err == nil {
                redCount += redNum
            }
        }

        colorGroupData.red = append(colorGroupData.red, redCount)

        blueRe := regexp.MustCompile("[0-9]+ blue")
        matchedBlue := blueRe.FindAllString(gameSet, -1)

        blueCount := 0
        for _, blue := range matchedBlue {
            blueNum, err := strconv.Atoi(numRe.FindString(blue));
            if err == nil {
                blueCount += blueNum
            }
        }

        colorGroupData.blue = append(colorGroupData.blue, blueCount)
    }

    return colorGroupData
}

func part1(scanner bufio.Scanner) {
    cubeCount := cubeStruct{red: 12, green: 13, blue: 14}

    accumulatedId := 0

    for scanner.Scan() {
        txt := scanner.Text()

        gameDataSplit := strings.Split(txt, ": ")
        numRe := regexp.MustCompile("[0-9]+")

        gameSets := strings.Split(gameDataSplit[1], ";")
        setByColor := findByColor(gameSets, numRe)

        validGame := true
        for _, red := range setByColor.red {
            if red > cubeCount.red {
                validGame = false
            }
        }

        for _, green := range setByColor.green {
            if green > cubeCount.green {
                validGame = false
            }
        }

        for _, blue := range setByColor.blue {
            if blue > cubeCount.blue {
                validGame = false
            }
        }

        if !validGame {
            continue
        }
        
        gameId, err := strconv.Atoi(numRe.FindString(gameDataSplit[0]));
        if err == nil {
            accumulatedId += gameId
        }
    }

    fmt.Println(accumulatedId)
}

func part2(scanner bufio.Scanner) {
    accumulated := 0

    for scanner.Scan() {
        txt := scanner.Text()

        gameDataSplit := strings.Split(txt, ": ")
        numRe := regexp.MustCompile("[0-9]+")

        gameSets := strings.Split(gameDataSplit[1], ";")
        setByColor := findByColor(gameSets, numRe)

        sort.Slice(setByColor.red, func(i, j int) bool {
            return setByColor.red[i] > setByColor.red[j]
        })

        sort.Slice(setByColor.blue, func(i, j int) bool {
            return setByColor.blue[i] > setByColor.blue[j]
        })

        sort.Slice(setByColor.green, func(i, j int) bool {
            return setByColor.green[i] > setByColor.green[j]
        })

        highestRed := setByColor.red[0]
        highestBlue := setByColor.blue[0]
        highestGreen := setByColor.green[0]

        accumulated += highestRed * highestBlue * highestGreen
    }

    fmt.Println(accumulated)
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
