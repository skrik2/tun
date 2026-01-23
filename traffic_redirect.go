package tun

// route table
// android vpn service
// apple network extension
type TrafficRedirector interface {
	Open() error   // 开始导流量到 TUN / Android VPN / Network Extension
	Close() error  // 停止导流量
	Update() error // 可选：动态更新规则或策略
}

// Linux, Windows, macOS
type RouteTableRedirector struct{}

// Android VPN Service
type AndroidVPNRedirector struct{}

// Network Extension
type NEPacketTunnelRedirector struct{}
