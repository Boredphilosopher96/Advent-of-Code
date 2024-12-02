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

func readFile(path string) {
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

func lowerBound(nums []int, target int) int {
    return sort.Search(len(loc2), func(i int) bool {
        return target <= nums[i] 
    })
}

func upperBound(nums []int, target int) int {
    return sort.Search(len(loc2), func(i int) bool {
        return target < nums[i] 
    })
}

func getSimilarityScore() int {
    var sim int
    sort.Ints(loc2)
    for i := 0; i < len(loc1); i++ {
        low := lowerBound(loc2, loc1[i])
        if low >= 0 {
            sim += loc1[i]*((upperBound(loc2, loc1[i]) - low) )
        }
    }
    return sim
}

func getDistance() int {
    var distance int
    sort.Ints(loc1)
    sort.Ints(loc2)
    for i := 0; i < len(loc1); i++ {
        distance += Abs(loc1[i], loc2[i])
    }
    return distance
}

func main () {
    var path string
    fmt.Scanln(&path)
    loc1 = make([]int, 0)
    loc2 = make([]int, 0)
    readFile(path)
    dist := getDistance()
    fmt.Println("Total distance is : ", dist)
    sim := getSimilarityScore()
    fmt.Println("Total similarity score is : ", sim)
}
