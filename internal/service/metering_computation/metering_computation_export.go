package metering_computation

import (
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterTenancyGraphs("metering_computation", meteringComputationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportMeteringComputationQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_query",
	DatasourceClass:        "oci_metering_computation_queries",
	IsDatasourceCollection: true,
	DatasourceItemsAttr:    "query_collection",
	ResourceAbbreviation:   "query",
	RequireResourceRefresh: true,
}

var exportMeteringComputationCustomTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_custom_table",
	DatasourceClass:        "oci_metering_computation_custom_tables",
	DatasourceItemsAttr:    "custom_table_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "custom_table",
	RequireResourceRefresh: true,
}

var exportMeteringComputationScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_schedule",
	DatasourceClass:        "oci_metering_computation_schedules",
	DatasourceItemsAttr:    "schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_metering_computation.ScheduleLifecycleStateActive),
	},
}

var meteringComputationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportMeteringComputationQueryHints},
		{TerraformResourceHints: exportMeteringComputationScheduleHints},
	},
}
