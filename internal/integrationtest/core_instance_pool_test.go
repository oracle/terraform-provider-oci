// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	InstancePoolRequiredOnlyResource = InstancePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, instancePoolRepresentation)

	InstancePoolResourceConfig = InstancePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, instancePoolRepresentation)

	instancePoolSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instancePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePoolDataSourceFilterRepresentation}}
	instancePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_instance_pool.test_instance_pool.id}`}},
	}

	instancePoolRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  acctest.RepresentationGroup{RepType: acctest.Required, Group: instancePoolPlacementConfigurationsRepresentation},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `Running`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: instancePoolLoadBalancersRepresentation},
	}
	instancePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":          acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}, Update: []string{`FAULT-DOMAIN-2`}},
		"secondary_vnic_subnets": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation},
	}
	instancePoolLoadBalancersRepresentation = map[string]interface{}{
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"vnic_selection":   acctest.Representation{RepType: acctest.Required, Create: `PrimaryVnic`},
	}
	instancePoolLoadBalancers2Representation = map[string]interface{}{
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set2.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer2.id}`},
		"port":             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"vnic_selection":   acctest.Representation{RepType: acctest.Required, Create: `PrimaryVnic`},
	}
	instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		//the display_name should be the same as in the instance configuration
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `backend-servers-pool`},
	}

	instanceConfigurationPoolRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceConfigurationInstanceDetailsPoolRepresentation},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	instanceConfigurationInstanceDetailsPoolRepresentation = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation},
		"launch_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		//the display_name should be the same as in the secondary_vnic_subnets
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers-pool`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"extended_metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"extendedMetadata": "extendedMetadata"}, Update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               acctest.Representation{RepType: acctest.Optional, Create: InstanceConfigurationVmShape},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceAgentConfigRepresentation},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceLaunchOptionsRepresentation},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `NATIVE`},
		"preferred_maintenance_action":        acctest.Representation{RepType: acctest.Optional, Create: `LIVE_MIGRATE`},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation = map[string]interface{}{
		"assign_public_ip":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	InstancePoolResourceDependenciesWithoutSecondaryVnic = SubnetResourceConfig + utils.OciImageIdsVariable + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationInstanceDetailsPoolRepresentation), []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))

	InstancePoolResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("instance_details.launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationPoolRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set2", acctest.Required, acctest.Create, backendSet2Representation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer2", acctest.Required, acctest.Create, loadBalancer2Representation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_pool.test_instance_pool"
	datasourceName := "data.oci_core_instance_pools.test_instance_pools"
	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InstancePoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, instancePoolRepresentation), "core", "instancePool", t)

	acctest.ResourceTest(t, testAccCheckCoreInstancePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependenciesWithoutSecondaryVnic +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.secondary_vnic_subnets.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(instancePoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "backend-servers-pool"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
		// verify attach
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(instancePoolRepresentation, map[string]interface{}{
					"load_balancers": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: instancePoolLoadBalancersRepresentation}, {RepType: acctest.Optional, Group: instancePoolLoadBalancers2Representation}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.1.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.1.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.1.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
		// verify detach
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.backend_set_name"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
		// verify stop the Instance Pool
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("state", acctest.Representation{RepType: acctest.Optional, Create: "Stopped"}, instancePoolRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
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
		// verify start the Instance Pool
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "3"),
				resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
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
		// verify datasource the state will be updated to RUNNING
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pools", "test_instance_pools", acctest.Optional, acctest.Update, instancePoolDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, instancePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.size", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_pools.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, instancePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.instance_pool_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.port", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancers.0.vnic_selection", "PrimaryVnic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + InstancePoolResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreInstancePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_pool" {
			noResourceFound = false
			request := oci_core.GetInstancePoolRequest{}

			tmp := rs.Primary.ID
			request.InstancePoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetInstancePool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InstancePoolLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreInstancePool") {
		resource.AddTestSweepers("CoreInstancePool", &resource.Sweeper{
			Name:         "CoreInstancePool",
			Dependencies: acctest.DependencyGraph["instancePool"],
			F:            sweepCoreInstancePoolResource,
		})
	}
}

func sweepCoreInstancePoolResource(compartment string) error {
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()
	instancePoolIds, err := getInstancePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, instancePoolId := range instancePoolIds {
		if ok := acctest.SweeperDefaultResourceId[instancePoolId]; !ok {
			terminateInstancePoolRequest := oci_core.TerminateInstancePoolRequest{}

			terminateInstancePoolRequest.InstancePoolId = &instancePoolId

			terminateInstancePoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateInstancePool(context.Background(), terminateInstancePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting InstancePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", instancePoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &instancePoolId, instancePoolSweepWaitCondition, time.Duration(3*time.Minute),
				instancePoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstancePoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstancePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()

	listInstancePoolsRequest := oci_core.ListInstancePoolsRequest{}
	listInstancePoolsRequest.CompartmentId = &compartmentId
	listInstancePoolsRequest.LifecycleState = oci_core.InstancePoolSummaryLifecycleStateRunning
	listInstancePoolsResponse, err := computeManagementClient.ListInstancePools(context.Background(), listInstancePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InstancePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instancePool := range listInstancePoolsResponse.Items {
		id := *instancePool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstancePoolId", id)
	}
	return resourceIds, nil
}

func instancePoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if instancePoolResponse, ok := response.Response.(oci_core.GetInstancePoolResponse); ok {
		return instancePoolResponse.LifecycleState != oci_core.InstancePoolLifecycleStateTerminated
	}
	return false
}

func instancePoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeManagementClient().GetInstancePool(context.Background(), oci_core.GetInstancePoolRequest{
		InstancePoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
