package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "strconv"
)

// split seeds value into array
func parseSeedRow (txt* string, seedsRe *regexp.Regexp) []string {
    seeds := seedsRe.ReplaceAllString(*txt, "")
    return strings.Split(seeds, " ")
}

// standard way to init a seed data
func makeSeedMap (seed string) map[string]int {
    seedId, err := strconv.Atoi(seed);
    seedIdentifier := 0
    if err == nil {
        seedIdentifier = seedId
    }

    var seedMap map[string]int
    seedMap = make(map[string]int)

    seedMap["seed"] = seedIdentifier
    seedMap["soil"] = seedIdentifier
    seedMap["fertilizer"] = -1
    seedMap["water"] = -1
    seedMap["light"] = -1
    seedMap["temperature"] = -1
    seedMap["humidity"] = -1
    seedMap["location"] = -1

    return seedMap
}

// standard way to go through input to get the corresponding seed conditions for planting
func standardParseSeeds(scanner bufio.Scanner, seedParser func(txt* string, seedsRe *regexp.Regexp) []map[string]int) {
    var seeds []map[string]int

    previousTarget := ""
    currentTarget := ""
    for scanner.Scan() {
        txt := scanner.Text()

        // for line breaks
        if txt == "" {
            continue
        }

        // detect seeds row
        seedsRe := regexp.MustCompile("^seeds: ")
        if seedsRe.MatchString(txt) {
            seeds = seedParser(&txt, seedsRe)
            currentTarget = "seed"
            continue
        }

        // detect start of maps
        mapRe := regexp.MustCompile("map:$")
        if mapRe.MatchString(txt) {
            targetMapString := strings.Split(mapRe.ReplaceAllString(txt, ""), "-")

            // if not in current map, then set the value to the previous target
            for _, seed := range seeds {
                if seed[currentTarget] == -1 {
                    seed[currentTarget] = seed[previousTarget]
                }
            }

            previousTarget = currentTarget
            currentTarget = strings.TrimSpace(targetMapString[len(targetMapString) - 1])
            continue
        }

        // for each value in map, split to
        // the destination start, source start, and count
        sourceValues := strings.Split(txt, " ")

        // convert to num for easy addition
        dstStart := 0
        dstStartNum, err := strconv.Atoi(sourceValues[0]);
        if err == nil {
            dstStart = dstStartNum
        }

        sourceStart := 0
        sourceStartNum, err := strconv.Atoi(sourceValues[1]);
        if err == nil {
            sourceStart = sourceStartNum
        }

        count := 0
        countNum, err := strconv.Atoi(sourceValues[2]);
        if err == nil {
            count = countNum
        }

        for _, seed := range seeds {
            previousTargetValue := seed[previousTarget]
            // if value is found in map,
            // then set the value to the corresponding destination value in the map
            if previousTargetValue >= sourceStart && previousTargetValue < sourceStart + count {
                seed[currentTarget] = dstStart + (previousTargetValue - sourceStart)
            }
        }
    }

    lowestLoc := 0
    for _, seed := range seeds {
        if seed[currentTarget] == -1 {
            seed[currentTarget] = seed[previousTarget]
        }

        if lowestLoc == 0 || seed["location"] < lowestLoc {
            lowestLoc = seed["location"]
        }
    }

    fmt.Println(lowestLoc)
}

// parse seeds for part 1
func part1SeedsParser (txt* string, seedsRe *regexp.Regexp) []map[string]int {
    var seeds []map[string]int

    
    seedsLs := parseSeedRow(txt, seedsRe)
    for _, seed := range seedsLs {
        parsedSeed := makeSeedMap(seed)
        seeds = append(seeds, parsedSeed)
    }

    return seeds
}

func part1(scanner bufio.Scanner) {
    standardParseSeeds(scanner, part1SeedsParser)
}

// parse seeds for part 1
func part2SeedsParser (txt* string, seedsRe *regexp.Regexp) []map[string]int {
    var seeds []map[string]int
    
    seedsLs := parseSeedRow(txt, seedsRe)

    for i := 0; i < len(seedsLs) / 2; i++ {
        seedNumber, err := strconv.Atoi(seedsLs[i * 2]);
        if err != nil {
            seedNumber = 0
        }
        seedRange, err := strconv.Atoi(seedsLs[(i * 2) + 1]);
        if err != nil {
            seedRange = 0
        }

        for j := seedNumber; j < (seedNumber + seedRange); j++ {
            parsedSeed := makeSeedMap(strconv.Itoa(j))
            seeds = append(seeds, parsedSeed)
        }
    }

    return seeds
}

// part 2, find the total number of cards
func part2(scanner bufio.Scanner) {
    standardParseSeeds(scanner, part2SeedsParser)
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
