// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/oracle/oci-go-sdk/v56/identity"
	"github.com/stretchr/testify/suite"
)

const (
	api_key                 = ""
	api_key_with_whitespace = ""
)

type ResourceIdentityAPIKeyTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityAPIKeyTestSuite) SetupTest() {
	_, tokenFn := acctest.TokenizeWithHttpReplay("api_key")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "automated test user"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)
	s.ResourceName = "oci_identity_api_key.t"
}

func (s *ResourceIdentityAPIKeyTestSuite) TestAccResourceIdentityAPIKey_basic() {
	_, tokenFn := acctest.TokenizeWithHttpReplay("api_key_2")
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

// issue-routing-tag: identity/default
func TestResourceIdentityAPIKeyTestSuite(t *testing.T) {
	t.Skip("Run manually with a valid api key")
	httpreplay.SetScenario("TestResourceIdentityAPIKeyTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
