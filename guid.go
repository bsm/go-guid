package guid

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"os"
	"sync/atomic"
	"time"
)

type GUID interface {
	CreatedAt() time.Time
	WriteTo(w io.Writer) (int64, error)
	Bytes() []byte
}

// --------------------------------------------------------------------

type guid96 [12]byte

var base96 = guid96{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// New96 creates a 96bit/12byte global identifier
func New96() GUID { return new96at(time.Now()) }

func new96at(ts time.Time) GUID {
	bininc := [4]byte{0, 0, 0, 0}
	encoder.PutUint32(bininc[:], nextInc())

	bytes := base96
	encoder.PutUint32(bytes[0:], uint32(ts.Unix()))
	copy(bytes[4:], hostpid[:])
	copy(bytes[9:], bininc[1:])

	return bytes
}

// Bytes returns a byte slice
func (g guid96) Bytes() []byte { return g[:] }

// WriteTo implements io.WriterTo
func (g guid96) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(g[:])
	return int64(n), err
}

// CreatedAt extract the timestap at which the GUID was created
func (g guid96) CreatedAt() time.Time {
	return time.Unix(int64(encoder.Uint32(g[0:])), 0)
}

// --------------------------------------------------------------------

type guid128 [16]byte

var base128 = guid128{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// New128 creates a 128bit/16byte global identifier
func New128() GUID { return new128at(time.Now()) }

func new128at(ts time.Time) GUID {
	bininc := [4]byte{0, 0, 0, 0}
	encoder.PutUint32(bininc[:], nextInc())

	bytes := base128
	encoder.PutUint64(bytes[0:], uint64(ts.Unix()))
	copy(bytes[8:], hostpid[:])
	copy(bytes[13:], bininc[1:])

	return bytes
}

// Bytes returns a byte slice
func (g guid128) Bytes() []byte { return g[:] }

// WriteTo implements io.WriterTo
func (g guid128) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(g[:])
	return int64(n), err
}

// CreatedAt extract the timestap at which the GUID was created
func (g guid128) CreatedAt() time.Time {
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
