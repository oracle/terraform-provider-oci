// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

func TestResourceDatabaseDBSystemIntelX7ToX9(t *testing.T) {
	const sourceDataBaseSystem = `
	resource "oci_database_db_system" "t" {
		availability_domain = "${oci_core_subnet.t.availability_domain}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		database_edition = "ENTERPRISE_EDITION"
		disk_redundancy = "NORMAL"
		shape = "VM.Standard2.4"
		cpu_core_count = "4"
		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
		domain = "${oci_core_subnet.t.subnet_domain_name}"
		hostname = "myOracleDB"
		data_storage_size_in_gb = "256"
		license_model = "LICENSE_INCLUDED"
		node_count = "1"
		display_name = "tfDbSystemTest"
		db_home {
			db_version = "21.8.0.0"
			display_name = "dbHome1"
			database {
				admin_password = "BEstrO0ng_#11"
				db_name = "aTFdb"
				character_set = "AL32UTF8"
				ncharacter_set = "AL16UTF16"
				db_workload = "OLTP"
				pdb_name = "pdbName"
				db_backup_config {
					auto_backup_enabled = false
				}
			}
		}
	}`

	const targetDataBaseSystem = `
    	resource "oci_database_db_system" "t" {
    		availability_domain = "${oci_core_subnet.t.availability_domain}"
    		compartment_id = "${var.compartment_id}"
    		subnet_id = "${oci_core_subnet.t.id}"
    		database_edition = "ENTERPRISE_EDITION"
    		disk_redundancy = "NORMAL"
    		shape = "VM.Standard3.Flex"
    		cpu_core_count = "4"
    		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
    		domain = "${oci_core_subnet.t.subnet_domain_name}"
    		hostname = "myOracleDB"
    		data_storage_size_in_gb = "256"
    		license_model = "LICENSE_INCLUDED"
    		node_count = "1"
    		display_name = "tfDbSystemTest"
    		db_home {
    			db_version = "21.8.0.0"
    			display_name = "dbHome1"
    			database {
    				admin_password = "BEstrO0ng_#11"
    				db_name = "aTFdb"
    				character_set = "AL32UTF8"
    				ncharacter_set = "AL16UTF16"
    				db_workload = "OLTP"
    				pdb_name = "pdbName"
    				db_backup_config {
    					auto_backup_enabled = false
    				}
    			}
    		}
    	}`
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: ResourceDatabaseBaseConfig + sourceDataBaseSystem + `
				data "oci_database_databases" "t" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}
				data "oci_database_db_homes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}`,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "storage_volume_performance_mode", "BALANCED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "VM.Standard2.4"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "21.8.0.0"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "db_home.0.display_name"),
			),
		},
		//Migration - Intel X7 to Intel X9
		{
			Config: ResourceDatabaseBaseConfig + targetDataBaseSystem + `
				data "oci_database_databases" "t" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}
				data "oci_database_db_homes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}`,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "storage_volume_performance_mode", "BALANCED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "VM.Standard3.Flex"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "21.8.0.0"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "db_home.0.display_name"),
			),
		},
	})
}
