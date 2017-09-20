// File : provider_test.go

// if you dont mention this line, package main , you will get the error below
//can't load package: package sysdig-provider:
//provider_test.go:2:1: expected 'package', found 'import'
//provider_test.go:3:3: expected ';', found 'STRING' "os"
package main

import (

  "testing"

  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

var testAccProvider *schema.Provider

func init() {
  testAccProvider = Provider().(*schema.Provider)
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
  // We will use this function later on to make sure our test environment is valid.
  // For example, you can make sure here that some environment variables are set.
}
