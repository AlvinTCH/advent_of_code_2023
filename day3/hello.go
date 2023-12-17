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


type surroundStruct struct {
    nums []int
    rangeIndex []int
}

func part2(scanner bufio.Scanner) {
    accum := 0

    previousRow := ""
    previousRowMatchIndex := []surroundStruct{}

    for scanner.Scan() {
        txt := scanner.Text()

        rowLastIndex := len(txt) - 1

        starRe := regexp.MustCompile("\\*")
        starValueIndex := starRe.FindAllStringIndex(txt, -1)

        numRe := regexp.MustCompile("[0-9]+")
        numValueIndex := numRe.FindAllStringIndex(txt, -1)

        currRowSurroundStore := []surroundStruct{}

        for _, starInd := range starValueIndex {
            starPosRangeStart := starInd[0]
            starPosRangeEnd := starInd[1]

            if starPosRangeStart > 0 {
                starPosRangeStart -= 1
            }

            if starPosRangeEnd <= rowLastIndex {
                starPosRangeEnd +=  1
            }

            // numbers surrounding the *
            surroundInt := []int{}
            

            // check if any number is on top of the star
            if previousRow != "" {
                previousRowMatchNumIndex := numRe.FindAllStringIndex(previousRow, -1)
                // there are nums on top of the star
                if numRe.MatchString(previousRow[starPosRangeStart:starPosRangeEnd]) {
                    // find the full number in the previous row
                    for _, prevRowNum := range previousRowMatchNumIndex {
                        // if start of number is within the range of the star  (3 digits)
                        if prevRowNum[0] >= starPosRangeStart && prevRowNum[0] < starPosRangeEnd ||
                            // if end of the number is within the range of the star (3 digits)
                            prevRowNum[1] > starPosRangeStart && prevRowNum[1] < starPosRangeEnd ||
                            // if number is encompassed the range of the star (> 4 digits)
                            prevRowNum[0] <= starPosRangeStart && prevRowNum[1] > starPosRangeEnd ||
                            // if number is encompassed by the range of the star (< 3 digits)
                            prevRowNum[0] >= starPosRangeStart && prevRowNum[1] < starPosRangeEnd {
                            numVal, err := strconv.Atoi(previousRow[prevRowNum[0]:prevRowNum[1]]);
                            if err == nil {
                                surroundInt = append(surroundInt, numVal)
                            }
                        }
                    }
                }
            }

            // if before and after the * is a number, then save the number to check for top and bottom
            if numRe.MatchString(txt[starPosRangeStart:starPosRangeEnd]) {
                for _, numInd := range numValueIndex {
                    // star is after the number
                    // or star is before the number
                    if (starInd[0] == numInd[1]) || (starInd[0] == numInd[0] - 1) {
                        numVal, err := strconv.Atoi(txt[numInd[0]:numInd[1]]);
                        if err == nil {
                            surroundInt = append(surroundInt, numVal)
                        }
                    }
                }
            }

            // save and check if any number is at the bottom of the star in the next row
            if len(surroundInt) <= 2 {
                data := surroundStruct{ nums: surroundInt, rangeIndex: []int{starPosRangeStart, starPosRangeEnd} }
                currRowSurroundStore = append(currRowSurroundStore, data)
            }
        }
        
        // check if the number is surrounded by symbols at the bottom in the previous row
        for _, prevRowMatch := range previousRowMatchIndex {
            prevRowMatchNums := prevRowMatch.nums

            // check if the number is at the bottom of the star
            if numRe.MatchString(txt[prevRowMatch.rangeIndex[0]:prevRowMatch.rangeIndex[1]]) {
                for _, currRowNum := range numValueIndex {
                    // if number is within the range of the star

                    // if start of number is within the range of the star  (3 digits)
                    if currRowNum[0] >= prevRowMatch.rangeIndex[0] && currRowNum[0] < prevRowMatch.rangeIndex[1] ||
                        // if end of the number is within the range of the star (3 digits)
                        currRowNum[1] > prevRowMatch.rangeIndex[0] && currRowNum[1] < prevRowMatch.rangeIndex[1] ||
                        // if number is encompassed the range of the star (> 4 digits)
                        currRowNum[0] <= prevRowMatch.rangeIndex[0] && currRowNum[1] > prevRowMatch.rangeIndex[1] ||
                        // if number is encompassed by the range of the star (< 3 digits)
                        currRowNum[0] >= prevRowMatch.rangeIndex[0] && currRowNum[1] < prevRowMatch.rangeIndex[1] {
                        numVal, err := strconv.Atoi(txt[currRowNum[0]:currRowNum[1]]);
                        if err == nil {
                            prevRowMatchNums = append(prevRowMatchNums, numVal)
                        }
                    }
                }
            }

            // add to sum only if star is surrounded by 2 numbers
            if len(prevRowMatchNums) == 2 {
                accum += prevRowMatchNums[0] * prevRowMatchNums[1]
            }
        }


        previousRowMatchIndex = currRowSurroundStore
        previousRow = txt
    }

    fmt.Println(accum)
}


type Schematic struct {
	partCandidates [][]Position
	symbols        [][]Position
	lineCount      int
}

type Position struct {
	lineRange []int
	value     string
}

func (p *Position) isAdjacent(otherPos *Position) bool {
	return p.lineRange[0] > otherPos.lineRange[0]-2 && p.lineRange[1] < otherPos.lineRange[1]+2
}

func (s *Schematic) gears() ([]int, int) {
	var gearRatios []int
	var ratioSum int

	for i, symbolLine := range s.symbols {
		for _, symbol := range symbolLine {
			adjacentParts := make([]int, 0)

			for _, candidateLine := range s.partCandidates[max(0, i-1):min(s.lineCount, i+2)] {
				for _, candidate := range candidateLine {
					if symbol.isAdjacent(&candidate) {
						intval, err := strconv.Atoi(candidate.value)
						if err != nil {
							log.Fatal(err)
						}
						adjacentParts = append(adjacentParts, intval)
					}
				}
			}
			if len(adjacentParts) == 2 {
				ratio := adjacentParts[0] * adjacentParts[1]
				gearRatios = append(gearRatios, ratio)
				ratioSum += ratio
			}
		}
	}

	return gearRatios, ratioSum
}

func parseSchematic(lines []string) *Schematic {
	schematic := Schematic{
		lineCount:      len(lines),
		partCandidates: make([][]Position, len(lines)),
		symbols:        make([][]Position, len(lines)),
	}

	reNum := regexp.MustCompile("(\\d+)")
	reSym := regexp.MustCompile("([^\\d.]+)")

	for i, line := range lines {
		matches := reNum.FindAllStringIndex(line, -1)
		matchesSym := reSym.FindAllStringIndex(line, -1)

		for _, candidate := range matches {
			part := Position{
				lineRange: candidate,
				value:     line[candidate[0]:candidate[1]],
			}
			if len(schematic.partCandidates) <= i {
				schematic.partCandidates[i] = make([]Position, len(matches))
			}
			schematic.partCandidates[i] = append(schematic.partCandidates[i], part)
		}
		for _, sym := range matchesSym {
			symbol := Position{
				lineRange: sym,
				value:     line[sym[0]:sym[1]],
			}
			schematic.symbols[i] = append(schematic.symbols[i], symbol)
		}
	}
	return &schematic
}

func main() {
    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    // part1(*scanner)
    part2(*scanner) // 81997870


    file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")

	schematic := parseSchematic(lines)
	_, ratioSum := schematic.gears()

	fmt.Println(ratioSum)
}
