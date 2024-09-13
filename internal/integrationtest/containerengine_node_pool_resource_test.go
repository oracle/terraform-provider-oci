// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NodePoolRegionalRequiredOnlyResource = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, nodePoolRegionalSubnetRepresentation)

	NodePoolRegionalResourceConfig = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRegionalSubnetRepresentation)

	NodePoolNonReginalResourceDependencies = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolNonRegionalSubnetRepresentation)

	nodePoolNonRegionalSubnetRepresentation = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_name":     acctest.Representation{RepType: acctest.Required, Create: `Oracle-Linux-7.6`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
		"quantity_per_subnet": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"subnet_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`}, Update: []string{`${oci_core_subnet.nodePool_Subnet_2.id}`}},
	}

	nodePoolRegionalSubnetRepresentation = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_name":     acctest.Representation{RepType: acctest.Required, Create: `Oracle-Linux-7.6`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"node_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodePoolNodeConfigDetailsRepresentation},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}

	nodePoolRegionalSubnetOnlyUpdateFaultDomainsRepresentation = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_name":     acctest.Representation{RepType: acctest.Required, Create: `Oracle-Linux-7.6`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"node_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodePoolNodeConfigDetailsOnlyUpdateFaultDomainsRepresentation},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}

	nodePoolNodeConfigDetailsRepresentation = map[string]interface{}{
		"placement_configs": acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolNodeConfigDetailsPlacementConfigsRepresentation},
		"size":              acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `4`},
	}

	nodePoolNodeConfigDetailsOnlyUpdateFaultDomainsRepresentation = map[string]interface{}{
		"placement_configs": acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolNodeConfigDetailsPlacementConfigsOnlyUpdateFaultDomainsRepresentation},
		"size":              acctest.Representation{RepType: acctest.Required, Create: `2`},
	}

	nodePoolNodeConfigDetailsPlacementConfigsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`, Update: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.node_pool_regional_subnet_1.id}`, Update: `${oci_core_subnet.node_pool_regional_subnet_2.id}`},
		"fault_domains":       acctest.Representation{RepType: acctest.Optional, Create: []string{"FAULT-DOMAIN-1"}, Update: []string{"FAULT-DOMAIN-1"}},
	}

	nodePoolNodeConfigDetailsPlacementConfigsOnlyUpdateFaultDomainsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.node_pool_regional_subnet_1.id}`},
		"fault_domains":       acctest.Representation{RepType: acctest.Optional, Create: []string{"FAULT-DOMAIN-1"}, Update: []string{"FAULT-DOMAIN-2"}},
	}

	nodePoolSingularDataSourceRepresentationForImageId = map[string]interface{}{
		"node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_node_pool.test_node_pool_imageId.id}`},
	}

	nodePoolSingularDataSourceRepresentationForNodeSourceDetails = map[string]interface{}{
		"node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_node_pool.test_node_pool_node_source_details.id}`},
	}

	nodePoolSingularDataSourceRepresentationForFlexShapes = map[string]interface{}{
		"node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_node_pool.test_node_pool_flexible_shapes.id}`},
	}

	nodePoolDataSourceRepresentationForImageId = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolDataSourceFilterRepresentationForImageId}}
	nodePoolDataSourceFilterRepresentationForImageId = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool_imageId.id}`}},
	}

	nodePoolDataSourceRepresentationForNodeSourceDetails = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolDataSourceFilterRepresentationForNodeSourceDetails}}
	nodePoolDataSourceFilterRepresentationForNodeSourceDetails = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool_node_source_details.id}`}},
	}

	nodePoolDataSourceRepresentationForFlexShapes = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `flexNodePool`, Update: `flexNodePool2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolDataSourceFilterRepresentationForFlexShapes}}
	nodePoolDataSourceFilterRepresentationForFlexShapes = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool_flexible_shapes.id}`}},
	}

	nodePoolRepresentationForImageId = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"quantity_per_subnet": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}
	nodePoolResourceConfigForVMStandard = utils.OciImageIdsVariable

	nodePoolRepresentationForNodeSourceDetails = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"node_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolNodeSourceDetailsRepresentation},
		"quantity_per_subnet": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
	}
	nodePoolResourceConfigForFlexShapes = utils.FlexVmImageIdsVariable
	nodePoolRepresentationForFlexShapes = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `flexNodePool`, Update: `flexNodePool2`},
		"node_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolNodeSourceDetailsRepresentationForFlexShapes},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E3.Flex`},
		"subnet_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"node_shape_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolNodeShapeConfigRepresentation},
		"quantity_per_subnet": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	nodePoolNodeSourceDetailsRepresentationForFlexShapes = map[string]interface{}{
		"image_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
	}
	nodePoolNodeSourceDetailsRepresentation = map[string]interface{}{
		"image_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `image`},
	}

	nodePoolNodeShapeConfigRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `2.0`},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `32.0`, Update: `36.0`},
	}

	NodePoolReginalResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreRouteTableRepresentation, map[string]interface{}{
			"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineRouteTableRouteRulesforNodePoolRepresentation},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
			"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesICMP2forNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesALLforNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesICMPforNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesTCPforNodePoolRepresentation}},
			"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineSecurityListEgressSecurityRulesAllforNodePoolRepresentation}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "node_pool_regional_subnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.24.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "node_pool_regional_subnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.25.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool2`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestResourceContainerengineNodePool_regionalsubnet(t *testing.T) {
	httpreplay.SetScenario("TestResourceContainerengineNodePool_regionalsubnet")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pools"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"availability_domain"}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"subnet_id"}),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.size", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				//resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.0", "FAULT-DOMAIN-1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRegionalSubnetRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"subnet_id"}),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.size", "4"),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.0", "FAULT-DOMAIN-1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools", acctest.Optional, acctest.Update, ContainerengineContainerengineNodePoolDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Datasource for NodePool created with Image Name
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
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_config_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_config_details.0.placement_configs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_config_details.0.placement_configs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_config_details.0.size", "4"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_id"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.subnet_ids.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Singular Datasource for NodePool created with Image Name
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config_details.0.placement_configs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config_details.0.size", "4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "1"),
				// "nodes" is not set until the instances in the node_pool are "Available" so we can't assert the nodes property
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes"),
			),
		},
	})
}

