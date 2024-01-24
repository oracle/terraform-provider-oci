// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/database"
)

// TestAccResourceDatabaseDBSystem_allVM tests DBsystems using Virtual Machines.
// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemAllVM(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_allVM") {
		t.Skip("Skipping suppressed DBSystem_allVM")
	}

	httpreplay.SetScenario("TestResourceDatabaseDBSystemAllVM")
	defer httpreplay.SaveScenario()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	var resId, resId2 string

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: ResourceDatabaseBaseConfig + kmsKeyIdVariableStr + kmsKeyVersionIdVariableStr + vaultIdVariableStr + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					//backup_subnet_id = "${oci_core_subnet.t2.id}" // this requires a specific shape
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
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							character_set = "AL32UTF8"
							defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
							freeform_tags = {"Department" = "Finance"}
							ncharacter_set = "AL16UTF16"
							db_workload = "OLTP"
							kms_key_id = "${var.kms_key_id}"
                            vault_id = "${var.vault_id}"
                            kms_key_version_id = "${var.kms_key_version_id}"
                            pdb_name = "pdbName"
							db_backup_config {
								auto_backup_enabled = true
								auto_backup_window = "SLOT_TWO"
								recovery_window_in_days = 10
							}
						}
					}
					data_collection_options {
						is_diagnostics_events_enabled = true
						is_health_monitoring_enabled = false
						is_incident_logs_enabled = true
					}
					db_system_options {
						storage_management = "LVM"
					}
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}
				data "oci_database_db_systems" "t" {
					compartment_id = "${var.compartment_id}"
					filter {
						name   = "id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_homes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_home" "t" {
					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
				}
				data "oci_database_databases" "t" {
					compartment_id = "${var.compartment_id}"
					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
					filter {
						name   = "db_name"
						values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
					}
				}
				data "oci_database_database" "t" {
					  database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
				data "oci_database_db_nodes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_node" "t" {
					db_node_id = "${data.oci_database_db_nodes.t.db_nodes.0.id}"
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "display_name", ResourceDatabaseToken),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_percentage", "80"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "reco_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "listener_port", "1521"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.display_name", "-tf-db-home"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.defined_tags.example-tag-namespace-all.example-tag", "originalValue"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_system_options.0.storage_management", "LVM"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "state"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),

				// Data Source tests
				// DBSystems Data Source
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.#"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.availability_domain"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.compartment_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.subnet_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.time_created"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.cpu_core_count"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.ssh_public_keys.0", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.display_name", ResourceDatabaseToken),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.db_system_options.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.data_collection_options.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.data_storage_percentage", "80"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.node_count", "1"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.reco_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.listener_port", "1521"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.defined_tags.example-tag-namespace-all.example-tag", "originalValue"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.nsg_ids.#", "1"),

				/* The following fields are null when retrieved via data source. Some were never populated, some nulls might be BM vs VM behavior.
				   maybe LIST summary vs GET behavior
					"backupSubnetId":null,
					"clusterName":null,
					"lastPatchHistoryEntryId":null,
					"lifecycleDetails":null,
					"scanDnsRecordId":null,
					"scanIpIds":null,
					"vipIds":null
				*/
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.data_storage_size_in_gb"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.node_count"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.state", string(database.DbSystemLifecycleStateAvailable)),

				// DB Systems nested DB Home fields are not supported on the data source, so tests like below wont work if/until fetching the sub resource is implemented
				//resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.0.db_home"),
				//resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.db_home.0.db_version", "12.1.0.2"),

				// DBHomes
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.#"),
				resource.TestCheckResourceAttr("data.oci_database_db_homes.t", "db_homes.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.compartment_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.db_version"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.id"),
				//resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.last_patch_history_entry_id"), // missing when null
				resource.TestCheckResourceAttr("data.oci_database_db_homes.t", "db_homes.0.state", string(database.DbHomeSummaryLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.0.db_system_id"),

				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_version"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "id"),
				//resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "last_patch_history_entry_id"), // missing when null
				resource.TestCheckResourceAttr("data.oci_database_db_home.t", "state", string(database.DbHomeSummaryLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_system_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.t", "display_name", "-tf-db-home"),

				// Databases
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.#"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.db_home_id"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.db_unique_name"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.id"),
				//resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.lifecycle_details"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.state"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.vault_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.connection_strings.0.cdb_default"),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.connection_strings.0.all_connection_strings.cdbDefault"),

				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_home_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "db_name", "aTFdb"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "db_unique_name"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				//resource.TestCheckResourceAttrSet("data.oci_database_database.t", "lifecycle_details"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "pdb_name", "pdbName"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "kms_key_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "kms_key_version_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "vault_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "connection_strings.0.cdb_default"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "connection_strings.0.all_connection_strings.cdbDefault"),

				// DB Nodes
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_system_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.#"),
				resource.TestCheckResourceAttr("data.oci_database_db_nodes.t", "db_nodes.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.db_node_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.db_system_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.hostname"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.state"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.fault_domain"),
				//resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.vnic_id"), // believe this is null when using FAKEHOSTSERIAL header
				//resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.0.backup_vnic_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_nodes.t", "db_nodes.0.software_storage_size_in_gb", "200"),

				// DB Node
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "db_node_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "db_system_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "hostname"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "state"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "time_created"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "fault_domain"),
				//resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "vnic_id"), // believe this is null when using FAKEHOSTSERIAL header
				//resource.TestCheckResourceAttrSet("data.oci_database_db_node.t", "backup_vnic_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_node.t", "software_storage_size_in_gb", "200"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, ResourceDatabaseResourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: ResourceDatabaseBaseConfig + compartmentIdUVariableStr + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id_for_update}"
					subnet_id = "${oci_core_subnet.t.id}"
					//backup_subnet_id = "${oci_core_subnet.t2.id}" // this requires a specific shape
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
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							character_set = "AL32UTF8"
							defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
							freeform_tags = {"Department" = "Finance"}
							ncharacter_set = "AL16UTF16"
							db_workload = "OLTP"
							pdb_name = "pdbName"
							db_backup_config {
								auto_backup_enabled = true
								auto_backup_window = "SLOT_TWO"
								recovery_window_in_days = 10
							}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "compartment_id", compartmentIdU),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// switch compartment back to use data sources in the next step
		{
			Config: ResourceDatabaseBaseConfig + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					//backup_subnet_id = "${oci_core_subnet.t2.id}" // this requires a specific shape
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
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							character_set = "AL32UTF8"
							defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
							freeform_tags = {"Department" = "Finance"}
							ncharacter_set = "AL16UTF16"
							db_workload = "OLTP"
							pdb_name = "pdbName"
							db_backup_config {
								auto_backup_enabled = true
								auto_backup_window = "SLOT_TWO"
								recovery_window_in_days = 10
							}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "compartment_id", compartmentId),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify auto_backup_window parameter update
		{
			Config: ResourceDatabaseBaseConfig + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					//backup_subnet_id = "${oci_core_subnet.t2.id}" // this requires a specific shape
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
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							character_set = "AL32UTF8"
							defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
							freeform_tags = {"Department" = "Finance"}
							ncharacter_set = "AL16UTF16"
							db_workload = "OLTP"
							pdb_name = "pdbName"
							db_backup_config {
								auto_backup_enabled = true
								auto_backup_window = "SLOT_THREE"
								recovery_window_in_days = 10
							}
						}
					}
					db_system_options {
						storage_management = "LVM"
					}
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_window", "SLOT_THREE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify Update
		{
			Config: ResourceDatabaseBaseConfig + ResourceDatabaseTokenFn(`
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					//backup_subnet_id = "${oci_core_subnet.t2.id}" // this requires a specific shape
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "VM.Standard2.2"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					display_name = "{{.token}}"
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB" // this will be lowercased server side
					data_storage_size_in_gb = "512"
					reco_storage_size_in_gb = "512"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "19.0.0.0"
						display_name = "-tf-db-home"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
							character_set = "AL32UTF8"
							defined_tags = "${map("example-tag-namespace-all.example-tag", "updateValue")}"
							freeform_tags = {"Department" = "Admin"}
							ncharacter_set = "AL16UTF16"
							db_workload = "OLTP"
							pdb_name = "pdbName"
							db_backup_config {
								auto_backup_enabled = false
								recovery_window_in_days = 10
							}
						}
					}
					data_collection_options {
						is_diagnostics_events_enabled = true
						is_health_monitoring_enabled = false
						is_incident_logs_enabled = true
					}
					db_system_options {
						storage_management = "LVM"
					}
					defined_tags = "${map("example-tag-namespace-all.example-tag", "updateValue")}"
					freeform_tags = {"Department" = "Admin"}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}", "${oci_core_network_security_group.test_network_security_group2.id}"]
				}
				data "oci_database_db_systems" "t" {
					compartment_id = "${var.compartment_id}"
					filter {
						name   = "id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_homes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_home" "t" {
					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
				}
				data "oci_database_databases" "t" {
					compartment_id = "${var.compartment_id}"
					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
					filter {
						name   = "db_name"
						values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
					}
				}
				data "oci_database_database" "t" {
					  database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
				data "oci_database_db_nodes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
					filter {
						name   = "db_system_id"
						values = ["${oci_database_db_system.t.id}"]
					}
				}
				data "oci_database_db_node" "t" {
					db_node_id = "${data.oci_database_db_nodes.t.db_nodes.0.id}"
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "cpu_core_count"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "display_name", ResourceDatabaseToken),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "512"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "reco_storage_size_in_gb", "512"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_percentage", "80"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "listener_port", "1521"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.display_name", "-tf-db-home"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.defined_tags.example-tag-namespace-all.example-tag", "updateValue"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.freeform_tags.Department", "Admin"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DbSystemLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "freeform_tags.Department", "Admin"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
	})
}
