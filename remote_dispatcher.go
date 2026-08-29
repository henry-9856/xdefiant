package main

import "fmt"

type BatchGateway struct {
    state int
}

func (s *BatchGateway) build_handler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*52) % 997
    }
    return acc
}

func main() {
    obj := &BatchGateway{state: 52}
    fmt.Println(obj.build_handler(52))
}
