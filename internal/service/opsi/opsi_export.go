package opsi

import (
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("opsi", opsiResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOpsiEnterpriseManagerBridgeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_enterprise_manager_bridge",
	DatasourceClass:        "oci_opsi_enterprise_manager_bridges",
	DatasourceItemsAttr:    "enterprise_manager_bridge_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "enterprise_manager_bridge",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	},
}

var exportOpsiDatabaseInsightHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_database_insight",
	DatasourceClass:        "oci_opsi_database_insights",
	DatasourceItemsAttr:    "database_insights_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_insight",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
	},
}

var exportOpsiHostInsightHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_host_insight",
	DatasourceClass:        "oci_opsi_host_insights",
	DatasourceItemsAttr:    "host_insight_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "host_insight",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
	},
}

var exportOpsiExadataInsightHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_exadata_insight",
	DatasourceClass:        "oci_opsi_exadata_insights",
	DatasourceItemsAttr:    "exadata_insight_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "exadata_insight",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.ExadataInsightLifecycleStateActive),
	},
}

var exportOpsiAwrHubHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_awr_hub",
	DatasourceClass:        "oci_opsi_awr_hubs",
	DatasourceItemsAttr:    "awr_hub_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "awr_hub",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.AwrHubLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_operations_insights_warehouse_user",
	DatasourceClass:        "oci_opsi_operations_insights_warehouse_users",
	DatasourceItemsAttr:    "operations_insights_warehouse_user_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "operations_insights_warehouse_user",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.OperationsInsightsWarehouseUserLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_operations_insights_warehouse",
	DatasourceClass:        "oci_opsi_operations_insights_warehouses",
	DatasourceItemsAttr:    "operations_insights_warehouse_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "operations_insights_warehouse",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateActive),
	},
}

var exportOpsiOperationsInsightsWarehouseDownloadWarehouseWalletHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_opsi_operations_insights_warehouse_download_warehouse_wallet",
	ResourceAbbreviation: "operations_insights_warehouse_download_warehouse_wallet",
}

var exportOpsiOperationsInsightsWarehouseRotateWarehouseWalletHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet",
	ResourceAbbreviation: "operations_insights_warehouse_rotate_warehouse_wallet",
}

var exportOpsiOperationsInsightsPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_operations_insights_private_endpoint",
	DatasourceClass:        "oci_opsi_operations_insights_private_endpoints",
	DatasourceItemsAttr:    "operations_insights_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "operations_insights_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateActive),
		string(oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateNeedsAttention),
	},
}

var exportOpsiOpsiConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_opsi_configuration",
	DatasourceClass:        "oci_opsi_opsi_configurations",
	DatasourceItemsAttr:    "opsi_configurations_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "opsi_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.OpsiConfigurationLifecycleStateActive),
	},
}

var exportOpsiNewsReportHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_news_report",
	DatasourceClass:        "oci_opsi_news_reports",
	DatasourceItemsAttr:    "news_report_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "news_report",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	},
}

var exportOpsiAwrHubSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opsi_awr_hub_source",
	DatasourceClass:        "oci_opsi_awr_hub_sources",
	DatasourceItemsAttr:    "awr_hub_source_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "awr_hub_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opsi.AwrHubSourceLifecycleStateActive),
	},
}

var opsiResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOpsiEnterpriseManagerBridgeHints},
		{TerraformResourceHints: exportOpsiDatabaseInsightHints},
		{TerraformResourceHints: exportOpsiHostInsightHints},
		{TerraformResourceHints: exportOpsiExadataInsightHints},
		{TerraformResourceHints: exportOpsiOperationsInsightsWarehouseHints},
		{TerraformResourceHints: exportOpsiOperationsInsightsPrivateEndpointHints},
		{TerraformResourceHints: exportOpsiOpsiConfigurationHints},
		{TerraformResourceHints: exportOpsiNewsReportHints},
	},
	"oci_opsi_awr_hub": {
		{
			TerraformResourceHints: exportOpsiAwrHubSourceHints,
			DatasourceQueryParams: map[string]string{
				"awr_hub_id": "id",
			},
		},
	},
	"oci_opsi_operations_insights_warehouse": {
		{
			TerraformResourceHints: exportOpsiAwrHubHints,
			DatasourceQueryParams: map[string]string{
				"operations_insights_warehouse_id": "id",
			},
		},
		{
			TerraformResourceHints: exportOpsiOperationsInsightsWarehouseUserHints,
			DatasourceQueryParams: map[string]string{
				"operations_insights_warehouse_id": "id",
			},
		},
	},
}
