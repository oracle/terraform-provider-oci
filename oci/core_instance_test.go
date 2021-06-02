// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v41/common"
	oci_core "github.com/oracle/oci-go-sdk/v41/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InstanceRequiredOnlyResource = SubnetResourceConfig + OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation)

	InstanceResourceConfig = InstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentation)

	instanceSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
	}
	instanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":               Representation{repType: Optional, create: `RUNNING`},
		"filter":              RepresentationGroup{Required, instanceDataSourceFilterRepresentation}}
	instanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance.test_instance.id}`}},
	}
	instanceRepresentation = map[string]interface{}{
		"availability_domain":  Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                Representation{repType: Required, create: `VM.Standard2.1`},
		"agent_config":         RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"availability_config":  RepresentationGroup{Optional, instanceAvailabilityConfigRepresentation},
		"create_vnic_details":  RepresentationGroup{Optional, instanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id": Representation{repType: Optional, create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"extended_metadata": Representation{repType: Optional, create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        Representation{repType: Optional, create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      Representation{repType: Optional, create: `hostnamelabel`},
		"instance_options":                    RepresentationGroup{Optional, instanceInstanceOptionsRepresentation},
		"image":                               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         Representation{repType: Optional, create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"metadata":                            Representation{repType: Optional, create: map[string]string{"user_data": "abcd"}, update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
		"source_details":                      RepresentationGroup{Optional, instanceSourceDetailsRepresentation},
		"subnet_id":                           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               Representation{repType: Optional, create: `STOPPED`, update: `RUNNING`},
	}
	instanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": Representation{repType: Optional, create: `false`, update: `false`},
		"is_management_disabled":   Representation{repType: Optional, create: `false`, update: `false`},
		"is_monitoring_disabled":   Representation{repType: Optional, create: `false`, update: `false`},
		"plugins_config":           RepresentationGroup{Optional, instanceAgentConfigPluginsConfigRepresentation},
	}
	instanceAvailabilityConfigRepresentation = map[string]interface{}{
		"is_live_migration_preferred": Representation{repType: Optional, create: `false`, update: `true`},
		"recovery_action":             Representation{repType: Optional, create: `RESTORE_INSTANCE`, update: `STOP_INSTANCE`},
	}
	instanceCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": Representation{repType: Optional, create: `true`},
		"assign_public_ip":          Representation{repType: Optional, create: `true`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `displayName`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Accounting"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":            Representation{repType: Optional, create: `hostnamelabel`},
		"nsg_ids":                   Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{}},
		"private_ip":                Representation{repType: Optional, create: `10.0.0.5`},
		"skip_source_dest_check":    Representation{repType: Optional, create: `false`},
		"subnet_id":                 Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
	}
	instanceInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": Representation{repType: Optional, create: `false`, update: `true`},
	}
	instanceLaunchOptionsRepresentation = map[string]interface{}{
		"boot_volume_type":                    Representation{repType: Optional, create: `ISCSI`},
		"firmware":                            Representation{repType: Optional, create: `UEFI_64`},
		"is_consistent_volume_naming_enabled": Representation{repType: Optional, create: `true`},
		"network_type":                        Representation{repType: Optional, create: `PARAVIRTUALIZED`},
		"remote_data_volume_type":             Representation{repType: Optional, create: `PARAVIRTUALIZED`},
	}
	instanceSubCorePlatformConfigRepresentation = map[string]interface{}{
		"type":                  Representation{repType: Required, create: `AMD_MILAN_BM`},
		"numa_nodes_per_socket": Representation{repType: Optional, create: `NPS0`},
	}
	instanceSourceDetailsRepresentation = map[string]interface{}{
		"source_id":               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"source_type":             Representation{repType: Required, create: `image`},
		"kms_key_id":              Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"boot_volume_size_in_gbs": Representation{repType: Optional, create: `60`, update: `70`},
	}
	instanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": Representation{repType: Required, create: `ENABLED`},
		"name":          Representation{repType: Required, create: `Compute Instance Monitoring`},
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
	InstanceResourceDependenciesWithoutDHV = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Required, Create,
			getUpdatedRepresentationCopy("cidr_block", Representation{repType: Required, create: `10.0.1.0/30`}, vlanRepresentation)) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig

	InstanceResourceDependencies = generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Update, dedicatedVmHostRepresentation) +
		InstanceResourceDependenciesWithoutDHV

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
	instanceDataSourceRepresentationForFlexShape = getUpdatedRepresentationCopy("availability_domain",
		Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`}, instanceDataSourceRepresentation)
	InstanceResourceConfigForFlexShape = InstanceResourceDependenciesWithoutDHV +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentationForFlexShape)
	instanceSourceDetailsRepresentationForFlexShape = getMultipleUpdatedRepresenationCopy(
		[]string{"source_id", "boot_volume_size_in_gbs"},
		[]interface{}{
			Representation{repType: Required, create: `${var.FlexInstanceImageOCID[var.region]}`},
			Representation{repType: Optional, create: `60`},
		},
		instanceSourceDetailsRepresentation)

	instanceShapeConfigRepresentationForFlexShape = map[string]interface{}{
		"baseline_ocpu_utilization": Representation{repType: Required, create: `BASELINE_1_8`, update: `BASELINE_1_2`},
		"memory_in_gbs":             Representation{repType: Required, create: `1.0`, update: `4.0`},
		"ocpus":                     Representation{repType: Required, create: `1`},
	}
	instanceLaunchOptionsRepresentationForFlexShape = getUpdatedRepresentationCopy("boot_volume_type",
		Representation{repType: Optional, create: `PARAVIRTUALIZED`}, instanceLaunchOptionsRepresentation)

	instanceRepresentationForFlexShape = representationCopyWithRemovedProperties(
		getMultipleUpdatedRepresenationCopy(
			[]string{"availability_domain", "shape", "image", "create_vnic_details", "launch_options", "source_details", "shape_config"},
			[]interface{}{
				Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
				Representation{repType: Required, create: InstanceConfigurationVmShapeForFlex},
				Representation{repType: Required, create: `${var.FlexInstanceImageOCID[var.region]}`},
				RepresentationGroup{Optional, instanceCreateVnicDetailsRepresentation},
				RepresentationGroup{Optional, instanceLaunchOptionsRepresentationForFlexShape},
				RepresentationGroup{Optional, instanceSourceDetailsRepresentationForFlexShape},
				RepresentationGroup{Required, instanceShapeConfigRepresentationForFlexShape},
			},
			instanceRepresentation),
		[]string{"dedicated_vm_host_id"},
	)

	// ------------- for capacity reservation -------------
	instanceSourceDetailsSansKmsRepresentation = map[string]interface{}{
		"source_id":               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"source_type":             Representation{repType: Required, create: `image`},
		"boot_volume_size_in_gbs": Representation{repType: Optional, create: `60`, update: `70`},
	}

	instanceWithCapacityReservationRepresentation = map[string]interface{}{
		// dedicated_vm_host_id is incompatible with capacity_reservation_id
		"availability_domain":     Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                   Representation{repType: Required, create: `VM.Standard2.1`},
		"agent_config":            RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"availability_config":     RepresentationGroup{Optional, instanceAvailabilityConfigRepresentation},
		"capacity_reservation_id": Representation{repType: Optional, create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"create_vnic_details":     RepresentationGroup{Optional, instanceCreateVnicDetailsRepresentation},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"extended_metadata": Representation{repType: Optional, create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}, update: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
			"other_string":  "stringD",
		}},
		"fault_domain":                        Representation{repType: Optional, create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      Representation{repType: Optional, create: `hostnamelabel`},
		"instance_options":                    RepresentationGroup{Optional, instanceInstanceOptionsRepresentation},
		"image":                               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         Representation{repType: Optional, create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"metadata":                            Representation{repType: Optional, create: map[string]string{"user_data": "abcd"}, update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
		"source_details":                      RepresentationGroup{Optional, instanceSourceDetailsSansKmsRepresentation},
		"subnet_id":                           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               Representation{repType: Optional, create: `STOPPED`, update: `RUNNING`},
	}

	instanceWithCapacityReservationResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	instanceWithCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain":     Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"capacity_reservation_id": Representation{repType: Optional, create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":                   Representation{repType: Optional, create: `RUNNING`},
		"filter":                  RepresentationGroup{Required, instanceDataSourceFilterRepresentation}}
)

func TestCoreInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + commonTestVariables()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+InstanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceRepresentation), "core", "instance", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to shape within the same family is not force new. Resizing can only be done to intances not using dedicated_vm_host_id
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, getUpdatedRepresentationCopy("shape", Representation{repType: Required, create: `VM.Standard2.2`}, instanceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies,
			},
			// verify create with is_pv_encryption_in_transit_enabled = true
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies + InstanceWithPVEncryptionInTransitEnabled,
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create,
						representationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", Optional, Update, instanceDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
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
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
			// verify updates to original parameters
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create,
						getUpdatedRepresentationCopy("source_details", RepresentationGroup{Optional,
							representationCopyWithRemovedProperties(instanceSourceDetailsRepresentation, []string{"boot_volume_size_in_gbs"})},
							instanceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "agent_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + InstanceResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// extended_metadata is set by import but service may potentially reorder map elements in imported JSON strings.
					// This is normally handled by diff suppress function but the Terraform import tests can't invoke diff suppression
					// and so it may complain that values are different.
					"extended_metadata",
					"hostname_label",
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

					expectedExtendedMetadataMap := instanceRepresentation["extended_metadata"].(Representation).update.(map[string]string)

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
						if err := checkJsonStringsEqual(expectedJsonString, actualJsonString); err != nil {
							return fmt.Errorf("%s: Attribute '%s' %s", resourceName, attributeKey, err)
						}
					}
					return nil
				},
				ResourceName: resourceName,
			},
		},
	})
}

