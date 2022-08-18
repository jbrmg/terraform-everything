package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var testAccProvider = New()()
var providerFactories = map[string]func() (*schema.Provider, error){
	"ikea": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func TestProvider(t *testing.T) {
	if err := New()().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IKEA_USERNAME"); v == "" {
		t.Fatal("IKEA_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("IKEA_PASSWORD"); v == "" {
		t.Fatal("IKEA_PASSWORD must be set for acceptance tests")
	}
}
