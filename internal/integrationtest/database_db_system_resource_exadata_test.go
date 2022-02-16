// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/oracle/oci-go-sdk/v58/database"
)

// TestAccResourceDatabaseDBSystem_Exadata tests DBsystems using Exadata
// issue-routing-tag: database/default
func TestResourceDatabaseDBSystemExaData(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemExaData")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_Exadata") {
		t.Skip("Skipping suppressed DBSystem_Exadata")
	}

	provider := acctest.TestAccProvider
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: ResourceDatabaseBaseConfig + `
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
					ssh_public_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV"]
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
				}`,

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "HIGH"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "Exadata.Quarter1.84"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "22"),
					resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "subnetexadata1.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.Department", "Finance"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "backup_network_nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window_details.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				),
			},
			// verify Update
			{
				Config: ResourceDatabaseBaseConfig + `
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
					ssh_public_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV"]
					domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					time_zone = "US/Pacific"
					backup_network_nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
					maintenance_window_details {
						preference = "CUSTOM_PREFERENCE"
    					days_of_week {
      						name = "TUESDAY"
    					}
    					hours_of_day = ["4"]
						lead_time_in_weeks = 11
						months {
      						name = "FEBRUARY"
						}
						months {
      						name = "MAY"
						}
						months {
      						name = "AUGUST"
						}
						months {
      						name = "NOVEMBER"
						}
    					weeks_of_month = ["2"]
					}
					nsg_ids = ["${oci_core_network_security_group.test_network_security_group2.id}"]
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
				}`,

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "HIGH"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "Exadata.Quarter1.84"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "22"),
					resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "subnetexadata1.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.Department", "Finance"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "backup_network_nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.days_of_week.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.hours_of_day.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.lead_time_in_weeks", "11"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.months.#", "4"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				),
			},
			// verify removing nsgIds and backupNsgIds trigger Update
			{
				Config: ResourceDatabaseBaseConfig + `
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
					ssh_public_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV"]
					domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					time_zone = "US/Pacific"
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
				}`,

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "HIGH"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "Exadata.Quarter1.84"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "22"),
					resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "subnetexadata1.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.freeform_tags.Department", "Finance"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "backup_network_nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "nsg_ids.#", "0"),
				),
			},
		},
	})
}
