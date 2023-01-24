// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMaintenanceRunRequiredOnlyResource = DatabaseMaintenanceRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, DatabaseMaintenanceRunRepresentation)

	DatabaseMaintenanceRunResourceConfig = DatabaseMaintenanceRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, DatabaseMaintenanceRunRepresentation)

	DatabaseDatabaseMaintenanceRunSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_maintenance_run.test_maintenance_run.id}`},
	}

	maintenanceRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"maintenance_subtype":  acctest.Representation{RepType: acctest.Optional, Create: `QUARTERLY`},
		"maintenance_type":     acctest.Representation{RepType: acctest.Optional, Create: `PLANNED`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_resource_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_target_resource.test_target_resource.id}`},
		"target_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `AUTONOMOUS_EXADATA_INFRASTRUCTURE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: maintenanceRunDataSourceFilterRepresentation},
	}

	maintenanceRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_maintenance_run.test_maintenance_run.id}`}},
	}

	mrTimeScheduledCreate = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)
	mrTimeScheduledUpdate = time.Now().UTC().AddDate(0, 0, 10).Truncate(time.Millisecond)

	DatabaseMaintenanceRunRepresentation = map[string]interface{}{
		"maintenance_run_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.maintenance_run_id}`},
		"current_custom_action_timeout_in_mins": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"custom_action_timeout_in_mins":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_custom_action_timeout_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_enabled":                            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_patch_now_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_resume_patching":                    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"patch_id":                              acctest.Representation{RepType: acctest.Optional, Create: `${var.maintenance_run_patch_id}`},
		"patching_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `ROLLING`, Update: `NONROLLING`},
		"target_db_server_version":              acctest.Representation{RepType: acctest.Optional, Create: `targetDbServerVersion`, Update: `targetDbServerVersion2`},
		"target_storage_server_version":         acctest.Representation{RepType: acctest.Optional, Create: `targetStorageServerVersion`, Update: `targetStorageServerVersion2`},
		"time_scheduled":                        acctest.Representation{RepType: acctest.Required, Create: mrTimeScheduledCreate.Format(time.RFC3339Nano), Update: mrTimeScheduledUpdate.Format(time.RFC3339Nano)},
	}

	DatabaseMaintenanceRunResourceDependencies = ""
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
	datasourceName := "data.oci_database_maintenance_run.test_maintenance_runs"
	singularDatasourceName := "data.oci_database_maintenance_run.test_maintenance_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseMaintenanceRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Create, DatabaseMaintenanceRunRepresentation), "database", "maintenanceRun", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + DatabaseMaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, DatabaseMaintenanceRunRepresentation),
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
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + DatabaseMaintenanceRunResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + DatabaseMaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Create, DatabaseMaintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "current_custom_action_timeout_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "custom_action_timeout_in_mins", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_custom_action_timeout_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_patch_now_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_resume_patching", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
				resource.TestCheckResourceAttr(resourceName, "patching_mode", "ROLLING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_db_server_version", "targetDbServerVersion"),
				resource.TestCheckResourceAttr(resourceName, "target_storage_server_version", "targetStorageServerVersion"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", "timeScheduled"),
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
			Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + DatabaseMaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, DatabaseMaintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "current_custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_patch_now_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_resume_patching", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
				resource.TestCheckResourceAttr(resourceName, "patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_db_server_version", "targetDbServerVersion2"),
				resource.TestCheckResourceAttr(resourceName, "target_storage_server_version", "targetStorageServerVersion2"),
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
		// TODO: remove if creates problem during testing
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_runs", "test_maintenance_runs", acctest.Optional, acctest.Update, DatabaseMaintenanceRunRepresentation) +
				compartmentIdVariableStr + DatabaseMaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, DatabaseMaintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "maintenance_type", "PLANNED"),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_type", "AUTONOMOUS_EXADATA_INFRASTRUCTURE"),

				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.current_custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.current_patching_component"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.description"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.estimated_component_patching_start_time"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.estimated_patching_time.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.maintenance_subtype"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.maintenance_type"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.patch_failure_count"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.patch_id"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.patching_end_time"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.patching_start_time"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.patching_status"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.peer_maintenance_run_id"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.target_db_server_version"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.target_resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.target_storage_server_version"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.time_ended"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_runs.0.time_scheduled", "timeScheduled2"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_runs.0.time_started"),
			),
		},

		// verify datasources
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_runs", "test_maintenance_runs", acctest.Optional, acctest.Update, maintenanceRunDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseMaintenanceRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Optional, acctest.Update, DatabaseMaintenanceRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_subtype", "QUARTERLY"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_type", "PLANNED"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "target_resource_type", "AUTONOMOUS_EXADATA_INFRASTRUCTURE"),

				resource.TestCheckResourceAttr(datasourceName, "maintenance_runs.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.maintenance_subtype"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.maintenance_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.patch_failure_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.patch_id"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_runs.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.peer_maintenance_run_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.target_resource_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.time_ended"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_runs.0.time_scheduled", "timeScheduled2"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_runs.0.time_started"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, DatabaseDatabaseMaintenanceRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + DatabaseMaintenanceRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_run_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_patching_component"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_action_timeout_in_mins", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_component_patching_start_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "estimated_patching_time.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_custom_action_timeout_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_subtype"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_failure_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patching_end_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patching_start_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patching_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_maintenance_run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_scheduled"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMaintenanceRunRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_enabled",
				"is_patch_now_enabled",
				"is_resume_patching",
				// In GET request `maintenance_run_id` is mapped to `id`
				"maintenance_run_id",
			},
			ResourceName: resourceName,
		},
	})
}
