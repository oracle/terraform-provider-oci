// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityGroupsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
	Token        string
	TokenFn      acctest.TokenFn
}

func (s *DatasourceIdentityGroupsTestSuite) SetupTest() {
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("identity_group_data_source")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + s.TokenFn(`
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "automated test Group"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)
	s.ResourceName = "data.oci_identity_groups.t"
}

func (s *DatasourceIdentityGroupsTestSuite) TestAccDatasourceIdentityGroups_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_groups" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.#"),
				),
			},
			// Test cascading filters
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_groups" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name   = "name"
						values = ["{{.token}}", "Administrators"]
					}
					filter {
						name   = "description"
						values = ["automated test Group"]
					}
				}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "groups.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.description", "automated test Group"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.state", string(identity.GroupLifecycleStateActive)),
					// TODO: This field is not being returned by the service call but is still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "groups.0.inactive_state"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: identity/default
func TestDatasourceIdentityGroupsTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityGroupsTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityGroupsTestSuite))
}
