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
	FleetAppsManagementRunbookImportSingularDataSourceRepresentation = map[string]interface{}{
		"runbook_id": acctest.Representation{RepType: acctest.Required, Create: `${var.runbook_id}`},
		"import_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.import_id}`},
	}
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookImportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookImportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	importId := utils.GetEnvSettingWithBlankDefault("import_id")
	importIdVariableStr := fmt.Sprintf("variable \"import_id\" { default = \"%s\" }\n", importId)
	runbookId := utils.GetEnvSettingWithBlankDefault("runbook_id_for_import")
	runbookIdVariableStr := fmt.Sprintf("variable \"runbook_id\" { default = \"%s\" }\n", runbookId)

	singularDatasourceName := "data.oci_fleet_apps_management_runbook_import.test_runbook_import"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				importIdVariableStr + runbookIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook_import", "test_runbook_import", acctest.Required, acctest.Create, FleetAppsManagementRunbookImportSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "import_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tracking_id"),
			),
		},
	})
}
