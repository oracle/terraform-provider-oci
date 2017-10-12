// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceDatabaseDBSystemTestSuite struct {
	suite.Suite
	Client    *baremetal.Client
	Config    string
	Provider  terraform.ResourceProvider
	Providers map[string]terraform.ResourceProvider
}

func (s *DatasourceDatabaseDBSystemTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	}

	resource "oci_database_db_system" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.t.id}"
		database_edition = "ENTERPRISE_EDITION"
		disk_redundancy = "NORMAL"
		shape = "BM.DenseIO1.36"
		cpu_core_count = "2"
		ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
		display_name = "-tf-db-system"
		domain = "mycompany.com"
		hostname = "myOracleDB"
		db_home {
			db_version = "12.1.0.2"
			display_name = "-tf-db-home"
			database {
				"admin_password" = "BEstrO0ng_#11"
				"db_name" = "aTFdb"
				character_set = "AL32UTF8"
				ncharacter_set = "AL16UTF16"
			}
		}
	}`
}

func (s *DatasourceDatabaseDBSystemTestSuite) TestAccDatasourceDatabaseDBSystem_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_database_db_systems" "t" {
					compartment_id = "${var.compartment_id}"
				}
				data "oci_database_db_homes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
				}
				data "oci_database_databases" "t" {
					compartment_id = "${var.compartment_id}"
					db_home_id = "${data.oci_database_db_homes.t.id}"
				}
				data "oci_database_database" "t" {
					  database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
				data "oci_database_db_nodes" "t" {
					compartment_id = "${var.compartment_id}"
					db_system_id = "${oci_database_db_system.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
					resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceDatabaseDBSystemTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceDatabaseDBSystemTestSuite))
}
