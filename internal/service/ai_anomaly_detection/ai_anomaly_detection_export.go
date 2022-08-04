package ai_anomaly_detection

import (
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ai_anomaly_detection", aiAnomalyDetectionResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAiAnomalyDetectionDataAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_anomaly_detection_data_asset",
	DatasourceClass:        "oci_ai_anomaly_detection_data_assets",
	DatasourceItemsAttr:    "data_asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.DataAssetLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_anomaly_detection_model",
	DatasourceClass:        "oci_ai_anomaly_detection_models",
	DatasourceItemsAttr:    "model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_anomaly_detection_project",
	DatasourceClass:        "oci_ai_anomaly_detection_projects",
	DatasourceItemsAttr:    "project_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "project",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.ProjectLifecycleStateActive),
	},
}

var exportAiAnomalyDetectionAiPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_anomaly_detection_ai_private_endpoint",
	DatasourceClass:        "oci_ai_anomaly_detection_ai_private_endpoints",
	DatasourceItemsAttr:    "ai_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ai_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateActive),
	},
}

var aiAnomalyDetectionResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiAnomalyDetectionDataAssetHints},
		{TerraformResourceHints: exportAiAnomalyDetectionModelHints},
		{TerraformResourceHints: exportAiAnomalyDetectionProjectHints},
		{TerraformResourceHints: exportAiAnomalyDetectionAiPrivateEndpointHints},
	},
}
