package tun

import (
	"context"
	"time"

	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

// gVisorLinkAdapter adapts a generic LinkEndpoint to gVisor.
type GVisorLinkEndpointProvider interface {
	NewLinkEndpoint() (stack.LinkEndpoint, stack.NICOptions, error)
}

// gVisorLinkInjector writes raw PacketBuffer to link.
// It MUST only be used by gVisor stack implementation.
// 允许 GVisorStack 在绕过 transport endpoint 的情况下，
// 直接将一个已经构造完成的 PacketBuffer 注入到 LinkEndpoint。
//
// 典型缺口包括：
//
//	ICMP 特殊应答
//	FullCone / symmetric NAT 的“伪回包”
//	UDP hole-punching
//	协议实验 / 扩展行为
type gVisorLinkInjector interface {
	InjectPacket(pkt *stack.PacketBuffer) error
}

type GVisorStack struct {
	ctx context.Context

	endpointProvider GVisorLinkEndpointProvider
	attachedEndpoint stack.LinkEndpoint

	linkInjector gVisorLinkInjector

	netstack *stack.Stack

	udpTimeout time.Duration
	handler    *Handler
}

type GVisorOptions struct {
	EndpointFactory GVisorLinkEndpointProvider
	LinkInjector    gVisorLinkInjector

	IdleTimeout time.Duration
	Handler     *Handler
}

func NewGVisorStack(ctx context.Context, opts GVisorOptions) (Stack, error)

func newGVisorNetstack(ep stack.LinkEndpoint, opts stack.NICOptions) (*stack.Stack, error)
