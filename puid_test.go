package guid

import (
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
		n := 40000
		if testing.Short() {
			n = 10000
		}

		set := make(map[PUID]int, n)
		for i := 0; i < n; i++ {
			p := NextPUID()
			set[p]++
		}
		Expect(len(set)).To(Equal(n))
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
