package cda

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cda": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv(environmentVariableMaps[cdaServer]); v == "" {
		t.Fatal("cda_server must be set for acceptance tests")
	}

	if v := os.Getenv(environmentVariableMaps[cdaUser]); v == "" {
		t.Fatal("user must be set for acceptance tests")
	}

	if v := os.Getenv(environmentVariableMaps[password]); v == "" {
		t.Fatal("password must be set for acceptance tests")
	}
}
