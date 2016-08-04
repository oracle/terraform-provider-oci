package main

import (
	"testing"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

const testProviderConfig = `

provider "baremetal" {
	tenancy_ocid = "ocid.tenancy.aaaa"
	user_ocid = "ocid.user.bbbbb"
	fingerprint = "xxxxxxxxxx"
	private_key_path = "/home/foo/private_key.pem"
	private_key_password = "password"
}

`

// This test runs the Provider sanity checks.
func TestProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &baremtlsdk.Client{}
	if err := Provider(func(d *schema.ResourceData) (BareMetalClient, error) {
		return client, nil
	}).(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
