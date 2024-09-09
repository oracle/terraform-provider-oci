// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementRunbookSingularDataSourceRepresentation = map[string]interface{}{
		// Runbooks are currently created by Oracle, and read-only. There is no Create API.
		"runbook_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_runbook_ocid}`},
	}

	FleetAppsManagementRunbookDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `OS (Compute) Patching`},
		"operation":         acctest.Representation{RepType: acctest.Optional, Create: `Patch`},
		"platform":          acctest.Representation{RepType: acctest.Optional, Create: `OS (Compute)`},
		"runbook_relevance": acctest.Representation{RepType: acctest.Optional, Create: `PRODUCT`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":              acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_DEFINED`},
	}

	FleetAppsManagementRunbookResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Runbooks are currently created by Oracle, and read-only. There is no Create API.
	runbookId := utils.GetEnvSettingWithBlankDefault("test_runbook_ocid")
	testRunbookStr := fmt.Sprintf("variable \"test_runbook_ocid\" { default = \"%s\" }\n", runbookId)

	datasourceName := "data.oci_fleet_apps_management_runbooks.test_runbooks"
	singularDatasourceName := "data.oci_fleet_apps_management_runbook.test_runbook"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbooks", "test_runbooks", acctest.Required, acctest.Create, FleetAppsManagementRunbookDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementRunbookResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "runbook_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookSingularDataSourceRepresentation) +
				testRunbookStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "associations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_relevance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
