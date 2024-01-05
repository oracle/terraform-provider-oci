// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/suite"
)

var (
	instancePlatformConfigRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `INTEL_VM`},
		"is_measured_boot_enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_secure_boot_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_trusted_platform_module_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	instanceBMMilanPlatformConfigRepresentation = map[string]interface{}{
		"type":                                           acctest.Representation{RepType: acctest.Required, Create: `AMD_MILAN_BM`},
		"are_virtual_instructions_enabled":               acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_access_control_service_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_input_output_memory_management_unit_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_measured_boot_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure_boot_enabled":                         acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_symmetric_multi_threading_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_trusted_platform_module_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"percentage_of_cores_enabled":                    acctest.Representation{RepType: acctest.Required, Create: `50`},
		"config_map":                                     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"numaNodesPerSocket": "NPS4"}},
	}
	instanceBMRomeShieldedPlatformConfigRepresentation = map[string]interface{}{
		"type":                                           acctest.Representation{RepType: acctest.Required, Create: `AMD_ROME_BM`},
		"are_virtual_instructions_enabled":               acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_access_control_service_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_input_output_memory_management_unit_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_measured_boot_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure_boot_enabled":                         acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_symmetric_multi_threading_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_trusted_platform_module_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"numa_nodes_per_socket":                          acctest.Representation{RepType: acctest.Required, Create: `NPS1`},
		"percentage_of_cores_enabled":                    acctest.Representation{RepType: acctest.Required, Create: `25`},
	}
	instanceBMIcelakePlatformConfigRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `INTEL_ICELAKE_BM`},
		"is_input_output_memory_management_unit_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_measured_boot_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure_boot_enabled":                         acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_symmetric_multi_threading_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_trusted_platform_module_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"numa_nodes_per_socket":                          acctest.Representation{RepType: acctest.Required, Create: `NPS1`},
		"percentage_of_cores_enabled":                    acctest.Representation{RepType: acctest.Required, Create: `25`},
	}
	instanceBMSkylakeShieldedPlatformConfigRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `INTEL_SKYLAKE_BM`},
		"is_measured_boot_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure_boot_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_trusted_platform_module_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	instanceVMIntelShieldedPlatformConfigRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `INTEL_VM`},
		"is_measured_boot_enabled":           acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_secure_boot_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_trusted_platform_module_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	instanceVMAmdShieldedPlatformConfigRepresentation = map[string]interface{}{
		"type":                               acctest.Representation{RepType: acctest.Required, Create: `AMD_VM`},
		"is_measured_boot_enabled":           acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_secure_boot_enabled":             acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_trusted_platform_module_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
	// instance representation for E4 Dense
	instanceRepresentationWithNvmes = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                               acctest.Representation{RepType: acctest.Required, Create: `VM.DenseIO.E4.Flex`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.image_id}`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceLaunchOptionsRepresentationWithNvmes},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceShapeConfigRepresentationForNvmeShape},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	instanceLaunchOptionsRepresentationWithNvmes = map[string]interface{}{
		"boot_volume_type":        acctest.Representation{RepType: acctest.Required, Create: `PARAVIRTUALIZED`},
		"firmware":                acctest.Representation{RepType: acctest.Required, Create: `UEFI_64`},
		"network_type":            acctest.Representation{RepType: acctest.Required, Create: `PARAVIRTUALIZED`},
		"remote_data_volume_type": acctest.Representation{RepType: acctest.Required, Create: `PARAVIRTUALIZED`},
	}

	instanceShapeConfigRepresentationForNvmeShape = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `128`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `8`},
		"nvmes":         acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	// instance representation for testing Update to launch_options and fault_domain
	instanceRepresentationCore_ForLaunchOptionsUpdate = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"launch_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentation_ForLaunchOptionsUpdate},
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`, Update: `FAULT-DOMAIN-2`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForLaunchOptionsUpdate},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceShapeConfigRepresentation_ForLaunchOptionsUpdate = map[string]interface{}{
		"ocpus": acctest.Representation{RepType: acctest.Optional, Create: "1", Update: "2"},
	}
	instanceRepresentationCore_ForFlexibleMemory = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`, Update: `FAULT-DOMAIN-2`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E3.Flex`},
		"image":          acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForFlexibleMemory},
		"source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceFlexSourceDetailsRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceShapeConfigRepresentation_ForFlexibleMemory = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: "2"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `10.0`, Update: `20.0`},
	}
	instanceShapeConfigRepresentation_ForFlexibleMemoryNoUpdate = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: "2"},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `10.0`, Update: `20.0`},
	}
	instanceRepresentationCore_ForFlexibleMemoryNoUpdate = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`, Update: `FAULT-DOMAIN-2`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E3.Flex`},
		"image":          acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_config":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation_ForFlexibleMemoryNoUpdate},
		"source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceFlexSourceDetailsRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceLaunchOptionsRepresentation_ForLaunchOptionsUpdate = acctest.RepresentationCopyWithNewProperties(CoreInstanceLaunchOptionsRepresentation, map[string]interface{}{
		"boot_volume_type":                    acctest.Representation{RepType: acctest.Optional, Create: `ISCSI`, Update: `PARAVIRTUALIZED`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"network_type":                        acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`, Update: `VFIO`},
	})
	instanceFlexSourceDetailsRepresentation = map[string]interface{}{
		"source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
		"kms_key_id":  acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}
	instanceWithBMMilanPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `BM.Standard.E4.128`},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceBMMilanPlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceWithBMRomeShieldedPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `BM.Standard.E3.128`},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceBMRomeShieldedPlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceWithBMIcelakePlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `BM.Optimized3.36`},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceBMIcelakePlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceWithBMSkylakeShieldedPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"shape":               acctest.Representation{RepType: acctest.Required, Create: `BM.Standard2.52`},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceBMSkylakeShieldedPlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceWithVMIntelPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceVMIntelShieldedPlatformConfigRepresentation},
	}), []string{
		"dedicated_vm_host_id",
	})
	instanceWithVMAmdPlatformConfigRepresentation = acctest.RepresentationCopyWithRemovedProperties(
		acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
			"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
			"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
			"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceVMAmdShieldedPlatformConfigRepresentation},
		}), []string{
			"dedicated_vm_host_id",
		})

	ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan = utils.DefinedShieldedImageOCIDs +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

type ResourceCoreInstanceTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName string
}

func (s *ResourceCoreInstanceTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "examplevcn"
	}

	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		dns_label = "examplesubnet"
	}

	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
	}
	` + utils.DefinedShieldedImageOCIDs + DefinedTagsDependencies + utils.FlexVmImageIdsVariable

	s.ResourceName = "oci_core_instance.t"
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_basic() {

	var instanceId string
	vnicResourceName := "data.oci_core_vnic.t"

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
					freeform_tags = { "Department" = "Accounting"}
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					timeouts {
						create = "15m"
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					// only set if specified
					resource.TestCheckNoResourceAttr(s.ResourceName, "ipxe_script"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "hostname1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "ZWNobyBoZWxsbw=="),
					resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
					resource.TestCheckResourceAttr(s.ResourceName, "extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "extended_metadata.keyA", "valA"),
					acctest.TestCheckJsonResourceAttr(s.ResourceName, "extended_metadata.keyB", "{\"keyB1\":\"valB1\",\"keyB2\":{\"keyB2\":\"valB2\"}}"),
					acctest.TestCheckJsonResourceAttr(s.ResourceName, "extended_metadata.keyC", "[\"valC1\",\"valC2\"]"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateRunning)),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "preserve_boot_volume"),
					func(ts *terraform.State) (err error) {
						instanceId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// Switching to create_vnic_details for subnet_id and hostname_label should not lead to a change.
			// Changing the letter case in the hostname_label of the instance should also not result in a change.
			// Changing the defined and freeform tags should
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						hostname_label = "hostNAME1"
					}
					image = "${var.InstanceImageOCID[var.region]}"
					hostname_label = "HOSTName1"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					timeouts {
						create = "15m"
					}
				}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// Switching to source_details for the same image ID should not lead to a change.
			// Also, check that source_type is case insensitive.
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "ImAgE"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					timeouts {
						create = "15m"
					}
				}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// verify Update - adds display name, Update tags
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value2"
									)}"
					freeform_tags = { "CostCenter" = "42"}
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("expected same instance ocid, got different")
						}
						return err
					},
				),
			},
			// Adding create_vnic_details with the same subnet_id and an updatable fields should cause an Update only.
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("expected same instance ocid, got different")
						}
						return err
					},
				),
			},
			// Adding create_vnic_details flags with default values should not lead to a change.
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// Update create_vnic_details
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						skip_source_dest_check = true
						hostname_label = "mytftesthostname"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.hostname_label", "mytftesthostname"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("Expected same instance ocid, got different.")
						}
						return err
					},
				),
			},
			// verify force new by setting non-updateable VNIC details and also add tags to the VNIC details
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
							)}"
						freeform_tags = { "Department" = "Accounting" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(vnicResourceName, "display_name", "-tf-vnic-2"),
					resource.TestCheckResourceAttr(vnicResourceName, "skip_source_dest_check", "true"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("expected new instance ocid, got the same")
						}
						instanceId = newId
						return err
					},
				),
			},
			// verify updating vnic tags result in an Update only
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
							)}"
						freeform_tags = { "Department" = "Finance" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId != instanceId {
							return fmt.Errorf("Expected same instance ocid, got different.")
						}
						return err
					},
				),
			},
			// changing order of map elements within JSON strings should not result in diff
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB2\": {\"keyB2\": \"valB2\"}, \"keyB1\": \"valB1\"}" # Order has been changed here, no diff expected
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
							)}"
						freeform_tags = { "Department" = "Finance" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				PlanOnly: true,
			},
		},
	})
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_customdiff() {

	var instanceId string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// Create a new instance with metadata interpolations so that no state exists
			{
				Config: s.Config + `
				locals {
				  nat_offset          = "4"
				}

				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					metadata = {
						should_observe_dependency = "${jsonencode(cidrhost(oci_core_subnet.t.cidr_block, local.nat_offset))}"
						this_should_also = "${oci_core_subnet.t.time_created}"
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						subnet_id = "${oci_core_subnet.t.id}"
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
							)}"
						freeform_tags = { "Department" = "Finance" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "6"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "ZWNobyBoZWxsbw=="),
					func(ts *terraform.State) (err error) {
						instanceId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// verify force new by changing ssh_authorized_keys and user_data in metadata
			{
				Config: s.Config + `
				locals {
				  nat_offset          = "4"
				}

				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					metadata = {
						should_observe_dependency = "${jsonencode(cidrhost(oci_core_subnet.t.cidr_block, local.nat_offset + 1))}"
						this_should_also = "${oci_core_subnet.t.time_created}"
						ssh_authorized_keys = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"
						user_data = "ZWNobyB3b3JsZA=="
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						subnet_id = "${oci_core_subnet.t.id}"
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						display_name = "-tf-vnic-2"
						assign_public_ip = false
						private_ip = "10.0.1.20"
						skip_source_dest_check = true
						defined_tags = "${map(
							"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue"
							)}"
						freeform_tags = { "Department" = "Finance" }
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "10.0.1.20"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.ssh_authorized_keys", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "ZWNobyB3b3JsZA=="),
					func(ts *terraform.State) (err error) {
						newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
						if newId == instanceId {
							return fmt.Errorf("expected new instance ocid, got the same")
						}
						instanceId = newId
						return err
					},
				),
			},
		},
	})
}

