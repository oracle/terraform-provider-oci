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
	FleetAppsManagementRunbookExportSingularDataSourceRepresentation = map[string]interface{}{
		"export_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.export_id}`},
		"runbook_id": acctest.Representation{RepType: acctest.Required, Create: `${var.runbook_id}`},
	}
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookExportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookExportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	exportId := utils.GetEnvSettingWithBlankDefault("export_id")
	exportIdVariableStr := fmt.Sprintf("variable \"export_id\" { default = \"%s\" }\n", exportId)
	runbookId := utils.GetEnvSettingWithBlankDefault("runbook_id_for_export")
	runbookIdVariableStr := fmt.Sprintf("variable \"runbook_id\" { default = \"%s\" }\n", runbookId)

	singularDatasourceName := "data.oci_fleet_apps_management_runbook_export.test_runbook_export"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				runbookIdVariableStr + exportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook_export", "test_runbook_export", acctest.Required, acctest.Create, FleetAppsManagementRunbookExportSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "export_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tracking_id"),
			),
		},
	})
}
