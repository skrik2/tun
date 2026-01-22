package tun

import "errors"

var (
	ErrTunNotOpen   = errors.New("tun not open")
	ErrDropByRule   = errors.New("drop by rule")
	ErrBypassByRule = errors.New("bypass by rule")
)

// Tun lifecycle:
//
//	New -> Closed
//	Closed --Open()--> Opened
//	Opened --Close()--> Closed
//
// Open must not be called concurrently with Close.
// Read/Write require Opened state.
type Tun interface {
	Open() error
	Close() error

	ReadPacket(pkt []byte) (int, error)
	WritePacket(pkt []byte) (int, error)

	Name() string
}

type TunOptions struct {
	Name string
	MTU  uint32
}

type LinuxTun interface {
	Tun
	BatchRead(pkts [][]byte) (int, error)
	BatchWrite(pkts [][]byte) (int, error)
	ChecksumOffload() bool
}
