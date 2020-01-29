// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/database"
)

var (
	DbSystemResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
		"route_table_id": Representation{repType: Optional, create: `${oci_core_route_table.test_route_table.id}`}})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Optional, Create, internetGatewayRepresentation)

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

	ResourceDatabaseBaseConfig = legacyTestProviderConfig() + DefinedTagsDependencies + `

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
	ResourceDatabaseToken, ResourceDatabaseTokenFn = tokenizeWithHttpReplay("database_db")
)

// TestAccResourceDatabaseDBSystem_basic tests creation of a DBSystem with the minimum required properties
// to assert expected default values are set
func TestResourceDatabaseDBSystemBasic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemBasic")
	defer httpreplay.SaveScenario()

	// This test is a subset of TestAccResourceDatabaseDBSystem_allXX. It tests omitting optional params.
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_basic") {
		t.Skip("Skipping suppressed DBSystem_basic")
	}

	var resId, resId2 string
	provider := testAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
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
        			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_database_db_system.t", "id")
						return err
					},
				),
			},
			// verify update without updating nsgIds
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
							db_name = "aTFdb"
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_database_db_system.t", "id")
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
							db_name = "aTFdb"
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "0"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_database_db_system.t", "id")
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
			// verify create without nsgIds and backupNsgIds
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
							db_name = "aTFdb"
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				),
			},
			// verify update without nsgIds and backupNsgIds
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
							db_name = "aTFdb"
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
				),
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("DatabaseDbSystem", &resource.Sweeper{
		Name:         "DatabaseDbSystem",
		Dependencies: DependencyGraph["dbSystem"],
		F:            sweepDatabaseDbSystemResource,
	})
}
