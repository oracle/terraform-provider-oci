// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDataGuardAssociationRequiredOnlyResource = DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
		acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Required, acctest.Create, dataGuardAssociationRepresentationExistingDbSystem)

	DatabaseDataGuardAssociationResourceConfig = DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
		acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, dataGuardAssociationRepresentationExistingDbSystem)

	DatabaseDataGuardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"data_guard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_data_guard_association.test_data_guard_association.id}`},
		"database_id":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
	}

	DatabaseDataGuardAssociationDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDataGuardAssociationDataSourceFilterRepresentation}}
	DatabaseDataGuardAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_data_guard_association.test_data_guard_association.id}`}},
	}

	DatabaseDataGuardAssociationRepresentationBase = map[string]interface{}{
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_db_system"}},
		"database_admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id":                      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"delete_standby_db_home_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"protection_mode":                  acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`},
		"transport_type":                   acctest.Representation{RepType: acctest.Required, Create: `ASYNC`},
		"cpu_core_count":                   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"data_collection_options":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: dataGuardAssociationDataCollectionOptionsRepresentation},
		"is_active_data_guard_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"storage_volume_performance_mode":  acctest.Representation{RepType: acctest.Optional, Create: `BALANCED`},
	}

	dataGuardAssociationRepresentationBaseForExadata = map[string]interface{}{
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_db_system.test_db_system"}},
		"database_admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id":                      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"delete_standby_db_home_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"protection_mode":                  acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
		"transport_type":                   acctest.Representation{RepType: acctest.Required, Create: `ASYNC`, Update: `SYNC`},
		"is_active_data_guard_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	dataGuardAssociationRepresentationExistingDbSystem = acctest.RepresentationCopyWithNewProperties(DatabaseDataGuardAssociationRepresentationBase, map[string]interface{}{
		"depends_on":        acctest.Representation{RepType: acctest.Required, Create: []string{`oci_database_db_system.test_db_system`, `oci_database_db_system.test_db_system2`}},
		"creation_type":     acctest.Representation{RepType: acctest.Required, Create: `ExistingDbSystem`},
		"peer_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system2.id}`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataGuardAssociationRepresentationExistingDbSystem},
	})

	dataGuardAssociationDataCollectionOptionsRepresentation = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ignoreDataGuardAssociationRepresentationExistingDbSystem = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`database_defined_tags`, `database_freeform_tags`, `db_system_defined_tags`, `db_system_freeform_tags`, `fault_domains`, `license_model`, `node_count`, `private_ip`, `time_zone`}},
	}

	dataGuardAssociationRepresentationNewDbSystem = acctest.RepresentationCopyWithNewProperties(DatabaseDataGuardAssociationRepresentationBase, map[string]interface{}{
		"creation_type":           acctest.Representation{RepType: acctest.Required, Create: `NewDbSystem`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `hostname`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.2`},
		"backup_network_nsg_ids":  acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"nsg_ids":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"database_defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "databaseDefinedTags1")}`},
		"database_freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"databaseFreeformTagsK": "databaseFreeformTagsV"}},
		"db_system_defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "dbSystemDefinedTags1")}`},
		"db_system_freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"dbSystemFreeformTagsK": "dbSystemFreeformTagsV"}},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-3`}},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `BRING_YOUR_OWN_LICENSE`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"private_ip":              acctest.Representation{RepType: acctest.Optional, Create: `10.0.2.223`},
		"time_zone":               acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
	})

	DatabaseDataGuardAssociationResourceDependenciesBase = AvailabilityDomainConfig + DefinedTagsDependencies + CoreVcnResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Optional, acctest.Update, CoreNetworkSecurityGroupRepresentation) +
		`
			#dataguard requires the some port to be open on the subnet
			resource "oci_core_subnet" "test_subnet" {
				availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
				cidr_block          = "10.0.2.0/24"
				display_name        = "TFADSubnet"
				dns_label           = "adsubnet"
				compartment_id      = "${var.compartment_id}"
				vcn_id              = "${oci_core_vcn.test_vcn.id}"
				security_list_ids   = ["${oci_core_security_list.test_security_list.id}"]
				route_table_id      = "${oci_core_vcn.test_vcn.default_route_table_id}"
				dhcp_options_id     = "${oci_core_vcn.test_vcn.default_dhcp_options_id}"
			}
	
			resource "oci_core_security_list" "test_security_list" {
				compartment_id = "${var.compartment_id}"
				vcn_id         = "${oci_core_vcn.test_vcn.id}"
				display_name   = "TFExampleSecurityList"

				// allow outbound tcp traffic on all ports
				egress_security_rules {
					destination = "0.0.0.0/0"
					protocol    = "6"
				}
				
				ingress_security_rules {
					protocol  = "6"
					source    = "0.0.0.0/0"
				}
			}
	
			data "oci_database_databases" "db" {
				compartment_id = "${var.compartment_id}"
				db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
			}
	
			data "oci_database_db_homes" "t" {
				compartment_id = "${var.compartment_id}"
				db_system_id = "${oci_database_db_system.test_db_system.id}"
				filter {
					name = "display_name"
					values = ["TFTestDbHome1"]
				}
			}
		`

	DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem = DatabaseDataGuardAssociationResourceDependenciesBase +
		`
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
				display_name = "TFTestDbSystemBM1"
				db_home {
					db_version = "12.1.0.2"
					display_name = "TFTestDbHome1"
					database {
						admin_password = "BEstrO0ng_#11"
						db_name = "tfDbName"
					}
				}
			}

			resource "oci_database_db_system" "test_db_system2" {
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
				display_name = "TFTestDbSystemBM2"
				db_home {
					db_version = "12.1.0.2"
					display_name = "TFTestDbHome1"
					database {
						admin_password = "BEstrO0ng_#11"
						db_name = "db2"
					}
				}
			}

			data "oci_database_db_systems" "t" {
				compartment_id = "${var.compartment_id}"
				filter {
					name   = "id"
					values = ["${oci_database_data_guard_association.test_data_guard_association.peer_db_system_id}"]
				}
				filter {
					name   = "state"
					values = ["AVAILABLE"]
				}
			}
		`

	DatabaseDataGuardAssociationResourceDependenciesNewDbSystem = DatabaseDataGuardAssociationResourceDependenciesBase +
		`
			resource "oci_database_db_system" "test_db_system" {
				availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
				compartment_id = "${var.compartment_id}"
				subnet_id = "${oci_core_subnet.test_subnet.id}"
				database_edition = "ENTERPRISE_EDITION"
				disk_redundancy = "NORMAL"
				shape = "VM.Standard2.2"
				cpu_core_count = "2"
				ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
				domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
				hostname = "myOracleDB"
				data_storage_size_in_gb = "256"
				license_model = "LICENSE_INCLUDED"
				node_count = "1"
				display_name = "TFTestDbSystemVM"
				db_home {
					db_version = "12.1.0.2"
					display_name = "TFTestDbHome1"
					database {
						admin_password = "BEstrO0ng_#11"
						db_name = "tfDbName"
					}
				}
			}
		`

	DataGuardAssociationResourceDependencies = DatabaseDataGuardAssociationResourceDependenciesBase + `
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
	display_name = "TFTestDbSystemBM1"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			admin_password = "BEstrO0ng_#11"
			db_name = "tfDbName"
		}
	}
}

resource "oci_database_db_system" "test_db_system2" {
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
	display_name = "TFTestDbSystemBM2"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			admin_password = "BEstrO0ng_#11"
			db_name = "db2"
		}
	}
}
`
	DataGuardAssociationResourceDependenciesNewDbSystem = DatabaseDataGuardAssociationResourceDependenciesBase + `
resource "oci_database_db_system" "test_db_system" {
	availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	database_edition = "ENTERPRISE_EDITION"
	disk_redundancy = "NORMAL"
	shape = "VM.Standard2.2"
	cpu_core_count = "2"
	ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
	domain = "${oci_core_subnet.test_subnet.subnet_domain_name}"
	hostname = "myOracleDB"
	data_storage_size_in_gb = "256"
	license_model = "LICENSE_INCLUDED"
	node_count = "1"
	display_name = "TFTestDbSystemVM"
	db_home {
		db_version = "12.1.0.2"
		display_name = "TFTestDbHome1"
		database {
			admin_password = "BEstrO0ng_#11"
			db_name = "tfDbName"
		}
	}
}
`
)

