// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	AutonomousDatabaseBackupResourceConfig = AutonomousDatabaseBackupResourceDependencies + `
resource "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	display_name = "${var.autonomous_database_backup_display_name}"
}
`
	AutonomousDatabaseBackupPropertyVariables = `
variable "autonomous_database_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_database_backup_state" { default = "ACTIVE" }

`
	AutonomousDatabaseBackupResourceDependencies = AutonomousDatabasePropertyVariables + AutonomousDatabaseResourceConfig
)

func TestDatabaseAutonomousDatabaseBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_backup.test_autonomous_database_backup"
	datasourceName := "data.oci_database_autonomous_database_backups.test_autonomous_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_database_backup.test_autonomous_database_backup"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + AutonomousDatabaseBackupPropertyVariables + compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "autonomous_database_backup_display_name" { default = "Monthly Backup" }
variable "autonomous_database_backup_state" { default = "ACTIVE" }

data "oci_database_autonomous_database_backups" "test_autonomous_database_backups" {

	#Optional
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	display_name = "${var.autonomous_database_backup_display_name}"
	state = "${var.autonomous_database_backup_state}"

    filter {
    	name = "id"
    	values = ["${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}"]
    }
}
                ` + compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "Monthly Backup"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.autonomous_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.0.display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.is_automatic"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "autonomous_database_backup_display_name" { default = "Monthly Backup" }

data "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_backup_id = "${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}"
}
                ` + compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_automatic", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "FULL"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
				ExpectNonEmptyPlan:      true,
			},
		},
	})
}
