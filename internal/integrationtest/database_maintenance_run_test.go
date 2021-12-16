// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	MaintenanceRunRequiredOnlyResource = MaintenanceRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, maintenanceRunRepresentation)

	MaintenanceRunResourceConfig = MaintenanceRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, maintenanceRunRepresentation)

	maintenanceRunSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_maintenance_run.test_maintenance_run.id}`},
	}

	mrTimeScheduledCreate = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)
	mrTimeScheduledUpdate = time.Now().UTC().AddDate(0, 0, 10).Truncate(time.Millisecond)

	maintenanceRunRepresentation = map[string]interface{}{
		"maintenance_run_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.maintenance_run_id}`},
		"is_enabled":           acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_patch_now_enabled": acctest.Representation{RepType: acctest.Optional},
		"patch_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.maintenance_run_patch_id}`},
		"patching_mode":        acctest.Representation{RepType: acctest.Optional, Create: `ROLLING`, Update: `NONROLLING`},
		"time_scheduled":       acctest.Representation{RepType: acctest.Required, Create: mrTimeScheduledCreate.Format(time.RFC3339Nano), Update: mrTimeScheduledUpdate.Format(time.RFC3339Nano)},
	}

	MaintenanceRunResourceDependencies = ""
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseMaintenanceRunResource_basic(t *testing.T) {
	t.Skip("Skip this test till DBaas provides a better way of testing this.")

	httpreplay.SetScenario("TestDatabaseMaintenanceRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	maintenanceRunId := utils.GetEnvSettingWithBlankDefault("maintenance_run_id")
	maintenanceRunIdVariableStr := fmt.Sprintf("variable \"maintenance_run_id\" { default = \"%s\" }\n", maintenanceRunId)

	patchId := utils.GetEnvSettingWithBlankDefault("maintenance_run_patch_id")
	patchIdVariableStr := fmt.Sprintf("variable \"maintenance_run_patch_id\" { default = \"%s\" }\n", patchId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_maintenance_run.test_maintenance_run"
	singularDatasourceName := "data.oci_database_maintenance_run.test_maintenance_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MaintenanceRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Create, maintenanceRunRepresentation), "database", "maintenanceRun", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, maintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Create, maintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, maintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, maintenanceRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
