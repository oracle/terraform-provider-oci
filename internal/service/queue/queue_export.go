package queue

import (
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("queue", queueResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportQueueQueueHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_queue_queue",
	DatasourceClass:        "oci_queue_queues",
	DatasourceItemsAttr:    "queue_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "queue",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_queue.QueueLifecycleStateActive),
	},
}

var queueResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportQueueQueueHints},
	},
}
