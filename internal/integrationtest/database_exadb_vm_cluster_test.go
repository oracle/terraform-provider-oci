// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExadbVmClusterRequiredOnlyResource = DatabaseExadbVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create, DatabaseExadbVmClusterRepresentation)

	DatabaseExadbVmClusterResourceConfig = DatabaseExadbVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Update, DatabaseExadbVmClusterRepresentation)

	DatabaseExadbVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id}`},
	}

	DatabaseExadbVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `TFExadbVmCluster`, Update: `TFExadbVmClusterUpdatedName`},
		"exascale_db_storage_vault_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExadbVmClusterDataSourceFilterRepresentation}}

	DatabaseExadbVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id}`}},
	}

	DatabaseExadbVmClusterRepresentation = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `TFExadbVmCluster`, Update: `TFExadbVmClusterUpdatedName`},
		"exascale_db_storage_vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"grid_image_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.grid_image_id}`},
		"hostname":                     acctest.Representation{RepType: acctest.Required, Create: `apollo`},
		"shape":                        acctest.Representation{RepType: acctest.Required, Create: `EXADBXS`},
		"ssh_public_keys":              acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadbxs_client_subnet.id}`},
		"backup_subnet_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadbxs_backup_subnet.id}`},
		"node_config":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeConfigRepresentation},
		"node_resource": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation1},
			{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation2},
		},
		"cluster_name":               acctest.Representation{RepType: acctest.Required, Create: `tfexadbxs`},
		"domain":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.exadbxs_client_subnet.subnet_domain_name}`},
		"nsg_ids":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.exadbxs_client_nsg.id}`}},
		"backup_network_nsg_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.exadbxs_backup_nsg.id}`}},
		"data_collection_options":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseExadbVmClusterDataCollectionOptionsRepresentation},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"scan_listener_port_tcp":     acctest.Representation{RepType: acctest.Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl": acctest.Representation{RepType: acctest.Optional, Create: `2484`},
		"time_zone":                  acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"example-tag-namespace-all.example-tag": "value"}, Update: map[string]string{"example-tag-namespace-all.example-tag": "updatedValue"}},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExadbVmClusterIgnoreDefinedTagsRepresentation},
	}

	DatabaseExadbVmClusterNodeConfigRepresentation = map[string]interface{}{
		"enabled_ecpu_count_per_node":              acctest.Representation{RepType: acctest.Required, Create: `8`, Update: `16`},
		"total_ecpu_count_per_node":                acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `32`},
		"vm_file_system_storage_size_gbs_per_node": acctest.Representation{RepType: acctest.Required, Create: `350`, Update: `400`},
	}

	DatabaseExadbVmClusterNodeResourceRepresentation1 = map[string]interface{}{
		"node_name": acctest.Representation{RepType: acctest.Required, Create: `node1`},
	}

	DatabaseExadbVmClusterNodeResourceRepresentation2 = map[string]interface{}{
		"node_name": acctest.Representation{RepType: acctest.Required, Create: `node2`},
	}

	DatabaseExadbVmClusterNodeResourceRepresentation3 = map[string]interface{}{
		"node_name": acctest.Representation{RepType: acctest.Required, Create: `node3`},
	}

	DatabaseExadbVmClusterDataCollectionOptionsRepresentation = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DatabaseExadbVmClusterIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	ExadbVmClusterNetwork = `
        resource "oci_core_virtual_network" "exadbxs_vcn" {
          compartment_id = "${var.compartment_id}"
          cidr_blocks    = ["10.1.0.0/16"]
          display_name   = "exadbxs-tf-vcn"
          dns_label      = "tfvcn"
        }
        
        resource "oci_core_internet_gateway" "exadbxs_igw" {
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          display_name   = "exadbxs-tf-igw"
        }
        
        resource "oci_core_route_table" "exadbxs_rt" {
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          display_name   = "exadbxs-tf-route-table"
          route_rules {
            destination       = "0.0.0.0/0"
            destination_type  = "CIDR_BLOCK"
            network_entity_id = "${oci_core_internet_gateway.exadbxs_igw.id}"
          }
        }
        
        resource "oci_core_subnet" "exadbxs_client_subnet" {
          cidr_block        = "10.1.20.0/24"
          compartment_id    = "${var.compartment_id}"
          vcn_id            = "${oci_core_virtual_network.exadbxs_vcn.id}"
          route_table_id    = "${oci_core_route_table.exadbxs_rt.id}"
          security_list_ids = ["${oci_core_virtual_network.exadbxs_vcn.default_security_list_id}", "${oci_core_security_list.exadbxs_security_list.id}"]
          dns_label         = "tfclientsub"
          display_name      = "exadbxs-tf-client-subnet"
        }
        
        resource "oci_core_subnet" "exadbxs_backup_subnet" {
          cidr_block     = "10.1.21.0/24"
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          route_table_id = "${oci_core_route_table.exadbxs_rt.id}"
          dns_label      = "tfbackupsub"
          display_name   = "exadbxs-tf-backup-subnet"
        }
        
        resource "oci_core_network_security_group" "exadbxs_client_nsg" {
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          display_name   = "exadbxs-client-nsg"
        }
        
        resource "oci_core_network_security_group" "exadbxs_backup_nsg" {
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          display_name   = "exadbxs-backup-nsg"
        }
        
        resource "oci_core_security_list" "exadbxs_security_list" {
          compartment_id = "${var.compartment_id}"
          vcn_id         = "${oci_core_virtual_network.exadbxs_vcn.id}"
          display_name   = "exadbxs-security-list"
        
          ingress_security_rules {
            source   = "10.1.22.0/24"
            protocol = "6"
          }
        
          ingress_security_rules {
            source   = "10.1.22.0/24"
            protocol = "1"
          }

          ingress_security_rules {
            source   = "0.0.0.0/0"
            protocol = "all"
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

	// Note: set env variable TF_VAR_grid_image_id before running this test
	GridImageIdDependency                      = `variable "grid_image_id" {}`
	DatabaseExadbVmClusterResourceDependencies = AvailabilityDomainConfig + ExadbVmClusterNetwork + GridImageIdDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation)
)

// issue-routing-tag: database/ExaCS
func TestDatabaseExadbVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadbVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_exadb_vm_cluster.test_exadb_vm_cluster"
	datasourceName := "data.oci_database_exadb_vm_clusters.test_exadb_vm_clusters"
	singularDatasourceName := "data.oci_database_exadb_vm_cluster.test_exadb_vm_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExadbVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Create, DatabaseExadbVmClusterRepresentation), "database", "exadbVmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExadbVmClusterDestroy, []resource.TestStep{
		// 0 verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create, DatabaseExadbVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// 1 delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies,
		},
		// 2 verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Create, DatabaseExadbVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_name"),
				// optional fields
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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

		// 3 verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseExadbVmClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_name"),
				// optional fields
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 4 verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Update, DatabaseExadbVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmClusterUpdatedName"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "32"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "400"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_name"),
				// optional fields
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// 5 verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_clusters", "test_exadb_vm_clusters", acctest.Optional, acctest.Update, DatabaseExadbVmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Update, DatabaseExadbVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TFExadbVmClusterUpdatedName"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.display_name", "TFExadbVmClusterUpdatedName"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.gi_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.grid_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.grid_image_type"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.hostname", "apollo"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.node_resource.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.node_config.0.enabled_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.node_config.0.total_ecpu_count_per_node", "32"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.node_config.0.vm_file_system_storage_size_gbs_per_node", "400"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.cluster_name"),
				// optional fields
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.domain"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.system_tags.%", "0"),
				// computed fields
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.listener_port"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.scan_dns_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.scan_dns_record_id"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_clusters.0.scan_ip_ids.#", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_clusters.0.state"),
			),
		},
		// 6 verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create, DatabaseExadbVmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadbVmClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadb_vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TFExadbVmClusterUpdatedName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_image_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config.0.enabled_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config.0.total_ecpu_count_per_node", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "400"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_name"),
				// optional fields
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_network_nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
				// computed fields
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scan_dns_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scan_dns_record_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_ip_ids.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// 7 verify resource import
		{
			Config:                  config + DatabaseExadbVmClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"node_resource.0.node_name", "node_resource.1.node_name"},
			ResourceName:            resourceName,
		},
	})
}

func TestDatabaseExadbVmClusterResource_add_remove_node(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadbVmClusterResource_add_remove_node")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadb_vm_cluster.test_exadb_vm_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExadbVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Optional, acctest.Create, DatabaseExadbVmClusterRepresentation), "database", "exadbVmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExadbVmClusterDestroy, []resource.TestStep{
		// 0 verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create, DatabaseExadbVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// 2 verify adding node3
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseExadbVmClusterRepresentation, map[string]interface{}{
						"node_resource": []acctest.RepresentationGroup{
							{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation1},
							{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation2},
							{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation3},
						},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 3 verify removing a node1
		{
			Config: config + compartmentIdVariableStr + DatabaseExadbVmClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadb_vm_cluster", "test_exadb_vm_cluster", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseExadbVmClusterRepresentation, map[string]interface{}{
						"node_resource": []acctest.RepresentationGroup{
							{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation2},
							{RepType: acctest.Required, Group: DatabaseExadbVmClusterNodeResourceRepresentation3},
						},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExadbVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gi_version"),
				resource.TestCheckResourceAttrSet(resourceName, "grid_image_id"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "apollo"),
				resource.TestCheckResourceAttr(resourceName, "shape", "EXADBXS"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "node_resource.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.enabled_ecpu_count_per_node", "8"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.total_ecpu_count_per_node", "16"),
				resource.TestCheckResourceAttr(resourceName, "node_config.0.vm_file_system_storage_size_gbs_per_node", "350"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.memory_size_in_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.snapshot_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "node_config.0.total_file_system_storage_size_gbs_per_node"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDatabaseExadbVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_exadb_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetExadbVmClusterRequest{}

			tmp := rs.Primary.ID
			request.ExadbVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExadbVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExadbVmClusterLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseExadbVmCluster") {
		resource.AddTestSweepers("DatabaseExadbVmCluster", &resource.Sweeper{
			Name:         "DatabaseExadbVmCluster",
			Dependencies: acctest.DependencyGraph["exadbVmCluster"],
			F:            sweepDatabaseExadbVmClusterResource,
		})
	}
}

func sweepDatabaseExadbVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	exadbVmClusterIds, err := getDatabaseExadbVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, exadbVmClusterId := range exadbVmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[exadbVmClusterId]; !ok {
			deleteExadbVmClusterRequest := oci_database.DeleteExadbVmClusterRequest{}

			deleteExadbVmClusterRequest.ExadbVmClusterId = &exadbVmClusterId

			deleteExadbVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExadbVmCluster(context.Background(), deleteExadbVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadbVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadbVmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &exadbVmClusterId, DatabaseExadbVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseExadbVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseExadbVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExadbVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExadbVmClustersRequest := oci_database.ListExadbVmClustersRequest{}
	listExadbVmClustersRequest.CompartmentId = &compartmentId
	listExadbVmClustersRequest.LifecycleState = oci_database.ExadbVmClusterSummaryLifecycleStateAvailable
	listExadbVmClustersResponse, err := databaseClient.ListExadbVmClusters(context.Background(), listExadbVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExadbVmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, exadbVmCluster := range listExadbVmClustersResponse.Items {
		id := *exadbVmCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExadbVmClusterId", id)
	}
	return resourceIds, nil
}

func DatabaseExadbVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exadbVmClusterResponse, ok := response.Response.(oci_database.GetExadbVmClusterResponse); ok {
		return exadbVmClusterResponse.LifecycleState != oci_database.ExadbVmClusterLifecycleStateTerminated
	}
	return false
}

func DatabaseExadbVmClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExadbVmCluster(context.Background(), oci_database.GetExadbVmClusterRequest{
		ExadbVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
