package guid

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"os"
	"sync/atomic"
	"time"
)

// GUID is an interface unifying identifiers of
// different lengths
type GUID interface {
	CreatedAt() time.Time
	WriteTo(w io.Writer) (int64, error)
	Bytes() []byte
}

// --------------------------------------------------------------------

// GUID96 is a 12-byte globally unique identifier
type GUID96 [12]byte

var base96 = GUID96{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// New96 creates a 96bit/12byte global identifier
func New96() GUID96 { return new96at(time.Now()) }

func new96at(ts time.Time) GUID96 {
	bininc := [4]byte{0, 0, 0, 0}
	encoder.PutUint32(bininc[:], nextInc())

	bytes := base96
	encoder.PutUint32(bytes[0:], uint32(ts.Unix()))
	copy(bytes[4:], hostpid[:])
	copy(bytes[9:], bininc[1:])

	return bytes
}

// Bytes returns a byte slice
func (g GUID96) Bytes() []byte { return g[:] }

// WriteTo implements io.WriterTo
func (g GUID96) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(g[:])
	return int64(n), err
}

// CreatedAt extract the timestamp at which the GUID was created
func (g GUID96) CreatedAt() time.Time {
	return time.Unix(int64(encoder.Uint32(g[0:])), 0)
}

// --------------------------------------------------------------------

// GUID128 is a 16-byte globally unique identifier
type GUID128 [16]byte

var base128 = GUID128{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// New128 creates a 128bit/16byte global identifier
func New128() GUID128 { return new128at(time.Now()) }

func new128at(ts time.Time) GUID128 {
	bininc := [4]byte{0, 0, 0, 0}
	encoder.PutUint32(bininc[:], nextInc())

	bytes := base128
	encoder.PutUint64(bytes[0:], uint64(ts.Unix()))
	copy(bytes[8:], hostpid[:])
	copy(bytes[13:], bininc[1:])

	return bytes
}

// Bytes returns a byte slice
func (g GUID128) Bytes() []byte { return g[:] }

// WriteTo implements io.WriterTo
func (g GUID128) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(g[:])
	return int64(n), err
}

// CreatedAt extract the timestamp at which the GUID was created
func (g GUID128) CreatedAt() time.Time {
	return time.Unix(int64(encoder.Uint64(g[0:])), 0)
}

// --------------------------------------------------------------------

const maxUint24 = (1 << 24) - 1

var inc uint32
var hostpid [5]byte
var encoder = binary.BigEndian

func init() {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "localhost"
	}

	hash := md5.Sum([]byte(hostname))
	copy(hostpid[:], hash[:3])
	encoder.PutUint16(hostpid[3:], uint16(os.Getpid()))
}

func nextInc() uint32 {
	num := atomic.AddUint32(&inc, 1)
	if num > maxUint24 {
		num = atomic.AddUint32(&inc, maxUint24+1)
	}
	return num
}
