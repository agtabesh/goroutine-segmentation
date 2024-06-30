package main

import (
        "sync"
)

func ProcessWithMutex(data []int, wg *sync.WaitGroup, size int) {
        var mu sync.Mutex
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(start int) {
                        defer wg.Done()
                        for j := start; j < start+size/10; j++ {
                                mu.Lock()
                                data[j] = j * 2
                                mu.Unlock()
                        }
                }(i * (size / 10))
        }
}
