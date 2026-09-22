// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package functions

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
	"github.com/oracle/terraform-provider-oci/internal/client"
)

func TestUnitCreateFunctionsInvokeFunctionReturnsEndpointClientError(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("cannot generate test private key: %v", err)
	}
	configuration := functionsTestConfiguration{privateKey: privateKey}
	baseClient, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(configuration, "https://functions.example.com")
	if err != nil {
		t.Fatalf("cannot construct test Functions client: %v", err)
	}
	clients := &client.OracleClients{SdkClientMap: map[string]interface{}{
		"oci_functions.FunctionsInvokeClient": &baseClient,
	}}
	d := schema.TestResourceDataRaw(t, FunctionsInvokeFunctionResource().Schema, map[string]interface{}{
		"function_id":     "ocid1.fnfunc.oc1.iad.example",
		"invoke_endpoint": "https://functions.example.com",
	})

	err = createFunctionsInvokeFunction(d, clients)
	if err == nil || !strings.Contains(err.Error(), "no configure client is registered") {
		t.Fatalf("createFunctionsInvokeFunction() error = %v, want endpoint-client configuration error", err)
	}
}

type functionsTestConfiguration struct {
	privateKey *rsa.PrivateKey
}

func (c functionsTestConfiguration) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return c.privateKey, nil
}

func (functionsTestConfiguration) KeyID() (string, error) {
	return "tenancy/user/fingerprint", nil
}

func (functionsTestConfiguration) TenancyOCID() (string, error) {
	return "tenancy", nil
}

func (functionsTestConfiguration) UserOCID() (string, error) {
	return "user", nil
}

func (functionsTestConfiguration) KeyFingerprint() (string, error) {
	return "fingerprint", nil
}

func (functionsTestConfiguration) Region() (string, error) {
	return "us-ashburn-1", nil
}

func (functionsTestConfiguration) AuthType() (oci_common.AuthConfig, error) {
	return oci_common.AuthConfig{AuthType: oci_common.UserPrincipal}, nil
}
