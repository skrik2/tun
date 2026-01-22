package tun

import (
	"time"
)

// Stack interface implement ip protocol stack, bridging raw network packets and data streams
type Stack interface {
	Open() error
	Close() error
}

// StackOptions for the stack implementation
type StackOptions struct {
	Tun        Tun
	UDPTimeout time.Duration
}
