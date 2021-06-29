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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_core "github.com/oracle/oci-go-sdk/v43/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ClusterNetworkRequiredOnlyResource = ClusterNetworkResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Required, Create, clusterNetworkRepresentation)

	ClusterNetworkResourceConfig = ClusterNetworkResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Optional, Update, clusterNetworkRepresentation)

	clusterNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_network_id": Representation{repType: Required, create: `${oci_core_cluster_network.test_cluster_network.id}`},
	}

	clusterNetworkDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `hpc-cluster-network`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, clusterNetworkDataSourceFilterRepresentation}}
	clusterNetworkDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_cluster_network.test_cluster_network.id}`}},
	}

	clusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_pools":          RepresentationGroup{Required, clusterNetworkInstancePoolsRepresentation},
		"placement_configuration": RepresentationGroup{Required, clusterNetworkPlacementConfigurationRepresentation},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{repType: Optional, create: `hpc-cluster-network`, update: `displayName2`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	clusterNetworkInstancePoolsRepresentation = map[string]interface{}{
		"instance_configuration_id": Representation{repType: Required, create: `${oci_core_instance_configuration.test_instance_configuration.id}`},
		"size":                      Representation{repType: Required, create: `1`, update: `2`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `hpc-cluster-network-pool`, update: `hpc-cluster-network-pool2`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	clusterNetworkPlacementConfigurationRepresentation = map[string]interface{}{
		"availability_domain":    Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"primary_subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"secondary_vnic_subnets": RepresentationGroup{Optional, clusterNetworkPlacementConfigurationSecondaryVnicSubnetsRepresentation},
	}
	clusterNetworkPlacementConfigurationSecondaryVnicSubnetsRepresentation = map[string]interface{}{
		"subnet_id":    Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"display_name": Representation{repType: Optional, create: `backend-servers`},
	}
	instanceConfigurationInstanceDetailsClusterNetworkRepresentation = map[string]interface{}{
		"instance_type":   Representation{repType: Required, create: `compute`},
		"secondary_vnics": RepresentationGroup{Optional, instanceConfigurationInstanceDetailsSecondaryVnicsRepresentation},
		"launch_details":  RepresentationGroup{Optional, instanceConfigurationInstanceDetailsLaunchDetailsClusterNetworkRepresentation},
	}

	availabilityDomainDataSourceClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":         RepresentationGroup{Required, availabilityDomainDataSourceFilterRepresentation}}
	availabilityDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${var.logical_ad}`}},
	}

	AvailabilityDomainClusterNetworkConfig = generateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", Required, Create, availabilityDomainDataSourceClusterNetworkRepresentation)

	instanceConfigurationInstanceDetailsLaunchDetailsClusterNetworkRepresentation = representationCopyWithRemovedProperties(getMultipleUpdatedRepresenationCopy(
		[]string{"shape", "source_details"}, []interface{}{
			Representation{repType: Optional, create: `BM.HPC2.36`},
			RepresentationGroup{Optional, getUpdatedRepresentationCopy("image_id", Representation{repType: Optional, create: `${var.image_id}`}, instanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsRepresentation)},
		}, instanceConfigurationInstanceDetailsLaunchDetailsRepresentation),
		[]string{"shape_config", "dedicated_vm_host_id", "is_pv_encryption_in_transit_enabled", "preferred_maintenance_action"})

	ClusterNetworkResourceRequiredOnlyDependencies = AvailabilityDomainClusterNetworkConfig + DefinedTagsDependencies + VcnResourceConfig + DhcpOptionsRequiredOnlyResource + AnotherSecurityListRequiredOnlyResource +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Optional, Update, getUpdatedRepresentationCopy("cidr_block", Representation{repType: Required, create: `10.0.2.0/24`}, subnetRepresentation)) +
		OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation)

	ClusterNetworkResourceDependencies = ClusterNetworkResourceRequiredOnlyDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create, getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional, instanceConfigurationInstanceDetailsClusterNetworkRepresentation}, instanceConfigurationRepresentation))

	ClusterNetworkResourceDependenciesWithoutSecondaryVnic = ClusterNetworkResourceRequiredOnlyDependencies +
		generateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", Optional, Create,
			getUpdatedRepresentationCopy("instance_details", RepresentationGroup{Optional,
				representationCopyWithRemovedProperties(instanceConfigurationInstanceDetailsClusterNetworkRepresentation, []string{"secondary_vnics"})}, instanceConfigurationPoolRepresentation))
)

func TestCoreClusterNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreClusterNetworkResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	logicalAd := getEnvSettingWithBlankDefault("logical_ad")
	logicalAdVariableStr := fmt.Sprintf("variable \"logical_ad\" { default = \"%s\" }\n", logicalAd)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	imageId := getEnvSettingWithBlankDefault("image_id")
	imageIdVariableStr := fmt.Sprintf("variable \"image_id\" { default = \"%s\" }\n", imageId)

	resourceName := "oci_core_cluster_network.test_cluster_network"
	datasourceName := "data.oci_core_cluster_networks.test_cluster_networks"
	singularDatasourceName := "data.oci_core_cluster_network.test_cluster_network"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create" step in the test.
	saveConfigContent(config+logicalAdVariableStr+compartmentIdVariableStr+imageIdVariableStr+ClusterNetworkResourceDependenciesWithoutSecondaryVnic+
		generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Required, Create, clusterNetworkRepresentation), "core", "clusterNetwork", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreClusterNetworkDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependenciesWithoutSecondaryVnic +
					generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Required, Create, clusterNetworkRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_pools.0.instance_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.0.size", "1"),
					resource.TestCheckResourceAttr(resourceName, "placement_configuration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "placement_configuration.0.primary_subnet_id"),

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
				Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Optional, Create, clusterNetworkRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "hpc-cluster-network"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.0.defined_tags.%", "1"),
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
					CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
						"display_name": "backend-servers",
					},
						[]string{
							"subnet_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				Config: config + logicalAdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Optional, Create,
						representationCopyWithNewProperties(clusterNetworkRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "hpc-cluster-network"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_pools.0.defined_tags.%", "1"),
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
					CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
						"display_name": "backend-servers",
					},
						[]string{
							"subnet_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				Config: config + logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Optional, Update, clusterNetworkRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					CheckResourceSetContainsElementWithProperties(resourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
						"display_name": "backend-servers",
					},
						[]string{
							"subnet_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					generateDataSourceFromRepresentationMap("oci_core_cluster_networks", "test_cluster_networks", Optional, Update, clusterNetworkDataSourceRepresentation) +
					logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Optional, Update, clusterNetworkRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					//resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "cluster_networks.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cluster_networks.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_core_cluster_network", "test_cluster_network", Required, Create, clusterNetworkSingularDataSourceRepresentation) +
					logicalAdVariableStr + compartmentIdVariableStr + imageIdVariableStr + ClusterNetworkResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_network_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "placement_configuration.0.secondary_vnic_subnets", map[string]string{
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
		},
	})
}

func testAccCheckCoreClusterNetworkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cluster_network" {
			noResourceFound = false
			request := oci_core.GetClusterNetworkRequest{}

			tmp := rs.Primary.ID
			request.ClusterNetworkId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreClusterNetwork") {
		resource.AddTestSweepers("CoreClusterNetwork", &resource.Sweeper{
			Name:         "CoreClusterNetwork",
			Dependencies: DependencyGraph["clusterNetwork"],
			F:            sweepCoreClusterNetworkResource,
		})
	}
}

func sweepCoreClusterNetworkResource(compartment string) error {
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient()
	clusterNetworkIds, err := getClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterNetworkId := range clusterNetworkIds {
		if ok := SweeperDefaultResourceId[clusterNetworkId]; !ok {
			terminateClusterNetworkRequest := oci_core.TerminateClusterNetworkRequest{}

			terminateClusterNetworkRequest.ClusterNetworkId = &clusterNetworkId

			terminateClusterNetworkRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeManagementClient.TerminateClusterNetwork(context.Background(), terminateClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterNetworkId, error)
				continue
			}
			waitTillCondition(testAccProvider, &clusterNetworkId, clusterNetworkSweepWaitCondition,
				time.Duration(7*time.Minute),
				clusterNetworkSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getClusterNetworkIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeManagementClient := GetTestClients(&schema.ResourceData{}).computeManagementClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ClusterNetworkId", id)
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

func clusterNetworkSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeManagementClient().GetClusterNetwork(context.Background(), oci_core.GetClusterNetworkRequest{
		ClusterNetworkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
