// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_database "github.com/oracle/oci-go-sdk/v54/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CloudVmClusterRequiredOnlyResource = CloudVmClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Required, Create, cloudVmClusterRepresentation)

	CloudVmClusterResourceConfig = CloudVmClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Update, cloudVmClusterRepresentation)

	cloudVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}

	cloudVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cloud_exadata_infrastructure_id": Representation{RepType: Optional, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"display_name":                    Representation{RepType: Optional, Create: `cloudVmCluster`, Update: `displayName2`},
		"state":                           Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                          RepresentationGroup{Required, cloudVmClusterDataSourceFilterRepresentation}}
	cloudVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`}},
	}

	cloudVmClusterRepresentation = map[string]interface{}{
		"backup_subnet_id":                Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet_backup.id}`},
		"cloud_exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"compartment_id":                  Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                  Representation{RepType: Required, Create: `4`},
		"display_name":                    Representation{RepType: Required, Create: `cloudVmCluster`, Update: `displayName2`},
		"gi_version":                      Representation{RepType: Required, Create: `19.0.0.0`},
		"hostname":                        Representation{RepType: Required, Create: `apollo`},
		"ssh_public_keys":                 Representation{RepType: Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                       Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet1.id}`},
		"domain":                          Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet1.subnet_domain_name}`},
		"backup_network_nsg_ids":          Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group_backup.id}`}},
		"cluster_name":                    Representation{RepType: Optional, Create: `clusterName`},
		"data_storage_percentage":         Representation{RepType: Optional, Create: `40`},
		"defined_tags":                    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":         Representation{RepType: Optional, Create: `true`},
		"is_sparse_diskgroup_enabled":     Representation{RepType: Optional, Create: `false`},
		"license_model":                   Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                         Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"ocpu_count":                      Representation{RepType: Optional, Create: `4.0`, Update: `4.0`},
		"scan_listener_port_tcp":          Representation{RepType: Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl":      Representation{RepType: Optional, Create: `2484`},
		"time_zone":                       Representation{RepType: Optional, Create: `US/Pacific`},
	}

	CloudVmClusterResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", Required, Create, cloudExadataInfrastructureRepresentation) +
		`
				data "oci_identity_availability_domains" "ADs" {
					compartment_id = "${var.compartment_id}"
				}

				data "oci_identity_availability_domain" "ad" {
  					compartment_id 		= "${var.compartment_id}"
  					ad_number      		= 1
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

				resource "oci_core_network_security_group" "test_network_security_group_backup" {
					compartment_id = "${var.compartment_id}"
					vcn_id            = "${oci_core_virtual_network.t.id}"
				}

				resource "oci_core_subnet" "test_subnet1" {
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

				resource "oci_core_subnet" "test_subnet_backup" {
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

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_cloud_vm_cluster.test_cloud_vm_cluster"
	datasourceName := "data.oci_database_cloud_vm_clusters.test_cloud_vm_clusters"
	singularDatasourceName := "data.oci_database_cloud_vm_cluster.test_cloud_vm_cluster"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+CloudVmClusterResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Create, cloudVmClusterRepresentation), "database", "cloudVmCluster", t)

	ResourceTest(t, testAccCheckDatabaseCloudVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Required, Create, cloudVmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Create, cloudVmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Create,
					RepresentationCopyWithNewProperties(cloudVmClusterRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Update, cloudVmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_cloud_vm_clusters", "test_cloud_vm_clusters", Optional, Update, cloudVmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + CloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Optional, Update, cloudVmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.backup_subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cpu_core_count", "4"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_storage_percentage", "40"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.disk_redundancy"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.domain"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.listener_port"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.node_count"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.ocpu_count", "4"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.scan_dns_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.shape"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.zone_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", Required, Create, cloudVmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudVmClusterResourceConfig + DefinedTagsDependencies + AvailabilityDomainConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "disk_redundancy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scan_dns_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_id"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceConfig + DefinedTagsDependencies + AvailabilityDomainConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"create_async",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseCloudVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_cloud_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetCloudVmClusterRequest{}

			tmp := rs.Primary.ID
			request.CloudVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetCloudVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.CloudVmClusterLifecycleStateTerminated): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DatabaseCloudVmCluster") {
		resource.AddTestSweepers("DatabaseCloudVmCluster", &resource.Sweeper{
			Name:         "DatabaseCloudVmCluster",
			Dependencies: DependencyGraph["cloudVmCluster"],
			F:            sweepDatabaseCloudVmClusterResource,
		})
	}
}

func sweepDatabaseCloudVmClusterResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	cloudVmClusterIds, err := getCloudVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudVmClusterId := range cloudVmClusterIds {
		if ok := SweeperDefaultResourceId[cloudVmClusterId]; !ok {
			deleteCloudVmClusterRequest := oci_database.DeleteCloudVmClusterRequest{}

			deleteCloudVmClusterRequest.CloudVmClusterId = &cloudVmClusterId

			deleteCloudVmClusterRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteCloudVmCluster(context.Background(), deleteCloudVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudVmClusterId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &cloudVmClusterId, cloudVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				cloudVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getCloudVmClusterIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "CloudVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listCloudVmClustersRequest := oci_database.ListCloudVmClustersRequest{}
	listCloudVmClustersRequest.CompartmentId = &compartmentId
	listCloudVmClustersRequest.LifecycleState = oci_database.CloudVmClusterSummaryLifecycleStateAvailable
	listCloudVmClustersResponse, err := databaseClient.ListCloudVmClusters(context.Background(), listCloudVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudVmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudVmCluster := range listCloudVmClustersResponse.Items {
		id := *cloudVmCluster.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudVmClusterId", id)
	}
	return resourceIds, nil
}

func cloudVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudVmClusterResponse, ok := response.Response.(oci_database.GetCloudVmClusterResponse); ok {
		return cloudVmClusterResponse.LifecycleState != oci_database.CloudVmClusterLifecycleStateTerminated
	}
	return false
}

func cloudVmClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetCloudVmCluster(context.Background(), oci_database.GetCloudVmClusterRequest{
		CloudVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
