// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dataGuardAssociationSingularExadataDataSourceRepresentation = map[string]interface{}{
		"data_guard_association_id": Representation{repType: Required, create: `${oci_database_data_guard_association.test_exadata_data_guard_association.id}`},
		"database_id":               Representation{repType: Required, create: `${data.oci_database_databases.exadb.databases.0.id}`},
	}

	dataGuardAssociationExadataDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{repType: Required, create: `${data.oci_database_databases.exadb.databases.0.id}`},
		"filter":      RepresentationGroup{Required, dataGuardAssociationExadataDataSourceFilterRepresentation}}
	dataGuardAssociationExadataDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_data_guard_association.test_exadata_data_guard_association.id}`}},
	}

	dataGuardAssociationRepresentationExistingExadataDbSystem = representationCopyWithNewProperties(dataGuardAssociationRepresentationExistingDbSystem, map[string]interface{}{
		"depends_on":        Representation{repType: Required, create: []string{"oci_database_db_system.test_exadata_db_system", `oci_database_db_system.test_exadata_db_system2`}},
		"database_id":       Representation{repType: Required, create: `${data.oci_database_databases.exadb.databases.0.id}`},
		"creation_type":     Representation{repType: Required, create: `ExistingDbSystem`},
		"peer_db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_exadata_db_system2.id}`},
	})

	dataGuardAssociationRepresentationExistingExadataDbHome = representationCopyWithNewProperties(dataGuardAssociationRepresentationExistingDbSystem, map[string]interface{}{
		"depends_on":        Representation{repType: Required, create: []string{"oci_database_db_system.test_exadata_db_system", `oci_database_db_system.test_exadata_db_system2`}},
		"database_id":       Representation{repType: Required, create: `${oci_database_db_system.test_exadata_db_system.db_home.0.database.0.id}`},
		"creation_type":     Representation{repType: Required, create: `ExistingDbSystem`},
		"peer_db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_exadata_db_system2.id}`},
		"peer_db_home_id":   Representation{repType: Required, create: `${oci_database_db_system.test_exadata_db_system2.db_home.0.id}`},
	})

	ExadataBaseDependenciesForDataGuardWithExistingVMCluster = ExadataBaseDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
			representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
				"activation_file": Representation{repType: Optional, update: activationFilePath},
				"shape":           Representation{repType: Required, create: `ExadataCC.Quarter2.92`},
			}),
		) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_exadata_vm_cluster", Required, Create,
			representationCopyWithNewProperties(vmClusterRepresentation, map[string]interface{}{
				"depends_on":   Representation{repType: Required, create: []string{`oci_database_vm_cluster_network.test_vm_cluster_network`}},
				"display_name": Representation{repType: Required, create: `vmClusterForPrimaryDB`},
			})) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_exadata_vm_cluster_for_standby_db", Required, Create,
			representationCopyWithNewProperties(vmClusterRepresentation, map[string]interface{}{
				"depends_on":   Representation{repType: Required, create: []string{`oci_database_vm_cluster_network.test_vm_cluster_network`}},
				"display_name": Representation{repType: Required, create: `vmClusterForStandbyDB`},
			}))

	dataGuardAssociationRepresentationExistingExadataVmCluster = representationCopyWithNewProperties(dataGuardAssociationRepresentationBase, map[string]interface{}{
		"depends_on":         Representation{repType: Required, create: []string{`oci_database_vm_cluster.test_exadata_vm_cluster`, `oci_database_vm_cluster.test_exadata_vm_cluster_for_standby_db`}},
		"database_id":        Representation{repType: Required, create: `${data.oci_database_databases.exadb.databases.0.id}`},
		"creation_type":      Representation{repType: Required, create: `ExistingVmCluster`},
		"peer_vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_exadata_vm_cluster_for_standby_db.id}`},
	})

	ExadataBaseDependencies = DefinedTagsDependencies + `
	#dataguard requires the port to be open on the subnet
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}
	data "oci_identity_availability_domain" "ad" {
		compartment_id 		= "${var.compartment_id}"
		ad_number      		= 1
	}
	resource "oci_core_subnet" "exadata_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.22.0/24"
		display_name        = "ExadataSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}", "${oci_core_security_list.exadata_shapes_security_list.id}"]
		dns_label           = "subnetexadata"
	}

	resource "oci_core_subnet" "exadata_backup_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.23.0/24"
		display_name        = "ExadataBackupSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "subnetexadata2"
	}

	resource "oci_core_security_list" "exadata_shapes_security_list" {
		compartment_id = "${var.compartment_id}"
		vcn_id         = "${oci_core_virtual_network.t.id}"
		display_name   = "ExadataSecurityList"

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "6"
		}

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "1"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "6"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "1"
		}
	}
`
)

