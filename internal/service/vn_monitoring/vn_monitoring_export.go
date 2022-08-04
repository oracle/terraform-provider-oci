package vn_monitoring

import (
	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("vn_monitoring", vnMonitoringResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportVnMonitoringPathAnalyzerTestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_vn_monitoring_path_analyzer_test",
	DatasourceClass:        "oci_vn_monitoring_path_analyzer_tests",
	DatasourceItemsAttr:    "path_analyzer_test_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "path_analyzer_test",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_vn_monitoring.PathAnalyzerTestLifecycleStateActive),
	},
}

var exportVnMonitoringPathAnalysiHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_vn_monitoring_path_analysi",
	ResourceAbbreviation: "path_analysi",
}

var vnMonitoringResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVnMonitoringPathAnalyzerTestHints},
	},
}
