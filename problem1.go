package main

import (
    "fmt"
    "math"
)

const (
    THREE = 3.0
    FIVE  = 5.0
)

func IsFactor(n, m float64) bool {
    _, frac := math.Modf(n / m)

    //fmt.Println(whole, frac)
    return frac == 0
}

func main() {

    total := 0
    for i := 1; i < 1000; i++ {
        if IsFactor(float64(i), THREE) || IsFactor(float64(i), FIVE) {
            total += i
        }
    }

    fmt.Println(total)
}
