package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var loc1 []int
var loc2 []int

func ReadFile(path string) {
    if len(path) == 0 {
        log.Fatal("Invalid file path")
        os.Exit(1)
    }

    f, err := os.Open(path)
    if err != nil {
        log.Fatal("Could not open file ",path)
        os.Exit(1)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        c := strings.Split(strings.TrimSpace(scanner.Text()), "   ")
        el1, err1 := strconv.Atoi(c[0])
        el2, err2 := strconv.Atoi(c[1])
        if err1 != nil || err2 != nil {
            log.Fatal("Error converting value in file to int")
            os.Exit(1)
        }
        loc1 = append(loc1, el1)
        loc2 = append(loc2, el2)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("Error scanning file")
        os.Exit(1)
    }
}

func abs(num1 int, num2 int) int {
    if num1 == num2 {
        return 0
    } else if num1 > num2 {
        return num1 - num2
    } else {
        return num2 - num1
    }
}

func GetDistance() int {
    var distance int
    sort.Ints(loc1)
    sort.Ints(loc2)
    for i := 0; i < len(loc1) - 1; i++ {
        distance += abs(loc1[i], loc2[i])
    }
    return distance
}

func main () {
    var path string
    fmt.Scanln(&path)
    loc1 = make([]int, 0)
    loc2 = make([]int, 0)
    ReadFile(path)
    dist := GetDistance()
    fmt.Println("Total distance is : ", dist)
}
