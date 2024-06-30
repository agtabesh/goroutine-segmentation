package main

import (
        "sync"
)

func ProcessWithSegmentation(data []int, wg *sync.WaitGroup, size int) {
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(start int) {
                        defer wg.Done()
                        for j := start; j < start+size/10; j++ {
                                data[j] = j * 2
                        }
                }(i * (size / 10))
        }
}
