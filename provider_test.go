package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

// This test runs Provider sanity checks
func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
