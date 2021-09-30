// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_core "github.com/oracle/oci-go-sdk/v48/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InstancePoolRequiredOnlyResource = InstancePoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation)

	InstancePoolResourceConfig = InstancePoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation)

	instancePoolSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id": Representation{RepType: Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instancePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `RUNNING`},
		"filter":         RepresentationGroup{Required, instancePoolDataSourceFilterRepresentation}}
	instancePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_instance_pool.test_instance_pool.id}`}},
	}

	instancePoolRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"instance_configuration_id": Representation{RepType: Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  RepresentationGroup{Required, instancePoolPlacementConfigurationsRepresentation},
		"size":                      Representation{RepType: Required, Create: `2`, Update: `3`},
		"state":                     Representation{RepType: Optional, Create: `Running`},
		"defined_tags":              Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{RepType: Optional, Create: `backend-servers-pool`, Update: `displayName2`},
		"freeform_tags":             Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            RepresentationGroup{Optional, instancePoolLoadBalancersRepresentation},
	}
	instancePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain":    Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":          Representation{RepType: Optional, Create: []string{`FAULT-DOMAIN-1`}, Update: []string{`FAULT-DOMAIN-2`}},
		"secondary_vnic_subnets": RepresentationGroup{Optional, instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation},
	}
	instancePoolLoadBalancersRepresentation = map[string]interface{}{
		"backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             Representation{RepType: Required, Create: `10`},
		"vnic_selection":   Representation{RepType: Required, Create: `PrimaryVnic`},
	}
	instancePoolLoadBalancers2Representation = map[string]interface{}{
		"backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set2.name}`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer2.id}`},
		"port":             Representation{RepType: Required, Create: `10`},
		"vnic_selection":   Representation{RepType: Required, Create: `PrimaryVnic`},
	}
	instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id": Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		//the display_name should be the same as in the instance configuration
		"display_name": Representation{RepType: Required, Create: `backend-servers-pool`},
	}

	instanceConfigurationPoolRepresentation = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"instance_details": RepresentationGroup{Required, instanceConfigurationInstanceDetailsPoolRepresentation},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{RepType: Optional, Create: `backend-servers`, Update: `displayName2`},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	instanceConfigurationInstanceDetailsPoolRepresentation = map[string]interface{}{
		"instance_type":   Representation{RepType: Required, Create: `compute`},
		"secondary_vnics": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation},
		"launch_details":  RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation = map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		//the display_name should be the same as in the secondary_vnic_subnets
		"display_name": Representation{RepType: Optional, Create: `backend-servers-pool`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation = map[string]interface{}{
		"compartment_id":                      Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		"defined_tags":                        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        Representation{RepType: Optional, Create: `backend-servers`},
		"extended_metadata":                   Representation{RepType: Optional, Create: map[string]string{"extendedMetadata": "extendedMetadata"}, Update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         Representation{RepType: Optional, Create: `ipxeScript`},
		"metadata":                            Representation{RepType: Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               Representation{RepType: Optional, Create: InstanceConfigurationVmShape},
		"source_details":                      RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"is_pv_encryption_in_transit_enabled": Representation{RepType: Optional, Create: `false`},
		"launch_mode":                         Representation{RepType: Optional, Create: `NATIVE`},
		"preferred_maintenance_action":        Representation{RepType: Optional, Create: `LIVE_MIGRATE`},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation = map[string]interface{}{
		"assign_public_ip":       Representation{RepType: Optional, Create: `true`},
		"display_name":           Representation{RepType: Optional, Create: `backend-servers`},
		"skip_source_dest_check": Representation{RepType: Optional, Create: `false`},
	}

	InstancePoolResourceDependenciesWithoutSecondaryVnic = SubnetResourceConfig + OciImageIdsVariable + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
			GetUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional,
				RepresentationCopyWithRemovedProperties(GetUpdatedRepresentationCopy("launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationInstanceDetailsPoolRepresentation), []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))

	InstancePoolResourceDependencies = OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create, GetUpdatedRepresentationCopy("instance_details.launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationPoolRepresentation)) +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set2", Required, Create, backendSet2Representation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer2", Required, Create, loadBalancer2Representation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_pool.test_instance_pool"
	datasourceName := "data.oci_core_instance_pools.test_instance_pools"
	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+InstancePoolResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create, instancePoolRepresentation), "core", "instancePool", t)

	ResourceTest(t, testAccCheckCoreInstancePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + InstancePoolResourceDependenciesWithoutSecondaryVnic +
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "2"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create,
					RepresentationCopyWithNewProperties(instancePoolRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, RepresentationCopyWithNewProperties(instancePoolRepresentation, map[string]interface{}{
					"load_balancers": []RepresentationGroup{{Optional, instancePoolLoadBalancersRepresentation}, {Optional, instancePoolLoadBalancers2Representation}},
				})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update,
					GetUpdatedRepresentationCopy("state", Representation{RepType: Optional, Create: "Stopped"}, instancePoolRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_core_instance_pools", "test_instance_pools", Optional, Update, instancePoolDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_pools.0.defined_tags.%", "1"),
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
				GenerateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstancePoolResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
	client := testAccProvider.Meta().(*OracleClients).computeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_pool" {
			noResourceFound = false
			request := oci_core.GetInstancePoolRequest{}

			tmp := rs.Primary.ID
			request.InstancePoolId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("CoreInstancePool") {
		resource.AddTestSweepers("CoreInstancePool", &resource.Sweeper{
			Name:         "CoreInstancePool",
			Dependencies: DependencyGraph["instancePool"],
			F:            sweepCoreInstancePoolResource,
		})
	}
}

func sweepCoreInstancePoolResource(compartment string) error {
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient()
	instancePoolIds, err := getInstancePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, instancePoolId := range instancePoolIds {
		if ok := SweeperDefaultResourceId[instancePoolId]; !ok {
			terminateInstancePoolRequest := oci_core.TerminateInstancePoolRequest{}

			terminateInstancePoolRequest.InstancePoolId = &instancePoolId

			terminateInstancePoolRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateInstancePool(context.Background(), terminateInstancePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting InstancePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", instancePoolId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &instancePoolId, instancePoolSweepWaitCondition, time.Duration(3*time.Minute),
				instancePoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstancePoolIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "InstancePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient()

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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "InstancePoolId", id)
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

func instancePoolSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeManagementClient().GetInstancePool(context.Background(), oci_core.GetInstancePoolRequest{
		InstancePoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
