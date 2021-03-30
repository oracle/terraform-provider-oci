// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/oracle/oci-go-sdk/v38/identity"
	"github.com/stretchr/testify/suite"
)

const (
	api_key = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtBLQAGmKJ7tpfzYJyqLG
ZDwHL51+d6T8Z00BnP9CFfzxZZZ48PcYSUHuTyCM8mR5JqYLyH6C8tZ/DKqwxUnc
ONgBytG3MM42bgxfHIhsZRj5rCz1oqWlSLuXvgww1kuqWnt6r+NtnXog439YsGTH
RotrTLTdEgOxH0EFP5uHUc9w/Uix7rWU7GB2ra060oeTB/hKpts5U70eI2EI6ec9
1sJdUIj7xNfBJeQQrz4CFUrkyzL06211CFvhmxH2hA9gBKOqC3rGL8XraHZBhGWn
mXlrQB7nNKsJrrv5fHwaPDrAY4iNP2W0q3LRpyNigJ6cgRuGJhHa82iHPmxgIx8m
fwIDAQAB
-----END PUBLIC KEY-----`
	api_key_with_whitespace = `  -----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtBLQAGmKJ7tpfzYJyqLGZDwHL51+d6T8Z00BnP9CFfzxZZZ48PcYSUHuTyCM8mR5JqYLyH6C8tZ/DKqwxUnc

ONgBytG3MM42bgxfHIhsZRj5rCz1oqWlSLuXvgww1kuqWnt6r+NtnXog439YsGTHRotrTLTdEgOxH0EFP5uHUc9w/Uix7rWU7GB2ra060oeTB/hKpts5U70eI2EI6ec9

1sJdUIj7xNfBJeQQrz4CFUrkyzL06211CFvhmxH2hA9gBKOqC3rGL8XraHZBhGWnmXlrQB7nNKsJrrv5fHwaPDrAY4iNP2W0q3LRpyNigJ6cgRuGJhHa82iHPmxgIx8m

fwIDAQAB
-----END PUBLIC KEY-----  `
)

type ResourceIdentityAPIKeyTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityAPIKeyTestSuite) SetupTest() {
	_, tokenFn := tokenizeWithHttpReplay("api_key")
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "automated test user"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)
	s.ResourceName = "oci_identity_api_key.t"
}

func (s *ResourceIdentityAPIKeyTestSuite) TestAccResourceIdentityAPIKey_basic() {
	_, tokenFn := tokenizeWithHttpReplay("api_key_2")
	tokenVars := map[string]string{
		"user_id": "${oci_identity_user.t.id}",
		"key_value": `<<EOF
` + api_key + `
EOF
`,
	}
	altTokenVars := map[string]string{
		"user_id": "${oci_identity_user.t.id}",
		"key_value": `<<EOF
` + api_key_with_whitespace + `
EOF
`,
	}
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// Verify construct resource
			{
				Config: tokenFn(s.Config+`
				resource "oci_identity_api_key" "t" {
				  user_id = "{{.user_id}}"
				  key_value = {{.key_value}}
				}`, tokenVars),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "fingerprint"),
					resource.TestCheckResourceAttr(s.ResourceName, "key_value", api_key),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.ApiKeyLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_status"),
				),
			},
			{
				Config: tokenFn(s.Config+`
				resource "oci_identity_api_key" "t" {
				  user_id = "{{.user_id}}"
				  key_value = {{.key_value}}
				}`, altTokenVars),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "fingerprint"),
					resource.TestCheckResourceAttr(s.ResourceName, "key_value", api_key), // Original api key, not 'alt'
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.ApiKeyLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_status"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestResourceIdentityAPIKeyTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceIdentityAPIKeyTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
