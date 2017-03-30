package guid

import (
	crypto "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"
)

const (
	tsBits   = 34
	tsOffset = 64 - tsBits
	maxTs    = 1<<tsBits - 1
	maxRn    = 1<<tsOffset - 1
)

var puidSource rand.Source

func init() {
	buf := make([]byte, 8)
	if _, err := crypto.Read(buf); err != nil {
		puidSource = rand.NewSource(time.Now().UnixNano())
	} else {
		n := int64(binary.BigEndian.Uint64(buf))
		puidSource = rand.NewSource(n)
	}
}

// PUID is a 64-bit pseudo unique identifier
// The generated IDs are meant to be unique but - unlike GUIDs -
// there is a tiny chance for collisions
type PUID uint64

// NextPUID creates a new pseudo unique identifier
func NextPUID() PUID {
	ts := uint64(time.Now().Unix()) % maxTs
	rn := uint64(puidSource.Int63()) % maxRn
	return PUID(ts<<tsOffset | rn)
}

// CreatedAt extract the timestap at which the PUID was created
func (p PUID) CreatedAt() time.Time {
	return time.Unix(int64(p>>tsOffset), 0)
}
