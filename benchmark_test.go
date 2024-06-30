package main

import (
        "sync"
        "testing"
)

const size = 1000000

func BenchmarkWithMutex(b *testing.B) {
        data := make([]int, size)
        var wg sync.WaitGroup

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                ProcessWithMutex(data, &wg, size)
                wg.Wait()
        }
}

func BenchmarkWithSegmentation(b *testing.B) {
        data := make([]int, size)
        var wg sync.WaitGroup

        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                ProcessWithSegmentation(data, &wg, size)
                wg.Wait()
        }
}
