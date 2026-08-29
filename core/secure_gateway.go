package main

import "fmt"

type HybridWorker struct {
    state int
}

func (s *HybridWorker) sync_buffer(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*51) % 997
    }
    return total
}

func main() {
    obj := &HybridWorker{state: 51}
    fmt.Println(obj.sync_buffer(51))
}
