package tools

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	defaultDNSResolver   = "8.8.8.8"
	dnsResolverProto     = "udp"
	dnsResolverTimeoutMs = 5000
)

type DNSResolver struct {
	Enable  bool   `json:"enable"`
	Server  string `json:"resolver,omitempty"`
	Proto   string `json:"proto,omitempty"`
	Timeout uint   `json:"timeout,omitempty"`
}

type CustomDNS func(ctx context.Context, network, addr string) (net.Conn, error)

func CustomDialer(conf DNSResolver) *net.Dialer {
	cnf := checkParamsForDNSResolver(conf)
	return &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(cnf.Timeout) * time.Millisecond,
				}
				return d.DialContext(ctx, cnf.Proto, fmt.Sprintf("%s:53", cnf.Server))
			},
		},
	}
}

func CustomDNSResolver(conf DNSResolver) CustomDNS {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return CustomDialer(conf).DialContext(ctx, network, addr)
	}
}

func checkParamsForDNSResolver(conf DNSResolver) DNSResolver {
	serverIP := defaultDNSResolver
	proto := dnsResolverProto
	timeout := dnsResolverTimeoutMs
	if conf.Server != "" {
		serverIP = conf.Server
	}
	if conf.Proto != "" {
		proto = conf.Proto
	}
	if conf.Timeout > 0 {
		timeout = int(conf.Timeout)
	}
	return DNSResolver{
		Server:  serverIP,
		Proto:   proto,
		Timeout: uint(timeout),
	}
}
