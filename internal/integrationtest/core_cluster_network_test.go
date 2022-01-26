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
	ClusterNetworkRequiredOnlyResource = ClusterNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Required, acctest.Create, clusterNetworkRepresentation)

	ClusterNetworkResourceConfig = ClusterNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Optional, acctest.Update, clusterNetworkRepresentation)

	clusterNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_network_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cluster_network.test_cluster_network.id}`},
	}

	clusterNetworkDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `hpc-cluster-network`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: clusterNetworkDataSourceFilterRepresentation}}
	clusterNetworkDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_cluster_network.test_cluster_network.id}`}},
	}

	clusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"instance_pools":          acctest.RepresentationGroup{RepType: acctest.Required, Group: clusterNetworkInstancePoolsRepresentation},
		"placement_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: clusterNetworkPlacementConfigurationRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `hpc-cluster-network`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	clusterNetworkInstancePoolsRepresentation = map[string]interface{}{
		"instance_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_configuration.test_instance_configuration.id}`, Update: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"size":                      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `hpc-cluster-network-pool`, Update: `hpc-cluster-network-pool2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	clusterNetworkPlacementConfigurationRepresentation = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"secondary_vnic_subnets": acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterNetworkPlacementConfigurationSecondaryVnicSubnetsRepresentation},
	}
	clusterNetworkPlacementConfigurationSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `backend-servers`},
	}
	instanceConfigurationInstanceDetailsClusterNetworkRepresentation = map[string]interface{}{
		"instance_type":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"secondary_vnics": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation},
		"launch_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsLaunchDetailsClusterNetworkRepresentation},
	}

	availabilityDomainDataSourceClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: availabilityDomainDataSourceFilterRepresentation}}
	availabilityDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.logical_ad}`}},
	}

	AvailabilityDomainClusterNetworkConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", acctest.Required, acctest.Create, availabilityDomainDataSourceClusterNetworkRepresentation)

	instanceConfigurationInstanceDetailsLaunchDetailsClusterNetworkRepresentation = acctest.RepresentationCopyWithRemovedProperties(acctest.GetMultipleUpdatedRepresenationCopy(
		[]string{"shape", "source_details"}, []interface{}{
			acctest.Representation{RepType: acctest.Optional, Create: `BM.HPC2.36`},
			acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.GetUpdatedRepresentationCopy("image_id", acctest.Representation{RepType: acctest.Optional, Create: `${var.image_id}`}, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)},
		}, instanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"shape_config", "dedicated_vm_host_id", "is_pv_encryption_in_transit_enabled", "preferred_maintenance_action"})

	ClusterNetworkResourceRequiredOnlyDependencies = AvailabilityDomainClusterNetworkConfig + DefinedTagsDependencies + VcnResourceConfig + DhcpOptionsRequiredOnlyResource + AnotherSecurityListRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("cidr_block", acctest.Representation{RepType: acctest.Required, Create: `10.0.2.0/24`}, subnetRepresentation)) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation)

	ClusterNetworkResourceDependencies = ClusterNetworkResourceRequiredOnlyDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceConfigurationInstanceDetailsClusterNetworkRepresentation}, instanceConfigurationRepresentation))

	ClusterNetworkResourceDependenciesWithoutSecondaryVnic = ClusterNetworkResourceRequiredOnlyDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("instance_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithRemovedProperties(instanceConfigurationInstanceDetailsClusterNetworkRepresentation, []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))
)

// issue-routing-tag: core/computeManagement
func TestCoreClusterNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreClusterNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	logicalAd := utils.GetEnvSettingWithBlankDefault("logical_ad")
	logicalAdVariableStr := fmt.Sprintf("variable \"logical_ad\" { default = \"%s\" }\n", logicalAd)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	imageId := utils.GetEnvSettingWithBlankDefault("image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	resourceName := "oci_core_cluster_network.test_cluster_network"
	datasourceName := "data.oci_core_cluster_networks.test_cluster_networks"
	singularDatasourceName := "data.oci_core_cluster_network.test_cluster_network"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+logicalAdVariableStr+compartmentIdVariableStr+imageIdVariableStr+ClusterNetworkResourceDependenciesWithoutSecondaryVnic+
		acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Required, acctest.Create, clusterNetworkRepresentation), "core", "clusterNetwork", t)

	acctest.ResourceTest(t, testAccCheckCoreClusterNetworkDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependenciesWithoutSecondaryVnic +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Required, acctest.Create, clusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.size", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.primary_subnet_id"),

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
			Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Optional, acctest.Create, clusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "hpc-cluster-network"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.display_name", "hpc-cluster-network-pool"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.placement_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.size", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.0.secondary_vnic_subnets.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
					"display_name": "backend-servers",
				},
					[]string{
						"subnet_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + logicalAdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(clusterNetworkRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "hpc-cluster-network"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.display_name", "hpc-cluster-network-pool"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.placement_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.size", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.0.secondary_vnic_subnets.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
					"display_name": "backend-servers",
				},
					[]string{
						"subnet_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Optional, acctest.Update, clusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.display_name", "hpc-cluster-network-pool2"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.placement_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_pools.0.size", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.primary_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_configuration.0.secondary_vnic_subnets.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
					"display_name": "backend-servers",
				},
					[]string{
						"subnet_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cluster_networks", "test_cluster_networks", acctest.Optional, acctest.Update, clusterNetworkDataSourceRepresentation) +
				logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Optional, acctest.Update, clusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.0.display_name", "hpc-cluster-network-pool2"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.instance_pools.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.instance_pools.0.instance_configuration_id"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.0.placement_configurations.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.instance_pools.0.size", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.instance_pools.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_networks.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", acctest.Required, acctest.Create, clusterNetworkSingularDataSourceRepresentation) +
				logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_network_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.display_name", "hpc-cluster-network-pool2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pools.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.load_balancers.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.placement_configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_pools.0.size", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pools.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pools.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configuration.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configuration.0.secondary_vnic_subnets.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
					"display_name": "backend-servers",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreClusterNetworkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cluster_network" {
			noResourceFound = false
			request := oci_core.GetClusterNetworkRequest{}

			tmp := rs.Primary.ID
			request.ClusterNetworkId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetClusterNetwork(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ClusterNetworkLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreClusterNetwork") {
		resource.AddTestSweepers("CoreClusterNetwork", &resource.Sweeper{
			Name:         "CoreClusterNetwork",
			Dependencies: acctest.DependencyGraph["clusterNetwork"],
			F:            sweepCoreClusterNetworkResource,
		})
	}
}

func sweepCoreClusterNetworkResource(compartment string) error {
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()
	clusterNetworkIds, err := getClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterNetworkId := range clusterNetworkIds {
		if ok := acctest.SweeperDefaultResourceId[clusterNetworkId]; !ok {
			terminateClusterNetworkRequest := oci_core.TerminateClusterNetworkRequest{}

			terminateClusterNetworkRequest.ClusterNetworkId = &clusterNetworkId

			terminateClusterNetworkRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateClusterNetwork(context.Background(), terminateClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterNetworkId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterNetworkId, clusterNetworkSweepWaitCondition,
				time.Duration(7*time.Minute),
				clusterNetworkSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getClusterNetworkIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeManagementClient()

	listClusterNetworksRequest := oci_core.ListClusterNetworksRequest{}
	listClusterNetworksRequest.CompartmentId = &compartmentId
	listClusterNetworksRequest.LifecycleState = oci_core.ClusterNetworkSummaryLifecycleStateRunning
	listClusterNetworksResponse, err := computeManagementClient.ListClusterNetworks(context.Background(), listClusterNetworksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterNetwork list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterNetwork := range listClusterNetworksResponse.Items {
		id := *clusterNetwork.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterNetworkId", id)
	}
	return resourceIds, nil
}

func clusterNetworkSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterNetworkResponse, ok := response.Response.(oci_core.GetClusterNetworkResponse); ok {
		return clusterNetworkResponse.LifecycleState != oci_core.ClusterNetworkLifecycleStateTerminated
	}
	return false
}

func clusterNetworkSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeManagementClient().GetClusterNetwork(context.Background(), oci_core.GetClusterNetworkRequest{
		ClusterNetworkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
