// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	AutonomousDataWarehouseBackupResourceConfig = AutonomousDataWarehouseBackupResourceDependencies + `
resource "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
}
`
	AutonomousDataWarehouseBackupPropertyVariables = `
variable "autonomous_data_warehouse_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_data_warehouse_backup_state" { default = "AVAILABLE" }

`
	AutonomousDataWarehouseBackupResourceDependencies = AutonomousDataWarehousePropertyVariables + AutonomousDataWarehouseResourceConfig
)

func TestDatabaseAutonomousDataWarehouseBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup"
	datasourceName := "data.oci_database_autonomous_data_warehouse_backups.test_autonomous_data_warehouse_backups"
	singularDatasourceName := "data.oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup"

	testResourceName := GenerateTestResourceName("adwdb1", 14)
	setEnvSetting("TF_VAR_autonomous_data_warehouse_db_name", testResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + AutonomousDataWarehouseBackupPropertyVariables + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "autonomous_data_warehouse_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_data_warehouse_backup_state" { default = "ACTIVE" }

data "oci_database_autonomous_data_warehouse_backups" "test_autonomous_data_warehouse_backups" {

	#Optional
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
	state = "${var.autonomous_data_warehouse_backup_state}"

    filter {
    	name = "id"
    	values = ["${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}"]
    }
}
                ` + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouse_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouse_backups.0.display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.is_automatic"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "autonomous_data_warehouse_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_data_warehouse_backup_state" { default = "AVAILABLE" }

data "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_backup_id = "${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}"
}
                ` + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_automatic", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + `
variable "autonomous_data_warehouse_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_data_warehouse_backup_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
