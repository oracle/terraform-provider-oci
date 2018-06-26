// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DbHomeRequiredOnlyResource = DbHomeResourceDependencies + `
resource "oci_database_db_home" "test_db_home" {
	#Required
	database {
		#Required
		admin_password = "${var.db_home_database_admin_password}"
		db_name = "${var.db_home_database_db_name}"
	}
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	db_version = "${var.db_home_db_version}"
}
`

	DbHomeResourceConfig = DbHomeResourceDependencies + `
resource "oci_database_db_home" "test_db_home" {
	#Required
	database {
		#Required
		admin_password = "${var.db_home_database_admin_password}"
		backup_id = "${oci_database_backup.test_backup.id}"
		backup_tde_password = "${var.db_home_database_backup_tde_password}"
		db_name = "${var.db_home_database_db_name}"

		#Optional
		character_set = "${var.db_home_database_character_set}"
		db_backup_config {

			#Optional
			auto_backup_enabled = "${var.db_home_database_db_backup_config_auto_backup_enabled}"
		}
		db_workload = "${var.db_home_database_db_workload}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.db_home_database_defined_tags_value}")}"
		freeform_tags = "${var.db_home_database_freeform_tags}"
		ncharacter_set = "${var.db_home_database_ncharacter_set}"
		pdb_name = "${var.db_home_database_pdb_name}"
	}
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	db_version = "${var.db_home_db_version}"

	#Optional
	display_name = "${var.db_home_display_name}"
	source = "${var.db_home_source}"
}
`
	DbHomePropertyVariables = `
variable "db_home_database_admin_password" { default = "BEstrO0ng_#11" }
variable "db_home_database_backup_tde_password" { default = "BEstrO0ng_#11" }
variable "db_home_database_character_set" { default = "AL32UTF8" }
variable "db_home_database_db_backup_config_auto_backup_enabled" { default = false }
variable "db_home_database_db_name" { default = "myTestDb" }
variable "db_home_database_db_workload" { default = "dbWorkload" }
variable "db_home_database_defined_tags_value" { default = "value" }
variable "db_home_database_freeform_tags" { default = {"Department"= "Finance"} }
variable "db_home_database_ncharacter_set" { default = "AL16UTF16" }
variable "db_home_database_pdb_name" { default = "pdbName" }
variable "db_home_display_name" { default = "createdDbHome" }
variable "db_home_source" { default = "DB_BACKUP" }
variable "db_home_db_version" { default = "12.1.0.2" }

`
	DbHomeResourceDependencies = BackupResourceConfig + BackupPropertyVariables + DefinedTagsDependencies
)

func TestDatabaseDbHomeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"
	datasourceName := "data.oci_database_db_homes.test_db_homes"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + DbHomePropertyVariables + compartmentIdVariableStr + DbHomeRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + DbHomePropertyVariables + compartmentIdVariableStr + DbHomeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "database.0.backup_tde_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_backup_config.0.auto_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_workload", "dbWorkload"),
					resource.TestCheckResourceAttr(resourceName, "database.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source", "DB_BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "db_home_database_admin_password" { default = "BEstrO0ng_#11" }
variable "db_home_database_backup_tde_password" { default = "BEstrO0ng_#11" }
variable "db_home_database_character_set" { default = "AL32UTF8" }
variable "db_home_database_db_backup_config_auto_backup_enabled" { default = false }
variable "db_home_database_db_name" { default = "myTestDb" }
variable "db_home_database_db_workload" { default = "dbWorkload" }
variable "db_home_database_defined_tags_value" { default = "value" }
variable "db_home_database_freeform_tags" { default = {"Department"= "Finance"} }
variable "db_home_database_ncharacter_set" { default = "AL16UTF16" }
variable "db_home_database_pdb_name" { default = "pdbName" }
variable "db_home_display_name" { default = "createdDbHome" }
variable "db_home_source" { default = "DB_BACKUP" }
variable "db_home_db_version" { default = "12.1.0.2" }

data "oci_database_db_homes" "test_db_homes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"

    filter {
    	name = "id"
    	values = ["${oci_database_db_home.test_db_home.id}"]
    }
}
                ` + compartmentIdVariableStr + DbHomeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),

					resource.TestCheckResourceAttr(datasourceName, "db_homes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.db_system_id"),
					resource.TestCheckResourceAttr(datasourceName, "db_homes.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(datasourceName, "db_homes.0.display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.state"),
				),
			},
		},
	})
}
