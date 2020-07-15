package helloworld

import (
	"context"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// pathConfig returns the path configuration for CRUD operations on the backend
// configuration.
func pathWorld(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: worldPath + "$",
		Fields: map[string]*framework.FieldSchema{
			"name": {
				Type:        framework.TypeString,
				Description: `Name must be a name, returned by hello world`,
				Default:     "World",
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.pathWorldRead,
		},

		HelpSynopsis:    worldHelpSyn,
		HelpDescription: worldHelpDesc,
	}
}

// pathConfigRead handles read commands to the config
func (b *backend) pathWorldRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	resp := &logical.Response{
		Data: map[string]interface{}{
			"hello": data.Get("name").(string),
		},
	}
	return resp, nil
}

const worldHelpSyn = `
Execute Hello world
`
const worldHelpDesc = `
This endpoint executes Vault's Hello World
`
