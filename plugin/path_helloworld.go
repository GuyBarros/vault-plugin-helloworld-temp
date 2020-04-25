package helloworld

import (
	"context"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const helloWorldPath string = "printf"

// pathHelloWorld prints out hello world
func pathHelloWorld(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: helloWorldPath + "$",
		Fields: map[string]*framework.FieldSchema{
			"name": {
				Type:        framework.TypeString,
				Description: `The name of which will be printed together with Hello World`,
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: b.pathHelloWorldUpdate,
			logical.UpdateOperation: b.pathHelloWorldUpdate,
			logical.ReadOperation:   b.pathHelloWorldRead,
		},
		HelpSynopsis:    queryHelpSyn,
		HelpDescription: queryHelpDesc,
	}
}

func (b *backend) pathHelloWorldUpdate(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	n := data.Get("name").(string)

	resp := &logical.Response{
		Data: map[string]interface{}{
			"helloworld": n,
		},
	}
	return resp, nil
}

func (b *backend) pathHelloWorldRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	//n := data.Get("name").(string)
	//panic(nil)
	//m := map[string]string{		"hello": "world",	}

	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		Family    `mapstructure:",squash"`
		Location  `mapstructure:",squash"`
		FirstName string
	}

	x := Person{
		FirstName: "Guy",
	}

	resp := &logical.Response{
		Data: map[string]interface{}{
			"FirstName": x.FirstName,
		},
	}
	return resp, nil
}

const queryHelpSyn = `
TODO
`
const queryHelpDesc = `
TODO
`
