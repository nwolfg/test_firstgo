package main

import (
    "fmt"
    "time"
)

func main() {
    i := 0
    for i <= 1000 {
        fmt.Printf("%d\n", i)
        time.Sleep(time.Duration(2)*time.Second)
        i += 2
    }
}
