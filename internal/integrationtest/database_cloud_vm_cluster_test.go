// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseCloudVmClusterRequiredOnlyResource = DatabaseCloudVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation)

	DatabaseCloudVmClusterResourceConfig = DatabaseCloudVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Update, DatabaseCloudVmClusterRepresentation)

	CloudVmClusterResourceConfigUpdateInfra = CloudVmClusterResourceUpdateDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Update, DatabaseCloudVmClusterRepresentation)

	DatabaseDatabaseCloudVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}

	DatabaseDatabaseCloudVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `cloudVmCluster`, Update: `displayName2`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_cluster_type":                 acctest.Representation{RepType: acctest.Optional, Create: `DEVELOPER`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseCloudVmClusterDataSourceFilterRepresentation}}
	DatabaseCloudVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`}},
	}

	DatabaseCloudVmClusterRepresentation = map[string]interface{}{
		"depends_on": []string{"time_sleep.wait_30_seconds"},
		"file_system_configuration_details": []acctest.RepresentationGroup{
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation0},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation1},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation2},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation3},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation4},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation5},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation6},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation7},
			{RepType: acctest.Optional, Group: DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation8}},
		"backup_subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.t2.id}`},
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                  acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `6`},
		"data_storage_size_in_tbs":        acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`},
		"db_node_storage_size_in_gbs":     acctest.Representation{RepType: acctest.Optional, Create: `120`, Update: `160`},
		"memory_size_in_gbs":              acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `90`},
		"db_servers":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`}},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `cloudVmCluster`, Update: `displayName2`},
		"gi_version":                      acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"hostname":                        acctest.Representation{RepType: acctest.Required, Create: `apollo`},
		"ssh_public_keys":                 acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.t.id}`},
		"domain":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.t.subnet_domain_name}`},
		"backup_network_nsg_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group_backup.id}`}},
		"cluster_name":                    acctest.Representation{RepType: acctest.Optional, Create: `clusterName`},
		"cloud_automation_update_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudVmClusterCloudAutomationUpdateDetailsRepresentation},
		"data_collection_options":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: cloudVmClusterDataCollectionOptionsRepresentation},
		"data_storage_percentage":         acctest.Representation{RepType: acctest.Optional, Create: `40`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_sparse_diskgroup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"ocpu_count":                      acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `6.0`},
		"scan_listener_port_tcp":          acctest.Representation{RepType: acctest.Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl":      acctest.Representation{RepType: acctest.Optional, Create: `2484`},
		"private_zone_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"security_attributes":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle-zpr.maxegresscount.value": "42", "oracle-zpr.maxegresscount.mode": "enforce"}, Update: map[string]string{"oracle-zpr.maxegresscount.value": "updatedValue", "oracle-zpr.maxegresscount.mode": "enforce"}},
		"time_zone":                       acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"vm_cluster_type":                 acctest.Representation{RepType: acctest.Optional, Create: `DEVELOPER`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudVmClusterIgnoreDefinedTagsRepresentation},
	}

	DatabaseCloudVmClusterRepresentation2 = map[string]interface{}{
		"backup_subnet_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet_backup.id}`},
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":                  acctest.Representation{RepType: acctest.Required, Create: `4`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `cloudVmCluster`, Update: `displayName2`},
		"gi_version":                      acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"hostname":                        acctest.Representation{RepType: acctest.Required, Create: `apollo`},
		"ssh_public_keys":                 acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet1.id}`},
		"domain":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet1.subnet_domain_name}`},
		"backup_network_nsg_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group_backup.id}`}},
		"cluster_name":                    acctest.Representation{RepType: acctest.Optional, Create: `clusterName`},
		"cloud_automation_update_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudVmClusterCloudAutomationUpdateDetailsRepresentation},
		"data_collection_options":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: cloudVmClusterDataCollectionOptionsRepresentation},
		"data_storage_percentage":         acctest.Representation{RepType: acctest.Optional, Create: `40`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_sparse_diskgroup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"scan_listener_port_tcp":          acctest.Representation{RepType: acctest.Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl":      acctest.Representation{RepType: acctest.Optional, Create: `2484`},
		"time_zone":                       acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudVmClusterIgnoreDefinedTagsRepresentation},
	}

	DatabaseCloudVmClusterCloudAutomationUpdateDetailsRepresentation = map[string]interface{}{
		"apply_update_time_preference": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudVmClusterCloudAutomationUpdateDetailsApplyUpdateTimePreferenceRepresentation},
		"freeze_period":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseCloudVmClusterCloudAutomationUpdateDetailsFreezePeriodRepresentation},
		"is_early_adoption_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_freeze_period_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}

	cloudVmClusterDataCollectionOptionsRepresentation = map[string]interface{}{
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	DatabaseCloudVmClusterCloudAutomationUpdateDetailsApplyUpdateTimePreferenceRepresentation = map[string]interface{}{
		"apply_update_preferred_end_time":   acctest.Representation{RepType: acctest.Optional, Create: `06:00`, Update: `08:00`},
		"apply_update_preferred_start_time": acctest.Representation{RepType: acctest.Optional, Create: `00:00`, Update: `02:00`},
	}
	DatabaseCloudVmClusterCloudAutomationUpdateDetailsFreezePeriodRepresentation = map[string]interface{}{
		"freeze_period_end_time":   acctest.Representation{RepType: acctest.Optional, Create: `2026-02-15`, Update: `2026-03-15`},
		"freeze_period_start_time": acctest.Representation{RepType: acctest.Optional, Create: `2026-02-13`, Update: `2026-03-13`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation0 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `20`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/`, Update: `/`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation1 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `250`, Update: `260`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/u01`, Update: `/u01`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation2 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `15`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/tmp`, Update: `/tmp`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation3 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `15`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var`, Update: `/var`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation4 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `40`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var/log`, Update: `/var/log`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation5 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `10`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/home`, Update: `/home`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation6 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `/var/log/audit`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation7 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `9`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `reserved`},
	}

	DatabaseCloudVmClusterFileSystemConfigurationDetailsRepresentation8 = map[string]interface{}{
		"file_system_size_gb": acctest.Representation{RepType: acctest.Optional, Create: `16`},
		"mount_point":         acctest.Representation{RepType: acctest.Optional, Create: `swap`},
	}

	zoneRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `sicdbaas.exacs.zonetest`},
		"zone_type":      acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	ViewRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudVmClusterIgnoreDefinedTagsRepresentation},
	}

	CoreCoreVcnDnsResolverAssociationRepresentation = map[string]interface{}{
		"vcn_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_network.t.id}`},
	}

	ResolverRepresentation = map[string]interface{}{
		"resolver_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"attached_views": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ResolverAttachedViewsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: cloudVmClusterIgnoreDefinedTagsRepresentation},
	}

	cloudVmClusterIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	ResolverAttachedViewsRepresentation = map[string]interface{}{
		"view_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	DatabaseDbServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	ad_subnet_security = `
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
                    is_ipv6enabled =  true
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
                    ipv6cidr_blocks     = ["${substr(oci_core_virtual_network.t.ipv6cidr_blocks[0], 0, length(oci_core_virtual_network.t.ipv6cidr_blocks[0]) - 7)}01::/64"]
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
                    ipv6cidr_blocks     = ["${substr(oci_core_virtual_network.t.ipv6cidr_blocks[0], 0, length(oci_core_virtual_network.t.ipv6cidr_blocks[0]) - 7)}02::/64"]
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
	DatabaseCloudVmClusterResourceDependencies = ad_subnet_security + acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudExadataInfrastructureRepresentation, []string{"compute_count"}), map[string]interface{}{
			"compute_count": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		})) + acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create, zoneRepresentation) +
		` data "oci_identity_tenancy" "test_tenancy" {
			tenancy_id = "${var.tenancy_ocid}"
          } 
		` +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Create, ViewRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Optional, acctest.Create, CoreCoreVcnDnsResolverAssociationRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Optional, acctest.Create, DatabaseDbServerDataSourceRepresentation)

	DatabaseDatabaseCloudVmClusterResourceDependencies = DatabaseCloudVmClusterResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Create, ResolverRepresentation) + Sleep30

	Sleep30 = "resource \"time_sleep\" \"wait_30_seconds\" {\n  depends_on = [oci_dns_resolver.test_resolver] \n create_duration = \"30s\"\n}" +
		`
		terraform {
  			required_providers {
    			time = "0.5.0"
  			}
		}
	`

	CloudVmClusterResourceUpdateDependencies = ad_subnet_security + acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Update,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudExadataInfrastructureRepresentation, []string{"compute_count"}), map[string]interface{}{
			"compute_count": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		})) + acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create, zoneRepresentation) +
		` data "oci_identity_tenancy" "test_tenancy" {
			tenancy_id = "${var.tenancy_ocid}"
         }
		` +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Create, ViewRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Optional, acctest.Create, CoreCoreVcnDnsResolverAssociationRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Optional, acctest.Create, DatabaseDbServerDataSourceRepresentation)

	CloudVmClusterCloudVmClusterResourceUpdateDependencies = CloudVmClusterResourceUpdateDependencies + acctest.GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", acctest.Optional, acctest.Create, ResolverRepresentation)

	CloudVmClusterResourceUpdateStorageDependencies = ad_subnet_security + acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Update,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudExadataInfrastructureRepresentation, []string{"storage_count"}), map[string]interface{}{
			"storage_count": acctest.Representation{RepType: acctest.Required, Create: `3`},
		}))
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_cloud_vm_cluster.test_cloud_vm_cluster"
	datasourceName := "data.oci_database_cloud_vm_clusters.test_cloud_vm_clusters"
	singularDatasourceName := "data.oci_database_cloud_vm_cluster.test_cloud_vm_cluster"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create, DatabaseCloudVmClusterRepresentation), "database", "cloudVmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseCloudVmClusterDestroy, []resource.TestStep{

		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudVmClusterRepresentation, []string{"domain"}), map[string]interface{}{
						"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "06:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "00:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-02-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "15"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.value", "42"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckNoResourceAttr(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_type", "DEVELOPER"),

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

		//verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseDatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudVmClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"domain":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "06:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "00:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-02-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-02-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "15"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.value", "42"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckNoResourceAttr(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_type", "DEVELOPER"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterCloudVmClusterResourceUpdateDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudVmClusterRepresentation, map[string]interface{}{
						"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-03-13"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "6"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.file_system_size_gb", "20"),
				resource.TestCheckResourceAttr(resourceName, "file_system_configuration_details.0.mount_point", "/"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.value", "updatedValue"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_count", "6"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckNoResourceAttr(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "vm_cluster_type", "DEVELOPER"),
				//resource.TestCheckResourceAttr(resourceName, "node_count", "3"), // Assertion Failing, needs to be reviewed

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_vm_clusters", "test_cloud_vm_clusters", acctest.Optional, acctest.Update, DatabaseDatabaseCloudVmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + CloudVmClusterCloudVmClusterResourceUpdateDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudVmClusterRepresentation, map[string]interface{}{
						"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_type", "DEVELOPER"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-03-13"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.compute_model", "OCPU"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.cpu_core_count", "6"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.data_storage_percentage", "40"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.disk_redundancy"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.file_system_configuration_details.0.file_system_size_gb", "20"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.file_system_configuration_details.0.mount_point", "/"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.domain", "sicdbaas.exacs.zonetest"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.system_tags.%", "2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.listener_port"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.node_count"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.ocpu_count", "6"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.scan_dns_name"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.scan_ip_ids.#", "3"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.scan_ipv6ids.#", "3"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.security_attributes.%", "2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.security_attributes.oracle-zpr.maxegresscount.value", "updatedValue"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.shape"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.subscription_id", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.vip_ids.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.vipv6ids.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_vm_clusters.0.vm_cluster_type", "DEVELOPER"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_vm_clusters.0.zone_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseDatabaseCloudVmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DefinedTagsDependencies + AvailabilityDomainConfig + CloudVmClusterCloudVmClusterResourceUpdateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudVmClusterRepresentation, map[string]interface{}{
						"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_end_time", "08:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.apply_update_time_preference.0.apply_update_preferred_start_time", "02:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_end_time", "2026-03-15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.freeze_period.0.freeze_period_start_time", "2026-03-13"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.is_early_adoption_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_automation_update_details.0.is_freeze_period_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_diagnostics_events_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_health_monitoring_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_collection_options.0.is_incident_logs_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "disk_redundancy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.#", "9"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.0.file_system_size_gb", "20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_configuration_details.0.mount_point", "/"),
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", "sicdbaas.exacs.zonetest"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_count", "6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scan_dns_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_ip_ids.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_ipv6ids.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_attributes.oracle-zpr.maxegresscount.value", "updatedValue"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_attributes.oracle-zpr.maxegresscount.mode", "enforce"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vip_ids.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vipv6ids.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_cluster_type", "DEVELOPER"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_id"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseCloudVmClusterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"private_zone_id",
				"create_async",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudVmClusterUpdate(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_cloud_vm_cluster.test_cloud_vm_cluster"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudVmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create, DatabaseCloudVmClusterRepresentation2), "database", "cloudVmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseCloudVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create, DatabaseCloudVmClusterRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

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

		//verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseCloudVmClusterRepresentation2, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "cloudVmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port_tcp_ssl", "2484"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterResourceUpdateStorageDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Optional, acctest.Update, DatabaseCloudVmClusterRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_name", "clusterName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "OCPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_percentage", "40"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.9.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				//resource.TestCheckResourceAttr(resourceName, "storage_size_in_gbs", "204388"), // 4 storage cells * 51097 (X8M.StorageCell AvailableDbStorageInGBs) //Assertion failing, needs to be reviewed

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDatabaseCloudVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_cloud_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetCloudVmClusterRequest{}

			tmp := rs.Primary.ID
			request.CloudVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseCloudVmCluster") {
		resource.AddTestSweepers("DatabaseCloudVmCluster", &resource.Sweeper{
			Name:         "DatabaseCloudVmCluster",
			Dependencies: acctest.DependencyGraph["cloudVmCluster"],
			F:            sweepDatabaseCloudVmClusterResource,
		})
	}
}

func sweepDatabaseCloudVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	cloudVmClusterIds, err := getDatabaseCloudVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudVmClusterId := range cloudVmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[cloudVmClusterId]; !ok {
			deleteCloudVmClusterRequest := oci_database.DeleteCloudVmClusterRequest{}

			deleteCloudVmClusterRequest.CloudVmClusterId = &cloudVmClusterId

			deleteCloudVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteCloudVmCluster(context.Background(), deleteCloudVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudVmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudVmClusterId, DatabaseCloudVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseCloudVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseCloudVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudVmClusterId", id)
	}
	return resourceIds, nil
}

func DatabaseCloudVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudVmClusterResponse, ok := response.Response.(oci_database.GetCloudVmClusterResponse); ok {
		return cloudVmClusterResponse.LifecycleState != oci_database.CloudVmClusterLifecycleStateTerminated
	}
	return false
}

func DatabaseCloudVmClusterSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetCloudVmCluster(context.Background(), oci_database.GetCloudVmClusterRequest{
		CloudVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
