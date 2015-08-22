package guid

import (
	"hash/crc64"
	"math/rand"
	"net"
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
	ifaces, err := net.Interfaces()
	if err != nil {
		puidSource = rand.NewSource(time.Now().UnixNano())
	} else {
		hash := crc64.New(crc64.MakeTable(crc64.ECMA))
		for _, iface := range ifaces {
			hash.Write(iface.HardwareAddr)
		}
		puidSource = rand.NewSource(int64(hash.Sum64()))
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
