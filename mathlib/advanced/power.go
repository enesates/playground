package advanced

import (
    "errors"
)

func multiply(x int, y int) int {
    return x * y
}

func Power(base, exp int) (int, error) {
    if exp < 0 {
        return 0, errors.New("invalid number")
    }
    if exp == 0 {
        return 1, nil
    }
    if exp == 1 {
        return base, nil
    }

    mul := base

    for i := 0; i < exp-1; i++ {
        mul = multiply(mul, base)
    }

    return mul, nil
}
