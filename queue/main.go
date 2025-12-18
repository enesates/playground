package queue

var q []int

func Push(el int) []int {
    return append(q, el)
}

func Pop() []int {
    return q[:len(q)-1]
}
