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
		generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunRepresentation)

	MaintenanceRunResourceConfig = MaintenanceRunResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Update, maintenanceRunRepresentation)

	maintenanceRunSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_run_id": Representation{repType: Required, create: `${oci_database_maintenance_run.test_maintenance_run.id}`},
	}

	mrTimeScheduledCreate = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)
	mrTimeScheduledUpdate = time.Now().UTC().AddDate(0, 0, 10).Truncate(time.Millisecond)

	maintenanceRunRepresentation = map[string]interface{}{
		"maintenance_run_id":   Representation{repType: Required, create: `${var.maintenance_run_id}`},
		"is_enabled":           Representation{repType: Required, create: `false`, update: `true`},
		"is_patch_now_enabled": Representation{repType: Optional, update: `true`},
		"patch_id":             Representation{repType: Optional, create: `${var.maintenance_run_patch_id}`},
		"time_scheduled":       Representation{repType: Optional, create: mrTimeScheduledCreate.Format(time.RFC3339Nano), update: mrTimeScheduledUpdate.Format(time.RFC3339Nano)},
	}

	MaintenanceRunResourceDependencies = ""
)

func TestDatabaseMaintenanceRunResource_basic(t *testing.T) {
	t.Skip("Skip this test till DBaas provides a better way of testing this.")

	httpreplay.SetScenario("TestDatabaseMaintenanceRunResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
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
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+MaintenanceRunResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Create, maintenanceRunRepresentation), "database", "maintenanceRun", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Create, maintenanceRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
					resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "time_scheduled", mrTimeScheduledCreate.Format(time.RFC3339Nano)),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Update, maintenanceRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_patch_now_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
					resource.TestCheckResourceAttrSet(resourceName, "patch_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "time_scheduled", mrTimeScheduledUpdate.Format(time.RFC3339Nano)),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Required, Create, maintenanceRunSingularDataSourceRepresentation) +
					compartmentIdVariableStr + maintenanceRunIdVariableStr + patchIdVariableStr + MaintenanceRunResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_run_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_subtype"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_type"),
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
		},
	})
}
