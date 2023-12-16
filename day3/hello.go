package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "strconv"
)

type saveIndexStruct struct{
    numIndex []int
    rangeIndex []int
}

func part1(scanner bufio.Scanner) {
    currRowIndex := 0

    accum := 0

    previousRow := ""
    previousRowMatchIndex := []saveIndexStruct{}
    
    symbolRe := regexp.MustCompile("[~!@#$%^&*\\(\\)_+=\\-`\\]\\[\\}\\{;:'\",\\<\\>\\/\\?\\|]+")
    
    for scanner.Scan() {
        txt := scanner.Text()

        rowLastIndex := len(txt) - 1

        numRe := regexp.MustCompile("[0-9]+")
        numValueIndex := numRe.FindAllStringIndex(txt, -1)

        remainingNumValueIndex := []saveIndexStruct{}

        for _, num := range numValueIndex {
            startIndex := num[0]
            endIndex := num[1]

            // limit of string. should surround the string
            var rangeStart int = startIndex
            var rangeEnd int = endIndex
            if startIndex > 0 {
                rangeStart -= 1
            }
            if endIndex <= rowLastIndex {
                rangeEnd +=  1
            }


            // if before and after the number is a symbol, then add the number to the accum
            if symbolRe.MatchString(txt[rangeStart:rangeEnd]) {
                numVal, err := strconv.Atoi(txt[num[0]:num[1]]);
                if err == nil {
                    accum += numVal
                }

                continue
            }


            // check if the number is surrounded by symbols at the top in the previous row
            if previousRow != "" && symbolRe.MatchString(previousRow[rangeStart:rangeEnd]) {
                numVal, err := strconv.Atoi(txt[num[0]:num[1]]);
                if err == nil {
                    accum += numVal
                }
                continue
            }


            // if not surrounded by left and right or top, then save the index for the next row
            // to wait for the next row to check if it is surrounded by symbols at the bottom
            rangeSlice := []int{rangeStart, rangeEnd}
            structToAppend := saveIndexStruct{numIndex: num, rangeIndex: rangeSlice}
            remainingNumValueIndex = append(remainingNumValueIndex, structToAppend)
        }

        // check if the number is surrounded by symbols at the bottom in the previous row
        if previousRow != "" {
            for _, prevRowMatch := range previousRowMatchIndex {
                if symbolRe.MatchString(txt[prevRowMatch.rangeIndex[0]:prevRowMatch.rangeIndex[1]]) {
                    numVal, err := strconv.Atoi(previousRow[prevRowMatch.numIndex[0]:prevRowMatch.numIndex[1]]);
                    if err == nil {
                        accum += numVal
                    }
                }
            }
        }

        previousRowMatchIndex = remainingNumValueIndex
        previousRow = txt
        currRowIndex += 1
    }
}

/*
func part2(scanner bufio.Scanner) {
    for scanner.Scan() {
        txt := scanner.Text()
    }
}
*/

func main() {
    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    // part1(*scanner)
    solution1(*scanner)
    // part2(*scanner)
}
