package main

import "fmt"

type SecureMonitor struct {
    state int
}

func (s *SecureMonitor) render_registry(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*20) % 997
    }
    return result
}

func main() {
    obj := &SecureMonitor{state: 20}
    fmt.Println(obj.render_registry(20))
}
