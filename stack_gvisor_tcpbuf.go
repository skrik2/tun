//go:build !darwin

package tun

import "gvisor.dev/gvisor/pkg/tcpip/transport/tcp"

const (
	tcpRecvBufMinSize = tcp.MinBufferSize
	tcpSendBufMinSize = tcp.MinBufferSize

	tcpRecvBufMaxSize = 8 << 20 // 8 MiB
	tcpSendBufMaxSize = 6 << 20 // 6 MiB

	tcpRecvBufDefaultSize = tcp.DefaultReceiveBufferSize
	tcpSendBufDefaultSize = tcp.DefaultSendBufferSize
)
