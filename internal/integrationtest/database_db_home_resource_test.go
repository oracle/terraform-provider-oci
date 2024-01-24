// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// issue-routing-tag: database/default
func TestAccResourceDatabaseDBHomeWithPointInTimeRecovery(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "timeStampForPointInTimeRecovery") {
		t.Skip("This test requires a source DB with automatic backups enabled. " +
			"There should be at least two automatic backups available." +
			"time_stamp_for_point_in_time_recovery time should be anytime after the start time of the 1st automatic backup and before the start time of the last automatic backup.")
	}

	const dbWaitConditionDuration = time.Duration(20 * time.Minute)
	const sourceDataBaseSystem = `
	resource "oci_database_db_system" "src_db_system" {
		availability_domain = "${oci_core_subnet.t.availability_domain}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		database_edition = "ENTERPRISE_EDITION"
		disk_redundancy = "NORMAL"
		shape = "BM.DenseIO2.52"
		cpu_core_count = "2"
		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
		domain = "${oci_core_subnet.t.subnet_domain_name}"
		hostname = "myOracleDB"
		data_storage_size_in_gb = "256"
		license_model = "LICENSE_INCLUDED"
		node_count = "1"
		display_name = "tfDbSystemTest"
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
				}
			}
		}
	}
	data "oci_database_db_homes" "t" {
		compartment_id = "${var.compartment_id}"
		db_system_id = "${oci_database_db_system.src_db_system.id}"
		filter {
			name = "display_name"
			values = ["dbHome1"]
		}
	}`

	var resId string
	resourceName := "oci_database_db_home.test_db_home_source_database"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: ResourceDatabaseBaseConfig + sourceDataBaseSystem + `
				data "oci_database_databases" "db" {
					compartment_id = "${var.compartment_id}"
					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "data.oci_database_databases.db", "databases.0.id")
					return err
				},
			),
		},
		// wait for backup and Create new db from it
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dbAutomaticBackupAvailableWaitCondition, dbWaitConditionDuration,
				listBackupsFetchOperation, "database", false),
			Config: ResourceDatabaseBaseConfig + sourceDataBaseSystem +
				`
				data "oci_database_databases" "db" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
				}

				data "oci_database_backups" "test_backups" {
					database_id = "${data.oci_database_databases.db.databases.0.id}"
				}

				resource "oci_database_db_home" "test_db_home_source_database" {
					database {
						admin_password = "BEstrO0ng_#11"
						backup_tde_password = "BEstrO0ng_#11"
						database_id = "${data.oci_database_databases.db.databases.0.id}"
						db_name = "dbDb"
						time_stamp_for_point_in_time_recovery = "${data.oci_database_backups.test_backups.backups.0.time_ended}"
					}
					db_system_id = "${oci_database_db_system.test_db_system.id}"
					source = "DATABASE"
				}
				`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database.0.database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2.200714"),
				resource.TestCheckResourceAttr(resourceName, "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "database.0.time_stamp_for_point_in_time_recovery"),
			),
		},
	})
}

// Creates a oci_database_db_home resource under a Cloud VM Cluster (also known as an ExaCS VM cluster).
// issue-routing-tag: database/default
func TestDatabaseDbHomeResource_createFromCloudVmCluster(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomeResource_createFromCloudVmCluster")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"
	dbHomeRepresentationSourceCloudVmClusterNew := acctest.GetUpdatedRepresentationCopy("vm_cluster_id",
		acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
		dbHomeRepresentationSourceVmClusterNew)

	var resId string

	acctest.ResourceTest(t, testAccCheckDatabaseCloudVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceCloudVmClusterNew),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "source", "VM_CLUSTER_NEW"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseDbHomeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"database.0.admin_password", "source"}, // Db passwords and Source of Db Home creation are not made visible by services
			ResourceName:            resourceName,
		},
	})
}
