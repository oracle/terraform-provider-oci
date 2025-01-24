// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	TenantmanagercontrolplaneLinkSingularDataSourceRepresentation = map[string]interface{}{
		"link_id": acctest.Representation{RepType: acctest.Required, Create: `${var.link_id}`},
	}

	TenantmanagercontrolplaneLinkDataSourceRepresentation = map[string]interface{}{
		"child_tenancy_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.child_tenancy_id}`},
		"parent_tenancy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.parent_tenancy_id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `${var.state}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneLinkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneLinkResource_basic")
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

	datasourceName := "data.oci_tenantmanagercontrolplane_links.test_links"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_link.test_link"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + childTenancyIdVariableStr + parentTenancyIdVariableStr + stateVariableStr + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_links", "test_links", acctest.Optional, acctest.Create, TenantmanagercontrolplaneLinkDataSourceRepresentation)
	singularDataSourceConfig := config + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_link", "test_link", acctest.Required, acctest.Create, TenantmanagercontrolplaneLinkSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "link_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "link_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
