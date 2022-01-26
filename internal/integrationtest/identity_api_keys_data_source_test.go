// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"regexp"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityAPIKeysTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         identity.ListApiKeysResponse
}

func (s *DatasourceIdentityAPIKeysTestSuite) SetupTest() {
	_, tokenFn := acctest.TokenizeWithHttpReplay("api_data_source")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + publicKeyVariableStr + publicKeyUpdateVariableStr + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.userName}}"
		description = "automated test user"
		compartment_id = "${var.tenancy_ocid}"
	}
	
	resource "oci_identity_api_key" "t" {
		user_id = "${oci_identity_user.t.id}"
		key_value = "${var.api_key_value}"
	}
	
	resource "oci_identity_api_key" "u" {
		user_id = "${oci_identity_user.t.id}"
		key_value = "${var.api_key_update_value}"
	}`, map[string]string{"userName": "user_" + utils.Timestamp()})
	s.ResourceName = "data.oci_identity_api_keys.t"
}

func (s *DatasourceIdentityAPIKeysTestSuite) TestAccDatasourceIdentityAPIKeys_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_api_keys" "t" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "2"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
				data "oci_identity_api_keys" "t" {
					user_id = "${oci_identity_user.t.id}"
					filter {
						name = "id"
						values = ["${oci_identity_api_key.t.id}"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.id", "oci_identity_api_key.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.fingerprint", "oci_identity_api_key.t", "fingerprint"),
					resource.TestMatchResourceAttr(s.ResourceName, "api_keys.0.key_value", regexp.MustCompile("-----BEGIN PUBL.*")),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.time_created", "oci_identity_api_key.t", "time_created"),
					// TODO: This field is not being returned by the service call but is showing up in the datasource
					//resource.TestCheckNoResourceAttr(s.ResourceName, "api_keys.0.inactive_status"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.0.state", string(identity.ApiKeyLifecycleStateActive)),
				),
			},
		},
	},
	)
}

// issue-routing-tag: identity/default
func TestDatasourceIdentityAPIKeysTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityAPIKeysTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityAPIKeysTestSuite))
}
