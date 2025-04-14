// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	nodePoolCyclingRepresentation = map[string]interface{}{
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
		"node_pool_cycling_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodePoolNodePoolCyclingDetailsRepresentation},
		"ssh_public_key":                   acctest.Representation{RepType: acctest.Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`},
		"node_config_details":              acctest.RepresentationGroup{RepType: acctest.Required, Group: nodeConfigDetailsCyclingRepresentation},
		"node_eviction_node_pool_settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nodeEvictionNodePoolSettingsCyclingRepresentation},
	}

	nodeConfigDetailsCyclingRepresentation = map[string]interface{}{
		"placement_configs":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: placementConfigsCyclingRepresentation},
		"size":                                acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
	}

	placementConfigsCyclingRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.nodePool_Subnet_1.id}`},
	}

	nodePoolNodePoolCyclingDetailsRepresentation = map[string]interface{}{
		"cycle_modes":             acctest.Representation{RepType: acctest.Optional, Create: []string{"INSTANCE_REPLACE"}, Update: []string{"BOOT_VOLUME_REPLACE"}},
		"is_node_cycling_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"maximum_surge":           acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"maximum_unavailable":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	nodeEvictionNodePoolSettingsCyclingRepresentation = map[string]interface{}{
		"eviction_grace_duration":              acctest.Representation{RepType: acctest.Optional, Create: `PT1H`, Update: `PT50M`},
		"is_force_action_after_grace_duration": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_force_delete_after_grace_duration": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	nodePoolResourceCyclingConfig = ContainerengineNodePoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolCyclingRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolCycling_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolCycling_basic")
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
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolCyclingRepresentation), "containerengine", "nodePool", t)

	acctest.ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Create, nodePoolCyclingRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_action_after_grace_duration", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.cycle_modes.#", "1"),
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
			Config: config + compartmentIdVariableStr + ContainerengineNodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("node_metadata", acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolCyclingRepresentation)),
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
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_action_after_grace_duration", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_name"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.cycle_modes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.is_node_cycling_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_pool_cycling_details.0.maximum_surge", "1"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", acctest.Optional, acctest.Update, nodePoolCyclingRepresentation),
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
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_eviction_node_pool_settings.0.is_force_action_after_grace_duration", "false"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_name"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.cycle_modes.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.is_node_cycling_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_pool_cycling_details.0.maximum_surge", "1"),
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
					ContainerengineContainerengineNodePoolSingularDataSourceRepresentation) + nodePoolResourceConfigForVMStandard + compartmentIdVariableStr + nodePoolResourceCyclingConfig,
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
				resource.TestCheckResourceAttr(singularDatasourceName, "node_eviction_node_pool_settings.0.is_force_action_after_grace_duration", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_eviction_node_pool_settings.0.is_force_delete_after_grace_duration", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.cycle_modes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.is_node_cycling_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_pool_cycling_details.0.maximum_surge", "1"),
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
