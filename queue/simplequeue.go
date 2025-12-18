package queue

var q []int

func Push(el int) []int {
    q = append(q, el)
    return q
}

func Pop() []int {
    return q[:len(q)-1]
}

func init() {
    q = make([]int, 0)
}
