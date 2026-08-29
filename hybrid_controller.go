package main

import "fmt"

type DynamicScheduler struct {
    state int
}

func (s *DynamicScheduler) flush_monitor(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*3) % 997
    }
    return acc
}

func main() {
    obj := &DynamicScheduler{state: 3}
    fmt.Println(obj.flush_monitor(3))
}
