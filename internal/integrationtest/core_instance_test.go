// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreInstanceRequiredOnlyResource = CoreSubnetResourceConfig + utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation)

	CoreInstanceResourceConfig = CoreInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, CoreInstanceRepresentation)

	CoreInstanceIpv6ResourceConfig = CoreInstanceResourceDependenciesIPV6 +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceWithAssignIPV6ResourceRepresentation)

	CoreCoreInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}
	CoreCoreInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceDataSourceFilterRepresentation}}
	CoreInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance.test_instance.id}`}},
	}
	instancePlatformConfigRepresentationForConfFlexShape = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `AMD_VM`},
		"is_memory_encryption_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	instanceShapeConfigRepresentationForConfFlexShape = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	CoreInstanceRepresentation = map[string]interface{}{
		"availability_domain":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"update_operation_constraint": acctest.Representation{RepType: acctest.Optional, Create: `ALLOW_DOWNTIME`, Update: `ALLOW_DOWNTIME`},
		"agent_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceSourceDetailsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	CoreFungibleInstanceRepresentation = map[string]interface{}{
		"availability_domain":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"update_operation_constraint": acctest.Representation{RepType: acctest.Optional, Create: `ALLOW_DOWNTIME`, Update: `ALLOW_DOWNTIME`},
		"agent_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation_FlexShape},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceSourceDetailsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	CoreInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"plugins_config":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigPluginsConfigRepresentation},
	}
	CoreInstanceAvailabilityConfigRepresentation = map[string]interface{}{
		"is_live_migration_preferred": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"recovery_action":             acctest.Representation{RepType: acctest.Optional, Create: `RESTORE_INSTANCE`, Update: `STOP_INSTANCE`},
	}
	CoreInstanceCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_ipv6ip":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"private_ip":                acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"skip_source_dest_check":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	CoreInstanceCreateVnicDetailsIpv6Representation = map[string]interface{}{
		"assign_ipv6ip":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"ipv6address_ipv6subnet_cidr_pair_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceIpv6AddressIpv6SubnetCidrPairRepresentation},
		"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"private_ip":             acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	CoreInstanceInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	CoreInstanceLaunchOptionsRepresentation = map[string]interface{}{
		"boot_volume_type":                    acctest.Representation{RepType: acctest.Optional, Create: `ISCSI`},
		"firmware":                            acctest.Representation{RepType: acctest.Optional, Create: `UEFI_64`},
		"is_consistent_volume_naming_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"network_type":                        acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
		"remote_data_volume_type":             acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
	CoreInstanceLaunchOptionsRepresentation_FlexShape = map[string]interface{}{
		"boot_volume_type":                    acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
		"firmware":                            acctest.Representation{RepType: acctest.Optional, Create: `UEFI_64`},
		"is_consistent_volume_naming_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"network_type":                        acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
		"remote_data_volume_type":             acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
	CoreInstanceWithIntelVmPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreFungibleInstanceRepresentation, map[string]interface{}{
		"fault_domain":    acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"image":           acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"source_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceFlexSourceDetailsRepresentation},
		"platform_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceIntelVmPlatformConfigRepresentationWithSMTDisabled},
	}), []string{
		"dedicated_vm_host_id",
	})
	CoreInstanceIntelVmPlatformConfigRepresentationWithSMTDisabled = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `INTEL_VM`},
		"is_symmetric_multi_threading_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
	CoreInstanceIntelVmPlatformConfigRepresentationWithSMTEnabled = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `INTEL_VM`},
		"is_symmetric_multi_threading_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	instanceSubCorePlatformConfigRepresentation = map[string]interface{}{
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `AMD_MILAN_BM`},
		"numa_nodes_per_socket": acctest.Representation{RepType: acctest.Optional, Create: `NPS0`},
	}
	CoreInstanceSourceDetailsRepresentation = map[string]interface{}{
		"source_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: `image`},
		"boot_volume_vpus_per_gb": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"kms_key_id":              acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `70`},
	}
	CoreInstanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `Compute Instance Monitoring`},
	}
	CoreInstanceSourceDetailsInstanceSourceImageFilterDetailsRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"operating_system": acctest.Representation{RepType: acctest.Optional, Create: `Oracle Linux`},
	}
	CoreInstanceIpv6AddressIpv6SubnetCidrPairRepresentation = map[string]interface{}{
		"ipv6subnet_cidr": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.ipv6cidr_blocks[0]}`},
		"ipv6address":     acctest.Representation{RepType: acctest.Optional, Create: `${substr(oci_core_subnet.test_subnet.ipv6cidr_blocks[0], 0, length(oci_core_subnet.test_subnet.ipv6cidr_blocks[0]) - 7)}${1000}`},
	}

	InstanceWithPVEncryptionInTransitEnabled = `
resource "oci_core_instance" "test_instance" {
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	image = "${var.InstanceImageOCID[var.region]}"
	is_pv_encryption_in_transit_enabled = "true"
	shape = "VM.Standard2.1"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
`
	CoreKeyResourceDependencyConfig = KmsKeyResourceDependencies + `
data "oci_kms_keys" "test_keys_dependency" {
    #Required
    compartment_id = "${var.tenancy_ocid}"
    management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
    algorithm = "AES"

    filter {
        name = "state"
        values = ["ENABLED", "UPDATING"]
    }
}
data "oci_kms_keys" "test_keys_dependency_RSA" {
    #Required
    compartment_id = "${var.tenancy_ocid}"
    management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
    algorithm = "RSA"

    filter {
        name = "state"
        values = ["ENABLED", "UPDATING"]
    }
}
`
	CoreInstanceResourceDependenciesWithoutDHV = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("cidr_block", acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/30`}, CoreVlanRepresentation)) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		CoreKeyResourceDependencyConfig

	CoreInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, CoreDedicatedVmHostRepresentation) +
		CoreInstanceResourceDependenciesWithoutDHV

	CoreInstanceResourceDependenciesIPV6 = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label":       acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Required, Create: []string{`${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("cidr_block", acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/30`}, CoreVlanRepresentation)) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	// ------------- For flex shape -------------
	InstanceWithPVEncryptionInTransitEnabledForFlexShape = `
	resource "oci_core_instance" "test_instance" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}"
		compartment_id = "${var.compartment_id}"
		image = "${var.FlexInstanceImageOCID[var.region]}"
		is_pv_encryption_in_transit_enabled = "true"
		shape = "VM.Standard.E3.Flex"
		subnet_id = "${oci_core_subnet.test_subnet.id}"
		shape_config {
			baseline_ocpu_utilization = "BASELINE_1_8"
			ocpus = 1
		}
	}
	`
	// We can not launch E3 flex instance in PHX ad1 eue to an temporary issue, use AD2 to get the test passed.
	// TODO: https://jira.oci.oraclecorp.com/browse/TERSI-674 to use AD1 for test of Flex shape
	instanceDataSourceRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("availability_domain",
		acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`}, CoreCoreInstanceDataSourceRepresentation)
	InstanceResourceConfigForFlexShape = CoreInstanceResourceDependenciesWithoutDHV +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationForFlexShape)
	instanceSourceDetailsRepresentationForFlexShape = acctest.GetMultipleUpdatedRepresenationCopy(
		[]string{"source_id", "boot_volume_size_in_gbs"},
		[]interface{}{
			acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
			acctest.Representation{RepType: acctest.Optional, Create: `60`},
		},
		CoreInstanceSourceDetailsRepresentation)

	instanceShapeConfigRepresentationForFlexShape = map[string]interface{}{
		"baseline_ocpu_utilization": acctest.Representation{RepType: acctest.Required, Create: `BASELINE_1_8`, Update: `BASELINE_1_2`},
		"memory_in_gbs":             acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `4.0`},
		"ocpus":                     acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	instanceLaunchOptionsRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("boot_volume_type",
		acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`}, CoreInstanceLaunchOptionsRepresentation_FlexShape)

	instanceRepresentationForFlexShape = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"availability_domain", "shape", "image", "create_vnic_details", "launch_options", "source_details", "shape_config"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
				acctest.Representation{RepType: acctest.Required, Create: InstanceConfigurationVmShapeForFlex},
				acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceShapeConfigRepresentationForFlexShape},
			},
			CoreInstanceRepresentation),
		[]string{"dedicated_vm_host_id"},
	)
	instanceShapeConfigRepresentation_ForFungibleShapeConfig = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: "1"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `10.0`},
	}
	instanceRepresentationCore_ForFugibleShapeConfig = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreFungibleInstanceRepresentation, map[string]interface{}{
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.AMD.Generic`},
		"image":          acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForFungibleShapeConfig},
		"source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceFlexSourceDetailsRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceShapeConfigRepresentation_ForFugibleShapeConfigUpdate = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Update: "2"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Update: `20.0`},
	}
	instanceRepresentationCore_ForFugibleShapeConfigUpdate = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreFungibleInstanceRepresentation, map[string]interface{}{
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.AMD.Generic`},
		"image":          acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForFugibleShapeConfigUpdate},
		"source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceFlexSourceDetailsRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceRepresentationForConfidentialFlexShape = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"agent_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentationForFlexShape},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceShapeConfigRepresentationForConfFlexShape},
		"platform_config":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePlatformConfigRepresentationForConfFlexShape},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsRepresentationForFlexShape},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	SourceDetailsRepresentationForBootVolumeUpdateKmsKeyOne = map[string]interface{}{
		"source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
		"kms_key_id": acctest.Representation{RepType: acctest.Optional,
			Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `70`},
	}

	SourceDetailsRepresentationForBootVolumeUpdateKmsKeyTwo = map[string]interface{}{
		"source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
		"kms_key_id": acctest.Representation{RepType: acctest.Optional,
			Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[1], "id")}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `70`},
	}

	InstanceWithNoKmsKeyInSourceDetailsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `tf_boot_volume_kms_key_update_test`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsSansKmsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	InstanceWithKmsKeyOneInSourceDetailsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `tf_boot_volume_kms_key_update_test`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: SourceDetailsRepresentationForBootVolumeUpdateKmsKeyOne},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	InstanceWithKmsKeyTwoInSourceDetailsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `tf_boot_volume_kms_key_update_test`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: SourceDetailsRepresentationForBootVolumeUpdateKmsKeyTwo},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	// ------------- for capacity reservation -------------
	instanceSourceDetailsSansKmsRepresentation = map[string]interface{}{
		"source_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: `image`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `70`},
	}

	instanceWithCapacityReservationRepresentation = map[string]interface{}{
		// dedicated_vm_host_id is incompatible with capacity_reservation_id
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"create_vnic_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsSansKmsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	instanceWithCapacityReservationResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, CoreComputeCapacityReservationRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	instanceWithCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceDataSourceFilterRepresentation}}

	// ---- Architecture Agnostic Launch with Instance Configuration ----
	InstanceRepresentationWithInstanceConfiguration = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"instance_configuration_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	InstanceConfigurationRepresentationWithImageFilters = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `example-instance-configuration`},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: InstanceConfigurationInstanceDetailsRepresentationWithImageFilters},
	}
	InstanceConfigurationInstanceDetailsRepresentationWithImageFilters = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `instance_options`},
		"options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: InstanceConfigurationInstanceDetailsOptionsRepresentationWithImageFilters},
	}
	InstanceConfigurationInstanceDetailsOptionsRepresentationWithImageFilters = map[string]interface{}{
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: InstanceConfigurationInstanceDetailsOptionsLaunchDetailsRepresentationWithImageFilters},
	}
	InstanceConfigurationInstanceDetailsOptionsLaunchDetailsRepresentationWithImageFilters = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `example-instance-configuration-instance`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}},
		"metadata":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}},
		"shape":          acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: InstanceConfigurationInstanceDetailsOptionsLaunchDetailsSourceDetailsRepresentationWithImageFilters},
	}
	InstanceConfigurationInstanceDetailsOptionsLaunchDetailsSourceDetailsRepresentationWithImageFilters = map[string]interface{}{
		"source_type":                          acctest.Representation{RepType: acctest.Required, Create: `image`},
		"instance_source_image_filter_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceSourceDetailsInstanceSourceImageFilterDetailsRepresentation},
	}
	InstanceResourceDependenciesWithInstanceConfiguration = CoreInstanceResourceDependenciesWithoutDHV +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, InstanceConfigurationRepresentationWithImageFilters)

	// ---- Launch with source detail image filters ----
	InstanceRepresentationWithImageFilters = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"update_operation_constraint":         acctest.Representation{RepType: acctest.Optional, Create: `ALLOW_DOWNTIME`, Update: `ALLOW_DOWNTIME`},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: InstanceSourceDetailsRepresentationWithImageFilters},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	InstanceSourceDetailsRepresentationWithImageFilters = map[string]interface{}{
		"source_type":                          acctest.Representation{RepType: acctest.Required, Create: `image`},
		"instance_source_image_filter_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceSourceDetailsInstanceSourceImageFilterDetailsRepresentation},
	}
	instanceWithAssignIPV6ResourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsIpv6Representation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, Update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsSansKmsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementEndpointStr+CoreInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, CoreInstanceRepresentation), "core", "instance", t)

	acctest.ResourceTest(t, testAccCheckCoreInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to shape within the same family is not force new. Resizing can only be done to intances not using dedicated_vm_host_id
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("shape", acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`}, CoreInstanceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies,
		},
		// verify Create with is_pv_encryption_in_transit_enabled = true
		{
			Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies + InstanceWithPVEncryptionInTransitEnabled,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "region"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, CoreInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.is_live_migration_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_ipv6ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "image"),
				resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "region"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_vpus_per_gb", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.is_live_migration_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_ipv6ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
				//resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "image"),
				resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "region"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_vpus_per_gb", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_operation_constraint", "ALLOW_DOWNTIME"),

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
			Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("update_operation_constraint", acctest.Representation{RepType: acctest.Optional, Update: `ALLOW_DOWNTIME`}, CoreInstanceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.is_live_migration_preferred", "true"),
				resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_ipv6ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "image"),
				resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "region"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_vpus_per_gb", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "70"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_operation_constraint", "ALLOW_DOWNTIME"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, CoreCoreInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, CoreInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

				resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.0.is_live_migration_preferred", "true"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.0.recovery_action", "STOP_INSTANCE"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.dedicated_vm_host_id"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.is_cross_numa_node"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.launch_mode"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_consistent_volume_naming_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.vcpus", "2"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.boot_volume_vpus_per_gb", "10"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.0.is_live_migration_preferred", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cross_numa_node"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "launch_mode"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.vcpus", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.boot_volume_vpus_per_gb", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "70"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
			),
		},
		// verify updates to original parameters
		{
			Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithRemovedProperties(CoreInstanceSourceDetailsRepresentation, []string{"boot_volume_size_in_gbs"})},
						CoreInstanceRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image"),
				resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "region"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "70"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "update_operation_constraint", "ALLOW_DOWNTIME"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + CoreInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// extended_metadata is set by import but service may potentially reorder map elements in imported JSON strings.
				// This is normally handled by diff suppress function but the Terraform import tests can't invoke diff suppression
				// and so it may complain that values are different.
				"extended_metadata",
				"hostname_label",
				"update_operation_constraint",
				"is_pv_encryption_in_transit_enabled",
				"create_vnic_details.0.assign_private_dns_record",
				"subnet_id",
				"source_details.0.kms_key_id", //TODO: Service is not returning this value, remove when the service returns it. COM-26394
			},
			ImportStateCheck: func(states []*terraform.InstanceState) error {
				var instanceState *terraform.InstanceState
				for _, state := range states {
					if state.ID == resId {
						instanceState = state
						break
					}
				}

				if instanceState == nil {
					return fmt.Errorf("could not find the imported instance state")
				}

				expectedExtendedMetadataMap := CoreInstanceRepresentation["extended_metadata"].(acctest.Representation).Create.(map[string]string)

				expectedValue := fmt.Sprintf("%d", len(expectedExtendedMetadataMap))
				if actualValue := instanceState.Attributes["extended_metadata.%"]; actualValue != expectedValue {
					return fmt.Errorf("expected 'extended_metadata' to have %s items, but got %s", expectedValue, actualValue)
				}

				for key, expectedJsonString := range expectedExtendedMetadataMap {
					attributeKey := fmt.Sprintf("extended_metadata.%s", key)
					actualJsonString, exists := instanceState.Attributes[attributeKey]
					if !exists {
						return fmt.Errorf("could not find expected attribute '%s' in imported state", attributeKey)
					}

					expectedJsonString = strings.Replace(expectedJsonString, "\\\"", "\"", -1)
					if err := acctest.CheckJsonStringsEqual(expectedJsonString, actualJsonString); err != nil {
						return fmt.Errorf("%s: Attribute '%s' %s", resourceName, attributeKey, err)
					}
				}
				return nil
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_updateBootVolumeKmsKey(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_updateBootVolumeKmsKey")
	defer httpreplay.SaveScenario()

	config := `
		provider oci {
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2, resId3 string
	// Save TF content to Create resource with optional properties.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementEndpointStr+CoreInstanceResourceDependenciesWithoutDHV+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
			InstanceWithNoKmsKeyInSourceDetailsRepresentation), "core", "instance", t)

	acctest.ResourceTest(t, testAccCheckCoreInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					InstanceWithNoKmsKeyInSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to add kms key id in source details
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					InstanceWithKmsKeyOneInSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.kms_key_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify Update to change kms key id in source details
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					InstanceWithKmsKeyTwoInSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttrSet(resourceName, "source_details.0.kms_key_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId3 != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_capacityReservation(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_capacityReservation")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify Create with optionals
			{
				Config: config +
					compartmentIdVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceWithCapacityReservationRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// Step 1: verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config +
					compartmentIdVariableStr +
					compartmentIdUVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(instanceWithCapacityReservationRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Step 2: verify updates to updatable parameters
			{
				Config: config +
					compartmentIdVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceWithCapacityReservationRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "70"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 3: verify datasource
			{
				Config: config +
					compartmentIdVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceWithCapacityReservationRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, instanceWithCapacityReservationDataSourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.capacity_reservation_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.launch_mode"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
			// Step 4: verify singular datasource
			{
				Config: config +
					compartmentIdVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceWithCapacityReservationRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "launch_mode"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "70"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_flexShape(t *testing.T) {
	httpreplay.SetScenario("TestCoreFlexInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// step 0 verify Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentationForFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					// currently E3 subcore is forced to use launch_mode = PARAVIRTUALIZED
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// step 1 delete before next Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable,
			},

			// step 2 verify Create with is_pv_encryption_in_transit_enabled = true
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable + InstanceWithPVEncryptionInTransitEnabledForFlexShape,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// step 3 delete before next Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable,
			},

			// step 4 verify Create with optionals
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentationForFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			//step 5: verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + compartmentIdUVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(instanceRepresentationForFlexShape, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// step 6: verify updates to updatable parameters
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationForFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					// resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					// resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					// resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_2"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "4"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// step 7: verify datasource
			{
				Config: acctest.ProviderTestConfig() +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, instanceDataSourceRepresentationForFlexShape) +
					compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationForFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.launch_mode"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.baseline_ocpu_utilization", "BASELINE_1_2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},

			// step 8: verify singular datasource
			{
				Config: acctest.ProviderTestConfig() + utils.FlexVmImageIdsVariable +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + managementEndpointStr + InstanceResourceConfigForFlexShape,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "availability_config.0.recovery_action", "STOP_INSTANCE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "launch_mode"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
				),
			},
		},
	})
}

func TestCoreInstanceResource_ConfidentialflexShape(t *testing.T) {
	httpreplay.SetScenario("TestCoreFlexInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// step 0 verify Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentationForConfidentialFlexShape),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "AMD_VM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_memory_encryption_enabled", "true"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if err != nil {
							fmt.Println(resId)
						}
						return err
					},
				),
			},
		},
	})
}

// fungible shape config update
func TestAccResourceCoreFungibleInstance_UpdateShapeConfig(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreFungibleInstance_UpdateShapeConfig")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider

	config := `
      provider oci {
         test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
      }
   ` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// step 0 verify Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentationCore_ForFugibleShapeConfig),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.AMD.Generic"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					// currently E3 subcore is forced to use launch_mode = PARAVIRTUALIZED
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to with changes in shape_config
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForFugibleShapeConfigUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.AMD.Generic"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "20"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "2"),
					// currently E3 subcore is forced to use launch_mode = PARAVIRTUALIZED
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

// platform config update
func TestAccResourceCoreInstance_UpdatePlatformConfig(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreFungibleInstance_UpdatePlatformConfig")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider

	config := `
      provider oci {
         test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
      }
   ` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{

			// Step 0 - verify create with platform_config = null
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance",
						acctest.Optional, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(CoreInstanceWithIntelVmPlatformConfigRepresentation, []string{"platform_config"})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// step 1 verify updates to a null platform_config object does not destroy the resource
			// (Only platform_config_type and platform_config.is_symmetric_multi_threading_enabled property updates for VM's only - Rest will ForceNew)
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance",
						acctest.Optional, acctest.Create, CoreInstanceWithIntelVmPlatformConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "INTEL_VM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify updates to existing platform_config.0.is_symmetric_multi_threading_enabled property does not destroy the resource
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("platform_config",
							acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceIntelVmPlatformConfigRepresentationWithSMTEnabled},
							CoreInstanceWithIntelVmPlatformConfigRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "INTEL_VM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify re-size operation updates the platform_config.0.type and retains the old SMT value
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
						acctest.GetMultipleUpdatedRepresenationCopy(
							[]string{"shape", "shape_config", "platform_config"},
							[]interface{}{
								acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
								acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForFungibleShapeConfig},
								acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceIntelVmPlatformConfigRepresentationWithSMTEnabled},
							}, CoreInstanceWithIntelVmPlatformConfigRepresentation),
					),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_aailInstanceConfiguration(t *testing.T) {

	httpreplay.SetScenario("TestCoreInstanceResource_aailInstanceConfiguration")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + managementEndpointStr + InstanceResourceDependenciesWithInstanceConfiguration +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, InstanceRepresentationWithInstanceConfiguration),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// Step 1: verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + managementEndpointStr + InstanceResourceDependenciesWithInstanceConfiguration +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, InstanceRepresentationWithInstanceConfiguration),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("3. Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 2: verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + managementEndpointStr + InstanceResourceDependenciesWithInstanceConfiguration +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, InstanceRepresentationWithInstanceConfiguration),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.instance_configuration_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
			// Step 3: verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + managementEndpointStr + InstanceResourceDependenciesWithInstanceConfiguration +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, InstanceRepresentationWithInstanceConfiguration),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_imageSourceWithImageFilters(t *testing.T) {

	httpreplay.SetScenario("TestCoreInstanceResource_imageSourceWithImageFilters")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, InstanceRepresentationWithImageFilters),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.0.operating_system", "Oracle Linux"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// Step 1: verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, InstanceRepresentationWithImageFilters),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.vcpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.instance_source_image_filter_details.0.operating_system", "Oracle Linux"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("3. Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_AssignIpv6(t *testing.T) {

	httpreplay.SetScenario("TestCoreInstanceResource_AssignIpv6")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesIPV6 +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceWithAssignIPV6ResourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.are_all_plugins_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.desired_state", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.plugins_config.0.name", "Compute Instance Monitoring"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.is_live_migration_preferred", "false"),
					resource.TestCheckResourceAttr(resourceName, "availability_config.0.recovery_action", "RESTORE_INSTANCE"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_ipv6ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"),
					resource.TestMatchResourceAttr(resourceName, "create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6address", regexp.MustCompile(".*1000$")),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_vpus_per_gb", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		},
	})
}

func testAccCheckCoreInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance" {
			noResourceFound = false
			request := oci_core.GetInstanceRequest{}

			tmp := rs.Primary.ID
			request.InstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InstanceLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreInstance") {
		resource.AddTestSweepers("CoreInstance", &resource.Sweeper{
			Name:         "CoreInstance",
			Dependencies: acctest.DependencyGraph["instance"],
			F:            sweepCoreInstanceResource,
		})
	}
}

func sweepCoreInstanceResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	instanceIds, err := getInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceId := range instanceIds {
		if ok := acctest.SweeperDefaultResourceId[instanceId]; !ok {
			terminateInstanceRequest := oci_core.TerminateInstanceRequest{}

			terminateInstanceRequest.InstanceId = &instanceId

			terminateInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.TerminateInstance(context.Background(), terminateInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting Instance %s %s, It is possible that the resource is already deleted. Please verify manually \n", instanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &instanceId, instanceSweepWaitCondition, time.Duration(3*time.Minute),
				instanceSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listInstancesRequest := oci_core.ListInstancesRequest{}
	listInstancesRequest.CompartmentId = &compartmentId
	listInstancesRequest.LifecycleState = oci_core.InstanceLifecycleStateRunning
	listInstancesResponse, err := computeClient.ListInstances(context.Background(), listInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Instance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instance := range listInstancesResponse.Items {
		id := *instance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstanceId", id)
	}
	return resourceIds, nil
}

func instanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if instanceResponse, ok := response.Response.(oci_core.GetInstanceResponse); ok {
		return instanceResponse.LifecycleState != oci_core.InstanceLifecycleStateTerminated
	}
	return false
}

func instanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetInstance(context.Background(), oci_core.GetInstanceRequest{
		InstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
