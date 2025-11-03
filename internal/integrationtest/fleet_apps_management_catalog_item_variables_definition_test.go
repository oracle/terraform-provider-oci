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

var (
	FleetAppsManagementCatalogItemVariablesDefinitionSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_item_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_catalog_item.test_catalog_item.id}`},
	}

	FleetAppsManagementCatalogItemVariablesDefinitionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Required, acctest.Create, FleetAppsManagementCatalogItemRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementCatalogItemVariablesDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementCatalogItemVariablesDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_fleet_apps_management_catalog_item_variables_definition.test_catalog_item_variables_definition"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_catalog_item_variables_definition", "test_catalog_item_variables_definition", acctest.Required, acctest.Create, FleetAppsManagementCatalogItemVariablesDefinitionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementCatalogItemVariablesDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_item_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "schema_document.#", "1"),
			),
		},
	})
}