// issue-routing-tag: containerengine/default
func TestResourceContainerengineNodePool_OnlyUpdateFaultDomain(t *testing.T) {
	httpreplay.SetScenario("TestResourceContainerengineNodePool_OnlyUpdateFaultDomain")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolRegionalSubnetOnlyUpdateFaultDomainsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"availability_domain"}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"subnet_id"}),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				//resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.0", "FAULT-DOMAIN-1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRegionalSubnetOnlyUpdateFaultDomainsRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"subnet_id"}),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.size", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.placement_configs.0.fault_domains.0", "FAULT-DOMAIN-2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Singular Datasource for NodePool created with Image Name
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_config_details.0.placement_configs.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "1"),
				// "nodes" is not set until the instances in the node_pool are "Available" so we can't assert the nodes property
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes"),
			),
		},
	})
}

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolResource_image(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_image")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameForImageId := "oci_containerengine_node_pool.test_node_pool_imageId"

	datasourceNameForImageId := "data.oci_containerengine_node_pools.test_node_pools_imageId"

	singularDatasourceNameForImageId := "data.oci_containerengine_node_pool.test_node_pool_imageId"

	var resIdCreatedWithImageId, resId2CreatedWithImageId string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Required, acctest.Create, nodePoolRepresentationForImageId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Id
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),
				//Asserting Resource created with Node Source Details

				func(s *terraform.State) (err error) {
					resIdCreatedWithImageId, err = acctest.FromInstanceState(s, resourceNameForImageId, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Optional, acctest.Create, nodePoolRepresentationForImageId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Id
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "quantity_per_subnet", "1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resIdCreatedWithImageId, err = acctest.FromInstanceState(s, resourceNameForImageId, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Optional, acctest.Update, nodePoolRepresentationForImageId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Id
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceNameForImageId, "node_image_id"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceNameForImageId, "subnet_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId2CreatedWithImageId, err = acctest.FromInstanceState(s, resourceNameForImageId, "id")
					if resIdCreatedWithImageId != resId2CreatedWithImageId {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools_imageId", acctest.Optional, acctest.Update, nodePoolDataSourceRepresentationForImageId) +
				compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Optional, acctest.Update, nodePoolRepresentationForImageId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Required, acctest.Create, nodePoolSingularDataSourceRepresentationForImageId) +
				compartmentIdVariableStr + ContainerengineNodePoolResourceConfig + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_imageId", acctest.Optional, acctest.Update, nodePoolRepresentationForImageId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolResource_nodeSourceDetails(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_nodeSourceDetails")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameForNodeSourceDetails := "oci_containerengine_node_pool.test_node_pool_node_source_details"

	datasourceNameForNodeSourceDetails := "data.oci_containerengine_node_pools.test_node_pools_node_source_details"

	singularDatasourceNameForNodeSourceDetails := "data.oci_containerengine_node_pool.test_node_pool_node_source_details"

	var resIdCreatedWithNodeSourceDetails, resId2CreatedWithNodeSourceDetails string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Required, acctest.Create, nodePoolRepresentationForNodeSourceDetails),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Node Source Details
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_name"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_id"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "subnet_ids.#", "2"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.source_type"),

				func(s *terraform.State) (err error) {
					resIdCreatedWithNodeSourceDetails, err = acctest.FromInstanceState(s, resourceNameForNodeSourceDetails, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Optional, acctest.Create, nodePoolRepresentationForNodeSourceDetails),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Node Source Details
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_id"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_name"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.source_type"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "quantity_per_subnet", "1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "subnet_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resIdCreatedWithNodeSourceDetails, err = acctest.FromInstanceState(s, resourceNameForNodeSourceDetails, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Optional, acctest.Update, nodePoolRepresentationForNodeSourceDetails),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Node Source Details
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_id"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_image_name"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "node_source_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.image_id"),
				resource.TestCheckResourceAttrSet(resourceNameForNodeSourceDetails, "node_source_details.0.source_type"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceNameForNodeSourceDetails, "subnet_ids.#", "2"),
				func(s *terraform.State) (err error) {
					resId2CreatedWithNodeSourceDetails, err = acctest.FromInstanceState(s, resourceNameForNodeSourceDetails, "id")
					if resIdCreatedWithNodeSourceDetails != resId2CreatedWithNodeSourceDetails {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools",
					"test_node_pools_node_source_details", acctest.Optional, acctest.Update, nodePoolDataSourceRepresentationForNodeSourceDetails) + nodePoolResourceConfigForVMStandard +
				compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Optional, acctest.Update, nodePoolRepresentationForNodeSourceDetails),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Datasource for NodePool created with Node Source Details
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "name", "name2"),

				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "node_pools.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "node_pools.0.id"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "node_pools.0.kubernetes_version"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "node_pools.0.node_image_id"),
				resource.TestCheckResourceAttrSet(datasourceNameForNodeSourceDetails, "node_pools.0.node_image_name"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.node_source.#", "1"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(datasourceNameForNodeSourceDetails, "node_pools.0.subnet_ids.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Required, acctest.Create, nodePoolSingularDataSourceRepresentationForNodeSourceDetails) +
				compartmentIdVariableStr + ContainerengineNodePoolResourceConfig + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_node_source_details", acctest.Optional, acctest.Update, nodePoolRepresentationForNodeSourceDetails),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Singular Datasource for NodePool created with Image Name
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "id"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "node_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForNodeSourceDetails, "node_image_name"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "node_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(singularDatasourceNameForNodeSourceDetails, "subnet_ids.#", "2"),
			),
		},
	})
}

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolResource_flexibleShapes(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_flexibleShapes")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameForFlexibleShapes := "oci_containerengine_node_pool.test_node_pool_flexible_shapes"

	datasourceNameForFlexibleShapes := "data.oci_containerengine_node_pools.test_node_pools_flexible_shapes"

	singularDatasourceNameForFlexibleShapes := "data.oci_containerengine_node_pool.test_node_pool_flexible_shapes"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify creation of flex node pool
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForFlexShapes + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_flexible_shapes", acctest.Required, acctest.Create, nodePoolRepresentationForFlexShapes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameForFlexibleShapes, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameForFlexibleShapes, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "name", "flexNodePool"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.0.memory_in_gbs", "32"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameForFlexibleShapes, "id")
					return err
				},
			),
		},

		// verify flex Update
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForFlexShapes +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_flexible_shapes", acctest.Optional, acctest.Update, nodePoolRepresentationForFlexShapes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameForFlexibleShapes, "cluster_id"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameForFlexibleShapes, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "name", "flexNodePool2"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(resourceNameForFlexibleShapes, "node_shape_config.0.memory_in_gbs", "36"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceNameForFlexibleShapes, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools_flexible_shapes", acctest.Optional, acctest.Update, nodePoolDataSourceRepresentationForFlexShapes) +
				compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForFlexShapes +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_flexible_shapes", acctest.Optional, acctest.Update, nodePoolRepresentationForFlexShapes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Datasource for NodePool created with Flexible Shape
				resource.TestCheckResourceAttrSet(datasourceNameForFlexibleShapes, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "name", "flexNodePool2"),

				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceNameForFlexibleShapes, "node_pools.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceNameForFlexibleShapes, "node_pools.0.id"),
				resource.TestCheckResourceAttrSet(datasourceNameForFlexibleShapes, "node_pools.0.kubernetes_version"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.name", "flexNodePool2"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.node_shape", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.node_shape_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.node_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(datasourceNameForFlexibleShapes, "node_pools.0.node_shape_config.0.memory_in_gbs", "36"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_flexible_shapes", acctest.Required, acctest.Create, nodePoolSingularDataSourceRepresentationForFlexShapes) +
				compartmentIdVariableStr + ContainerengineNodePoolResourceConfig + nodePoolResourceConfigForFlexShapes +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_flexible_shapes", acctest.Optional, acctest.Update, nodePoolRepresentationForFlexShapes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Singular Datasource for NodePool created with Flex Shape
				resource.TestCheckResourceAttrSet(singularDatasourceNameForFlexibleShapes, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForFlexibleShapes, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForFlexibleShapes, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceNameForFlexibleShapes, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "name", "flexNodePool2"),
				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "node_shape", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "node_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "node_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(singularDatasourceNameForFlexibleShapes, "node_shape_config.0.memory_in_gbs", "36"),
			),
		},
	})
}

// issue-routing-tag: containerengine/default
func TestResourceContainerengineNodePool_qps(t *testing.T) {
	httpreplay.SetScenario("TestResourceContainerengineNodePool_qps")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool_qps"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pool_qps"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool_qps"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NodePoolNonReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_qps", acctest.Optional, acctest.Create, nodePoolNonRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NodePoolNonReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_qps", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolNonRegionalSubnetRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Resource created with Image Name
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pool_qps", acctest.Optional, acctest.Update, ContainerengineContainerengineNodePoolDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolNonReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_qps", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolNonRegionalSubnetRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Datasource for NodePool created with Image Name
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
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.quantity_per_subnet", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_qps", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolNonReginalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool_qps", acctest.Optional, acctest.Update, nodePoolNonRegionalSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//Asserting Singular Datasource for NodePool created with Image Name
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
				resource.TestCheckResourceAttr(singularDatasourceName, "node_image_name", "Oracle-Linux-7.6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "quantity_per_subnet", "2"),
				// "nodes" is not set until the instances in the node_pool are "Available" so we can't assert the nodes property
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "nodes"),
			),
		},
	})
}
