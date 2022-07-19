package proxy

import (
	serverpb "github.com/hashicorp/boundary/internal/gen/controller/servers/services"
)

// Option - how Options are passed as arguments.
type Option func(*Options)

// GetOpts - iterate the inbound Options and return a struct.
func GetOpts(opt ...Option) Options {
	opts := getDefaultOptions()
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Options = how options are represented
type Options struct {
	WithInjectedApplicationCredentials []*serverpb.Credential
}

func getDefaultOptions() Options {
	return Options{
		WithInjectedApplicationCredentials: nil,
	}
}

// WithInjectedApplicationCredentials provides an optional injected application
// credentials to use when establishing a proxy
func WithInjectedApplicationCredentials(creds []*serverpb.Credential) Option {
	return func(o *Options) {
		o.WithInjectedApplicationCredentials = creds
	}
}
