// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseDbSystemPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.t.id}`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	datasourceName := "data.oci_database_db_system_patch_history_entries.test_db_system_patch_history_entries"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: ResourceDatabaseBaseConfig +
				`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "2"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "BRING_YOUR_OWN_LICENSE"
					node_count = "1"
					fault_domains = ["FAULT-DOMAIN-1"]
        			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							tde_wallet_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
						}
					}
				}` +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_system_patch_history_entries", "test_db_system_patch_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseDbSystemPatchHistoryEntryDataSourceRepresentation),
			//compartmentIdVariableStr + DatabaseDbSystemPatchHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "patch_history_entries.#", "0"),
			),
		},
	})
}
