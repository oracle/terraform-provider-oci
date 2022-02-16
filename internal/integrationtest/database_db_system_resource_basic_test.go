// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/database"
)

var (
	DbSystemResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Optional, acctest.Create, internetGatewayRepresentation)

	DbSystemResourceConfig = DbSystemResourceDependencies + AvailabilityDomainConfig + DefinedTagsDependencies + `

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
		db_home {
			db_version = "12.1.0.2"
			display_name = "dbHome1"
			database {
				admin_password = "BEstrO0ng_#11"
				tde_wallet_password = "BEstrO0ng_#11"
				db_name = "tfDbName"
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

	ResourceDatabaseBaseConfig = acctest.LegacyTestProviderConfig() + DefinedTagsDependencies + `

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}

	resource "oci_core_route_table" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.t.id}"
		}
	}
	resource "oci_core_internet_gateway" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "-tf-internet-gateway"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.20.0/24"
		display_name        = "TFSubnet1"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet"
	}
	resource "oci_core_subnet" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.21.0/24"
		display_name        = "TFSubnet2"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet2"
	}
	resource "oci_core_network_security_group" "test_network_security_group" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.t.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.t.id}"
	}
    `

	ResourceDatabaseResourceName                   = "oci_database_db_system.t"
	ResourceDatabaseToken, ResourceDatabaseTokenFn = acctest.TokenizeWithHttpReplay("database_db")
)

// issue-routing-tag: database/default
func TestAccResourceDatabaseDBSystemFromDatabase(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	const dbWaitConditionDuration = time.Duration(6 * time.Minute)
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

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: ResourceDatabaseBaseConfig + sourceDataBaseSystem + `
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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dbBackupAvailableWaitCondition, dbWaitConditionDuration,
				listBackupsFetchOperation, "core", false),
			Config: ResourceDatabaseBaseConfig + sourceDataBaseSystem + `
				data "oci_database_databases" "t" {
  					compartment_id = "${var.compartment_id}"
  					db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
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
					source = "DATABASE"
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							backup_tde_password = "BEstrO0ng_#11"
							database_id = "${data.oci_database_databases.t.databases.0.id}"
							db_name = "dbarclog"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.#", "1"),
				//resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.0", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "display_name", "tfDbSystemTestFromBackup"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "dbarclog"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.0.last_backup_timestamp"),
			),
		},
	})
}

// issue-routing-tag: database/default
func TestAccResourceDatabaseDBSystemWithPointInTimeRecovery(t *testing.T) {
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
					//database_id = "ocid1.database.oc1.phx.abyhqljsz2exfqnuimsq4hmio5pgbk2ifsqphexknlmnbnqpc7lcauvejmqq"
					database_id = "${data.oci_database_databases.db.databases.0.id}"
				}

				resource "oci_database_db_system" "t" {
					//availability_domain = "NyKp:PHX-AD-1"
					availability_domain = "${oci_core_subnet.t.availability_domain}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "2"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myoracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					display_name = "tfDbSystemTestFromBackup"
					source = "DATABASE"
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							backup_tde_password = "BEstrO0ng_#11"
							database_id = "${data.oci_database_databases.db.databases.0.id}"
							db_name = "dbarclog"
							time_stamp_for_point_in_time_recovery = "${data.oci_database_backups.test_backups.backups.0.time_ended}"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.#", "1"),
				//resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.0", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "display_name", "tfDbSystemTestFromBackup"),
				//resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "db_home.0.database.0.time_stamp_for_point_in_time_recovery"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				//resource.TestCheckResourceAttrSet("data.oci_database_databases.db", "databases.0.last_backup_timestamp"),
			),
		},
	})
}

func dbAutomaticBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if listBackupResponse, ok := response.Response.(database.ListBackupsResponse); ok {
		if len(listBackupResponse.Items) > 1 {
			return listBackupResponse.Items[1].LifecycleState != database.BackupSummaryLifecycleStateActive
		}
		return true
	}
	return false
}

// TestAccResourceDatabaseDBSystem_basic tests creation of a DBSystem with the minimum required properties
// to assert expected default values are set
// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemBasic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemBasic")
	defer httpreplay.SaveScenario()

	// This test is a subset of TestAccResourceDatabaseDBSystem_allXX. It tests omitting optional params.
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: ResourceDatabaseBaseConfig + `
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
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					return err
				},
			),
		},
		// verify Update without updating nsgIds
		{
			Config: ResourceDatabaseBaseConfig + `
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "4"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify removing nsgIds
		{
			Config: ResourceDatabaseBaseConfig + `
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "4"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							tde_wallet_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "0"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify updateing adminPassword and tdeWalletPassword
		{
			Config: ResourceDatabaseBaseConfig + `
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "4"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#12"
							tde_wallet_password = "BEstrO0ng_#12"
							db_name = "aTFdb"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "0"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		{
			Config: ResourceDatabaseBaseConfig,
		},
		// verify Create without nsgIds and backupNsgIds
		{
			Config: ResourceDatabaseBaseConfig + `
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
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							tde_wallet_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
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
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
			),
		},
		// verify Update without nsgIds and backupNsgIds
		{
			Config: ResourceDatabaseBaseConfig + `
				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					database_edition = "ENTERPRISE_EDITION"
					disk_redundancy = "NORMAL"
					shape = "BM.DenseIO2.52"
					cpu_core_count = "4"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					fault_domains = ["FAULT-DOMAIN-1"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							tde_wallet_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
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
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "4"),
				resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "tfsubnet.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	resource.AddTestSweepers("DatabaseDbSystem", &resource.Sweeper{
		Name:         "DatabaseDbSystem",
		Dependencies: acctest.DependencyGraph["dbSystem"],
		F:            sweepDatabaseDbSystemResource,
	})
}
