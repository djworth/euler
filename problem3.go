package main

import (
    "fmt"
    "math"
)

const (
    TARGET = 600851475143
)

func IsPrime(n float64) bool {

    for i := 2.0; i < n; i++ {

        _, frac := math.Modf(n / i)

        if frac != 0 && i != n {
            return false
        }
    }

    return true

}

func main() {

    i := 0
    for {

        if IsPrime(float64(i)) {

        }

        i++
    }

    fmt.Println("Hello")
}
