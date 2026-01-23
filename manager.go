package tun

// Manager coordinates the lifecycle of tun, routing and IP stack.
//
// Responsibilities:
//   - Open and close the tun device.
//   - Apply and cleanup routing rules.
//   - Initialize and shutdown the IP stack.
//
// Non-responsibilities:
//   - Traffic policy decisions.
//   - Per-connection or per-packet processing.
//   - Runtime configuration mutation beyond lifecycle control.
//
// Manager is a composition root and MUST NOT become a god object.
type Manager interface {
	Open() error
	Close() error
}

type DefaultManager struct {
	Tun        Tun
	Redirector TrafficRedirector
	Stack      Stack

	opened bool
}
