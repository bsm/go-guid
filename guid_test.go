package guid

import (
	"bytes"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GUID", func() {

	BeforeEach(func() {
		inc = 0
		hostpid = [5]byte{91, 93, 95, 97, 99}
	})

	It("should increment correctly", func() {
		Expect(New96().Bytes()[9:]).To(Equal([]byte{0, 0, 1}))
		inc = 16777214
		Expect(New96().Bytes()[9:]).To(Equal([]byte{255, 255, 255}))
		Expect(New96().Bytes()[9:]).To(Equal([]byte{0, 0, 0}))
	})

	Describe("96bit", func() {

		It("should generate", func() {
			g1, g2 := New96(), New96()
			Expect(g1).NotTo(Equal(g2))
		})

		It("should extract creation time", func() {
			g := New96()
			Expect(g.CreatedAt()).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should export bytes", func() {
			g := new96at(time.Unix(1414141414, 123456))
			Expect(g.Bytes()).To(Equal([]byte{84, 74, 21, 230, 91, 93, 95, 97, 99, 0, 0, 1}))
		})

		It("should implement io.WriterTo", func() {
			n, err := New96().WriteTo(&bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(12)))
		})

		It("should avoid collisions", func() {
			set := make(map[GUID96]int)
			for i := 0; i < 1000000; i++ {
				set[New96()]++
			}
			Expect(len(set)).To(Equal(1000000))
		})

	})

	Describe("128bit", func() {

		It("should generate", func() {
			g1, g2 := New128(), New128()
			Expect(g1).NotTo(Equal(g2))
		})

		It("should extract creation time", func() {
			g := New128()
			Expect(g.CreatedAt()).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should export bytes", func() {
			g := new128at(time.Unix(1414141414, 123456))
			Expect(g.Bytes()).To(Equal([]byte{0, 0, 0, 0, 84, 74, 21, 230, 91, 93, 95, 97, 99, 0, 0, 1}))
		})

		It("should implement io.WriterTo", func() {
			n, err := New128().WriteTo(&bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(16)))
		})

		It("should avoid collisions", func() {
			set := make(map[GUID128]int)
			for i := 0; i < 1000000; i++ {
				set[New128()]++
			}
			Expect(len(set)).To(Equal(1000000))
		})

	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "github.com/bsm/guid")
}

// --------------------------------------------------------------------

func BenchmarkNew96(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New96()
	}
}

func BenchmarkNew128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New128()
	}
}
