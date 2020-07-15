package helloworld

import (
	"context"
	"fmt"
	"os"
	"testing"

	logicaltest "github.com/hashicorp/vault/helper/testhelpers/logical"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/mitchellh/mapstructure"
)

func TestBackend(t *testing.T) {
	if os.Getenv("VAULT_ACC") == "" {
		//	t.SkipNow()
	}
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	b, err := Factory(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}

	logicaltest.Test(t, logicaltest.TestCase{
		LogicalBackend: b,
		Steps: []logicaltest.TestStep{
			//testAccStepWorldRead(t),
			testAccStepWorldReadName(t, "HashiCorp"),
		},
	})

}

func testAccStepWorldRead(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.ReadOperation,
		Path:      "world",
		Check: func(resp *logical.Response) error {
			var d struct {
				Hello string `mapstructure:"hello"`
			}
			if err := mapstructure.Decode(resp.Data, &d); err != nil {
				return err
			}

			n := "World"
			if d.Hello != n {
				return fmt.Errorf("got %v: want %v", d.Hello, n)
			}

			return nil
		},
	}
}

func testAccStepWorldReadName(t *testing.T, n string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.ReadOperation,
		Path:      "world",
		Data: map[string]interface{}{
			"name": n,
		},
		Check: func(resp *logical.Response) error {
			var d struct {
				Hello string `mapstructure:"hello"`
			}
			if err := mapstructure.Decode(resp.Data, &d); err != nil {
				return err
			}

			if d.Hello != n {
				return fmt.Errorf("got %v: want %v", d.Hello, n)
			}

			return nil
		},
	}
}
