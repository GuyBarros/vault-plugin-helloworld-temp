package helloworld

import (
	"context"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const worldPath string = "world"

type backend struct {
	*framework.Backend
}

// Factory returns a new backend as logical.Backend.
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := newBackend()

	return b, nil
}

// Backend implements the CCP Secrets Engine.
func newBackend() *backend {
	var b = &backend{}

	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Help:        strings.TrimSpace(backendHelp),

		Paths: framework.PathAppend(
			[]*framework.Path{
				pathWorld(b),
			},
		),
	}

	return b
}

const backendHelp = `
HellowWorld is a minimal Vault Secrets Engine
`
