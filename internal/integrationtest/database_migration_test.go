// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	migrationRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.t.id}`},
	}

	// Copy from TestAccResourceDatabaseDBSystem_Exadata
	MigrationResourceDependencies = ResourceDatabaseBaseConfig + `
				data "oci_identity_availability_domain" "ad" {
  					compartment_id 		= "${var.compartment_id}"
  					ad_number      		= 3
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
					dns_label           = "subnetexadata1"
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

				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domain.ad.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.exadata_subnet.id}"
					backup_subnet_id = "${oci_core_subnet.exadata_backup_subnet.id}"
					database_edition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
					disk_redundancy = "HIGH"
					shape = "Exadata.Quarter1.84"
					cpu_core_count = "22"
					ssh_public_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"]
					domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					time_zone = "US/Pacific"
					backup_network_nsg_ids = ["${oci_core_network_security_group.test_network_security_group2.id}"]
					maintenance_window_details {
						preference = "NO_PREFERENCE"
					}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
					db_home {
						db_version = "12.1.0.2"
						defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
						freeform_tags = {
							"Department" = "Finance"
						}
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
						}
					}
				}`
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_migration.test_migration"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: MigrationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_migration", "test_migration", acctest.Required, acctest.Create, migrationRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
					resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(resourceName, "cloud_vm_cluster_id"),
				),
			},
			// remove db system config and apply
			{
				Config: config + compartmentIdVariableStr,
				Check:  acctest.ComposeAggregateTestCheckFuncWrapper(),
			},
		},
	})
}
