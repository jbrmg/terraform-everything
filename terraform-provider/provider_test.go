package main

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var testAccProvider = New("dev")()
var providerFactories = map[string]func() (*schema.Provider, error){
	"everything": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("EVERYTHING_USERNAME"); v == "" {
		t.Fatal("EVERYTHING_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("EVERYTHING_PASSWORD"); v == "" {
		t.Fatal("EVERYTHING_PASSWORD must be set for acceptance tests")
	}
}
