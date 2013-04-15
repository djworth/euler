package main

import (
    "fmt"
    "math"
)

const (
    LARGEST = 4000000
)

func IsEven(n float64) bool {

    _, frac := math.Modf(n / 2.0)

    return frac == 0

}

func Fib() func() int {
    x, y, z := 0, 1, 0
    return func() int {
        z, x, y = x, y, x+y
        return z
    }
}

func main() {

    total := 0

    f := Fib()
    for {

        i := f()

        if i >= LARGEST {
            break
        }

        if IsEven(float64(i)) {
            total += i
        }

    }

    fmt.Println(total)
}
