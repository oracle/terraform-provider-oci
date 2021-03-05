// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instancePoolInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_pool_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"display_name":     Representation{repType: Optional, create: `displayName`},
		"filter":           RepresentationGroup{Required, instancePoolInstanceDataSourceFilterRepresentation}}
	instancePoolInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance_pool_instance.test_instance_pool_instance.id}`}},
	}

	instancePoolInstanceRepresentation = map[string]interface{}{
		"instance_id":      Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"instance_pool_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instanceForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain":  Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                Representation{repType: Required, create: `VM.Standard2.1`},
		"agent_config":         RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"availability_config":  RepresentationGroup{Optional, instanceAvailabilityConfigRepresentation},
		"create_vnic_details":  RepresentationGroup{Optional, instanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id": Representation{repType: Optional, create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":         Representation{repType: Optional, create: `displayName`},
		"extended_metadata": Representation{repType: Optional, create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}},
		"fault_domain":                        Representation{repType: Required, create: `FAULT-DOMAIN-1`},
		"freeform_tags":                       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}},
		"hostname_label":                      Representation{repType: Optional, create: `hostnamelabel`},
		"instance_options":                    RepresentationGroup{Optional, instanceInstanceOptionsRepresentation},
		"image":                               Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         Representation{repType: Optional, create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"metadata":                            Representation{repType: Optional, create: map[string]string{"user_data": "abcd"}},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
		"source_details":                      RepresentationGroup{Optional, instanceSourceDetailsRepresentation},
		"subnet_id":                           Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               Representation{repType: Required, create: `RUNNING`},
	}

	instancePoolForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_configuration_id": Representation{repType: Required, create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  RepresentationGroup{Required, instancePoolPlacementConfigurationsForAttachInstanceRepresentation},
		"size":                      Representation{repType: Required, create: `0`},
		"state":                     Representation{repType: Required, create: `Running`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              Representation{repType: Optional, create: `backend-servers-pool`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}},
		"load_balancers":            RepresentationGroup{Required, instancePoolLoadBalancersRepresentation},
		"lifecycle":                 RepresentationGroup{Required, ignoreInstancePoolSizeChanges},
	}

	// Needs to ignore this size because attach/detach will internally modify the size of the instance pool
	ignoreInstancePoolSizeChanges = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`size`}},
	}

	instancePoolPlacementConfigurationsForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":   Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":       Representation{repType: Required, create: []string{`FAULT-DOMAIN-1`}},
	}

	instanceConfigurationFromInstanceForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `backend-servers`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"instance_id":    Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"source":         Representation{repType: Required, create: `INSTANCE`},
	}

	InstancePoolInstanceResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolForAttachInstanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Required, Create, instanceConfigurationFromInstanceForAttachInstanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceForAttachInstanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestCoreInstancePoolInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_pool_instance.test_instance_pool_instance"
	datasourceName := "data.oci_core_instance_pool_instances.test_instance_pool_instances"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+InstancePoolInstanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation), "core", "instancePoolInstance", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_pool_id"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_pool_instances", "test_instance_pool_instances", Required, Create, instancePoolInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					// only verify the number of instance pool instances because after detach, there will be no instance pool instances
					resource.TestCheckResourceAttrSet(datasourceName, "instances.#"),
					resource.TestCheckResourceAttr(datasourceName, "instances.#", "0"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"instance_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
