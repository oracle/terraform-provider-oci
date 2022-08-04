package data_labeling_service

import (
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v65/datalabelingservice"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("data_labeling_service", dataLabelingServiceResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDataLabelingServiceDatasetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_labeling_service_dataset",
	DatasourceClass:        "oci_data_labeling_service_datasets",
	DatasourceItemsAttr:    "dataset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dataset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_labeling_service.DatasetLifecycleStateActive),
		string(oci_data_labeling_service.DatasetLifecycleStateNeedsAttention),
	},
}

var dataLabelingServiceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataLabelingServiceDatasetHints},
	},
}
