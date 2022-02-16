// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/database"
)

// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemFromBackup(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	httpreplay.SetScenario("TestResourceDatabaseDBSystemFromBackup")
	defer httpreplay.SaveScenario()
	const DBWaitConditionDuration = time.Duration(20 * time.Minute)
	const DataBaseSystemWithBackup = `
	resource "oci_database_db_system" "test_db_system" {
		availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.test_subnet.id}"
		database_edition = "ENTERPRISE_EDITION"
		disk_redundancy = "NORMAL"
		shape = "BM.DenseIO2.52"
		cpu_core_count = "2"
		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
		domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
		hostname = "myOracleDB"
		data_storage_size_in_gb = "256"
		license_model = "LICENSE_INCLUDED"
		node_count = "1"
		display_name = "tfDbSystemTest"
        nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
		db_home {
			db_version = "12.1.0.2"
			display_name = "dbHome1"
			database {
				admin_password = "BEstrO0ng_#11"
				db_name = "aTFdb"
				character_set = "AL32UTF8"
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
	}
	data "oci_database_db_homes" "t" {
		compartment_id = "${var.compartment_id}"
		db_system_id = "${oci_database_db_system.test_db_system.id}"
		filter {
			name = "display_name"
			values = ["dbHome1"]
		}
	}`

	var resId string

	var ResourceDatabaseResourceNameImg string = "oci_database_db_system.tImg"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: ResourceDatabaseBaseConfig + DbSystemResourceDependencies + DataBaseSystemWithBackup + AvailabilityDomainConfig + `
				data "oci_database_databases" "t" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "data.oci_database_databases.t", "databases.0.id")
					return err
				},
			),
		},
		// wait for backup and Create new db from it
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dbBackupAvailableWaitCondition, DBWaitConditionDuration,
				listBackupsFetchOperation, "core", false),
			Config: ResourceDatabaseBaseConfig + DbSystemResourceDependencies + DataBaseSystemWithBackup + AvailabilityDomainConfig + `
				data "oci_database_databases" "t" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}
				data "oci_database_backups" "test_backups"{
					database_id = "${data.oci_database_databases.t.databases.0.id}"

				}
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
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					display_name = "tfDbSystemTestFromBackup"
					source = "DB_BACKUP"
        			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							backup_tde_password = "BEstrO0ng_#11"
							backup_id = "${data.oci_database_backups.test_backups.backups.0.id}"
							db_name = "dbback"
						}
					}
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "display_name", "tfDbSystemTestFromBackup"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "dbback"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
			),
		},

		{
			Config: ResourceDatabaseBaseConfig + DbSystemResourceDependencies + DataBaseSystemWithBackup + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageDataSourceRepresentationForVmBmShape) + `
                data "oci_database_databases" "t" {
                    compartment_id = "${var.compartment_id}"
                    db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
                }
                data "oci_database_backups" "test_backups"{
                    database_id = "${data.oci_database_databases.t.databases.0.id}"

                }
                resource "oci_database_db_system" "tImg" {
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
                    license_model = "LICENSE_INCLUDED"
                    node_count = "1"
                    display_name = "tfDbSystemTestFromBackupImg"
                    source = "DB_BACKUP"
                    nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
                    db_home {
                        db_version = "12.1.0.2"
                        database_software_image_id = "${oci_database_database_software_image.test_database_software_image.id}"
                        database {
                            admin_password = "BEstrO0ng_#11"
                            backup_tde_password = "BEstrO0ng_#11"
                            backup_id = "${data.oci_database_backups.test_backups.backups.0.id}"
                            db_name = "dbback"
                        }
                    }
                }`,

			Check: resource.ComposeAggregateTestCheckFunc(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "time_created"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "db_home.0.database_software_image_id"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "database_edition", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "shape", "BM.DenseIO2.52"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "display_name", "tfDbSystemTestFromBackupImg"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceNameImg, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "db_home.0.database.0.db_name", "dbback"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceNameImg, "nsg_ids.#", "1"),
			),
		},
	})
}

func dbBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if listBackupResponse, ok := response.Response.(database.ListBackupsResponse); ok {
		if len(listBackupResponse.Items) > 0 {
			return listBackupResponse.Items[0].LifecycleState != database.BackupSummaryLifecycleStateActive
		}
		return true
	}
	return false
}

func listBackupsFetchOperation(client *client.OracleClients, databaseId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.DatabaseClient().ListBackups(context.Background(), database.ListBackupsRequest{
		DatabaseId: databaseId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
