package helloworld

import (
	"context"
	"fmt"
	"log"
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
			testHelloWorldRead(t),
		},
	})

}

func testHelloWorldRead(t *testing.T) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.ReadOperation,
		Path:      "printf",
		Check: func(resp *logical.Response) error {
			var d struct {
				Content string `mapstructure:"content"`
			}
			if err := mapstructure.Decode(resp.Data, &d); err != nil {
				return err
			}
			if len(d.Content) == 0 {
				return fmt.Errorf("Error retrieving content")
			}
			log.Printf("[WARN] Retrieved credentials: %v", d)

			return nil
		},
	}
}
