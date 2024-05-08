package cloud_guard

import (
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("cloud_guard", cloudGuardResourceGraph)
	tf_export.RegisterTenancyGraphs("cloud_guard_tenancy", cloudGuardTenancyResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportCloudGuardTargetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_target",
	DatasourceClass:        "oci_cloud_guard_targets",
	DatasourceItemsAttr:    "target_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "target",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardManagedListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_managed_list",
	DatasourceClass:        "oci_cloud_guard_managed_lists",
	DatasourceItemsAttr:    "managed_list_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "managed_list",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardResponderRecipeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_responder_recipe",
	DatasourceClass:        "oci_cloud_guard_responder_recipes",
	DatasourceItemsAttr:    "responder_recipe_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "responder_recipe",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardDataMaskRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_data_mask_rule",
	DatasourceClass:        "oci_cloud_guard_data_mask_rules",
	DatasourceItemsAttr:    "data_mask_rule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_mask_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardDetectorRecipeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_detector_recipe",
	DatasourceClass:        "oci_cloud_guard_detector_recipes",
	DatasourceItemsAttr:    "detector_recipe_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "detector_recipe",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardSecurityRecipeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_security_recipe",
	DatasourceClass:        "oci_cloud_guard_security_recipes",
	DatasourceItemsAttr:    "security_recipe_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "security_recipe",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardSecurityZoneHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_security_zone",
	DatasourceClass:        "oci_cloud_guard_security_zones",
	DatasourceItemsAttr:    "security_zone_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "security_zone",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardDataSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_data_source",
	DatasourceClass:        "oci_cloud_guard_data_sources",
	DatasourceItemsAttr:    "data_source_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardWlpAgentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_wlp_agent",
	DatasourceClass:        "oci_cloud_guard_wlp_agents",
	DatasourceItemsAttr:    "wlp_agent_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "wlp_agent",
	RequireResourceRefresh: true,
}

var exportCloudGuardAdhocQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_adhoc_query",
	DatasourceClass:        "oci_cloud_guard_adhoc_queries",
	DatasourceItemsAttr:    "adhoc_query_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "adhoc_query",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var exportCloudGuardSavedQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_guard_saved_query",
	DatasourceClass:        "oci_cloud_guard_saved_queries",
	DatasourceItemsAttr:    "saved_query_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "saved_query",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_guard.LifecycleStateActive),
	},
}

var cloudGuardResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCloudGuardTargetHints},
		{TerraformResourceHints: exportCloudGuardManagedListHints},
		{TerraformResourceHints: exportCloudGuardResponderRecipeHints},
		{TerraformResourceHints: exportCloudGuardDetectorRecipeHints},
		{TerraformResourceHints: exportCloudGuardSecurityRecipeHints},
		{TerraformResourceHints: exportCloudGuardSecurityZoneHints},
		{TerraformResourceHints: exportCloudGuardDataSourceHints},
		{TerraformResourceHints: exportCloudGuardWlpAgentHints},
		{TerraformResourceHints: exportCloudGuardAdhocQueryHints},
		{TerraformResourceHints: exportCloudGuardSavedQueryHints},
	},
}

var cloudGuardTenancyResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportCloudGuardDataMaskRuleHints},
	},
}
