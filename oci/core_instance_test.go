// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InstanceRequiredOnlyResource = SubnetResourceConfig + InstanceCommonVariables +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group1", Required, Create, networkSecurityGroupRepresentation) +
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
		"image":                               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         Representation{repType: Optional, create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"metadata":                            Representation{repType: Optional, create: map[string]string{"user_data": "abcd"}, update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"source_details":                      RepresentationGroup{Optional, instanceSourceDetailsRepresentation},
		"subnet_id":                           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               Representation{repType: Optional, create: `STOPPED`, update: `RUNNING`},
	}
	instanceAgentConfigRepresentation = map[string]interface{}{
		"is_monitoring_disabled": Representation{repType: Optional, create: `false`, update: `true`},
	}
	instanceCreateVnicDetailsRepresentation = map[string]interface{}{
		"subnet_id":              Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"assign_public_ip":       Representation{repType: Optional, create: `true`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{repType: Optional, create: `displayName`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Accounting"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":         Representation{repType: Optional, create: `hostnamelabel`},
		"nsg_ids":                Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}, update: []string{}},
		"private_ip":             Representation{repType: Optional, create: `10.0.0.5`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
	}
	instanceSourceDetailsRepresentation = map[string]interface{}{
		"source_id":   Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"source_type": Representation{repType: Required, create: `image`},
		"kms_key_id":  Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}

	InstanceCommonVariables = `
variable "InstanceImageOCID" {
	type = "map"
	default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
	}
}

`
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
	InstanceResourceDependencies = generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Update, dedicatedVmHostRepresentation) + SubnetResourceConfig + InstanceCommonVariables +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group1", Required, Create, networkSecurityGroupRepresentation)
)

func TestCoreInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + commonTestVariables() + KeyResourceDependencyConfig

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

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
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
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
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
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
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
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "false"),
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
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
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
					resource.TestCheckResourceAttr(resourceName, "agent_config.0.is_monitoring_disabled", "true"),
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
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "is_pv_encryption_in_transit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.source_id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
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
					resource.TestCheckResourceAttr(datasourceName, "instances.0.agent_config.0.is_monitoring_disabled", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.dedicated_vm_host_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.launch_mode"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.launch_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.system_tags"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_config.0.is_monitoring_disabled", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "launch_mode"),
					resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "system_tags"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
				),
			},
			// verify updates to original parameters
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceRepresentation),
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

func testAccCheckCoreInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
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
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient
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
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient

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
	_, err := client.computeClient.GetInstance(context.Background(), oci_core.GetInstanceRequest{
		InstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