func TestResourceDatabaseDataGuardAssociation_Exadata(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDataGuardAssociation_Exadata")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DataGuardAssociation_Exadata") {
		t.Skip("Skipping suppressed DataGuardAssociation_Exadata")
	}

	provider := testAccProvider
	config := testProviderConfig() + ExadataBaseDependencies + `
	data "oci_database_databases" "exadb" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
	}

	data "oci_database_db_homes" "t" {
		compartment_id = "${var.compartment_id}"
		db_system_id = "${oci_database_db_system.test_exadata_db_system.id}"
		filter {
			name = "display_name"
			values = ["TFTestDbHome1"]
		}
	}

	resource "oci_database_db_system" "test_exadata_db_system" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.exadata_subnet.id}"
		backup_subnet_id = "${oci_core_subnet.exadata_backup_subnet.id}"
		database_edition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
		disk_redundancy = "NORMAL"
		shape = "Exadata.Quarter2.92"
		cpu_core_count = "22"
		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
		domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
		hostname = "myOracleDB"
		data_storage_size_in_gb = "256"
		license_model = "LICENSE_INCLUDED"
		node_count = "1"
		time_zone = "US/Pacific"
		display_name = "TFTestExadataDbSystemVM"
		db_home {
			db_version = "12.1.0.2"
			display_name = "TFTestDbHome1"
			database {
				admin_password = "BEstrO0ng_#11"
				db_name = "aTFdb"
			}
		}
	}
	` +
		`resource "oci_database_db_system" "test_exadata_db_system2" {
					availability_domain = "${data.oci_identity_availability_domain.ad.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.exadata_subnet.id}"
					backup_subnet_id = "${oci_core_subnet.exadata_backup_subnet.id}"
					database_edition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
					disk_redundancy = "NORMAL"
					shape = "Exadata.Quarter2.92"
					cpu_core_count = "22"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					time_zone = "US/Pacific"
					display_name = "TFTestExadataDbSystemVM2"
					db_home {
						db_version = "12.1.0.2"
						display_name = "TFTestDbHome1"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb2"
						}
					}
				}

`

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_exadata_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_exadata_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_exadata_data_guard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals Existing DbSystem
			{
				Config: config + compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Create, dataGuardAssociationRepresentationExistingExadataDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingDbSystem"),
					resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_exadata_data_guard_associations", Optional, Update, dataGuardAssociationExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Required, Create, dataGuardAssociationSingularExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataDbSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
				),
			},
		},
	})
}

