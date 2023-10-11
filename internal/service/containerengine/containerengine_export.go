package containerengine

import (
	"fmt"
	"log"

	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportContainerengineAddonHints.GetIdFn = getContainerengineAddonId
	exportContainerengineClusterWorkloadMappingHints.GetIdFn = getContainerengineClusterWorkloadMappingId
	exportContainerengineNodePoolHints.ProcessDiscoveredResourcesFn = processContainerengineNodePool
	tf_export.RegisterCompartmentGraphs("containerengine", containerengineResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processContainerengineNodePool(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, nodePool := range resources {
		// subnet_ids and quantity_per_subnet are deprecated and conflict with node_config_details
		if _, exists := nodePool.SourceAttributes["node_config_details"]; exists {
			if _, ok := nodePool.SourceAttributes["subnet_ids"]; ok {
				delete(nodePool.SourceAttributes, "subnet_ids")
			}
			if _, ok := nodePool.SourceAttributes["quantity_per_subnet"]; ok {
				delete(nodePool.SourceAttributes, "quantity_per_subnet")
			}
		}
	}
	return resources, nil
}

func getContainerengineAddonId(resource *tf_export.OCIResource) (string, error) {

	addonName, ok := resource.SourceAttributes["addon_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find addonName for Containerengine Addon")
	}
	clusterId := resource.Parent.Id
	return GetAddonCompositeId(addonName, clusterId), nil
}

func getContainerengineClusterWorkloadMappingId(resource *tf_export.OCIResource) (string, error) {

	clusterId := resource.Parent.Id
	workloadMappingId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find workloadMappingId for Containerengine ClusterWorkloadMapping")
	}
	return GetClusterWorkloadMappingCompositeId(clusterId, workloadMappingId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportContainerengineClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_containerengine_cluster",
	DatasourceClass:        "oci_containerengine_clusters",
	DatasourceItemsAttr:    "clusters",
	ResourceAbbreviation:   "cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.ClusterLifecycleStateActive),
	},
}

var exportContainerengineNodePoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_containerengine_node_pool",
	DatasourceClass:        "oci_containerengine_node_pools",
	DatasourceItemsAttr:    "node_pools",
	ResourceAbbreviation:   "node_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.NodePoolLifecycleStateActive),
		string(oci_containerengine.NodePoolLifecycleStateNeedsAttention),
	},
}

var exportContainerengineVirtualNodePoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_containerengine_virtual_node_pool",
	DatasourceClass:        "oci_containerengine_virtual_node_pools",
	DatasourceItemsAttr:    "virtual_node_pools",
	ResourceAbbreviation:   "virtual_node_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.VirtualNodePoolLifecycleStateActive),
		string(oci_containerengine.VirtualNodePoolLifecycleStateNeedsAttention),
	},
}

var exportContainerengineAddonHints = &tf_export.TerraformResourceHints{
	GetIdFn:                 getContainerengineAddonId,
	ResourceClass:           "oci_containerengine_addon",
	DatasourceClass:         "oci_containerengine_addons",
	DatasourceItemsAttr:     "addons",
	ResourceAbbreviation:    "addon",
	RequireResourceRefresh:  true,
	FindResourcesOverrideFn: findAddonResourcesOverride,
	DefaultValuesForMissingAttributes: map[string]interface{}{
		"remove_addon_resources_on_delete": true,
	},
	ProcessDiscoveredResourcesFn: processContainerengineAddon,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.AddonLifecycleStateActive),
		string(oci_containerengine.AddonLifecycleStateNeedsAttention),
	},
}

var exportContainerengineClusterWorkloadMappingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_containerengine_cluster_workload_mapping",
	DatasourceClass:      "oci_containerengine_cluster_workload_mappings",
	DatasourceItemsAttr:  "workload_mappings",
	ResourceAbbreviation: "cluster_workload_mapping",
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.WorkloadMappingLifecycleStateActive),
	},
}

var containerengineResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportContainerengineClusterHints},
		{TerraformResourceHints: exportContainerengineNodePoolHints},
		{TerraformResourceHints: exportContainerengineVirtualNodePoolHints},
	},
	"oci_containerengine_cluster": {
		{
			TerraformResourceHints: exportContainerengineAddonHints,
			DatasourceQueryParams: map[string]string{
				"cluster_id": "id",
			},
		},
		{
			TerraformResourceHints: exportContainerengineClusterWorkloadMappingHints,
			DatasourceQueryParams: map[string]string{
				"cluster_id": "id",
			},
		},
	},
}

func findAddonResourcesOverride(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	//when try to listAddons for a non-enhanced cluster, server return 400 addon-management not supported, this override export resources by clusterId function to return empty list instead
	if clusterType, ok := parent.SourceAttributes["type"]; ok {
		if clusterType != "ENHANCED_CLUSTER" {
			log.Printf("[DEBUG] cluster is not enhanced cluster, skip listAddons and return empty addons: %v\n", parent.Id)
			results := []*tf_export.OCIResource{}
			return results, nil
		}
	}

	return tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
}

func processContainerengineAddon(context *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	//since `remove_addon_resources_on_delete` is required field but not part of resource attribute, so it is not returned by the list/get API, populate default value true here to make sure import resource work.
	for _, resource := range resources {
		if _, ok := resource.SourceAttributes["remove_addon_resources_on_delete"]; !ok {
			resource.SourceAttributes["remove_addon_resources_on_delete"] = true
		}
	}

	return resources, nil
}
