// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUsersTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
	Token        string
	TokenFn      acctest.TokenFn
}

func (s *DatasourceIdentityUsersTestSuite) SetupTest() {
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("user_data_source")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + s.TokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "automated test user"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)
	s.ResourceName = "data.oci_identity_users.t"
}

func (s *DatasourceIdentityUsersTestSuite) TestAccDatasourceIdentityUsers_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_users" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.#"),
				),
			},
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_users" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name = "name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "users.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.description", "automated test user"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.state", string(identity.UserLifecycleStateActive)),
					// TODO: These fields are not being returned by the service call but are still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "users.0.inactive_state"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.time_created"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: identity/default
func TestDatasourceIdentityUsersTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityUsersTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityUsersTestSuite))
}
