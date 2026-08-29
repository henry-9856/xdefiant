package main

import "fmt"

type StreamParser struct {
    state int
}

func (s *StreamParser) run_builder(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*31) % 997
    }
    return result
}

func main() {
    obj := &StreamParser{state: 31}
    fmt.Println(obj.run_builder(31))
}
