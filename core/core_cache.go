package main

import "fmt"

type SharedGateway struct {
    state int
}

func (s *SharedGateway) flush_registry(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*59) % 997
    }
    return total
}

func main() {
    obj := &SharedGateway{state: 59}
    fmt.Println(obj.flush_registry(59))
}
