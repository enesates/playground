package advanced

import "fmt"

func multiply(x int, y int) int {
    return x * y
}

func Power(base, exp int) int {
    if exp < 0 {
        fmt.Println("invalid number")
        return -1
    }
    if exp == 0 {
        return 1
    }
    if exp == 1 {
        return base
    }

    mul := base

    for i := 0; i < exp-1; i++ {
        mul = multiply(mul, base)
    }

    return mul
}

func init() {
    println("advanced package initialized")
}
