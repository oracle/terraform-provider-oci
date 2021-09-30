// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MaintenanceRunRequiredOnlyResource = MaintenanceRunResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunRepresentation)

	MaintenanceRunResourceConfig = MaintenanceRunResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Update, maintenanceRunRepresentation)

	maintenanceRunSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_run_id": Representation{RepType: Required, Create: `${oci_database_maintenance_run.test_maintenance_run.id}`},
	}

	mrTimeScheduledCreate = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)
	mrTimeScheduledUpdate = time.Now().UTC().AddDate(0, 0, 10).Truncate(time.Millisecond)

	maintenanceRunRepresentation = map[string]interface{}{
		"maintenance_run_id":   Representation{RepType: Required, Create: `${var.maintenance_run_id}`},
		"is_enabled":           Representation{RepType: Required, Create: `true`},
		"is_patch_now_enabled": Representation{RepType: Optional},
		"patch_id":             Representation{RepType: Optional, Create: `${var.maintenance_run_patch_id}`},
		"patching_mode":        Representation{RepType: Optional, Create: `ROLLING`, Update: `NONROLLING`},
		"time_scheduled":       Representation{RepType: Required, Create: mrTimeScheduledCreate.Format(time.RFC3339Nano), Update: mrTimeScheduledUpdate.Format(time.RFC3339Nano)},
	}

	MaintenanceRunResourceDependencies = ""
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseMaintenanceRunResource_basic(t *testing.T) {
	t.Skip("Skip this test till DBaas provides a better way of testing this.")

	httpreplay.SetScenario("TestDatabaseMaintenanceRunResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	maintenanceRunId := getEnvSettingWithBlankDefault("maintenance_run_id")
	maintenanceRunIdVariableStr := fmt.Sprintf("variable \"maintenance_run_id\" { default = \"%s\" }\n", maintenanceRunId)

	patchId := getEnvSettingWithBlankDefault("maintenance_run_patch_id")
	patchIdVariableStr := fmt.Sprintf("variable \"maintenance_run_patch_id\" { default = \"%s\" }\n", patchId)

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_maintenance_run.test_maintenance_run"
	singularDatasourceName := "data.oci_database_maintenance_run.test_maintenance_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+MaintenanceRunResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Create, maintenanceRunRepresentation), "database", "maintenanceRun", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Create, maintenanceRunRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
				resource.TestCheckResourceAttr(resourceName, "patching_mode", "ROLLING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", mrTimeScheduledCreate.Format(time.RFC3339Nano)),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Update, maintenanceRunRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_patch_now_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
				resource.TestCheckResourceAttr(resourceName, "patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_scheduled"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_run_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_subtype"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_failure_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_maintenance_run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_scheduled"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_enabled",
				"is_patch_now_enabled",
				// In GET request `maintenance_run_id` is mapped to `id`
				"maintenance_run_id",
			},
			ResourceName: resourceName,
		},
	})
}
