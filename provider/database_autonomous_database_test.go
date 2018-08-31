// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

const (
	AutonomousDatabaseRequiredOnlyResource = AutonomousDatabaseResourceDependencies + `
resource "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	admin_password = "${var.autonomous_database_admin_password}"
	compartment_id = "${var.compartment_id}"
	cpu_core_count = "${var.autonomous_database_cpu_core_count}"
	data_storage_size_in_tbs = "${var.autonomous_database_data_storage_size_in_tbs}"
	db_name = "${var.autonomous_database_db_name}"
}
`

	AutonomousDatabaseResourceConfig = AutonomousDatabaseResourceDependencies + `
resource "oci_database_autonomous_database" "test_autonomous_database" {
	#Required

	admin_password = "${var.autonomous_database_admin_password}"
	compartment_id = "${var.compartment_id}"
	cpu_core_count = "${var.autonomous_database_cpu_core_count}"
	data_storage_size_in_tbs = "${var.autonomous_database_data_storage_size_in_tbs}"
	db_name = "${var.autonomous_database_db_name}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.autonomous_database_defined_tags_value}")}"
	display_name = "${var.autonomous_database_display_name}"
	freeform_tags = "${var.autonomous_database_freeform_tags}"
	license_model = "${var.autonomous_database_license_model}"
}
`
	AutonomousDatabasePropertyVariables = `
variable "autonomous_database_admin_password" { default = "BEstrO0ng_#11" }
variable "autonomous_database_cpu_core_count" { default = 1 }
variable "autonomous_database_data_storage_size_in_tbs" { default = 1 }
variable "autonomous_database_db_name" { default = "adatabasedb1" }
variable "autonomous_database_defined_tags_value" { default = "value" }
variable "autonomous_database_display_name" { default = "example_autonomous_database" }
variable "autonomous_database_freeform_tags" { default = {"Department"= "Finance"} }
variable "autonomous_database_license_model" { default = "LICENSE_INCLUDED" }
variable "autonomous_database_state" { default = "AVAILABLE" }

`
	AutonomousDatabaseResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseAutonomousDatabaseResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"
	singularDatasourceName := "data.oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + AutonomousDatabasePropertyVariables + compartmentIdVariableStr + AutonomousDatabaseRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "adatabasedb1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + AutonomousDatabasePropertyVariables + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "adatabasedb1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			// @CODEGEN 09/2018 - Updating password and other fields except cpu_core_count and data_storage_size_in_tbs as service does not support scaling and password update at same time
			{
				Config: config + `
variable "autonomous_database_admin_password" { default = "BEstrO0ng_#12" }
variable "autonomous_database_cpu_core_count" { default = 1 }
variable "autonomous_database_data_storage_size_in_tbs" { default = 1 }
variable "autonomous_database_db_name" { default = "adatabasedb1" }
variable "autonomous_database_defined_tags_value" { default = "updatedValue" }
variable "autonomous_database_display_name" { default = "displayName2" }
variable "autonomous_database_freeform_tags" { default = {"Department"= "Accounting"} }
variable "autonomous_database_license_model" { default = "LICENSE_INCLUDED" }
variable "autonomous_database_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "adatabasedb1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters - scaling properties: cpu_core_count and data_storage_size_in_tbs
			// @CODEGEN 09/2018 - Despite scaling the ADB successfully, cpu_core_count and data_storage_size_in_tbs values returned may still return original values resulting in plan diff
			{
				Config: config + `
variable "autonomous_database_admin_password" { default = "BEstrO0ng_#12" }
variable "autonomous_database_cpu_core_count" { default = 2 }
variable "autonomous_database_data_storage_size_in_tbs" { default = 2 }
variable "autonomous_database_db_name" { default = "adatabasedb1" }
variable "autonomous_database_defined_tags_value" { default = "updatedValue" }
variable "autonomous_database_display_name" { default = "displayName2" }
variable "autonomous_database_freeform_tags" { default = {"Department"= "Accounting"} }
variable "autonomous_database_license_model" { default = "LICENSE_INCLUDED" }
variable "autonomous_database_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpu_core_count"),
					resource.TestCheckResourceAttrSet(resourceName, "data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "adatabasedb1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
				ExpectNonEmptyPlan: true,
			},
			// verify datasource
			{
				Config: config + `
variable "autonomous_database_admin_password" { default = "BEstrO0ng_#12" }
variable "autonomous_database_cpu_core_count" { default = 2 }
variable "autonomous_database_data_storage_size_in_tbs" { default = 2 }
variable "autonomous_database_db_name" { default = "adatabasedb1" }
variable "autonomous_database_defined_tags_value" { default = "updatedValue" }
variable "autonomous_database_display_name" { default = "displayName2" }
variable "autonomous_database_freeform_tags" { default = {"Department"= "Accounting"} }
variable "autonomous_database_license_model" { default = "LICENSE_INCLUDED" }
variable "autonomous_database_state" { default = "AVAILABLE" }

data "oci_database_autonomous_databases" "test_autonomous_databases" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.autonomous_database_display_name}"
	state = "${var.autonomous_database_state}"

    filter {
    	name = "id"
    	values = ["${oci_database_autonomous_database.test_autonomous_database.id}"]
    }
}
                ` + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.cpu_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.db_name", "adatabasedb1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_databases.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_databases.0.state"),
				),
				ExpectNonEmptyPlan: true,
			},
			// verify singular datasource
			{
				Config: config + `
variable "autonomous_database_admin_password" { default = "BEstrO0ng_#12" }
variable "autonomous_database_cpu_core_count" { default = 2 }
variable "autonomous_database_data_storage_size_in_tbs" { default = 2 }
variable "autonomous_database_db_name" { default = "adatabasedb1" }
variable "autonomous_database_defined_tags_value" { default = "updatedValue" }
variable "autonomous_database_display_name" { default = "displayName2" }
variable "autonomous_database_freeform_tags" { default = {"Department"= "Accounting"} }
variable "autonomous_database_license_model" { default = "LICENSE_INCLUDED" }
variable "autonomous_database_state" { default = "AVAILABLE" }

data "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
}
                ` + compartmentIdVariableStr + AutonomousDatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_strings.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_name", "adatabasedb1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
				ExpectNonEmptyPlan: true,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_password",
					"lifecycle_details",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatabaseAutonomousDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_database" {
			noResourceFound = false
			request := oci_database.GetAutonomousDatabaseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDatabaseId = &tmp

			response, err := client.GetAutonomousDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousDatabaseLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
