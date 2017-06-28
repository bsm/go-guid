package guid

import (
	"sync"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PUID", func() {

	It("should generate", func() {
		p1, p2 := NextPUID(), NextPUID()
		Expect(p1).NotTo(Equal(p2))
	})

	It("should extract creation time", func() {
		p := NextPUID()
		Expect(p.CreatedAt()).To(BeTemporally("~", time.Now(), time.Second))
	})

	It("should avoid collisions", func() {
		set := make(map[PUID]int)
		mu := new(sync.Mutex)
		wg := new(sync.WaitGroup)
		now := time.Now()

		for i := 0; i < 20; i++ {
			wg.Add(1)

			go func() {
				defer GinkgoRecover()
				defer wg.Done()

				src := NewPUIDSource().(*puidSource)
				for i := 0; i < 5000; i++ {
					p := src.NextAt(now)
					mu.Lock()
					set[p]++
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		Expect(len(set)).To(Equal(100000))
	})

})

// --------------------------------------------------------------------

func BenchmarkNextPUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextPUID()
	}
}

func BenchmarkNextPUID_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NextPUID()
		}
	})
}
