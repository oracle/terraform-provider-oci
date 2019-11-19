// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

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

	maintenanceRunRepresentation = map[string]interface{}{
		"maintenance_run_id": Representation{repType: Required, create: `${var.maintenance_run_id}`},
		"is_enabled":         Representation{repType: Required, create: `false`, update: `true`},
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

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_maintenance_run.test_maintenance_run"
	singularDatasourceName := "data.oci_database_maintenance_run.test_maintenance_run"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceDependencies +
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
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Create, maintenanceRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_scheduled"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_maintenance_run", "test_maintenance_run", Optional, Update, maintenanceRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "maintenance_run_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_scheduled"),

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
					compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceConfig,
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
				Config: config + compartmentIdVariableStr + maintenanceRunIdVariableStr + MaintenanceRunResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"is_enabled",
					// In GET request `maintenance_run_id` is mapped to `id`
					"maintenance_run_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
