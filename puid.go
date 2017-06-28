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
	tsBits   = 34
	tsOffset = 64 - tsBits
	maxTs    = 1<<tsBits - 1
	maxRn    = 1<<tsOffset - 1

	numPUIDSlots = 32
)

var puidSlot int64

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
	si := atomic.AddInt64(&puidSlot, 1) % numPUIDSlots
	slot := puidSlots[si]

	slot.Lock()
	n := slot.Int63()
	slot.Unlock()

	ts := uint64(time.Now().Unix()) % maxTs
	rn := uint64(n) % maxRn
	return PUID(ts<<tsOffset | rn)
}

// CreatedAt extract the timestamp at which the PUID was created
func (p PUID) CreatedAt() time.Time {
	return time.Unix(int64(p>>tsOffset), 0)
}
