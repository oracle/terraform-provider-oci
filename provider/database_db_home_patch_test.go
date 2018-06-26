// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DbHomePatchResourceConfig = DbHomePatchResourceDependencies

	DbHomePatchResourceDependencies = SubnetRequiredOnlyResource + SubnetPropertyVariables + `
resource "oci_database_db_system" "test_db_system" {
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	database_edition = "ENTERPRISE_EDITION"
	disk_redundancy = "NORMAL"
	shape = "BM.DenseIO1.36"
	cpu_core_count = "2"
	ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
	domain = "${var.subnet_dns_label}.oraclevcn.com"
	hostname = "myOracleDB"
	data_storage_size_in_gb = "256"
	license_model = "LICENSE_INCLUDED"
	node_count = "1"
	display_name = "tfDbSystemTest"
	db_home {
		db_version = "12.1.0.2"
		display_name = "dbHome1"
		database {
			"admin_password" = "BEstrO0ng_#11"
			"db_name" = "tfDbName"
		}
	}
}

data "oci_database_db_homes" "t" {
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	filter {
		name = "display_name"
		values = ["dbHome1"]
	}
}
`
)

func TestDatabaseDbHomePatchResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_home_patches.test_db_home_patches"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_database_db_home_patches" "test_db_home_patches" {
	#Required
	db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}
                ` + compartmentIdVariableStr + DbHomePatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
				),
			},
		},
	})
}
