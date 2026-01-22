package tun

// NewTun 只做 结构构造 + 参数校验
// Open 才触发 任何系统级副作用
func NewTun(opts TunOptions, routeOpts RouteOptions) (Tun, error) {}
