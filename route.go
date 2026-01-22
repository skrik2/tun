package tun

type RouteController interface {
	Apply() error
	Cleanup() error
}
