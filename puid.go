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
	puidRnB = 18
	puidInB = 64 - puidTsB - puidRnB

	puidTsO = puidRnB + puidInB
	puidRnO = puidInB

	puidMaxTS = 1<<puidTsB - 1
	puidMaxRN = 1<<puidRnB - 1
	puidMaxIN = 1<<puidInB - 1

	puidNumSlots = 16
)

// PUID is a 64-bit pseudo unique identifier
// The generated IDs are meant to be unique but - unlike GUIDs -
// there is a tiny chance for collisions
type PUID uint64

// CreatedAt extract the timestamp at which the PUID was created
func (p PUID) CreatedAt() time.Time {
	return time.Unix(int64(p>>puidTsO), 0)
}

var puidDefault = NewPUIDSource()

// NextPUID creates a new pseudo unique identifier
func NextPUID() PUID { return puidDefault.Next() }

// PUIDSource is a factory for PUIDs
type PUIDSource interface {
	// Next returns a PUID
	Next() PUID
}

// NewPUIDSource builds a new PUID source
func NewPUIDSource() PUIDSource {
	var n int64
	buf := make([]byte, 8)
	if _, err := crypto.Read(buf); err != nil {
		n = time.Now().UnixNano()
	} else {
		n = int64(binary.BigEndian.Uint64(buf))
	}

	src := new(puidSource)
	for i := 0; i < puidNumSlots; i++ {
		src.slots[i].Source = rand.NewSource(n)
	}
	return src
}

type puidSource struct {
	inc   uint64
	slots [puidNumSlots]struct {
		rand.Source
		sync.Mutex
	}
}

func (p *puidSource) Next() PUID {
	inc := atomic.AddUint64(&p.inc, 1)
	slot := p.slots[inc%puidNumSlots]

	slot.Lock()
	rnn := slot.Int63()
	slot.Unlock()

	ts := uint64(time.Now().Unix()) % puidMaxTS
	rn := uint64(rnn) % puidMaxRN
	in := inc % puidMaxIN

	return PUID(ts<<puidTsO | rn<<puidRnO | in)
}
