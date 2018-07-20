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
		CheckDestroy: testAccCheckDatabaseBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + BackupPropertyVariables + compartmentIdVariableStr + BackupResourceConfig,
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
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}

func testAccCheckDatabaseBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_backup" {
			noResourceFound = false
			request := oci_database.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			_, err := client.GetBackup(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
