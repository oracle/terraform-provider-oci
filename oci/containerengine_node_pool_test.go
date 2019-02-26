// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

var (
	NodePoolRequiredOnlyResource = NodePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation)

	NodePoolResourceConfig = NodePoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation)

	nodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_id": Representation{repType: Required, create: `${oci_containerengine_node_pool.test_node_pool.id}`},
	}

	nodePoolSingularDataSourceRepresentationForImageId = map[string]interface{}{
		"node_pool_id": Representation{repType: Required, create: `${oci_containerengine_node_pool.test_node_pool_imageId.id}`},
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

	nodePoolDataSourceRepresentationForImageId = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cluster_id":     Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           Representation{repType: Optional, create: `name`, update: `name2`},
		"filter":         RepresentationGroup{Required, nodePoolDataSourceFilterRepresentationForImageId}}
	nodePoolDataSourceFilterRepresentationForImageId = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_containerengine_node_pool.test_node_pool_imageId.id}`}},
	}

	nodePoolRepresentation = map[string]interface{}{
		"cluster_id":          Representation{repType: Required, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"kubernetes_version":  Representation{repType: Required, create: `${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions.0}`},
		"name":                Representation{repType: Required, create: `name`, update: `name2`},
		"node_image_name":     Representation{repType: Required, create: `Oracle-Linux-7.4`},
		"node_shape":          Representation{repType: Required, create: `VM.Standard2.1`},
		"subnet_ids":          Representation{repType: Required, create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": RepresentationGroup{Optional, nodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       Representation{repType: Optional, create: map[string]string{"nodeMetadata": "nodeMetadata"}, update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"quantity_per_subnet": Representation{repType: Optional, create: `1`, update: `2`},
		"ssh_public_key":      Representation{repType: Optional, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}
	nodePoolRepresentationForImageId = map[string]interface{}{
		"cluster_id":          Representation{repType: Required, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"kubernetes_version":  Representation{repType: Required, create: `${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions.0}`},
		"name":                Representation{repType: Required, create: `name`, update: `name2`},
		"node_image_id":       Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"node_shape":          Representation{repType: Required, create: `VM.Standard2.1`},
		"subnet_ids":          Representation{repType: Required, create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": RepresentationGroup{Optional, nodePoolInitialNodeLabelsRepresentation},
		"quantity_per_subnet": Representation{repType: Optional, create: `1`, update: `2`},
		"ssh_public_key":      Representation{repType: Optional, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}
	nodePoolInitialNodeLabelsRepresentation = map[string]interface{}{
		"key":   Representation{repType: Optional, create: `key`, update: `key2`},
		"value": Representation{repType: Optional, create: `value`, update: `value2`},
	}

	// @CODEGEN: OKE does not support regional subnets
	NodePoolResourceDependencies = ClusterRequiredOnlyResource + InstanceCommonVariables +
		generateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_1", Optional, Create,
			getMultipleUpdatedRepresenationCopy(
				[]string{"cidr_block", "dns_label"},
				[]interface{}{Representation{repType: Optional, create: `10.0.22.0/24`}, Representation{repType: Optional, create: `nodepool1`}},
				subnetRepresentation)) +
		generateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_2", Optional, Create,
			getMultipleUpdatedRepresenationCopy(
				[]string{"cidr_block", "dns_label"},
				[]interface{}{Representation{repType: Optional, create: `10.0.23.0/24`}, Representation{repType: Optional, create: `nodepool2`}},
				subnetRepresentation)) +
		generateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Required, Create, nodePoolOptionSingularDataSourceRepresentation)
)

func TestContainerengineNodePoolResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	resourceNameForImageId := "oci_containerengine_node_pool.test_node_pool_imageId"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pools"
	datasourceNameForImageId := "data.oci_containerengine_node_pools.test_node_pools_imageId"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"
	singularDatasourceNameForImageId := "data.oci_containerengine_node_pool.test_node_pool_imageId"

	var resId, resId2, resIdCreatedWithImageId, resId2CreatedWithImageId string

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
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation) +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Required, Create, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Resource created with Image Name
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					//Asserting Resource created with Image Id
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						resIdCreatedWithImageId, err = fromInstanceState(s, resourceNameForImageId, "id")
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
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Create, nodePoolRepresentation) +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Optional, Create, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Resource created with Image Name
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					//Asserting Resource created with Image Id
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.key", "key"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "quantity_per_subnet", "1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						resIdCreatedWithImageId, err = fromInstanceState(s, resourceNameForImageId, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, getUpdatedRepresentationCopy("node_metadata", Representation{repType: Optional, update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRepresentation)) +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Optional, Update, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Resource created with Image Name
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					//Asserting Resource created with Image Id
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						resId2CreatedWithImageId, err = fromInstanceState(s, resourceNameForImageId, "id")
						if resIdCreatedWithImageId != resId2CreatedWithImageId {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify updates to node_metadata should create new resource
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation) +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Optional, Update, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Resource created with Image Name
					resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					//Asserting Resource created with Image Id
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name2"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_name"),
					resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						resId2CreatedWithImageId, err = fromInstanceState(s, resourceNameForImageId, "id")
						if resIdCreatedWithImageId != resId2CreatedWithImageId {
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
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools_imageId", Optional, Update, nodePoolDataSourceRepresentationForImageId) +
					compartmentIdVariableStr + NodePoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation) +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Optional, Update, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Datasource for NodePool created with Image Name
					resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

					resource.TestCheckResourceAttr(datasourceName, "node_pools.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.kubernetes_version"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(datasourceName, "node_pools.0.subnet_ids.#", "2"),
					//Asserting Datasource for NodePool created with Image Id
					resource.TestCheckResourceAttrSet(datasourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "name", "name2"),

					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceNameForImageId, "node_pools.0.cluster_id"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(datasourceNameForImageId, "node_pools.0.kubernetes_version"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.name", "name2"),
					resource.TestCheckResourceAttrSet(datasourceNameForImageId, "node_pools.0.node_image_name"),
					resource.TestCheckResourceAttrSet(datasourceNameForImageId, "node_pools.0.node_image_id"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(datasourceNameForImageId, "node_pools.0.subnet_ids.#", "2"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Required, Create, nodePoolSingularDataSourceRepresentationForImageId) +
					compartmentIdVariableStr + NodePoolResourceConfig +
					generateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", Optional, Update, nodePoolRepresentationForImageId),
				Check: resource.ComposeAggregateTestCheckFunc(
					//Asserting Singular Datasource for NodePool created with Image Name
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_image_name", "Oracle-Linux-7.4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "2"),
					//Asserting Singular Datasource for NodePool created with Image Id
					resource.TestCheckResourceAttrSet(singularDatasourceNameForImageId, "cluster_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceNameForImageId, "node_pool_id"),

					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "initial_node_labels.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "initial_node_labels.0.key", "key2"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "initial_node_labels.0.value", "value2"),
					resource.TestCheckResourceAttrSet(singularDatasourceNameForImageId, "kubernetes_version"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "name", "name2"),
					resource.TestCheckResourceAttrSet(singularDatasourceNameForImageId, "node_image_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceNameForImageId, "node_image_id"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "node_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "quantity_per_subnet", "2"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
					resource.TestCheckResourceAttr(singularDatasourceNameForImageId, "subnet_ids.#", "2"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NodePoolResourceConfig,
			},
			// verify resource import
			{
				Config:                  config + compartmentIdVariableStr + NodePoolResourceConfig,
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
	client := testAccProvider.Meta().(*OracleClients).containerEngineClient
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
