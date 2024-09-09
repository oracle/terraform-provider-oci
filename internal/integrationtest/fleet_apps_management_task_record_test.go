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
	FleetAppsManagementTaskRecordSingularDataSourceRepresentation = map[string]interface{}{
		// TaskRecords are currently created by Oracle, and read-only. There is no Create API.
		"task_record_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_task_record_ocid}`},
	}

	FleetAppsManagementTaskRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Patch Operating System`},
		"platform":       acctest.Representation{RepType: acctest.Optional, Create: `OS (Compute)`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_DEFINED`},
	}

	FleetAppsManagementTaskRecordResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementTaskRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementTaskRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// TaskRecords are currently created by Oracle, and read-only. There is no Create API.
	taskRecordId := utils.GetEnvSettingWithBlankDefault("test_task_record_ocid")
	taskRecordIdVariableStr := fmt.Sprintf("variable \"test_task_record_ocid\" { default = \"%s\" }\n", taskRecordId)

	datasourceName := "data.oci_fleet_apps_management_task_records.test_task_records"
	singularDatasourceName := "data.oci_fleet_apps_management_task_record.test_task_record"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_task_records", "test_task_records", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "task_record_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordSingularDataSourceRepresentation) +
				taskRecordIdVariableStr + compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_record_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
