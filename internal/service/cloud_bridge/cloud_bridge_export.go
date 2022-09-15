package cloud_bridge

import (
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportCloudBridgeAgentPluginHints.GetIdFn = getCloudBridgeAgentPluginId
	tf_export.RegisterCompartmentGraphs("cloud_bridge", cloudBridgeResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getCloudBridgeAgentPluginId(resource *tf_export.OCIResource) (string, error) {

	agentId := resource.Parent.Id
	pluginName := resource.Parent.Id
	return GetAgentPluginCompositeId(agentId, pluginName), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportCloudBridgeAgentPluginHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_cloud_bridge_agent_plugin",
	DatasourceClass:      "oci_cloud_bridge_agent_plugin",
	ResourceAbbreviation: "agent_plugin",
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.PluginLifecycleStateActive),
		string(oci_cloud_bridge.PluginLifecycleStateNeedsAttention),
	},
}

var exportCloudBridgeAgentDependencyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_agent_dependency",
	DatasourceClass:        "oci_cloud_bridge_agent_dependencies",
	DatasourceItemsAttr:    "agent_dependency_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "agent_dependency",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.AgentDependencyLifecycleStateActive),
	},
}

var exportCloudBridgeEnvironmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_environment",
	DatasourceClass:        "oci_cloud_bridge_environments",
	DatasourceItemsAttr:    "environment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "environment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.EnvironmentLifecycleStateActive),
	},
}

var exportCloudBridgeAgentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_agent",
	DatasourceClass:        "oci_cloud_bridge_agents",
	DatasourceItemsAttr:    "agent_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "agent",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.AgentLifecycleStateActive),
	},
}

var exportCloudBridgeAssetSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_asset_source",
	DatasourceClass:        "oci_cloud_bridge_asset_sources",
	DatasourceItemsAttr:    "asset_source_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "asset_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.AssetSourceLifecycleStateActive),
		string(oci_cloud_bridge.AssetSourceLifecycleStateNeedsAttention),
	},
}

var exportCloudBridgeDiscoveryScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_discovery_schedule",
	DatasourceClass:        "oci_cloud_bridge_discovery_schedules",
	DatasourceItemsAttr:    "discovery_schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "discovery_schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.DiscoveryScheduleLifecycleStateActive),
	},
}

var exportCloudBridgeAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_asset",
	DatasourceClass:        "oci_cloud_bridge_assets",
	DatasourceItemsAttr:    "asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.AssetLifecycleStateActive),
	},
}

var exportCloudBridgeInventoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_bridge_inventory",
	DatasourceClass:        "oci_cloud_bridge_inventories",
	DatasourceItemsAttr:    "inventory_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "inventory",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_bridge.InventoryLifecycleStateActive),
	},
}

var cloudBridgeResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCloudBridgeAgentDependencyHints},
		{TerraformResourceHints: exportCloudBridgeEnvironmentHints},
		{TerraformResourceHints: exportCloudBridgeAgentHints},
		{TerraformResourceHints: exportCloudBridgeAssetSourceHints},
		{TerraformResourceHints: exportCloudBridgeDiscoveryScheduleHints},
		{TerraformResourceHints: exportCloudBridgeAssetHints},
		{TerraformResourceHints: exportCloudBridgeInventoryHints},
	},
	"oci_cloud_bridge_agent": {
		{
			TerraformResourceHints: exportCloudBridgeAgentPluginHints,
			DatasourceQueryParams: map[string]string{
				"agent_id":    "id",
				"plugin_name": "name",
			},
		},
	},
}