// Tests preserve boot volume and attach behavior using source details
func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_preserveBootVolume() {

	var instanceId string
	var preservedBootVolumeId string

	// This is a reference to the TestSteps. We will use this reference to change the TF configs while test steps are
	// being run. This is necessary because some configs require a computed boot volume ID from a previous test step.
	// We cannot set the boot volume id here (it will be nil), so we have to do it within a function closure that gets
	// executed at test step execution time.
	var testStepsReference []resource.TestStep

	testSteps := []resource.TestStep{
		// verify Create of an instance with source_details and that we can get a boot volume id
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
				resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
				// only set if specified
				resource.TestCheckNoResourceAttr(s.ResourceName, "ipxe_script"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
				resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "hostname1"),
				resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "2"),
				resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "ZWNobyBoZWxsbw=="),
				resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateRunning)),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.boot_volume_size_in_gbs"),
				resource.TestCheckNoResourceAttr(s.ResourceName, "preserve_boot_volume"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "boot_volume_id"),
				// Store the instance ID for future verification
				func(ts *terraform.State) (err error) {
					instanceId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
					return err
				},
			),
		},
		// Switching from source_details back to image ID should not lead to a change.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			ExpectNonEmptyPlan: false,
			PlanOnly:           true,
		},
		// verify the preserve_boot_volume setting can be applied and doesn't cause a ForceNew instance
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					preserve_boot_volume = "true"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "true"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "boot_volume_id"),
				// Verify that we didn't get a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
					if newId != instanceId {
						return fmt.Errorf("expected same instance ocid, got different")
					}
					return err
				},
				// Store the boot volume id, so we can use it for attaching to an Instance later
				// Also Update all the config test steps to use the computed boot volume ID
				func(ts *terraform.State) (err error) {
					preservedBootVolumeId, err = acctest.FromInstanceState(ts, s.ResourceName, "boot_volume_id")

					_, tokenFn := acctest.TokenizeWithHttpReplay("instance_resource")
					for idx := range testStepsReference {
						testStepsReference[idx].Config = tokenFn(testStepsReference[idx].Config, map[string]string{"preservedBootVolumeId": preservedBootVolumeId})
					}

					return err
				},
			),
		},
		// ForceNew an instance by changing its hostname_label and boot volume size
		// Verify that the boot volume was preserved and can be attached to the new instance as a data volume.
		// Also verify that the new boot volume size is being used.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
						boot_volume_size_in_gbs = "60"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}

				resource "oci_core_volume_attachment" "volume_attach" {
					attachment_type = "iscsi"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "{{.preservedBootVolumeId}}"
				}
				`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "false"),
				acctest.TestCheckResourceAttributesEqual("oci_core_volume_attachment.volume_attach", "instance_id", s.ResourceName, "id"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
				// Verify that we got a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
					if newId == instanceId {
						return fmt.Errorf("expected different instance ocid, got same")
					}

					instanceId = newId
					return err
				},
				// Verify that the volume attachment's ID is the same as the preserved boot volume
				func(ts *terraform.State) (err error) {
					volumeId, err := acctest.FromInstanceState(ts, "oci_core_volume_attachment.volume_attach", "volume_id")
					if preservedBootVolumeId != volumeId {
						return fmt.Errorf("expected attached volume id to be same as preserved boot volume, got different one")
					}

					return err
				},
			),
		},
		// Detach the boot volume and force a new instance by attaching preserved boot volume in the source details.
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "bootVolume"
						source_id = "{{.preservedBootVolumeId}}"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "false"),
				// Verify that we got a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
					if newId == instanceId {
						return fmt.Errorf("expected different instance ocid, got same")
					}

					instanceId = newId
					return err
				},
				// Verify that the boot volume attachment is the same as the preserved boot volume
				func(ts *terraform.State) (err error) {
					volumeId, err := acctest.FromInstanceState(ts, s.ResourceName, "boot_volume_id")
					if preservedBootVolumeId != volumeId {
						return fmt.Errorf("expected attached boot volume ID to be same as preserved boot volume, got different one")
					}

					return err
				},
			),
		},
		// Verify updating boot_volume_size_in_gbs without recreating the instance
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname2"
					source_details {
						source_type = "bootVolume"
						source_id = "{{.preservedBootVolumeId}}"
						boot_volume_size_in_gbs = "60"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(s.ResourceName, "preserve_boot_volume", "false"),
				resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
				// Verify that we got a new Instance
				func(ts *terraform.State) (err error) {
					newId, err := acctest.FromInstanceState(ts, s.ResourceName, "id")
					if newId != instanceId {
						return fmt.Errorf("expected same instance ocid, got different")
					}

					instanceId = newId
					return err
				},
				// Verify that the boot volume attachment is the same as the preserved boot volume
				func(ts *terraform.State) (err error) {
					volumeId, err := acctest.FromInstanceState(ts, s.ResourceName, "boot_volume_id")
					if preservedBootVolumeId != volumeId {
						return fmt.Errorf("expected attached boot volume ID to be same as preserved boot volume, got different one")
					}

					return err
				},
			),
		},
		// to verify reattaching to the old boot volume resource should be terminated before the waiting for boot volume condition
		{
			Config: s.Config,
		},
		// ForceNew an instance by changing hostname_label and try reattach to the old boot volume,
		// We didn't set preserve flag in the previous step, so the boot volume should be deleted and
		// this should result in an error from service.
		{
			PreConfig: func() {
				acctest.WaitTillCondition(acctest.TestAccProvider, &preservedBootVolumeId, CoreBootVolumeSweepWaitCondition, time.Duration(3*time.Minute),
					CoreBootVolumeSweepResponseFetchOperation, "core", true)
			},
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "bootVolume"
						source_id = "{{.preservedBootVolumeId}}"
					}
					preserve_boot_volume = "false"
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					timeouts {
						create = "15m"
					}
				}`,
			ExpectError: regexp.MustCompile("One or more of the specified volumes are not found"),
		},
	}

	testStepsReference = testSteps
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps:     testSteps,
	})
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_failedByTimeout() {

	testSteps := []resource.TestStep{
		// verify Create of an instance with source_details and that we can get a boot volume id
		{
			Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					source_details {
						source_type = "image"
						source_id = "${var.InstanceImageOCID[var.region]}"
					}
					shape = "VM.Standard2.1"
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg=="
					}
					timeouts {
						create = "15s"
					}
				}`,
			ExpectError: regexp.MustCompile("timeout while waiting for state"),
		},
	}

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps:     testSteps,
	})
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_fetchVnicWhenStopped() {

	resourceName := "oci_core_instance.t"
	config := s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.t.id}"
					hostname_label = "hostname1"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					defined_tags = "${map(
									"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
									)}"
					freeform_tags = { "Department" = "Accounting"}
					metadata = {
						ssh_authorized_keys = "${var.ssh_public_key}"
						user_data = "ZWNobyBoZWxsbw=="
					}
					extended_metadata = {
						keyA = "valA"
						keyB = "{\"keyB1\": \"valB1\", \"keyB2\": {\"keyB2\": \"valB2\"}}"
						keyC = "[\"valC1\", \"valC2\"]"
					}
					timeouts {
						create = "15m"
					}
					state = "STOPPED"
				}`

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify fetching vnic details for an instance that is in stopped state
			{
				Config: config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "image"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					// only set if specified
					resource.TestCheckNoResourceAttr(s.ResourceName, "ipxe_script"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "hostname1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.user_data", "ZWNobyBoZWxsbw=="),
					resource.TestCheckResourceAttrSet(s.ResourceName, "metadata.ssh_authorized_keys"),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "region"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.hostname_label"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "create_vnic_details.0.private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InstanceLifecycleStateStopped)),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "source_details.0.source_id"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "preserve_boot_volume"),
					func(ts *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// verify resource import when instance state is STOPPED
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// TODO: extended_metadata intentionally not set in resource Gets, even though supported
					// by GetInstance calls. Remove this when the issue is resolved.
					"extended_metadata",
					"hostname_label",
					"is_pv_encryption_in_transit_enabled",
					"subnet_id",
					"source_details.0.kms_key_id", //TODO: Service is not returning this value, remove when the service returns it. COM-26394
				},
				ResourceName: resourceName,
			},
		},
	})
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_updateAssignPublicIp() {

	var resId, resId2 string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// Create with assign_public_ip
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// Update assign_public_ip to false
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = false
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "public_ip", ""),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "false"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Update assign_public_ip to true with freeform_tags
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.InstanceImageOCID[var.region]}"
					shape = "VM.Standard2.1"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
						freeform_tags = { "Department" = "Accounting"}
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
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

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_flexVMShape() {

	var resId, resId2 string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// Create with flex shape and shape config
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.FlexInstanceImageOCID[var.region]}"
					shape = "VM.Standard.E3.Flex"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
					shape_config {
						ocpus = "1"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// Update shape config
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.FlexInstanceImageOCID[var.region]}"
					shape = "VM.Standard.E3.Flex"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
					shape_config {
						ocpus = "2"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.0.ocpus", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Update shape_config and displayName
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.FlexInstanceImageOCID[var.region]}"
					shape = "VM.Standard.E3.Flex"
					display_name = "-tf-instance-1"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
					shape_config {
						ocpus = "1"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance-1"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Update displayName
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.FlexInstanceImageOCID[var.region]}"
					shape = "VM.Standard.E3.Flex"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
					shape_config {
						ocpus = "1"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Update shape and shape_config
			{
				Config: s.Config + `
				resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					image = "${var.FlexInstanceImageOCID[var.region]}"
					shape = "VM.Standard2.2"
					display_name = "-tf-instance"
					subnet_id = "${oci_core_subnet.t.id}"
					create_vnic_details {
						subnet_id = "${oci_core_subnet.t.id}"
						skip_source_dest_check = false
						assign_public_ip = true
					}
					shape_config {
						ocpus = "2"
					}
				}
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shape_config.0.ocpus", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
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
func TestAccResourceCoreInstance_BM_Milan_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_BM_Milan_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_BM_Milan_instance_resource")
	}

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

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create with platform config
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create,
						instanceWithBMMilanPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shape", "BM.Standard.E4.128"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "AMD_MILAN_BM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.percentage_of_cores_enabled", "50"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.config_map.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.config_map.numaNodesPerSocket", "NPS4"),

					func(ts *terraform.State) (err error) {
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Required, acctest.Create, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMMilanPlatformConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "BM.Standard.E4.128"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.type", "AMD_MILAN_BM"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.percentage_of_cores_enabled", "50"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.config_map.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.config_map.numaNodesPerSocket", "NPS4"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMMilanPlatformConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "BM.Standard.E4.128"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.numa_nodes_per_socket", "NPS4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.percentage_of_cores_enabled", "50"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.type", "AMD_MILAN_BM"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.config_map.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.config_map.numaNodesPerSocket", "NPS4"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_BM_Rome_shielded_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_BM_Rome_shielded_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_BM_Rome_shielded_instance_resource")
	}

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

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create with platform config
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMRomeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shape", "BM.Standard.E3.128"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "AMD_ROME_BM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.percentage_of_cores_enabled", "25"),

					func(ts *terraform.State) (err error) {
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Required, acctest.Create, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMRomeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "BM.Standard.E3.128"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.type", "AMD_ROME_BM"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.percentage_of_cores_enabled", "25"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMRomeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "BM.Standard.E3.128"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.are_virtual_instructions_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_access_control_service_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.percentage_of_cores_enabled", "25"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.type", "AMD_ROME_BM"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_BM_Icelake_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_BM_Icelake_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_BM_Icelake_instance_resource")
	}

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

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create with platform config
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create,
						instanceWithBMIcelakePlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shape", "BM.Optimized3.36"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "INTEL_ICELAKE_BM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.percentage_of_cores_enabled", "25"),

					func(ts *terraform.State) (err error) {
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Required, acctest.Create, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMIcelakePlatformConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "BM.Optimized3.36"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.type", "INTEL_ICELAKE_BM"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.percentage_of_cores_enabled", "25"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMIcelakePlatformConfigRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "BM.Optimized3.36"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_input_output_memory_management_unit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_symmetric_multi_threading_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.numa_nodes_per_socket", "NPS1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.percentage_of_cores_enabled", "25"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.type", "INTEL_ICELAKE_BM"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_BM_Skylake_shielded_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_BM_Skylake_shielded_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_BM_Skylake_shielded_instance_resource")
	}

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

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create with platform config
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMSkylakeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shape", "BM.Standard2.52"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "INTEL_SKYLAKE_BM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),

					func(ts *terraform.State) (err error) {
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Required, acctest.Create, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMSkylakeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "BM.Standard2.52"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.type", "INTEL_SKYLAKE_BM"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_trusted_platform_module_enabled", "true"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithBMSkylakeShieldedPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "BM.Standard2.52"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_measured_boot_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.type", "INTEL_SKYLAKE_BM"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_VM_Intel_shielded_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_VM_Intel_shielded_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_VM_Intel_shielded_instance_resource")
	}

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
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create with platform config
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithVMIntelPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.type", "INTEL_VM"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_measured_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),

					func(ts *terraform.State) (err error) {
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Required, acctest.Create, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithVMIntelPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.type", "INTEL_VM"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_measured_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.platform_config.0.is_trusted_platform_module_enabled", "true"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreCoreInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithVMIntelPlatformConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.memory_in_gbs", "15"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_measured_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_secure_boot_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.is_trusted_platform_module_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_config.0.type", "INTEL_VM"),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_VM_Amd_shielded_instance_resource(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestAccResourceCoreInstance_VM_Amd_shielded_instance_resource") {
		t.Skip("Skipping suppressed TestAccResourceCoreInstance_VM_Amd_shielded_instance_resource")
	}

	provider := acctest.TestAccProvider

	config := `
        provider oci {
            test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
        }
    ` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config + compartmentIdVariableStr + ShieldedInstanceResourceDependenciesWithoutDVHWithoutVlan +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceWithVMAmdPlatformConfigRepresentation),
				ExpectError: regexp.MustCompile("VM.Standard2.1 does not support the provided platform configuration"),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestResourceCoreInstanceTestSuite(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip TestResourceCoreInstanceTestSuite in HttpReplay mode.")
	}
	suite.Run(t, new(ResourceCoreInstanceTestSuite))
}

func (s *ResourceCoreInstanceTestSuite) TestAccResourceCoreInstance_launchOptions() {
	httpreplay.SetScenario("TestAccResourceCoreInstance_launchOptions")
	defer httpreplay.SaveScenario()

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

	var resId, resId2 string

	resource.Test(s.T(), resource.TestCase{
		Providers:    s.Providers,
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentationCore_ForLaunchOptionsUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
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
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForLaunchOptionsUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
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
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "2"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
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
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForLaunchOptionsUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.launch_mode"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
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
					compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForLaunchOptionsUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "launch_mode"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.boot_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_consistent_volume_naming_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.ocpus", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
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
func TestAccResourceCoreInstance_nvmeVMShape(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreInstance_nvmeVMShape")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := `
      provider oci {
         test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
      }
   ` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	imageId := utils.GetEnvSettingWithBlankDefault("image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	resourceName := "oci_core_instance.test_instance"

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Create E4 Dense shape and shape config
			{
				Config: config + compartmentIdVariableStr + imageIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentationWithNvmes),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.DenseIO.E4.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "8"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.local_disks", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "128"),
					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestAccResourceCoreInstance_FlexibleMemory(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreInstance_FlexibleMemory")
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

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentationCore_ForFlexibleMemory),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
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
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters but no change in shape_config
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForFlexibleMemoryNoUpdate),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
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
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "20"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForFlexibleMemory),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
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
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "20"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
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

			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, CoreCoreInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForFlexibleMemory),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.memory_in_gbs", "20"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
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
					compartmentIdVariableStr + CoreInstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
					acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationCore_ForFlexibleMemory),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_management_disabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.gpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.local_disks_total_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.max_vnic_attachments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape_config.0.memory_in_gbs", "20"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.networking_bandwidth_in_gbps"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
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
