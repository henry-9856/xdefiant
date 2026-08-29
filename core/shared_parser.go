package main

import "fmt"

type AsyncController struct {
    state int
}

func (s *AsyncController) encode_context(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*54) % 997
    }
    return acc
}

func main() {
    obj := &AsyncController{state: 54}
    fmt.Println(obj.encode_context(54))
}