func TestResourceDatabaseDataGuardAssociation_ExadataExistingVMCluster(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDataGuardAssociation_ExadataExistingVMCluster")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig() + ExadataBaseDependenciesForDataGuardWithExistingVMCluster + `
	data "oci_database_databases" "exadb" {
		compartment_id = "${var.compartment_id}"
		db_home_id = "${oci_database_db_home.test_db_home_vm_cluster.id}"
	}

	resource "oci_database_db_home" "test_db_home_vm_cluster" {
	  vm_cluster_id = "${oci_database_vm_cluster.test_exadata_vm_cluster.id}"
	
	  database {
		admin_password = "BEstrO0ng_#11"
		db_name        = "dbVMClus"
		character_set  = "AL32UTF8"
		ncharacter_set = "AL16UTF16"
		db_workload    = "OLTP"
		pdb_name       = "pdbName"
	
		freeform_tags = {
		  "Department" = "Finance"
		}
	  }
	
	  source       = "VM_CLUSTER_NEW"
	  db_version   = "12.1.0.2"
	  display_name = "TFTestDbHome1"
	}
`

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_exadata_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_exadata_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_exadata_data_guard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals Existing VM Cluster
			{
				Config: config + compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Create, dataGuardAssociationRepresentationExistingExadataVmCluster),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingVmCluster"),
					resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_vm_cluster_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_exadata_data_guard_associations", Optional, Update, dataGuardAssociationExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataVmCluster),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Required, Create, dataGuardAssociationSingularExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataVmCluster),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
				),
			},
		},
	})
}

func TestResourceDatabaseDataGuardAssociation_ExadataExistingDBHome(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDataGuardAssociation_ExadataExistingVMCluster")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig() + `

	data "oci_identity_availability_domain" "ad" {
		compartment_id 		= "${var.compartment_id}"
		ad_number      		= 1
	}

	resource "oci_core_vcn" "vcn" {
	  cidr_block     = "10.1.0.0/16"
	  compartment_id = var.compartment_id
	  display_name   = "TFExampleVCNDBSystem"
	  dns_label      = "tfexvcndbsys"
	}
	
	resource "oci_core_subnet" "subnet" {
	  availability_domain = data.oci_identity_availability_domain.ad.name
	  cidr_block          = "10.1.20.0/24"
	  display_name        = "TFExampleSubnetDBSystem"
	  dns_label           = "tfexsubdbsys"
	  security_list_ids   = [oci_core_security_list.ExampleSecurityList.id]
	  compartment_id      = var.compartment_id
	  vcn_id              = oci_core_vcn.vcn.id
	  route_table_id      = oci_core_route_table.route_table.id
	  dhcp_options_id     = oci_core_vcn.vcn.default_dhcp_options_id
	}
	
	resource "oci_core_subnet" "subnet_backup" {
	  availability_domain = data.oci_identity_availability_domain.ad.name
	  cidr_block          = "10.1.1.0/24"
	  display_name        = "TFExampleSubnetDBSystemBackup"
	  dns_label           = "tfexsubdbsysbp"
	  security_list_ids   = [oci_core_security_list.ExampleSecurityList.id]
	  compartment_id      = var.compartment_id
	  vcn_id              = oci_core_vcn.vcn.id
	  route_table_id      = oci_core_route_table.route_table_backup.id
	  dhcp_options_id     = oci_core_vcn.vcn.default_dhcp_options_id
	}
	
	resource "oci_core_internet_gateway" "internet_gateway" {
	  compartment_id = var.compartment_id
	  display_name   = "TFExampleIGDBSystem"
	  vcn_id         = oci_core_vcn.vcn.id
	}
	
	resource "oci_core_route_table" "route_table" {
	  compartment_id = var.compartment_id
	  vcn_id         = oci_core_vcn.vcn.id
	  display_name   = "TFExampleRouteTableDBSystem"
	
	  route_rules {
		destination       = "0.0.0.0/0"
		destination_type  = "CIDR_BLOCK"
		network_entity_id = oci_core_internet_gateway.internet_gateway.id
	  }
	}
	
	resource "oci_core_route_table" "route_table_backup" {
	  compartment_id = var.compartment_id
	  vcn_id         = oci_core_vcn.vcn.id
	  display_name   = "TFExampleRouteTableDBSystemBackup"
	
	  route_rules {
		destination       = "0.0.0.0/0"
		destination_type  = "CIDR_BLOCK"
		network_entity_id = oci_core_internet_gateway.internet_gateway.id
	  }
	}
	
	resource "oci_core_security_list" "ExampleSecurityList" {
	  compartment_id = var.compartment_id
	  vcn_id         = oci_core_vcn.vcn.id
	  display_name   = "TFExampleSecurityList"
	
	  // allow outbound tcp traffic on all ports
	  egress_security_rules {
		destination = "0.0.0.0/0"
		protocol    = "6"
	  }
	
	  // allow outbound udp traffic on a port range
	  egress_security_rules {
		destination = "0.0.0.0/0"
		protocol    = "17" // udp
		stateless   = true
	  }
	
	  egress_security_rules {
		destination = "0.0.0.0/0"
		protocol    = "1"
		stateless   = true
	  }
	
	  // allow inbound ssh traffic from a specific port
	  ingress_security_rules {
		protocol  = "6" // tcp
		source    = "0.0.0.0/0"
		stateless = false
	  }
	
	  // allow inbound icmp traffic of a specific type
	  ingress_security_rules {
		protocol  = 1
		source    = "0.0.0.0/0"
		stateless = true
	  }
	}
	
	resource "oci_core_network_security_group" "test_network_security_group" {
	  compartment_id = var.compartment_id
	  vcn_id         = oci_core_vcn.vcn.id
	  display_name   = "displayName"
	}
	
	resource "oci_core_network_security_group" "test_network_security_group_backup" {
	  compartment_id = var.compartment_id
	  vcn_id         = oci_core_vcn.vcn.id
	  display_name   = "displayName"
	}

	data "oci_database_databases" "exadb" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${oci_database_db_system.test_exadata_db_system.db_home.0.id}"
	}

	variable "db_system_shape" {
	  default = "Exadata.Quarter1.84"
	}
	
	variable "cpu_core_count" {
	  default = "22"
	}
	
	variable "db_edition" {
	  default = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	}
	
	variable "db_admin_password" {
	  default = "BEstrO0ng_#12"
	}
	
	variable "db_version" {
	  default = "19.0.0.0"
	}
	
	variable "db_disk_redundancy" {
	  default = "HIGH"
	}
	
	variable "sparse_diskgroup" {
	  default = true
	}
	
	variable "hostname" {
	  default = "myoracledb"
	}
	
	variable "host_user_name" {
	  default = "opc"
	}
	
	variable "n_character_set" {
	  default = "AL16UTF16"
	}
	
	variable "character_set" {
	  default = "AL32UTF8"
	}
	
	variable "db_workload" {
	  default = "OLTP"
	}
	
	variable "pdb_name" {
	  default = "pdbName"
	}
	
	variable "data_storage_size_in_gb" {
	  default = "256"
	}
	
	variable "license_model" {
	  default = "LICENSE_INCLUDED"
	}
	
	variable "node_count" {
	  default = "2"
	}
	
	variable "data_storage_percentage" {
	  default = "40"
	}
	
	variable "time_zone" {
	  default = "US/Pacific"
	}
	
	data "oci_database_db_system_shapes" "test_db_system_shapes" {
	  availability_domain = data.oci_identity_availability_domain.ad.name
	  compartment_id      = var.compartment_id
	
	  filter {
		name   = "shape"
		values = [var.db_system_shape]
	  }
	}

	resource "oci_database_db_system" "test_exadata_db_system" {
	  availability_domain = data.oci_identity_availability_domain.ad.name
	  compartment_id      = var.compartment_id
	  cpu_core_count      = var.cpu_core_count
	  database_edition    = var.db_edition
	  time_zone           = var.time_zone
	
	  db_home {
		database {
		  admin_password = var.db_admin_password
		  db_name        = "TFdbExa1"
		  character_set  = var.character_set
		  ncharacter_set = var.n_character_set
		  db_workload    = var.db_workload
		  pdb_name       = var.pdb_name
	
		  db_backup_config {
			auto_backup_enabled = false
		  }
		}
	
		db_version   = var.db_version
		display_name = "MyTFDBHomeExa1"
	  }
	
	  maintenance_window_details {
		preference = "CUSTOM_PREFERENCE"
	
		days_of_week {
		  name = "MONDAY"
		}
	
		hours_of_day       = ["4"]
		lead_time_in_weeks = 2
	
		months {
		  name = "APRIL"
		}
	
		weeks_of_month = ["2"]
	  }
	
	  shape            = var.db_system_shape
	  subnet_id        = oci_core_subnet.subnet.id
	  backup_subnet_id = oci_core_subnet.subnet_backup.id
	  ssh_public_keys  = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNpUvNPYAJUoH/C7sS90htOdIqSPG/YHJzQdUniBJTOe5TtI+wNQ7xFLc6rPp430kbt/KtQ3YyaTkWIiOHcuInBirGUGF1JPdjovmppBA8FJIz69YghLit1uLV6HxVBuEqbSEsh8zA7ZjyR1qDtA5pvjzvSBjxFKBrRhq+HD+CHMTUufbyZzk1oWItdJF5GqJtMZoDw5EwRrvll8PqHkNUCONrSTgZC85oxsgaDdseXNPRT5fzf8i5BWmkvLcq9gx0Hvk/pt7USnI0qW4jo877qljxA8TqLFipvBa9s+xRJKGzgdaSdfKaEPwDRkjKP5WVH+RYZPrEjn9vW+IsCRWcsmBsrkfrCCZ1QEWxzI3MxRrICdQ1v/++o3oD2Ksp4pMZHEI/RSGo0rZW8znerD8+WoEtHvyQAmJnmFBKoAiqLHWCgeHjXB1+UMSebLhy1LG1PFcw4bTf1vD66dkSvUOIj1lLz67N4rlmFz7bkTOj2WvYAGlqMrpBTVCj4qvKqGj9eSi8Mk2MydTEMgxIrVUAYp2+e2fgBm7Nopu23lPYwa/2gKpkNfaOjxAro0R5E6nweFCVqxA71UvNWCWI4NEBz7PQFqpY65COGVt/okNLZy0U154foYJNGYhXBpIeXvpeJU8sdmiSe4BbK0VR+LwZHHlAhOk/64n160fzTH8Cbw== govindrao.kulkarni@oracle.com"]
	  display_name     = "MyTFDBSystem1"
	  sparse_diskgroup = var.sparse_diskgroup
	
	  hostname                = var.hostname
	  data_storage_percentage = var.data_storage_percentage
	
	  #data_storage_size_in_gb = var.data_storage_size_in_gb
	  node_count             = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
	}
	
	resource "oci_database_db_system" "test_exadata_db_system2" {
	  availability_domain = data.oci_identity_availability_domain.ad.name
	  compartment_id      = var.compartment_id
	  cpu_core_count      = var.cpu_core_count
	  database_edition    = var.db_edition
	  time_zone           = var.time_zone
	
	  db_home {
		database {
		  admin_password = var.db_admin_password
		  db_name        = "TFdbExa2"
		  character_set  = var.character_set
		  ncharacter_set = var.n_character_set
		  db_workload    = var.db_workload
		  pdb_name       = var.pdb_name
	
		  db_backup_config {
			auto_backup_enabled = false
		  }
		}
	
		db_version   = "19.0.0.0"
		display_name = "MyTFDBHomeExa2"
	  }
	
	  maintenance_window_details {
		preference = "CUSTOM_PREFERENCE"
	
		days_of_week {
		  name = "MONDAY"
		}
	
		hours_of_day       = ["4"]
		lead_time_in_weeks = 2
	
		months {
		  name = "APRIL"
		}
	
		weeks_of_month = ["2"]
	  }
	
	  shape            = var.db_system_shape
	  subnet_id        = oci_core_subnet.subnet.id
	  backup_subnet_id = oci_core_subnet.subnet_backup.id
	  ssh_public_keys  = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNpUvNPYAJUoH/C7sS90htOdIqSPG/YHJzQdUniBJTOe5TtI+wNQ7xFLc6rPp430kbt/KtQ3YyaTkWIiOHcuInBirGUGF1JPdjovmppBA8FJIz69YghLit1uLV6HxVBuEqbSEsh8zA7ZjyR1qDtA5pvjzvSBjxFKBrRhq+HD+CHMTUufbyZzk1oWItdJF5GqJtMZoDw5EwRrvll8PqHkNUCONrSTgZC85oxsgaDdseXNPRT5fzf8i5BWmkvLcq9gx0Hvk/pt7USnI0qW4jo877qljxA8TqLFipvBa9s+xRJKGzgdaSdfKaEPwDRkjKP5WVH+RYZPrEjn9vW+IsCRWcsmBsrkfrCCZ1QEWxzI3MxRrICdQ1v/++o3oD2Ksp4pMZHEI/RSGo0rZW8znerD8+WoEtHvyQAmJnmFBKoAiqLHWCgeHjXB1+UMSebLhy1LG1PFcw4bTf1vD66dkSvUOIj1lLz67N4rlmFz7bkTOj2WvYAGlqMrpBTVCj4qvKqGj9eSi8Mk2MydTEMgxIrVUAYp2+e2fgBm7Nopu23lPYwa/2gKpkNfaOjxAro0R5E6nweFCVqxA71UvNWCWI4NEBz7PQFqpY65COGVt/okNLZy0U154foYJNGYhXBpIeXvpeJU8sdmiSe4BbK0VR+LwZHHlAhOk/64n160fzTH8Cbw== govindrao.kulkarni@oracle.com"]
	  display_name     = "MyTFDBSystem2"
	  sparse_diskgroup = var.sparse_diskgroup
	
	  hostname                = var.hostname
	  data_storage_percentage = var.data_storage_percentage
	
	  #data_storage_size_in_gb = var.data_storage_size_in_gb
	  node_count             = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
	}
`

	compartmentId := getEnvSettingWithBlankDefault("compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_data_guard_association.test_exadata_data_guard_association"
	datasourceName := "data.oci_database_data_guard_associations.test_exadata_data_guard_associations"
	singularDatasourceName := "data.oci_database_data_guard_association.test_exadata_data_guard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals Existing DbSystem
			{
				Config: config + compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Create, dataGuardAssociationRepresentationExistingExadataDbHome),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "creation_type", "ExistingDbSystem"),
					resource.TestCheckResourceAttr(resourceName, "database_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_db_home_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "role"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "transport_type", "ASYNC"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_associations", "test_exadata_data_guard_associations", Optional, Update, dataGuardAssociationExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataDbHome),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.peer_role"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.role"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_guard_associations.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "data_guard_associations.0.transport_type", "ASYNC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Required, Create, dataGuardAssociationSingularExadataDataSourceRepresentation) +
					compartmentIdVariableStr +
					generateResourceFromRepresentationMap("oci_database_data_guard_association", "test_exadata_data_guard_association", Optional, Update, dataGuardAssociationRepresentationExistingExadataDbHome),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_guard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_system_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_db_home_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "transport_type", "ASYNC"),
				),
			},
		},
	})
}
