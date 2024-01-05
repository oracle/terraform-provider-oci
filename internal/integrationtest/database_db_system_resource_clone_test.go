// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestAccResourceDatabaseDBSystem_clone tests DBsystems using Virtual Machines.
// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemClone(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_clone") {
		t.Skip("Skipping suppressed DBSystem_clone")
	}

	httpreplay.SetScenario("TestResourceDatabaseDBSystemClone")
	defer httpreplay.SaveScenario()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	cloneDatabaseDbSystemResourceName := "oci_database_db_system.clone"

	provider := acctest.TestAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify clone VM DbSystem launch
			{
				Config: ResourceDatabaseBaseConfig + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + compartmentIdUVariableStr + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "source" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "VM.Standard2.1"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					display_name = "{{.token}}"
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB" // this will be lowercased server side
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							kms_key_version_id = "${var.kms_key_version_id}"
                            vault_id = "${var.vault_id}"
                            kms_key_id = "${var.kms_key_id}"
							db_name = "aTFdb"
							freeform_tags = {"Department" = "Finance"}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					freeform_tags = {"Department" = "Finance"}
				}`, nil) +
					ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "clone" {
				    source              = "DB_SYSTEM"
                    source_db_system_id = "${oci_database_db_system.source.id}"
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
							kms_key_version_id = "${var.kms_key_version_id}"
                            vault_id = "${var.vault_id}"
                            kms_key_id = "${var.kms_key_id}"
							db_name = "aTFdb"
							freeform_tags = {"Department" = "Finance"}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					freeform_tags = {"Department" = "Finance"}
				}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "id"),
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "source_db_system_id"),
					resource.TestCheckResourceAttr(cloneDatabaseDbSystemResourceName, "source", "DB_SYSTEM"),
					resource.TestCheckResourceAttr(cloneDatabaseDbSystemResourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(cloneDatabaseDbSystemResourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "db_home.0.database.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "db_home.0.database.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(cloneDatabaseDbSystemResourceName, "db_home.0.database.0.vault_id"),
				),
			},
		},
	})
}
