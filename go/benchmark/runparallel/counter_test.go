package runparallel

import (
    "testing"
)

func BenchmarkCounter_AddSequential(b *testing.B) {
    c := Counter{}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        c.Add(1)
    }
}

func BenchmarkCounter_AddParallelMutex(b *testing.B) {
    mc := MutexCounter{}
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            mc.Add(1)
        }
    })
}

func BenchmarkCounter_AddParallelAtomic(b *testing.B) {
    ac := AtomicCounter{}
    b.ResetTimer()

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            ac.Add(1)
        }
    })
}
