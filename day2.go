package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var count int64
var dcount int64

func isValidWithDampener(line string, wg *sync.WaitGroup) {
    defer wg.Done()
    c := strings.Split(strings.TrimSpace(line), " ")
    incr := 0
    var p int
    for i, v := range c {
        n, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal("Could not convert to int")
            os.Exit(1)
        }
        if i != 0 {
            if n == p || Abs(n, p) > 3 {
                atomic.AddInt64(&dcount, -1)
                break;
            } else if n > p {
                if incr == 0 {
                    incr = 1
                } else if incr == -1 {
                    atomic.AddInt64(&dcount, -1)
                    break;
                }
            } else {
                if n < p {
                    if incr == 0 {
                        incr = -1
                    } else if incr == 1 {
                        atomic.AddInt64(&dcount, -1)
                        break;
                    }
                }
            }
        }
        p = n
    }
    atomic.AddInt64(&dcount, 1)
}

func isValid(line string, wg *sync.WaitGroup) {
    defer wg.Done()
    c := strings.Split(strings.TrimSpace(line), " ")
    incr := 0
    var p int
    for i, v := range c {
        n, err := strconv.Atoi(v)
        if err != nil {
            log.Fatal("Could not convert to int")
            os.Exit(1)
        }
        if i != 0 {
            if n == p || Abs(n, p) > 3 {
                atomic.AddInt64(&count, -1)
                break;
            } else if n > p {
                if incr == 0 {
                    incr = 1
                } else if incr == -1 {
                    atomic.AddInt64(&count, -1)
                    break;
                }
            } else {
                if n < p {
                    if incr == 0 {
                        incr = -1
                    } else if incr == 1 {
                        atomic.AddInt64(&count, -1)
                        break;
                    }
                }
            }
        }
        p = n
    }
    atomic.AddInt64(&count, 1)
}

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
    wg := &sync.WaitGroup{}
    for scanner.Scan() {
        wg.Add(2)
        go isValid(scanner.Text(), wg)
        go isValidWithDampener(scanner.Text(), wg)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("Error scanning file")
        os.Exit(1)
    }
    wg.Wait()
}

func main () {
    var path string
    fmt.Scanln(&path)
    readFile(path)
    fmt.Println("Number of valid arrays ",count)
    fmt.Println("Number of valid arrays with a dampener ",dcount)
}
