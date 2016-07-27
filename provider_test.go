package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// This test runs Provider sanity checks.
func TestProvider(t *testing.T) {
	client := &MockClient{}
	if err := Provider(client).(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
