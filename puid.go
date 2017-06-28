package guid

import (
	crypto "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	puidTsB = 34
	puidRnB = 10
	puidInB = 64 - puidTsB - puidRnB

	puidTsO = puidRnB + puidInB
	puidRnO = puidInB

	puidMaxTS = 1<<puidTsB - 1
	puidMaxRN = 1<<puidRnB - 1
	puidMaxIN = 1<<puidInB - 1

	numPUIDSlots = 32
)

var puidInc uint64

var puidSlots [numPUIDSlots]struct {
	rand.Source
	sync.Mutex
}

func init() {
	buf := make([]byte, 8)
	for i := 0; i < numPUIDSlots; i++ {
		var n int64
		if _, err := crypto.Read(buf); err != nil {
			n = time.Now().UnixNano()
		} else {
			n = int64(binary.BigEndian.Uint64(buf))
		}
		puidSlots[i].Source = rand.NewSource(n)
	}
}

// PUID is a 64-bit pseudo unique identifier
// The generated IDs are meant to be unique but - unlike GUIDs -
// there is a tiny chance for collisions
type PUID uint64

// NextPUID creates a new pseudo unique identifier
func NextPUID() PUID {
	inc := atomic.AddUint64(&puidInc, 1)
	slot := puidSlots[inc%numPUIDSlots]

	slot.Lock()
	rnn := slot.Int63()
	slot.Unlock()

	ts := uint64(time.Now().Unix()) % puidMaxTS
	rn := uint64(rnn) % puidMaxRN
	in := inc % puidMaxIN

	return PUID(ts<<puidTsO | rn<<puidRnO | in)
}

// CreatedAt extract the timestamp at which the PUID was created
func (p PUID) CreatedAt() time.Time {
	return time.Unix(int64(p>>puidTsO), 0)
}
