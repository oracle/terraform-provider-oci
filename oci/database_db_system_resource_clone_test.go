// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// TestAccResourceDatabaseDBSystem_clone tests DBsystems using Virtual Machines.
func TestResourceDatabaseDBSystemClone(t *testing.T) {
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_clone") {
		t.Skip("Skipping suppressed DBSystem_clone")
	}

	httpreplay.SetScenario("TestResourceDatabaseDBSystemClone")
	defer httpreplay.SaveScenario()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	source_db_system_id := getEnvSettingWithBlankDefault("source_db_system_id")
	sourceDbSystemIdVariableStr := fmt.Sprintf("variable \"source_db_system_id\" { default = \"%s\" }\n", source_db_system_id)

	provider := testAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify clone VM DbSystem launch
			{
				Config: ResourceDatabaseBaseConfig + sourceDbSystemIdVariableStr + compartmentIdUVariableStr + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
				    source              = "DB_SYSTEM"
                    source_db_system_id = "${var.source_db_system_id}"
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id_for_update}"
					subnet_id = "${oci_core_subnet.t.id}"
					shape = "VM.Standard2.1"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					display_name = "{{.token}}"
					hostname = "myOracleDB" // this will be lowercased server side
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							freeform_tags = {"Department" = "Finance"}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					freeform_tags = {"Department" = "Finance"}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "source_db_system_id"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "source", "DB_SYSTEM"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "compartment_id", compartmentIdU),
				),
			},
		},
	})
}
