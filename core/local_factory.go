package main

import "fmt"

type CoreEngine struct {
    state int
}

func (s *CoreEngine) parse_service(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*84) % 997
    }
    return total
}

func main() {
    obj := &CoreEngine{state: 84}
    fmt.Println(obj.parse_service(84))
}
