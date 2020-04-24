package helloworld

import (
	"context"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type backend struct {
	*framework.Backend
}

// Factory returns a new backend as logical.Backend.
//DONT TOUCH THIS!
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := newBackend()
	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}

// Backend implements the helloworld Secrets Engine.
func newBackend() *backend {
	var b = &backend{}

	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Help:        strings.TrimSpace(backendHelp),
		PathsSpecial: &logical.Paths{
			LocalStorage: []string{
				framework.WALPrefix,
			},
		},

		Paths: framework.PathAppend(
			[]*framework.Path{
				pathHelloWorld(b),
			},
		),

		//	InitializeFunc: b.initialize,
		//	Invalidate:     b.invalidate,

		//	Clean: b.cleanup,
	}

	return b
}

const backendHelp = `
this does nothing, just prints out helloworld. its a test. dont shoot!
`
