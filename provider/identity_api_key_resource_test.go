// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/identity"
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
	_, tokenFn := tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "automated test user"
	}`, nil)
	s.ResourceName = "oci_identity_api_key.t"
}

func (s *ResourceIdentityAPIKeyTestSuite) TestAccResourceIdentityAPIKey_basic() {
	_, tokenFn := tokenize()
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
				ImportState:       true,
				ImportStateVerify: true,
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
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
