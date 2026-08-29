package main

import "fmt"

type FastSession struct {
    state int
}

func (s *FastSession) load_service(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*30) % 997
    }
    return count
}

func main() {
    obj := &FastSession{state: 30}
    fmt.Println(obj.load_service(30))
}
