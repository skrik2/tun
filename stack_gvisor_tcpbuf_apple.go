//go:build darwin

package tun

import "gvisor.dev/gvisor/pkg/tcpip/transport/tcp"

// NetworkExtension has memory usage limits.
// See https://developer.apple.com/forums/thread/73148

const (
	tcpRecvBufMinSize = tcp.MinBufferSize
	tcpSendBufMinSize = tcp.MinBufferSize

	tcpRecvBufMaxSize = tcp.MaxBufferSize
	tcpSendBufMaxSize = tcp.MaxBufferSize

	tcpRecvBufDefaultSize = tcp.DefaultReceiveBufferSize
	tcpSendBufDefaultSize = tcp.DefaultSendBufferSize
)
