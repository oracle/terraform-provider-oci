// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v25/containerengine"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NodePoolRequiredOnlyResource = NodePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation)

	NodePoolResourceConfig = NodePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation)

	nodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_id": Representation{repType: Required, create: `${oci_containerengine_node_pool.test_node_pool.id}`},
	}

	nodePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cluster_id":     Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           Representation{repType: Optional, create: `name`, update: `name2`},
		"filter":         RepresentationGroup{Required, nodePoolDataSourceFilterRepresentation}}
	nodePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_containerengine_node_pool.test_node_pool.id}`}},
	}

	// check the resource test file for other node sources
	nodePoolRepresentation = map[string]interface{}{
		"cluster_id":          Representation{repType: Required, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"kubernetes_version":  Representation{repType: Required, create: `${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions.0}`},
		"name":                Representation{repType: Required, create: `name`, update: `name2`},
		"node_image_name":     Representation{repType: Required, create: `${data.oci_containerengine_node_pool_option.test_node_pool_option.images.0}`},
		"node_shape":          Representation{repType: Required, create: `VM.Standard2.1`, update: `VM.Standard2.2`},
		"subnet_ids":          Representation{repType: Required, create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": RepresentationGroup{Optional, nodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       Representation{repType: Optional, create: map[string]string{"nodeMetadata": "nodeMetadata"}, update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"quantity_per_subnet": Representation{repType: Optional, create: `1`, update: `2`},
		"ssh_public_key":      Representation{repType: Optional, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}
	nodePoolInitialNodeLabelsRepresentation = map[string]interface{}{
		"key":   Representation{repType: Optional, create: `key`, update: `key2`},
		"value": Representation{repType: Optional, create: `value`, update: `value2`},
	}

	NodePoolResourceDependencies = generateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Required, Create, nodePoolOptionSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_1", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.22.0/24`}, "dns_label": Representation{repType: Required, create: `nodepool1`}})) +
		generateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_2", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.23.0/24`}, "dns_label": Representation{repType: Required, create: `nodepool2`}})) +
		generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.20.0/24`}, "dns_label": Representation{repType: Required, create: `cluster1`}})) +
		generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.21.0/24`}, "dns_label": Representation{repType: Required, create: `cluster2`}})) +
		AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		}))
)

func TestContainerengineNodePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pools"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckContainerengineNodePoolDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Create, nodePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
					resource.TestCheckResourceAttrSet(resourceName, "node_image_name"),
					resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, getUpdatedRepresentationCopy("node_metadata", Representation{repType: Optional, update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
					resource.TestCheckResourceAttrSet(resourceName, "node_image_name"),
					resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

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
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools", Optional, Update, nodePoolDataSourceRepresentation) +
					compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

					resource.TestCheckResourceAttr(datasourceName, "node_pools.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.kubernetes_version"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_name"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_source.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.subnet_ids.#", "2"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NodePoolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_source.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "2"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceConfig,
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

func testAccCheckContainerengineNodePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).containerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_node_pool" {
			noResourceFound = false
			request := oci_containerengine.GetNodePoolRequest{}

			tmp := rs.Primary.ID
			request.NodePoolId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "containerengine")

			_, err := client.GetNodePool(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !inSweeperExcludeList("ContainerengineNodePool") {
		resource.AddTestSweepers("ContainerengineNodePool", &resource.Sweeper{
			Name:         "ContainerengineNodePool",
			Dependencies: DependencyGraph["nodePool"],
			F:            sweepContainerengineNodePoolResource,
		})
	}
}

func sweepContainerengineNodePoolResource(compartment string) error {
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()
	nodePoolIds, err := getNodePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, nodePoolId := range nodePoolIds {
		if ok := SweeperDefaultResourceId[nodePoolId]; !ok {
			deleteNodePoolRequest := oci_containerengine.DeleteNodePoolRequest{}

			deleteNodePoolRequest.NodePoolId = &nodePoolId

			deleteNodePoolRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteNodePool(context.Background(), deleteNodePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting NodePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", nodePoolId, error)
				continue
			}
		}
	}
	return nil
}

func getNodePoolIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "NodePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()

	listNodePoolsRequest := oci_containerengine.ListNodePoolsRequest{}
	listNodePoolsRequest.CompartmentId = &compartmentId
	listNodePoolsResponse, err := containerEngineClient.ListNodePools(context.Background(), listNodePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NodePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, nodePool := range listNodePoolsResponse.Items {
		id := *nodePool.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "NodePoolId", id)
	}
	return resourceIds, nil
}