func TestCoreInstanceResource_capacityReservation(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_capacityReservation")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + commonTestVariables()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 0: verify create with optionals
			{
				Config: config +
					compartmentIdVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceWithCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// Step 1: verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config +
					compartmentIdVariableStr +
					compartmentIdUVariableStr +
					instanceWithCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create,
						representationCopyWithNewProperties(instanceWithCapacityReservationRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceWithCapacityReservationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceWithCapacityReservationRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", Optional, Update, instanceWithCapacityReservationDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
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
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceWithCapacityReservationRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceSingularDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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

func TestCoreInstanceResource_flexShape(t *testing.T) {
	httpreplay.SetScenario("TestCoreFlexInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"
	singularDatasourceName := "data.oci_core_instance.test_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// step 0 verify create
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentationForFlexShape),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.memory_in_gbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					// currently E3 subcore is forced to use launch_mode = PARAVIRTUALIZED
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "PARAVIRTUALIZED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// step 1 delete before next create
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable,
			},

			// step 2 verify create with is_pv_encryption_in_transit_enabled = true
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable + InstanceWithPVEncryptionInTransitEnabledForFlexShape,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// step 3 delete before next create
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable,
			},

			// step 4 verify create with optionals
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceRepresentationForFlexShape),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			//step 5: verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: testProviderConfig() + compartmentIdVariableStr + compartmentIdUVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create,
						representationCopyWithNewProperties(instanceRepresentationForFlexShape, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// step 6: verify updates to updatable parameters
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentationForFlexShape),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
					// resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.boot_volume_size_in_gbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

			// step 7: verify datasource
			{
				Config: testProviderConfig() +
					generateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", Optional, Update, instanceDataSourceRepresentationForFlexShape) +
					compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + FlexVmImageIdsVariable +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Update, instanceRepresentationForFlexShape),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
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
				Config: testProviderConfig() + FlexVmImageIdsVariable +
					generateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceResourceConfigForFlexShape,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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

func testAccCheckCoreInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance" {
			noResourceFound = false
			request := oci_core.GetInstanceRequest{}

			tmp := rs.Primary.ID
			request.InstanceId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreInstance") {
		resource.AddTestSweepers("CoreInstance", &resource.Sweeper{
			Name:         "CoreInstance",
			Dependencies: DependencyGraph["instance"],
			F:            sweepCoreInstanceResource,
		})
	}
}

func sweepCoreInstanceResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	instanceIds, err := getInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceId := range instanceIds {
		if ok := SweeperDefaultResourceId[instanceId]; !ok {
			terminateInstanceRequest := oci_core.TerminateInstanceRequest{}

			terminateInstanceRequest.InstanceId = &instanceId

			terminateInstanceRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeClient.TerminateInstance(context.Background(), terminateInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting Instance %s %s, It is possible that the resource is already deleted. Please verify manually \n", instanceId, error)
				continue
			}
			waitTillCondition(testAccProvider, &instanceId, instanceSweepWaitCondition, time.Duration(3*time.Minute),
				instanceSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstanceIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "InstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "InstanceId", id)
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

func instanceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeClient().GetInstance(context.Background(), oci_core.GetInstanceRequest{
		InstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
