package main

import (
    "bufio"
    "fmt"
    "os"
)

func countLines(dir string) (int, int, int, error) {
    var allLinesCnt int
    var emptyLinesCnt int

    file, err := os.Open(dir)
    if err != nil {
        return 0, 0, 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        allLinesCnt++

        if len(line) == 0 {
            emptyLinesCnt++
        }
    }

    return allLinesCnt, emptyLinesCnt, allLinesCnt - emptyLinesCnt, nil
}

func main() {
    readDir := "./files/big.log"

    allLinesCnt, emptyLinesCnt, linesWithDataCnt, err := countLines(readDir)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Number of lines: ", allLinesCnt)
    fmt.Println("Number of empty lines: ", emptyLinesCnt)
    fmt.Println("Number of lines with data: ", linesWithDataCnt)
}
