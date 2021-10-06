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
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"instance_pool_id": Representation{RepType: Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"display_name":     Representation{RepType: Optional, Create: `displayName`},
		"filter":           RepresentationGroup{Required, instancePoolInstanceDataSourceFilterRepresentation}}
	instancePoolInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_instance_pool_instance.test_instance_pool_instance.id}`}},
	}

	instancePoolInstanceRepresentation = map[string]interface{}{
		"instance_id":      Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"instance_pool_id": Representation{RepType: Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instanceForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain":  Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       Representation{RepType: Required, Create: `${var.compartment_id}`},
		"shape":                Representation{RepType: Required, Create: `VM.Standard2.1`},
		"agent_config":         RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"availability_config":  RepresentationGroup{Optional, instanceAvailabilityConfigRepresentation},
		"create_vnic_details":  RepresentationGroup{Optional, instanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id": Representation{RepType: Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":         Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":         Representation{RepType: Optional, Create: `displayName`},
		"extended_metadata": Representation{RepType: Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}},
		"fault_domain":                        Representation{RepType: Required, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":                       Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":                      Representation{RepType: Optional, Create: `hostnamelabel`},
		"instance_options":                    RepresentationGroup{Optional, instanceInstanceOptionsRepresentation},
		"image":                               Representation{RepType: Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         Representation{RepType: Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": Representation{RepType: Optional, Create: `false`},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"metadata":                            Representation{RepType: Optional, Create: map[string]string{"user_data": "abcd"}},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
		"source_details":                      RepresentationGroup{Optional, instanceSourceDetailsRepresentation},
		"subnet_id":                           Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               Representation{RepType: Required, Create: `RUNNING`},
	}

	instancePoolForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": Representation{RepType: Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  RepresentationGroup{Required, instancePoolPlacementConfigurationsForAttachInstanceRepresentation},
		"size":                      Representation{RepType: Required, Create: `0`},
		"state":                     Representation{RepType: Required, Create: `Running`},
		"defined_tags":              Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              Representation{RepType: Optional, Create: `backend-servers-pool`},
		"freeform_tags":             Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}},
		"load_balancers":            RepresentationGroup{Required, instancePoolLoadBalancersRepresentation},
		"lifecycle":                 RepresentationGroup{Required, ignoreInstancePoolSizeChanges},
	}

	// Needs to ignore this size because attach/detach will internally modify the size of the instance pool
	ignoreInstancePoolSizeChanges = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`size`}},
	}

	instancePoolPlacementConfigurationsForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":   Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":       Representation{RepType: Required, Create: []string{`FAULT-DOMAIN-1`}},
	}

	instanceConfigurationFromInstanceForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_id":    Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"source":         Representation{RepType: Required, Create: `INSTANCE`},
	}

	InstancePoolInstanceResourceDependencies = OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolForAttachInstanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Required, Create, instanceConfigurationFromInstanceForAttachInstanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceForAttachInstanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_pool_instance.test_instance_pool_instance"
	datasourceName := "data.oci_core_instance_pool_instances.test_instance_pool_instances"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+InstancePoolInstanceResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation), "core", "instancePoolInstance", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pool_id"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_instance_pool_instances", "test_instance_pool_instances", Required, Create, instancePoolInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", Required, Create, instancePoolInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
