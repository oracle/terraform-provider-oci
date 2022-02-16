// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v58/containerengine"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	NodePoolRequiredOnlyResource = NodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, nodePoolRepresentation)

	NodePoolResourceConfig = NodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRepresentation)

	nodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_node_pool.test_node_pool.id}`},
	}

	nodePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolDataSourceFilterRepresentation}}
	nodePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool.id}`}},
	}

	nodePoolRepresentation = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_name":     acctest.Representation{RepType: acctest.Required, Create: `Oracle-Linux-7.6`},
		"node_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"subnet_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"quantity_per_subnet": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"ssh_public_key":      acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreNodePoolSystemTagsChangesRep},
	}
	nodePoolInitialNodeLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	ignoreNodePoolSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`}},
	}

	routeTableRouteRulesforNodePoolRepresentation = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
	}

	securityListIngressSecurityRulesALLforNodePoolRepresentation = map[string]interface{}{
		"source":           acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	securityListIngressSecurityRulesICMPforNodePoolRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	securityListEgressSecurityRulesAllforNodePoolRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	securityListIngressSecurityRulesTCPforNodePoolRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesTcpOptionsforNodePoolRepresentation},
	}

	securityListIngressSecurityRulesTcpOptionsforNodePoolRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}

	securityListIngressSecurityRulesICMP2forNodePoolRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	NodePoolResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", acctest.Required, acctest.Create, nodePoolOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(routeTableRepresentation, map[string]interface{}{
			"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: routeTableRouteRulesforNodePoolRepresentation},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(securityListRepresentation, map[string]interface{}{
			"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: securityListIngressSecurityRulesICMP2forNodePoolRepresentation}, {RepType: acctest.Required, Group: securityListIngressSecurityRulesALLforNodePoolRepresentation}, {RepType: acctest.Required, Group: securityListIngressSecurityRulesICMPforNodePoolRepresentation}, {RepType: acctest.Required, Group: securityListIngressSecurityRulesTCPforNodePoolRepresentation}},
			"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: securityListEgressSecurityRulesAllforNodePoolRepresentation}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.22.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.23.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool2`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, clusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, clusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pools"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NodePoolResourceDependencies+nodePoolResourceConfigForVMStandard+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolRepresentation), "containerengine", "nodePool", t)

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, nodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

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
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools", acctest.Optional, acctest.Update, nodePoolDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

				resource.TestCheckResourceAttr(datasourceName, "node_pools.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.freeform_tags.%", "1"),
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
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool",
					acctest.Required, acctest.Create,
					nodePoolSingularDataSourceRepresentation) + nodePoolResourceConfigForVMStandard + compartmentIdVariableStr + NodePoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
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
	})
}

func testAccCheckContainerengineNodePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_node_pool" {
			noResourceFound = false
			request := oci_containerengine.GetNodePoolRequest{}

			tmp := rs.Primary.ID
			request.NodePoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ContainerengineNodePool") {
		resource.AddTestSweepers("ContainerengineNodePool", &resource.Sweeper{
			Name:         "ContainerengineNodePool",
			Dependencies: acctest.DependencyGraph["nodePool"],
			F:            sweepContainerengineNodePoolResource,
		})
	}
}

func sweepContainerengineNodePoolResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	nodePoolIds, err := getNodePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, nodePoolId := range nodePoolIds {
		if ok := acctest.SweeperDefaultResourceId[nodePoolId]; !ok {
			deleteNodePoolRequest := oci_containerengine.DeleteNodePoolRequest{}

			deleteNodePoolRequest.NodePoolId = &nodePoolId

			deleteNodePoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
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
	ids := acctest.GetResourceIdsToSweep(compartment, "NodePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listNodePoolsRequest := oci_containerengine.ListNodePoolsRequest{}
	listNodePoolsRequest.CompartmentId = &compartmentId
	listNodePoolsResponse, err := containerEngineClient.ListNodePools(context.Background(), listNodePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NodePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, nodePool := range listNodePoolsResponse.Items {
		id := *nodePool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NodePoolId", id)
	}
	return resourceIds, nil
}
