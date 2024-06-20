package globally_distributed_database

import (
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("globally_distributed_database", globallyDistributedDatabaseResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGloballyDistributedDatabaseShardedDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_globally_distributed_database_sharded_database",
	DatasourceClass:        "oci_globally_distributed_database_sharded_databases",
	DatasourceItemsAttr:    "sharded_database_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sharded_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateActive),
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateInactive),
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateNeedsAttention),
	},
}

var exportGloballyDistributedDatabasePrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_globally_distributed_database_private_endpoint",
	DatasourceClass:        "oci_globally_distributed_database_private_endpoints",
	DatasourceItemsAttr:    "private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_globally_distributed_database.PrivateEndpointLifecycleStateActive),
	},
}

var globallyDistributedDatabaseResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGloballyDistributedDatabaseShardedDatabaseHints},
		{TerraformResourceHints: exportGloballyDistributedDatabasePrivateEndpointHints},
	},
}
