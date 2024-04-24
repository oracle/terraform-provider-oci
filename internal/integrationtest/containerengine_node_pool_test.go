// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NodePoolRequiredOnlyResource = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, nodePoolRepresentation)

	ContainerengineNodePoolResourceConfig = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolRepresentation)

	ContainerengineContainerengineNodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_node_pool.test_node_pool.id}`},
	}

	ContainerengineContainerengineNodePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePoolDataSourceFilterRepresentation}}
	nodePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool.id}`}},
	}

	nodePoolRepresentation = map[string]interface{}{
		"cluster_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":               acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                             acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"node_image_name":                  acctest.Representation{RepType: acctest.Required, Create: `Oracle-Linux-7.6`},
		"node_shape":                       acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`, Update: `VM.Standard2.1`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"initial_node_labels":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolInitialNodeLabelsRepresentation},
		"node_metadata":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"node_pool_cycling_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineNodePoolNodePoolCyclingDetailsRepresentation},
		"ssh_public_key":                   acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
		"node_config_details":              acctest.RepresentationGroup{RepType: acctest.Required, Group: nodeConfigDetailsRepresentation},
		"node_eviction_node_pool_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodePoolNodeEvictionNodePoolSettingsRepresentation},
	}

	nodeConfigDetailsRepresentation = map[string]interface{}{
		"placement_configs":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: placementConfigsRepresentation},
		"size":                                acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	placementConfigsRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.nodePool_Subnet_1.id}`},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Update: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
	}

	nodePreemptibleNodeConfigRepresentation = map[string]interface{}{
		"preemption_action": acctest.RepresentationGroup{RepType: acctest.Required, Group: nodePreemptibleNodeConfigPreemptionActionRepresentation},
	}
	nodePreemptibleNodeConfigPreemptionActionRepresentation = map[string]interface{}{
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `TERMINATE`},
		"is_preserve_boot_volume": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	ContainerengineNodePoolInitialNodeLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	nodePoolNodeEvictionNodePoolSettingsRepresentation = map[string]interface{}{
		"eviction_grace_duration":              acctest.Representation{RepType: acctest.Optional, Create: `PT1H`, Update: `PT50M`},
		"is_force_delete_after_grace_duration": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ContainerengineNodePoolNodePoolCyclingDetailsRepresentation = map[string]interface{}{
		"is_node_cycling_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"maximum_surge":           acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"maximum_unavailable":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	ContainerengineRouteTableRouteRulesforNodePoolRepresentation = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
	}

	ContainerengineSecurityListIngressSecurityRulesALLforNodePoolRepresentation = map[string]interface{}{
		"source":           acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ContainerengineSecurityListIngressSecurityRulesICMPforNodePoolRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreSecurityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ContainerengineSecurityListEgressSecurityRulesAllforNodePoolRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ContainerengineSecurityListIngressSecurityRulesTCPforNodePoolRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityListIngressSecurityRulesTcpOptionsforNodePoolRepresentation},
	}

	securityListIngressSecurityRulesTcpOptionsforNodePoolRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"min":               acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreSecurityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}

	ContainerengineSecurityListIngressSecurityRulesICMP2forNodePoolRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreSecurityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	ContainerengineNodePoolResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", acctest.Required, acctest.Create, ContainerengineContainerengineNodePoolOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreRouteTableRepresentation, map[string]interface{}{
			"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineRouteTableRouteRulesforNodePoolRepresentation},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
			"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesICMP2forNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesALLforNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesICMPforNodePoolRepresentation}, {RepType: acctest.Required, Group: ContainerengineSecurityListIngressSecurityRulesTCPforNodePoolRepresentation}},
			"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineSecurityListEgressSecurityRulesAllforNodePoolRepresentation}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.22.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.23.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `nodepool2`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreComputeCapacityReservationRepresentation, map[string]interface{}{
			"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`}, "instance_reservation_configs": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(CoreComputeCapacityReservationInstanceReservationConfigsRepresentation, map[string]interface{}{
				"instance_shape": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`}, "fault_domain": acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`}, "reserved_count": acctest.Representation{RepType: acctest.Required, Create: `6`}, "cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
			})}}))
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineNodePoolResourceDependencies+nodePoolResourceConfigForVMStandard+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolRepresentation), "containerengine", "nodePool", t)

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Required, acctest.Create, nodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"capacity_reservation_id"}),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.eviction_grace_duration", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.is_node_cycling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.maximum_surge", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.maximum_unavailable", "0"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "node_config_details.0.placement_configs", nil, []string{"capacity_reservation_id"}),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.size", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.is_pv_encryption_in_transit_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.eviction_grace_duration", "PT50M"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_name"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.is_node_cycling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.maximum_surge", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.maximum_unavailable", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.size", "2"),
				resource.TestCheckResourceAttr(resourceName, "node_config_details.0.is_pv_encryption_in_transit_enabled", "true"),

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
				compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
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
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_eviction_node_pool_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_eviction_node_pool_settings.0.eviction_grace_duration", "PT50M"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_name"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.is_node_cycling_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.maximum_surge", "2"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.maximum_unavailable", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_source.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool",
					acctest.Required, acctest.Create,
					ContainerengineContainerengineNodePoolSingularDataSourceRepresentation) + nodePoolResourceConfigForVMStandard + compartmentIdVariableStr + ContainerengineNodePoolResourceConfig,
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
				resource.TestCheckResourceAttr(singularDatasourceName, "node_eviction_node_pool_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_eviction_node_pool_settings.0.eviction_grace_duration", "PT50M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.is_node_cycling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.maximum_surge", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.maximum_unavailable", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
			),
		},
		// verify resource import
		{
			Config:                  config + NodePoolRequiredOnlyResource,
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

			response, err := client.GetNodePool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.NodePoolLifecycleStateDeleted): true,
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &nodePoolId, nodePoolSweepWaitCondition, time.Duration(3*time.Minute),
				nodePoolSweepResponseFetchOperation, "containerengine", true)
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
	listNodePoolsRequest.LifecycleState = []oci_containerengine.NodePoolLifecycleStateEnum{oci_containerengine.NodePoolLifecycleStateNeedsAttention}
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

func nodePoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if nodePoolResponse, ok := response.Response.(oci_containerengine.GetNodePoolResponse); ok {
		return nodePoolResponse.LifecycleState != oci_containerengine.NodePoolLifecycleStateDeleted
	}
	return false
}

func nodePoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetNodePool(context.Background(), oci_containerengine.GetNodePoolRequest{
		NodePoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
