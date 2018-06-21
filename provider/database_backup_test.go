// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BackupResourceConfig = BackupResourceDependencies + `

resource "oci_database_backup" "test_backup" {
	#Required
	database_id = "${data.oci_database_databases.db.databases.0.id}"
	display_name = "${var.backup_display_name}"
}
`
	BackupPropertyVariables = `
variable "backup_display_name" { default = "Monthly Backup" }

`
	BackupResourceDependencies = DbHomePatchResourceDependencies + `
data "oci_database_databases" "db" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}`
)

func TestDatabaseBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup.test_backup"
	datasourceName := "data.oci_database_backups.test_backups"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + BackupPropertyVariables + compartmentIdVariableStr + BackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "backup_display_name" { default = "Monthly Backup" }

data "oci_database_backups" "test_backups" {

	#Optional
	database_id = "${data.oci_database_databases.db.databases.0.id}"

    filter {
    	name = "id"
    	values = ["${oci_database_backup.test_backup.id}"]
    }
}
                ` + compartmentIdVariableStr + BackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),
					resource.TestCheckResourceAttr(datasourceName, "backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_id"),
					resource.TestCheckResourceAttr(datasourceName, "backups.0.display_name", "Monthly Backup"),
				),
			},
		},
	})
}
