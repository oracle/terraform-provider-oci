// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

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
	CoreInstanceConfigurationRequiredOnlyResource = CoreInstanceConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, CoreInstanceConfigurationRepresentation)

	CoreInstanceConfigurationResourceConfig = CoreInstanceConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, CoreInstanceConfigurationRepresentation)

	CoreCoreInstanceConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
	}

	CoreCoreInstanceConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationDataSourceFilterRepresentation}}
	CoreInstanceConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_configuration.test_instance_configuration.id}`}},
	}

	CoreInstanceConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsRepresentation},
		"source":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	CoreInstanceConfigurationRepresentationWithOptions = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsWithOptionsRepresentation},
		"source":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	CoreInstanceConfigurationRepresentationImageFilters = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsRepresentation},
		"source":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	CoreInstanceConfigurationRepresentationIPv6 = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsRepresentationIpv6},
		"source":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	CoreInstanceConfigurationFromInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"source":         acctest.Representation{RepType: acctest.Optional, Create: `INSTANCE`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchRepresentation = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
	}

	CoreInstanceConfigurationInstanceDetailsLaunchRepresentationWithFilterDetails = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationImageFilters},
	}

	CoreInstanceConfigurationInstanceDetailsLaunchRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("launch_details",
		acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationForFlexShape},
		CoreInstanceConfigurationInstanceDetailsLaunchRepresentation)

	CoreInstanceConfigurationInstanceDetailsLaunchRepresentationForDenseShape = acctest.GetUpdatedRepresentationCopy("launch_details",
		acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationForDenseShape},
		CoreInstanceConfigurationInstanceDetailsLaunchRepresentation)

	CoreInstanceConfigurationInstanceDetailsBlockRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"block_volumes": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsParavirtualizedBlockRepresentation = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"block_volumes ": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsRepresentation = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationInstanceDetailsSecondaryVnicsRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsWithOptionsRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `instance_options`},
		"options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsOptionsRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsWithOptionsAndFilterDetailsRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `instance_options`},
		"options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsOptionsWithFilterDetailsRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsWithOptionsAndBlockVolumesRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `instance_options`},
		"options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsOptionsAndBlockVolumesRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsRepresentationIpv6 = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationInstanceDetailsSecondaryVnicsRepresentationWithIpv6SubnetCidrAndRange},
	}
	CoreInstanceConfigurationInstanceDetailsBlockVolumesRepresentation = map[string]interface{}{
		"create_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	CoreInstanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation = map[string]interface{}{
		"attach_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	CoreInstanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachRepresentation = map[string]interface{}{
		"attach_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	CoreInstanceShapeConfigRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: "1"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: "15"},
	}
	CoreInstanceConfigurationInstanceLaunchOptionsRepresentation = map[string]interface{}{
		"network_type": acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
	CoreInstanceConfigurationAvailabilityConfigRepresentation = map[string]interface{}{
		"is_live_migration_preferred": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"recovery_action":             acctest.Representation{RepType: acctest.Optional, Create: `RESTORE_INSTANCE`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation},
		"availability_config":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationAvailabilityConfigRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"extended_metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"extendedMetadata": "extendedMetadata"}, Update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShape},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceLaunchOptionsRepresentation},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceOptionsRepresentation},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"dedicated_vm_host_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"launch_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `NATIVE`},
		"preferred_maintenance_action":        acctest.Representation{RepType: acctest.Optional, Create: `LIVE_MIGRATE`},
		"security_attributes":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]any{"Oracle-DataSecurity-ZPR": map[string]any{"MaxEgressCount": map[string]string{"value": "42", "mode": "audit"}}}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationImageFilters = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"shape":               acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShape},
		"source_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationImageFilters},
		"shape_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationForFlexShape = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"shape", "source_details", "shape_config"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShapeForFlex},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentationForFlexShape},
			},
			CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"dedicated_vm_host_id", "preferred_maintenance_action"},
	)
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationForDenseShape = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"shape", "source_details", "shape_config"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShapeForDense},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForDenseShape},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentationForNvmeShape},
			},
			CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"dedicated_vm_host_id", "preferred_maintenance_action"},
	)
	CoreInstanceConfigurationInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	CoreInstanceConfigurationInstanceDetailsOptionsRepresentation = map[string]interface{}{
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsOptionsWithFilterDetailsRepresentation = map[string]interface{}{
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentationImageFilters},
	}
	CoreInstanceConfigurationInstanceDetailsOptionsAndBlockVolumesRepresentation = map[string]interface{}{
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
		"block_volumes":  acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsRepresentation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}
	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsRepresentationWithIpv6SubnetCidrAndRange = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentationIpv6SubnetCidrAndRange},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}
	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsRepresentationIpv6WithIpv6SubnetCidr = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentationIpv6SubnetCidr},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}
	CoreInstanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `iscsi`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"is_read_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"use_chap":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	CoreInstanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachDetailsRepresentation = map[string]interface{}{
		"type":                                acctest.Representation{RepType: acctest.Required, Create: `paravirtualized`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"device":                              acctest.Representation{RepType: acctest.Optional, Create: `server`},
		"is_read_only":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_shareable":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	CoreInstanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation = map[string]interface{}{
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"backup_policy_id":           acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"block_volume_replicas":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationBlockVolumeReplicaDetailsRepresentation},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"compartment_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_tune_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"size_in_gbs":                acctest.Representation{RepType: acctest.Optional, Create: `50`},
		"source_details":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation},
		"vpus_per_gb":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	CoreInstanceConfigurationBlockVolumeReplicaDetailsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `display_name`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_ip":                acctest.Representation{RepType: acctest.Optional, Create: `privateIp`},
		"security_attributes":       acctest.Representation{RepType: acctest.Optional, Create: map[string]any{"Oracle-DataSecurity-ZPR": map[string]any{"MaxEgressCount": map[string]string{"value": "42", "mode": "audit"}}}},
		"skip_source_dest_check":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation = map[string]interface{}{
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: `image`},
		"image_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `55`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationImageFilters = map[string]interface{}{
		"source_type":                          acctest.Representation{RepType: acctest.Required, Create: `image`},
		"instance_source_image_filter_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsOptionsLaunchDetailsSourceDetailsInstanceSourceImageFilterDetailsRepresentation},
	}
	CoreInstanceConfigurationInstanceDetailsOptionsLaunchDetailsSourceDetailsInstanceSourceImageFilterDetailsRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"operating_system": acctest.Representation{RepType: acctest.Optional, Create: `Oracle Linux`},
	}
	CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("image_id",
		acctest.Representation{RepType: acctest.Optional, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)

	CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForDenseShape = acctest.GetUpdatedRepresentationCopy("image_id",
		acctest.Representation{RepType: acctest.Optional, Create: `${var.image_id}`},
		CoreInstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)

	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_ip":                acctest.Representation{RepType: acctest.Optional, Create: `privateIp`},
		"skip_source_dest_check":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentationIpv6SubnetCidrAndRange = map[string]interface{}{
		"assign_ipv6ip":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"ipv6address_ipv6subnet_cidr_pair_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationIpv6AddressIpv6SubnetCidrPairRepresentation},
		"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_ip":             acctest.Representation{RepType: acctest.Optional, Create: `privateIp`},
		"security_attributes":    acctest.Representation{RepType: acctest.Optional, Create: map[string]any{"Oracle-DataSecurity-ZPR": map[string]any{"MaxEgressCount": map[string]string{"value": "42", "mode": "audit"}}}},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	CoreInstanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentationIpv6SubnetCidr = map[string]interface{}{
		"assign_ipv6ip":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"ipv6address_ipv6subnet_cidr_pair_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationIpv6AddressIpv6SubnetCidrPairRepresentationClusterNetwork},
		"nsg_ids":                acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_ip":             acctest.Representation{RepType: acctest.Optional, Create: `privateIp`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	CoreInstanceConfigurationIpv6AddressIpv6SubnetCidrPairRepresentationClusterNetwork = map[string]interface{}{
		"ipv6subnet_cidr": acctest.Representation{RepType: acctest.Optional, Create: `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`},
	}
	CoreInstanceConfigurationIpv6AddressIpv6SubnetCidrPairRepresentation = map[string]interface{}{
		"ipv6subnet_cidr": acctest.Representation{RepType: acctest.Optional, Create: `2607:f590:0000:1600:0000:0000:0000:0000/64`},
		"ipv6address":     acctest.Representation{RepType: acctest.Optional, Create: `2607:f590:0000:1600:3000:0000:0000:0000`},
	}
	CoreInstanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `volume`},
		"id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}

	CoreInstanceConfigurationResourceDependenciesWithoutKms = acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Required, acctest.Create, CoreBootVolumeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, CoreDedicatedVmHostRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		utils.VolumeBackupPolicyDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	CoreInstanceConfigurationResourceDependencies = CoreInstanceConfigurationResourceDependenciesWithoutKms +
		KeyResourceDependencyConfig
	InstanceConfigurationVmShape         = `VM.Standard2.1`
	InstanceConfigurationVmShapeForFlex  = `VM.Standard.E3.Flex`
	InstanceConfigurationVmShapeForDense = `VM.DenseIO.E4.Flex`

	CoreInstanceConfigurationResourceImageConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation))
)

// issue-routing-tag: core/computeManagement
func TestCoreInstanceConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_ocid\" { default = \"%s\" }\n", kmsKeyId)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	imageId := utils.GetEnvSettingWithBlankDefault("image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"
	datasourceName := "data.oci_core_instance_configurations.test_instance_configurations"
	singularDatasourceName := "data.oci_core_instance_configuration.test_instance_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+vaultIdVariableStr+kmsKeyIdVariableStr+compartmentIdVariableStr+CoreInstanceConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentation), "core", "instanceConfiguration", t)
	acctest.ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},
		// verify create regular instance configuration with imageFilter details
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentationWithFilterDetails}, CoreInstanceConfigurationRepresentationImageFilters)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.instance_source_image_filter_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.instance_source_image_filter_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.instance_source_image_filter_details.0.operating_system", "Oracle Linux"),
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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},
		// verify Create from instance_id
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationFromInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},

		// verify Create with optionals launch_details for E3 flex micro shape
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies + utils.FlexVmImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentationForFlexShape}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.is_live_migration_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.recovery_action", "RESTORE_INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.extended_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", InstanceConfigurationVmShapeForFlex),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.boot_volume_size_in_gbs", "55"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.source_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.memory_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
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

		// delete before next Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},

		// verify Create with optionals launch_details for E4 dense shape
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + imageIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentationForDenseShape}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.extended_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", InstanceConfigurationVmShapeForDense),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.boot_volume_size_in_gbs", "55"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.ocpus", "8"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.memory_in_gbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.nvmes", "1"),
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

		// delete before next Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},
		// verify Create with optionals launch_details
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.is_live_migration_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.recovery_action", "RESTORE_INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.extended_metadata.%", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", InstanceConfigurationVmShape),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.boot_volume_size_in_gbs", "55"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.source_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.source_type", "image"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_options.0.network_type", "PARAVIRTUALIZED"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.dedicated_vm_host_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.preferred_maintenance_action", "LIVE_MIGRATE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape_config.0.memory_in_gbs", "15"),
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
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsLaunchRepresentation}, CoreInstanceConfigurationRepresentation),
					map[string]interface{}{"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.is_live_migration_preferred", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.availability_config.0.recovery_action", "RESTORE_INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.extended_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", InstanceConfigurationVmShape),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.source_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.source_type", "image"),
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
		// verify recreate with optionals block_volumes.create_details
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsBlockRepresentation}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.0.type", "volume"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.vpus_per_gb", "10"),
				// resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.is_auto_tune_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.block_volume_replicas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.block_volume_replicas.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.block_volume_replicas.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.volume_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify recreate with optionals block_volumes.create_details to block_volumes.attach_details
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.GetUpdatedRepresentationCopy("block_volumes", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation}, CoreInstanceConfigurationInstanceDetailsBlockRepresentation)}, CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.type", "iscsi"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.use_chap", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.volume_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource was not recreated")
					}
					return err
				},
			),
		},
		// verify recreate with optionals block_volumes.create_details to block_volumes.attach_details-paravirtualized
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.GetUpdatedRepresentationCopy("block_volumes", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachRepresentation},
						CoreInstanceConfigurationInstanceDetailsBlockRepresentation)},
						CoreInstanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.device", "server"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.is_read_only", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.type", "paravirtualized"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.attach_details.0.is_shareable", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.volume_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource was not recreated")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},
		// verify Create with optionals secondary_vnics
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, CoreInstanceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
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
		// verify datasource
		{
			Config: config +
				vaultIdVariableStr + kmsKeyIdVariableStr + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", acctest.Optional, acctest.Update, CoreCoreInstanceConfigurationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, CoreInstanceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "instance_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_configurations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_configurations.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				vaultIdVariableStr + kmsKeyIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, CoreCoreInstanceConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstanceConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "deferred_fields.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + CoreInstanceConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"instance_id",
				"source",
				"cluster_placement_group_id",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: core/computeManagement
func TestCoreInstanceConfigurationResourceOptions_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceConfigurationResourceOptions_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreInstanceConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentationWithOptions), "core", "instanceConfiguration", t)
	acctest.ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// verify Create with optionals to test instance options
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentationWithOptions),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "instance_options"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.security_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.display_name", "backend-servers"),
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
		// verify Update to the compartment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsWithOptionsRepresentation}, CoreInstanceConfigurationRepresentationWithOptions),
					map[string]interface{}{"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "instance_options"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.security_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.display_name", "backend-servers"),
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
		// verify recreate with optionals block_volumes.create_details
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsWithOptionsAndBlockVolumesRepresentation}, CoreInstanceConfigurationRepresentationWithOptions)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "instance_options"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.source_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.source_details.0.type", "volume"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.block_volumes.0.create_details.0.vpus_per_gb", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies,
		},
		// verify Create with imageFilter details for instance Options
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceConfigurationInstanceDetailsWithOptionsAndFilterDetailsRepresentation}, CoreInstanceConfigurationRepresentationWithOptions)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "instance_options"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.security_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.options.0.launch_details.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.source_details.0.instance_source_image_filter_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.source_details.0.instance_source_image_filter_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.options.0.launch_details.0.source_details.0.instance_source_image_filter_details.0.operating_system", "Oracle Linux"),
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
	})
}

// issue-routing-tag: core/computeManagement
func TestCoreInstanceConfigurationResourceIpv6_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceConfigurationResourceIpv6_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"
	singularDatasourceName := "data.oci_core_instance_configuration.test_instance_configuration"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreInstanceConfigurationResourceDependenciesWithoutKms+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentation), "core", "instanceConfiguration", t)
	acctest.ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// verify Create with optionals secondary_vnics to test ipv6 support
		{
			Config: config + compartmentIdVariableStr + CoreInstanceConfigurationResourceDependenciesWithoutKms +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentationIPv6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6address", "2607:f590:0000:1600:3000:0000:0000:0000"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr", "2607:f590:0000:1600:0000:0000:0000:0000/64"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.security_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
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

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, CoreCoreInstanceConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstanceConfigurationResourceDependenciesWithoutKms +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstanceConfigurationRepresentationIPv6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "deferred_fields.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_ipv6ip", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6address", "2607:f590:0000:1600:3000:0000:0000:0000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr", "2607:f590:0000:1600:0000:0000:0000:0000/64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.security_attributes.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

func testAccCheckCoreInstanceConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_configuration" {
			noResourceFound = false
			request := oci_core.GetInstanceConfigurationRequest{}

			tmp := rs.Primary.ID
			request.InstanceConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetInstanceConfiguration(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("CoreInstanceConfiguration") {
		resource.AddTestSweepers("CoreInstanceConfiguration", &resource.Sweeper{
			Name:         "CoreInstanceConfiguration",
			Dependencies: acctest.DependencyGraph["instanceConfiguration"],
			F:            sweepCoreInstanceConfigurationResource,
		})
	}
}

func sweepCoreInstanceConfigurationResource(compartment string) error {
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()
	instanceConfigurationIds, err := getCoreInstanceConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceConfigurationId := range instanceConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[instanceConfigurationId]; !ok {
			deleteInstanceConfigurationRequest := oci_core.DeleteInstanceConfigurationRequest{}

			deleteInstanceConfigurationRequest.InstanceConfigurationId = &instanceConfigurationId

			deleteInstanceConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeManagementClient.DeleteInstanceConfiguration(context.Background(), deleteInstanceConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting InstanceConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", instanceConfigurationId, error)
				continue
			}
		}
	}
	return nil
}

func getCoreInstanceConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstanceConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()

	listInstanceConfigurationsRequest := oci_core.ListInstanceConfigurationsRequest{}
	listInstanceConfigurationsRequest.CompartmentId = &compartmentId
	listInstanceConfigurationsResponse, err := computeManagementClient.ListInstanceConfigurations(context.Background(), listInstanceConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InstanceConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instanceConfiguration := range listInstanceConfigurationsResponse.Items {
		id := *instanceConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstanceConfigurationId", id)
	}
	return resourceIds, nil
}
