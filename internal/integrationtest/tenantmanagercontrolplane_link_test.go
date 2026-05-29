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

	TenantmanagercontrolplaneLinkTenancyNameSingularDataSourceRepresentation = map[string]interface{}{
		"link_id": acctest.Representation{RepType: acctest.Required, Create: `${var.link_id}`},
	}

	TenantmanagercontrolplaneLinkDataSourceRepresentation = map[string]interface{}{
		"child_tenancy_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.child_tenancy_id}`},
		"feature":           acctest.Representation{RepType: acctest.Optional, Create: `${var.feature}`},
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

	feature := utils.GetEnvSettingWithBlankDefault("feature")
	featureVariableStr := fmt.Sprintf("variable \"feature\" { default = \"%s\" }\n", feature)

	singularDatasourceName := "data.oci_tenantmanagercontrolplane_link.test_link"
	linkTenancyNameSingularDatasourceName := "data.oci_tenantmanagercontrolplane_link_tenancy_name.test_link_tenancy_name"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + childTenancyIdVariableStr + parentTenancyIdVariableStr + stateVariableStr + featureVariableStr + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_links", "test_links", acctest.Optional, acctest.Create, TenantmanagercontrolplaneLinkDataSourceRepresentation)
	singularDataSourceConfig := config + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_link", "test_link", acctest.Required, acctest.Create, TenantmanagercontrolplaneLinkSingularDataSourceRepresentation)
	linkTenancyNameSingularDataSourceConfig := config + linkIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_link_tenancy_name", "test_link_tenancy_name", acctest.Required, acctest.Create, TenantmanagercontrolplaneLinkTenancyNameSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)
	fmt.Printf("Link Tenancy Name Singular Data Source Config: %s\n", linkTenancyNameSingularDataSourceConfig)

	testSteps := []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
		},
	}

	if linkId != "" {
		testSteps = append(testSteps, resource.TestStep{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "link_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "feature"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		})

		testSteps = append(testSteps, resource.TestStep{
			Config: linkTenancyNameSingularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "link_id"),

				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "child_tenancy_id"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "child_tenancy_name"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "feature"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "parent_tenancy_id"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "parent_tenancy_name"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(linkTenancyNameSingularDatasourceName, "time_updated"),
			),
		})
	}

	acctest.ResourceTest(t, nil, testSteps)
}
