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
