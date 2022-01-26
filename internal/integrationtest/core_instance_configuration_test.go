// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	InstanceConfigurationRequiredOnlyResource = InstanceConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, instanceConfigurationRepresentation)

	InstanceConfigurationResourceConfig = InstanceConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, instanceConfigurationRepresentation)

	instanceConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
	}

	instanceConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceConfigurationDataSourceFilterRepresentation}}
	instanceConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_configuration.test_instance_configuration.id}`}},
	}

	instanceConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsRepresentation},
		"source":           acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}
	instanceConfigurationFromInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"source":         acctest.Representation{RepType: acctest.Optional, Create: `INSTANCE`},
	}
	instanceConfigurationInstanceDetailsLaunchRepresentation = map[string]interface{}{
		"instance_type":  acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"launch_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
	}

	instanceConfigurationInstanceDetailsLaunchRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("launch_details",
		acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsRepresentationForFlexShape},
		instanceConfigurationInstanceDetailsLaunchRepresentation)

	instanceConfigurationInstanceDetailsBlockRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"block_volumes": acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	instanceConfigurationInstanceDetailsParavirtualizedBlockRepresentation = map[string]interface{}{
		"instance_type": acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"block_volumes": acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	instanceConfigurationInstanceDetailsRepresentation = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation},
	}
	instanceConfigurationInstanceDetailsBlockVolumesRepresentation = map[string]interface{}{
		"create_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation = map[string]interface{}{
		"attach_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	instanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachRepresentation = map[string]interface{}{
		"attach_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachDetailsRepresentation},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	instanceShapeConfigRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: "1"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: "15"},
	}
	instanceConfigurationInstanceLaunchOptionsRepresentation = map[string]interface{}{
		"network_type": acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsRepresentation = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"extended_metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"extendedMetadata": "extendedMetadata"}, Update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShape},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigRepresentation},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceLaunchOptionsRepresentation},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceOptionsRepresentation},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"dedicated_vm_host_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"launch_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `NATIVE`},
		"preferred_maintenance_action":        acctest.Representation{RepType: acctest.Optional, Create: `LIVE_MIGRATE`},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsRepresentationForFlexShape = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"shape", "source_details", "shape_config"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShapeForFlex},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentationForFlexShape},
			},
			instanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"dedicated_vm_host_id", "preferred_maintenance_action"},
	)
	instanceConfigurationInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `iscsi`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"is_read_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"use_chap":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	instanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachDetailsRepresentation = map[string]interface{}{
		"type":                                acctest.Representation{RepType: acctest.Required, Create: `paravirtualized`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"device":                              acctest.Representation{RepType: acctest.Optional, Create: `server`},
		"is_read_only":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_shareable":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"backup_policy_id":    acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"size_in_gbs":         acctest.Representation{RepType: acctest.Optional, Create: `50`},
		"source_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation},
		"vpus_per_gb":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"kms_key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_ip":                acctest.Representation{RepType: acctest.Optional, Create: `privateIp`},
		"skip_source_dest_check":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation = map[string]interface{}{
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: `image`},
		"image_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.InstanceImageOCID[var.region]}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `55`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("image_id",
		acctest.Representation{RepType: acctest.Optional, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)

	instanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation = map[string]interface{}{
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
	instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `volume`},
		"id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}

	InstanceConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", acctest.Required, acctest.Create, bootVolumeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		utils.VolumeBackupPolicyDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
	InstanceConfigurationVmShape        = `VM.Standard2.1`
	InstanceConfigurationVmShapeForFlex = `VM.Standard.E3.Flex`

	InstanceConfigurationResourceImageConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation))
)

// issue-routing-tag: core/computeManagement
func TestCoreInstanceConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"
	datasourceName := "data.oci_core_instance_configurations.test_instance_configurations"
	singularDatasourceName := "data.oci_core_instance_configuration.test_instance_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InstanceConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, instanceConfigurationRepresentation), "core", "instanceConfiguration", t)
	acctest.ResourceTest(t, testAccCheckCoreInstanceConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, instanceConfigurationRepresentation),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
		},

		// verify Create from instance_id
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, instanceConfigurationFromInstanceRepresentation),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
		},

		// verify Create with optionals launch_details for E3 flex micro shape
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies + utils.FlexVmImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchRepresentationForFlexShape}, instanceConfigurationRepresentation)),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
		},
		// verify Create with optionals launch_details
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation)),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsBlockRepresentation}, instanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.GetUpdatedRepresentationCopy("block_volumes", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation}, instanceConfigurationInstanceDetailsBlockRepresentation)}, instanceConfigurationRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.GetUpdatedRepresentationCopy("block_volumes", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsParavirtualizedBlockVolumeAttachRepresentation},
						instanceConfigurationInstanceDetailsParavirtualizedBlockRepresentation)},
						instanceConfigurationRepresentation)),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
		},
		// verify Create with optionals secondary_vnics
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, instanceConfigurationRepresentation),
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
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, instanceConfigurationRepresentation),
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
				compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", acctest.Optional, acctest.Update, instanceConfigurationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Update, instanceConfigurationRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, instanceConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceConfigurationResourceConfig,
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + InstanceConfigurationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"instance_id",
				"source",
			},
			ResourceName: resourceName,
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
	instanceConfigurationIds, err := getInstanceConfigurationIds(compartment)
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

func getInstanceConfigurationIds(compartment string) ([]string, error) {
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
