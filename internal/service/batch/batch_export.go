package batch

import (
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("batch", batchResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportBatchBatchTaskEnvironmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_batch_batch_task_environment",
	DatasourceClass:        "oci_batch_batch_task_environments",
	DatasourceItemsAttr:    "batch_task_environment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "batch_task_environment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_batch.BatchTaskEnvironmentLifecycleStateActive),
	},
}

var exportBatchBatchTaskProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_batch_batch_task_profile",
	DatasourceClass:        "oci_batch_batch_task_profiles",
	DatasourceItemsAttr:    "batch_task_profile_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "batch_task_profile",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_batch.BatchTaskProfileLifecycleStateActive),
	},
}

var exportBatchBatchContextHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_batch_batch_context",
	DatasourceClass:        "oci_batch_batch_contexts",
	DatasourceItemsAttr:    "batch_context_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "batch_context",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_batch.BatchContextLifecycleStateActive),
		string(oci_batch.BatchContextLifecycleStateNeedsAttention),
	},
}

var exportBatchBatchJobPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_batch_batch_job_pool",
	DatasourceClass:        "oci_batch_batch_job_pools",
	DatasourceItemsAttr:    "batch_job_pool_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "batch_job_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_batch.BatchJobPoolLifecycleStateActive),
		string(oci_batch.BatchJobPoolLifecycleStateNeedsAttention),
	},
}

var batchResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBatchBatchTaskEnvironmentHints},
		{TerraformResourceHints: exportBatchBatchTaskProfileHints},
		{TerraformResourceHints: exportBatchBatchContextHints},
		{TerraformResourceHints: exportBatchBatchJobPoolHints},
	},
}
