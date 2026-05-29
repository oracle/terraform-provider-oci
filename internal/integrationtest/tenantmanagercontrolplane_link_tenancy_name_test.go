// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneLinkTenancyNameResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneLinkTenancyNameResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	linkId := utils.GetEnvSettingWithBlankDefault("link_id")
	linkIdVariableStr := fmt.Sprintf("variable \"link_id\" { default = \"%s\" }\n", linkId)

	childTenancyId := utils.GetEnvSettingWithBlankDefault("child_tenancy_id")
	childTenancyIdVariableStr := fmt.Sprintf("variable \"child_tenancy_id\" { default = \"%s\" }\n", childTenancyId)

	parentTenancyId := utils.GetEnvSettingWithBlankDefault("parent_tenancy_id")
	parentTenancyIdVariableStr := fmt.Sprintf("variable \"parent_tenancy_id\" { default = \"%s\" }\n", parentTenancyId)

	state := utils.GetEnvSettingWithBlankDefault("state")
	stateVariableStr := fmt.Sprintf("variable \"state\" { default = \"%s\" }\n", state)

	feature := utils.GetEnvSettingWithBlankDefault("feature")
	featureVariableStr := fmt.Sprintf("variable \"feature\" { default = \"%s\" }\n", feature)

	singularDatasourceName := "data.oci_tenantmanagercontrolplane_link_tenancy_name.test_link_tenancy_name"

	acctest.SaveConfigContent("", "", "", t)

	if linkId == "" {
		dataSourceConfig := config + childTenancyIdVariableStr + parentTenancyIdVariableStr + stateVariableStr + featureVariableStr +
			acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_links", "test_links", acctest.Optional, acctest.Create, TenantmanagercontrolplaneLinkDataSourceRepresentation) + `

data "oci_tenantmanagercontrolplane_link_tenancy_name" "test_link_tenancy_name" {
  link_id = data.oci_tenantmanagercontrolplane_links.test_links.link_collection[0].items[0].id
}
`
		fmt.Printf("Data Source Derived Singular Data Source Config: %s\n", dataSourceConfig)

		acctest.ResourceTest(t, nil, []resource.TestStep{
			// verify singular datasource using an existing service-owned link
			{
				Config: dataSourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "link_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "child_tenancy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "child_tenancy_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "feature"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_tenancy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_tenancy_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
		})
		return
	}

	singularDataSourceConfig := config + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_link_tenancy_name", "test_link_tenancy_name", acctest.Required, acctest.Create, TenantmanagercontrolplaneLinkTenancyNameSingularDataSourceRepresentation)

	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "link_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "child_tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "child_tenancy_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "feature"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_tenancy_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
