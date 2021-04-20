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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_core "github.com/oracle/oci-go-sdk/v40/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InstancePoolRequiredOnlyResource = InstancePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation)

	InstancePoolResourceConfig = InstancePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation)

	instancePoolSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
	}

	instancePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `backend-servers-pool`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `RUNNING`},
		"filter":         RepresentationGroup{Required, instancePoolDataSourceFilterRepresentation}}
	instancePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance_pool.test_instance_pool.id}`}},
	}

	instancePoolRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_configuration_id": Representation{repType: Required, create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"placement_configurations":  RepresentationGroup{Required, instancePoolPlacementConfigurationsRepresentation},
		"size":                      Representation{repType: Required, create: `2`, update: `3`},
		"state":                     Representation{repType: Optional, create: `Running`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `backend-servers-pool`, update: `displayName2`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"load_balancers":            RepresentationGroup{Optional, instancePoolLoadBalancersRepresentation},
	}
	instancePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain":    Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"fault_domains":          Representation{repType: Optional, create: []string{`FAULT-DOMAIN-1`}, update: []string{`FAULT-DOMAIN-2`}},
		"secondary_vnic_subnets": RepresentationGroup{Optional, instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation},
	}
	instancePoolLoadBalancersRepresentation = map[string]interface{}{
		"backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             Representation{repType: Required, create: `10`},
		"vnic_selection":   Representation{repType: Required, create: `PrimaryVnic`},
	}
	instancePoolPlacementConfigurationsSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id": Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		//the display_name should be the same as in the instance configuration
		"display_name": Representation{repType: Required, create: `backend-servers-pool`},
	}

	instanceConfigurationPoolRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_details": RepresentationGroup{Required, instanceConfigurationInstanceDetailsPoolRepresentation},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":     Representation{repType: Optional, create: `backend-servers`, update: `displayName2`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	instanceConfigurationInstanceDetailsPoolRepresentation = map[string]interface{}{
		"instance_type":   Representation{repType: Required, create: `compute`},
		"secondary_vnics": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation},
		"launch_details":  RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation},
	}
	instanceConfigurationInstanceDetailsSecondaryVnicsPoolRepresentation = map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		//the display_name should be the same as in the secondary_vnic_subnets
		"display_name": Representation{repType: Optional, create: `backend-servers-pool`},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsPoolRepresentation = map[string]interface{}{
		"compartment_id":                      Representation{repType: Optional, create: `${var.compartment_id}`},
		"create_vnic_details":                 RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation},
		"defined_tags":                        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        Representation{repType: Optional, create: `backend-servers`},
		"extended_metadata":                   Representation{repType: Optional, create: map[string]string{"extendedMetadata": "extendedMetadata"}, update: map[string]string{"extendedMetadata2": "extendedMetadata2"}},
		"freeform_tags":                       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ipxe_script":                         Representation{repType: Optional, create: `ipxeScript`},
		"metadata":                            Representation{repType: Optional, create: map[string]string{"metadata": "metadata"}, update: map[string]string{"metadata2": "metadata2"}},
		"shape":                               Representation{repType: Optional, create: InstanceConfigurationVmShape},
		"source_details":                      RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation},
		"agent_config":                        RepresentationGroup{Optional, instanceAgentConfigRepresentation},
		"launch_options":                      RepresentationGroup{Optional, instanceLaunchOptionsRepresentation},
		"is_pv_encryption_in_transit_enabled": Representation{repType: Optional, create: `false`},
		"launch_mode":                         Representation{repType: Optional, create: `NATIVE`},
		"preferred_maintenance_action":        Representation{repType: Optional, create: `LIVE_MIGRATE`},
		"shape_config":                        RepresentationGroup{Optional, instanceShapeConfigRepresentation},
	}
	instanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsPoolRepresentation = map[string]interface{}{
		"assign_public_ip":       Representation{repType: Optional, create: `true`},
		"display_name":           Representation{repType: Optional, create: `backend-servers`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
	}

	InstancePoolResourceDependenciesWithoutSecondaryVnic = SubnetResourceConfig + OciImageIdsVariable + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}` +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
			getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional,
				representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationInstanceDetailsPoolRepresentation), []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))

	InstancePoolResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create, getUpdatedRepresentationCopy("instance_details.launch_details.launch_options", instanceLaunchOptionsRepresentationForInstanceConfiguration, instanceConfigurationPoolRepresentation)) +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestCoreInstancePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance_pool.test_instance_pool"
	datasourceName := "data.oci_core_instance_pools.test_instance_pools"
	singularDatasourceName := "data.oci_core_instance_pool.test_instance_pool"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+InstancePoolResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create, instancePoolRepresentation), "core", "instancePool", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstancePoolDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependenciesWithoutSecondaryVnic +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "size", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Create,
						representationCopyWithNewProperties(instancePoolRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update,
						getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "Stopped"}, instancePoolRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_instance_pools", "test_instance_pools", Optional, Update, instancePoolDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Optional, Update, instancePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if !inSweeperExcludeList("CoreInstancePool") {
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

			terminateInstancePoolRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateInstancePool(context.Background(), terminateInstancePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting InstancePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", instancePoolId, error)
				continue
			}
			waitTillCondition(testAccProvider, &instancePoolId, instancePoolSweepWaitCondition, time.Duration(3*time.Minute),
				instancePoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getInstancePoolIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "InstancePoolId")
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
		addResourceIdToSweeperResourceIdMap(compartmentId, "InstancePoolId", id)
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
