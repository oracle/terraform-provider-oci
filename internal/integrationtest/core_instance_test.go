// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

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
	InstanceRequiredOnlyResource = SubnetResourceConfig + utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation)

	InstanceResourceConfig = InstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentation)

	instanceSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}
	instanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceDataSourceFilterRepresentation}}
	instanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance.test_instance.id}`}},
	}
	instanceRepresentation = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigRepresentation},
		"availability_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAvailabilityConfigRepresentation},
		"create_vnic_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceCreateVnicDetailsRepresentation},
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
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	instanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"plugins_config":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigPluginsConfigRepresentation},
	}
	instanceAvailabilityConfigRepresentation = map[string]interface{}{
		"is_live_migration_preferred": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"recovery_action":             acctest.Representation{RepType: acctest.Optional, Create: `RESTORE_INSTANCE`, Update: `STOP_INSTANCE`},
	}
	instanceCreateVnicDetailsRepresentation = map[string]interface{}{
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
	instanceInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	instanceLaunchOptionsRepresentation = map[string]interface{}{
		"boot_volume_type":                    acctest.Representation{RepType: acctest.Optional, Create: `ISCSI`},
		"firmware":                            acctest.Representation{RepType: acctest.Optional, Create: `UEFI_64`},
		"is_consistent_volume_naming_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"network_type":                        acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
		"remote_data_volume_type":             acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`},
	}
	instanceSubCorePlatformConfigRepresentation = map[string]interface{}{
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `AMD_MILAN_BM`},
		"numa_nodes_per_socket": acctest.Representation{RepType: acctest.Optional, Create: `NPS0`},
	}
	instanceSourceDetailsRepresentation = map[string]interface{}{
		"source_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: `image`},
		"kms_key_id":              acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `70`},
	}
	instanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `Compute Instance Monitoring`},
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
	InstanceResourceDependenciesWithoutDHV = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("cidr_block", acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/30`}, vlanRepresentation)) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig

	InstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation) +
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
	instanceDataSourceRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("availability_domain",
		acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`}, instanceDataSourceRepresentation)
	InstanceResourceConfigForFlexShape = InstanceResourceDependenciesWithoutDHV +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentationForFlexShape)
	instanceSourceDetailsRepresentationForFlexShape = acctest.GetMultipleUpdatedRepresenationCopy(
		[]string{"source_id", "boot_volume_size_in_gbs"},
		[]interface{}{
			acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
			acctest.Representation{RepType: acctest.Optional, Create: `60`},
		},
		instanceSourceDetailsRepresentation)

	instanceShapeConfigRepresentationForFlexShape = map[string]interface{}{
		"baseline_ocpu_utilization": acctest.Representation{RepType: acctest.Required, Create: `BASELINE_1_8`, Update: `BASELINE_1_2`},
		"memory_in_gbs":             acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `4.0`},
		"ocpus":                     acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	instanceLaunchOptionsRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("boot_volume_type",
		acctest.Representation{RepType: acctest.Optional, Create: `PARAVIRTUALIZED`}, instanceLaunchOptionsRepresentation)

	instanceRepresentationForFlexShape = acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{"availability_domain", "shape", "image", "create_vnic_details", "launch_options", "source_details", "shape_config"},
			[]interface{}{
				acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
				acctest.Representation{RepType: acctest.Required, Create: InstanceConfigurationVmShapeForFlex},
				acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceCreateVnicDetailsRepresentation},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsRepresentationForFlexShape},
				acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceShapeConfigRepresentationForFlexShape},
			},
			instanceRepresentation),
		[]string{"dedicated_vm_host_id"},
	)

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
		"agent_config":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigRepresentation},
		"availability_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAvailabilityConfigRepresentation},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"create_vnic_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceCreateVnicDetailsRepresentation},
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
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsSansKmsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}

	instanceWithCapacityReservationResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, computeCapacityReservationRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	instanceWithCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceDataSourceFilterRepresentation}}
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

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentation), "core", "instance", t)

	acctest.ResourceTest(t, testAccCheckCoreInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation),
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
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("shape", acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`}, instanceRepresentation)),
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
			Config: config + compartmentIdVariableStr + InstanceResourceDependencies,
		},
		// verify Create with is_pv_encryption_in_transit_enabled = true
		{
			Config: config + compartmentIdVariableStr + InstanceResourceDependencies + InstanceWithPVEncryptionInTransitEnabled,
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
			Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceRepresentation),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
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
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnamelabel"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentation),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", acctest.Optional, acctest.Update, instanceDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Update, instanceRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceResourceConfig,
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
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithRemovedProperties(instanceSourceDetailsRepresentation, []string{"boot_volume_size_in_gbs"})},
						instanceRepresentation)),
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

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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

				expectedExtendedMetadataMap := instanceRepresentation["extended_metadata"].(acctest.Representation).Update.(map[string]string)

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
		Providers: map[string]terraform.ResourceProvider{
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceSingularDataSourceRepresentation),
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

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// step 0 verify Create
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
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
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable,
			},

			// step 2 verify Create with is_pv_encryption_in_transit_enabled = true
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable + InstanceWithPVEncryptionInTransitEnabledForFlexShape,
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
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable,
			},

			// step 4 verify Create with optionals
			{
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
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
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + compartmentIdUVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
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
				Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
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
					compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV + utils.FlexVmImageIdsVariable +
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceResourceConfigForFlexShape,
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
