package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"drone": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("DRONE_SERVER"); err == "" {
		t.Fatal("DRONE_SERVER must be set for acceptance tests")
	}

	if err := os.Getenv("DRONE_TOKEN"); err == "" {
		t.Fatal("DRONE_TOKEN must be set for acceptance tests")
	}
}
