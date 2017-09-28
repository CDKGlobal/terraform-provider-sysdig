// File : provider_test.go

package sysdig

import (

  "testing"
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "os"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
  testAccProvider = Provider().(*schema.Provider)
  testAccProviders = map[string]terraform.ResourceProvider{
    "sysdig": testAccProvider,
  }

}

func TestProvider(t *testing.T) {
  if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
    t.Fatalf("err: %s", err)
  }
}

func TestProvider_impl(t *testing.T) {
  var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
  if v := os.Getenv("token"); v == "" {
    t.Fatal("Sysdig token must be set for acceptance tests")
  }

}
