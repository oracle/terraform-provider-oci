// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	instancePoolInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePoolInstanceDataSourceFilterRepresentation}}
	instancePoolInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_pool_instance.test_instance_pool_instance.id}`}},
	}

	instancePoolInstanceRepresentation = map[string]interface{}{
		"instance_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"instance_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instanceForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigRepresentation},
		"availability_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAvailabilityConfigRepresentation},
		"create_vnic_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"extended_metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{
			"some_string":   "stringA",
			"nested_object": "{\\\"some_string\\\": \\\"stringB\\\", \\\"object\\\": {\\\"some_string\\\": \\\"stringC\\\"}}",
		}},
		"fault_domain":                        acctest.Representation{RepType: acctest.Required, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceSourceDetailsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Required, Create: `RUNNING`},
	}

	instancePoolForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePoolPlacementConfigurationsForAttachInstanceRepresentation},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `0`},
		"state":                     acctest.Representation{RepType: acctest.Required, Create: `Running`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"load_balancers":            acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePoolLoadBalancersRepresentation},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreInstancePoolSizeChanges},
	}

	// Needs to ignore this size because attach/detach will internally modify the size of the instance pool
	ignoreInstancePoolSizeChanges = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`size`}},
	}

	instancePoolPlacementConfigurationsForAttachInstanceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":       acctest.Representation{RepType: acctest.Required, Create: []string{`FAULT-DOMAIN-1`}},
	}

	instanceConfigurationFromInstanceForAttachInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"source":         acctest.Representation{RepType: acctest.Required, Create: `INSTANCE`},
	}

	InstancePoolInstanceResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, instancePoolForAttachInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Required, acctest.Create, instanceConfigurationFromInstanceForAttachInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceForAttachInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_pool_instance.test_instance_pool_instance"
	datasourceName := "data.oci_core_instance_pool_instances.test_instance_pool_instances"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InstancePoolInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", acctest.Required, acctest.Create, instancePoolInstanceRepresentation), "core", "instancePoolInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", acctest.Required, acctest.Create, instancePoolInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pool_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pool_instances", "test_instance_pool_instances", acctest.Required, acctest.Create, instancePoolInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool_instance", "test_instance_pool_instance", acctest.Required, acctest.Create, instancePoolInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
