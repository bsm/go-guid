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
		n := 1000
		if testing.Short() {
			n = 100
		}

		set := make(map[PUID]struct{}, n)
		for i := 0; i < n; i++ {
			p := NextPUID()
			Expect(set).NotTo(HaveKey(p))
			set[p] = struct{}{}
		}
	})

})

// --------------------------------------------------------------------

func BenchmarkNextPUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextPUID()
	}
}