// issue-routing-tag: database/default
func TestDatabaseDataGuardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDataGuardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_data_guard_association"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem+
		acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Create, dataGuardAssociationRepresentationExistingDbSystem), "database", "dataGuardAssociation", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create NewDbSystem
		{
			Config: config + compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesNewDbSystem +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Create, dataGuardAssociationRepresentationNewDbSystem),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr("data.oci_database_db_systems.t", "db_systems.0.backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "NewDbSystem"),
				resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				resource.TestCheckResourceAttr(resourceName, "db_system_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_system_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip", "10.0.2.223"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_active_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesNewDbSystem,
		},
		// verify Create with optionals on Existing DbSystem
		{
			Config: config + compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Create, dataGuardAssociationRepresentationExistingDbSystem),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingDbSystem"),
				resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_active_data_guard_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_volume_performance_mode", "BALANCED"),
				resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, dataGuardAssociationRepresentationExistingDbSystem),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingDbSystem"),
				resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "delete_standby_db_home_on_delete", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_volume_performance_mode", "BALANCED"),
				resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_data_guard_associations", acctest.Optional, acctest.Update, DatabaseDataGuardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, dataGuardAssociationRepresentationExistingDbSystem),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Required, acctest.Create, DatabaseDataGuardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDataGuardAssociationResourceDependenciesExistingDbSystem +
				acctest.GenerateResourceFromRepresentationMap("oci_database_data_guard_association", "test_data_guard_association", acctest.Optional, acctest.Update, dataGuardAssociationRepresentationExistingDbSystem),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_data_guard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
			),
		},
	})
}
