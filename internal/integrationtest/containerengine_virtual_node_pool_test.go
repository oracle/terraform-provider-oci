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

const (
	ContainerengineDefinedTagsDependencies = `

resource "oci_identity_tag" "tag2" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}
`
)

var (
	ContainerengineVirtualNodePoolRequiredOnlyResource = ContainerengineVirtualNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Required, acctest.Create, ContainerengineVirtualNodePoolRepresentation)

	ContainerengineVirtualNodePoolResourceConfig = ContainerengineVirtualNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Optional, acctest.Update, ContainerengineVirtualNodePoolRepresentation)

	ContainerengineVirtualNodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_node_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_virtual_node_pool.test_virtual_node_pool.id}`},
	}
	// Virtual Node Pools
	ContainerengineVirtualNodePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolDataSourceFilterRepresentation}}
	ContainerengineVirtualNodePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_virtual_node_pool.test_virtual_node_pool.id}`}},
	}
	ContainerengineVirtualNodePoolRepresentation = map[string]interface{}{
		"cluster_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"placement_configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolPlacementConfigurationsRepresentation},
		//  We have determined that these arguments do not successfully update in Integ, this is on the server side, so we will create a bug and resolve separately
		//	"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//	"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//	"nsg_ids":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"initial_virtual_node_labels": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolInitialVirtualNodeLabelsRepresentation},
		"pod_configuration":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolPodConfigurationRepresentation},
		"size":                        acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"taints":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolTaintsRepresentation},
		"virtual_node_tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolVirtualNodeTagsRepresentation},
	}
	ContainerengineVirtualNodePoolPlacementConfigurationsRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`, Update: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"fault_domain":        acctest.Representation{RepType: acctest.Required, Create: []string{`FAULT-DOMAIN-2`}, Update: []string{`FAULT-DOMAIN-2`}},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	ContainerengineVirtualNodePoolInitialVirtualNodeLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	ContainerengineVirtualNodePoolPodConfigurationRepresentation = map[string]interface{}{
		"shape":     acctest.Representation{RepType: acctest.Required, Create: `Pod.Standard.E4.Flex`, Update: `Pod.Standard.E3.Flex`},
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
	}
	ContainerengineVirtualNodePoolTaintsRepresentation = map[string]interface{}{
		"effect": acctest.Representation{RepType: acctest.Optional, Create: `NoSchedule`},
		"key":    acctest.Representation{RepType: acctest.Optional, Create: `taint1`, Update: `taint2`},
		"value":  acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}
	ContainerengineVirtualNodePoolVirtualNodeTagsRepresentation = map[string]interface{}{
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	// Cluster
	ContainerengineVirtualNodePoolClusterRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`, Update: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"vcn_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"cluster_pod_network_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolClusterPodNetworkOptionsRepresentation},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"endpoint_config":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolClusterEndpointConfigRepresentation},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_policy_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolClusterImagePolicyConfigRepresentation},
		"kms_key_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"options":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolClusterOptionsAddOnsRepresentation},
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `ENHANCED_CLUSTER`, Update: `ENHANCED_CLUSTER`},
	}
	ContainerengineVirtualNodePoolClusterPodNetworkOptionsRepresentation = map[string]interface{}{
		"cni_type": acctest.Representation{RepType: acctest.Required, Create: `OCI_VCN_IP_NATIVE`},
	}
	ContainerengineVirtualNodePoolClusterEndpointConfigRepresentation = map[string]interface{}{
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	ContainerengineVirtualNodePoolClusterImagePolicyConfigRepresentation = map[string]interface{}{
		"is_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineVirtualNodePoolClusterImagePolicyConfigKeyDetailsRepresentation},
	}
	ContainerengineVirtualNodePoolClusterImagePolicyConfigKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
	}
	ContainerengineVirtualNodePoolClusterOptionsAddOnsRepresentation = map[string]interface{}{
		"is_kubernetes_dashboard_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_tiller_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	ContainerengineVirtualNodePoolClusterOptionSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_option_id": acctest.Representation{RepType: acctest.Required, Create: `all`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
	// Network
	// Ingress
	ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesAllRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
	ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`}, // ICMP
		"source":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"source_type":  acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"stateless":    acctest.Representation{RepType: acctest.Required, Create: `false`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreSecurityListIngressSecurityRulesIcmpOptionsRepresentation},
	}
	ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`}, // TCP
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesTcpOptionsRepresentation},
	}
	// Egress
	ContainerengineVirtualNodePoolSecurityListEgressSecurityRulesAllRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"destination_type": acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless":        acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
	// TCP Options
	ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               acctest.Representation{RepType: acctest.Required, Create: `6443`},
		"min":               acctest.Representation{RepType: acctest.Required, Create: `6443`},
		"source_port_range": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreSecurityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	// Resource Dependencies
	ContainerengineVirtualNodePoolResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineVirtualNodePoolClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
			"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesAllRepresentation}, {RepType: acctest.Required, Group: ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesICMPRepresentation}, {RepType: acctest.Required, Group: ContainerengineVirtualNodePoolSecurityListIngressSecurityRulesTCPRepresentation}},
			"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineVirtualNodePoolSecurityListEgressSecurityRulesAllRepresentation}},
		})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineVirtualNodePoolClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies +
		ContainerengineDefinedTagsDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineVirtualNodePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineVirtualNodePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_virtual_node_pool.test_virtual_node_pool"
	datasourceName := "data.oci_containerengine_virtual_node_pools.test_virtual_node_pools"
	singularDatasourceName := "data.oci_containerengine_virtual_node_pool.test_virtual_node_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineVirtualNodePoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Optional, acctest.Create, ContainerengineVirtualNodePoolRepresentation), "containerengine", "virtualNodePool", t)

	acctest.ResourceTest(t, testAccCheckContainerengineVirtualNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Required, acctest.Create, ContainerengineVirtualNodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domain.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.0.shape", "Pod.Standard.E4.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "pod_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Optional, acctest.Create, ContainerengineVirtualNodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				//	resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domain.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.0.shape", "Pod.Standard.E4.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "pod_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "1"),
				resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoSchedule"),
				resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taint1"),
				resource.TestCheckResourceAttr(resourceName, "taints.0.value", "true"),
				resource.TestCheckResourceAttr(resourceName, "virtual_node_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "virtual_node_tags.0.freeform_tags.%", "1"),

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
			Config: config + compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Optional, acctest.Update, ContainerengineVirtualNodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"), // VNPs cannot be moved to a different compartment
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				//	resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_virtual_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.fault_domain.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "pod_configuration.0.shape", "Pod.Standard.E3.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "pod_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "size", "1"),
				resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "taints.0.effect"),
				resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taint2"),
				resource.TestCheckResourceAttr(resourceName, "taints.0.value", "false"),
				resource.TestCheckResourceAttr(resourceName, "virtual_node_tags.#", "1"),
				//	resource.TestCheckResourceAttr(resourceName, "virtual_node_tags.0.freeform_tags.%", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_virtual_node_pools", "test_virtual_node_pools", acctest.Optional, acctest.Update, ContainerengineVirtualNodePoolDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool", acctest.Optional, acctest.Update, ContainerengineVirtualNodePoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.compartment_id", compartmentId),
				//	resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.display_name", "displayName2"),
				//	resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.initial_virtual_node_labels.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.initial_virtual_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.initial_virtual_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.kubernetes_version"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.placement_configurations.0.fault_domain.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.placement_configurations.0.subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.pod_configuration.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.pod_configuration.0.shape", "Pod.Standard.E3.Flex"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.pod_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.size", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.taints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.taints.0.effect", "NoSchedule"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.taints.0.key", "taint2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_node_pools.0.taints.0.value", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_node_pools.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_virtual_node_pool", "test_virtual_node_pool",
					acctest.Required, acctest.Create,
					ContainerengineVirtualNodePoolSingularDataSourceRepresentation) + compartmentIdVariableStr + ContainerengineVirtualNodePoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_node_pool_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				//	resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_virtual_node_labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_virtual_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_virtual_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "placement_configurations.0.fault_domain.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "placement_configurations.0.subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pod_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pod_configuration.0.shape", "Pod.Standard.E3.Flex"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pod_configuration.0.subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "taints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "taints.0.effect", "NoSchedule"),
				resource.TestCheckResourceAttr(singularDatasourceName, "taints.0.key", "taint2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "taints.0.value", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineVirtualNodePoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineVirtualNodePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_virtual_node_pool" {
			noResourceFound = false
			request := oci_containerengine.GetVirtualNodePoolRequest{}

			tmp := rs.Primary.ID
			request.VirtualNodePoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetVirtualNodePool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.VirtualNodePoolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineVirtualNodePool") {
		resource.AddTestSweepers("ContainerengineVirtualNodePool", &resource.Sweeper{
			Name:         "ContainerengineVirtualNodePool",
			Dependencies: acctest.DependencyGraph["virtualNodePool"],
			F:            sweepContainerengineVirtualNodePoolResource,
		})
	}
}

func sweepContainerengineVirtualNodePoolResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	virtualNodePoolIds, err := getContainerengineVirtualNodePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualNodePoolId := range virtualNodePoolIds {
		if ok := acctest.SweeperDefaultResourceId[virtualNodePoolId]; !ok {
			deleteVirtualNodePoolRequest := oci_containerengine.DeleteVirtualNodePoolRequest{}

			deleteVirtualNodePoolRequest.VirtualNodePoolId = &virtualNodePoolId

			deleteVirtualNodePoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteVirtualNodePool(context.Background(), deleteVirtualNodePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualNodePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualNodePoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &virtualNodePoolId, ContainerengineVirtualNodePoolSweepWaitCondition, time.Duration(20*time.Minute),
				ContainerengineVirtualNodePoolSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineVirtualNodePoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VirtualNodePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listVirtualNodePoolsRequest := oci_containerengine.ListVirtualNodePoolsRequest{}
	listVirtualNodePoolsRequest.CompartmentId = &compartmentId
	listVirtualNodePoolsRequest.LifecycleState = []oci_containerengine.VirtualNodePoolLifecycleStateEnum{oci_containerengine.VirtualNodePoolLifecycleStateNeedsAttention}
	listVirtualNodePoolsResponse, err := containerEngineClient.ListVirtualNodePools(context.Background(), listVirtualNodePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualNodePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualNodePool := range listVirtualNodePoolsResponse.Items {
		id := *virtualNodePool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VirtualNodePoolId", id)
	}
	return resourceIds, nil
}

func ContainerengineVirtualNodePoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 0 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualNodePoolResponse, ok := response.Response.(oci_containerengine.GetVirtualNodePoolResponse); ok {
		return virtualNodePoolResponse.LifecycleState != oci_containerengine.VirtualNodePoolLifecycleStateDeleted
	}
	return false
}

func ContainerengineVirtualNodePoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetVirtualNodePool(context.Background(), oci_containerengine.GetVirtualNodePoolRequest{
		VirtualNodePoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
