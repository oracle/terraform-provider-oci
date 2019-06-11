// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InstanceConfigurationRequiredOnlyResource = InstanceConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Required, Create, instanceConfigurationRepresentation)

	InstanceConfigurationResourceConfig = InstanceConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Update, instanceConfigurationRepresentation)

	instanceConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"instance_configuration_id": Representation{repType: Required, create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
	}

	instanceConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, instanceConfigurationDataSourceFilterRepresentation}}
	instanceConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance_configuration.test_instance_configuration.id}`}},
	}

	instanceConfigurationRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_details": RepresentationGroup{Required, instanceConfigurationInstanceDetailsRepresentation},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{repType: Optional, create: `backend-servers`, update: `displayName2`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	instanceConfigurationInstanceDetailsLaunchRepresentation = map[string]interface{}{
		"instance_type":  Representation{repType: Required, create: `compute`},
		"launch_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsRepresentation},
	}
	instanceConfigurationInstanceDetailsBlockRepresentation = map[string]interface{}{
		"instance_type": Representation{repType: Required, create: `compute`},
		"block_volumes": RepresentationGroup{Required, instanceConfigurationInstanceDetailsBlockVolumesRepresentation},
	}
	instanceConfigurationInstanceDetailsRepresentation = map[string]interface{}{
		"instance_type":   Representation{repType: Required, create: `compute`},
		"secondary_vnics": RepresentationGroup{Required, instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation},
	}
	instanceConfigurationInstanceDetailsBlockVolumesRepresentation = map[string]interface{}{
		"create_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation},
		"volume_id":      Representation{repType: Optional, create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation = map[string]interface{}{
		"attach_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation},
		"volume_id":      Representation{repType: Optional, create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Optional, create: `${var.compartment_id}`},
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `backend-servers`},
		"extended_metadata":   Representation{repType: Optional, create: map[string]string{"extendedMetadata": "extendedMetadata"}, update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"fault_domain":        Representation{repType: Optional, create: `FAULT-DOMAIN-2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":         Representation{repType: Optional, create: `ipxeScript`},
		"metadata":            Representation{repType: Optional, create: map[string]string{"metadata": "metadata"}, update: map[string]string{"metadata2": "metadata2"}},
		"shape":               Representation{repType: Optional, create: InstanceConfigurationVmShape},
		"source_details":      RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation = map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation},
		"display_name":        Representation{repType: Optional, create: `backend-servers`},
		"nic_index":           Representation{repType: Optional, create: `0`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesAttachDetailsRepresentation = map[string]interface{}{
		"type":         Representation{repType: Required, create: `iscsi`},
		"display_name": Representation{repType: Optional, create: `backend-servers`},
		"is_read_only": Representation{repType: Optional, create: `false`},
		"use_chap":     Representation{repType: Optional, create: `false`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"backup_policy_id":    Representation{repType: Optional, create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"compartment_id":      Representation{repType: Optional, create: `${var.compartment_id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `backend-servers`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"size_in_gbs":         Representation{repType: Optional, create: `10`},
		"source_details":      RepresentationGroup{Optional, instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip":       Representation{repType: Optional, create: `false`},
		"display_name":           Representation{repType: Optional, create: `backend-servers`},
		"hostname_label":         Representation{repType: Optional, create: `hostnameLabel`},
		"private_ip":             Representation{repType: Optional, create: `privateIp`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
		"subnet_id":              Representation{repType: Optional, create: `${oci_core_subnet.test_subnet.id}`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation = map[string]interface{}{
		"source_type":             Representation{repType: Required, create: `image`},
		"image_id":                Representation{repType: Optional, create: `${var.InstanceImageOCID[var.region]}`},
		"boot_volume_size_in_gbs": Representation{repType: Optional, create: `55`},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip":       Representation{repType: Optional, create: `false`},
		"display_name":           Representation{repType: Optional, create: `backend-servers`},
		"hostname_label":         Representation{repType: Optional, create: `hostnameLabel`},
		"private_ip":             Representation{repType: Optional, create: `privateIp`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
		"subnet_id":              Representation{repType: Optional, create: `${oci_core_subnet.test_subnet.id}`},
	}
	instanceConfigurationInstanceDetailsBlockVolumesCreateDetailsSourceDetailsRepresentation = map[string]interface{}{
		"type": Representation{repType: Required, create: `volume`},
		"id":   Representation{repType: Optional, create: `${oci_core_boot_volume.test_boot_volume.id}`},
	}

	InstanceConfigurationResourceDependencies = InstanceRequiredOnlyResource
	InstanceConfigurationVmShape              = `VM.Standard2.1`

	InstanceConfigurationResourceImageConfig = generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
		getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation))
)

func TestCoreInstanceConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_configuration.test_instance_configuration"
	datasourceName := "data.oci_core_instance_configurations.test_instance_configurations"
	singularDatasourceName := "data.oci_core_instance_configuration.test_instance_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceConfigurationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Required, Create, instanceConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.instance_type", "compute"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
			},
			// verify create with optionals launch_details
			{
				Config: config + compartmentIdVariableStr + ImageRequiredOnlyResource +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
						getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchRepresentation}, instanceConfigurationRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.extended_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.shape", InstanceConfigurationVmShape),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.boot_volume_size_in_gbs", "55"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.launch_details.0.source_details.0.image_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.launch_details.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify recreate with optionals block_volumes.create_details
			{
				Config: config + compartmentIdVariableStr + BootVolumeRequiredOnlyResource +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
						getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationInstanceDetailsBlockRepresentation}, instanceConfigurationRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.size_in_gbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.0.id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.block_volumes.0.create_details.0.source_details.0.type", "volume"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.block_volumes.0.volume_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify recreate with optionals block_volumes.create_details to block_volumes.attach_details
			{
				Config: config + compartmentIdVariableStr + BootVolumeRequiredOnlyResource +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
						getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional,
							getUpdatedRepresentationCopy("block_volumes", RepresentationGroup{Optional, instanceConfigurationInstanceDetailsBlockVolumesAttachRepresentation}, instanceConfigurationInstanceDetailsBlockRepresentation)}, instanceConfigurationRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("resource was not recreated")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies,
			},
			// verify create with optionals secondary_vnics
			{
				Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create, instanceConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + InstanceConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Update, instanceConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(resourceName, "instance_details.0.secondary_vnics.0.nic_index", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_instance_configurations", "test_instance_configurations", Optional, Update, instanceConfigurationDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Update, instanceConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "instance_configurations.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instance_configurations.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_configurations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_configurations.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Required, Create, instanceConfigurationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "deferred_fields.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.display_name", "backend-servers"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_details.0.secondary_vnics.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreInstanceConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeManagementClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_configuration" {
			noResourceFound = false
			request := oci_core.GetInstanceConfigurationRequest{}

			tmp := rs.Primary.ID
			request.InstanceConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreInstanceConfiguration") {
		resource.AddTestSweepers("CoreInstanceConfiguration", &resource.Sweeper{
			Name:         "CoreInstanceConfiguration",
			Dependencies: DependencyGraph["instanceConfiguration"],
			F:            sweepCoreInstanceConfigurationResource,
		})
	}
}

func sweepCoreInstanceConfigurationResource(compartment string) error {
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient
	instanceConfigurationIds, err := getInstanceConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceConfigurationId := range instanceConfigurationIds {
		if ok := SweeperDefaultResourceId[instanceConfigurationId]; !ok {
			deleteInstanceConfigurationRequest := oci_core.DeleteInstanceConfigurationRequest{}

			deleteInstanceConfigurationRequest.InstanceConfigurationId = &instanceConfigurationId

			deleteInstanceConfigurationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
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
	ids := getResourceIdsToSweep(compartment, "InstanceConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient

	listInstanceConfigurationsRequest := oci_core.ListInstanceConfigurationsRequest{}
	listInstanceConfigurationsRequest.CompartmentId = &compartmentId
	listInstanceConfigurationsResponse, err := computeManagementClient.ListInstanceConfigurations(context.Background(), listInstanceConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InstanceConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instanceConfiguration := range listInstanceConfigurationsResponse.Items {
		id := *instanceConfiguration.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "InstanceConfigurationId", id)
	}
	return resourceIds, nil
}
