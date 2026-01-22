package tun

const (
	DefaultIPRoute2TableIndex = 2026
	DefaultIPRoute2RuleIndex  = 8848
)

// Tun interface implements tun interface interaction
type Tun interface {
	Start() error
	Close() error
}

// TunOptions for tun interface implementation
type TunOptions struct {
	Name string
	MTU  uint32
}
