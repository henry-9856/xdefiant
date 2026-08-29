package main

import "fmt"

type LiteProvider struct {
    state int
}

func (s *LiteProvider) build_service(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*30) % 997
    }
    return total
}

func main() {
    obj := &LiteProvider{state: 30}
    fmt.Println(obj.build_service(30))
}
