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


func findNumInStr(txt* string) int {
    re := regexp.MustCompile("[0-9]+")
    allNums := re.FindAllString(*txt, -1)

    var numVal string = strings.Join(allNums, "")
    var chars string = ""

    if len(numVal) > 0 {
        chars += string(numVal[0])

        if len(numVal) > 1 {
            chars += string(numVal[len(numVal)-1])
        } else {
            chars += string(numVal[0])
        }
    }

    s, err := strconv.Atoi(chars);
    if err == nil {
        return s
    }

    return 0
}

func part1(scanner bufio.Scanner) {
    accum := 0
    for scanner.Scan() {
        txt := scanner.Text()

        accum += findNumInStr(&txt)
    }

    fmt.Println(accum)
}

type indexLsStruct struct {
    val string
    replace string
}

func part2(scanner bufio.Scanner) {
    accum := 0

    // replace all the words with words with a num in it
    var replacerStruct = []indexLsStruct{
        {val: "one", replace: "o1e"},
        {val: "two", replace: "t2o"},
        {val: "three", replace: "t3e"},
        {val: "four", replace: "f4r"},
        {val: "five", replace: "f5e"},
        {val: "six", replace: "s6x"},
        {val: "seven", replace: "s7n"},
        {val: "eight", replace: "e8t"},
        {val: "nine", replace: "n9e"},
    }

    for scanner.Scan() {
        txt := scanner.Text()

        for _, replacer := range replacerStruct {
            txt = strings.ReplaceAll(txt, replacer.val, replacer.replace)
        }

        accum += findNumInStr(&txt)
    }

    fmt.Println("my solution", accum)
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
