package distributed_database

import (
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("distributed_database", distributedDatabaseResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDistributedDatabaseDistributedDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_distributed_database_distributed_database",
	DatasourceClass:        "oci_distributed_database_distributed_databases",
	DatasourceItemsAttr:    "distributed_database_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "distributed_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_distributed_database.DistributedDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedDatabaseLifecycleStateNeedsAttention),
	},
}

var exportDistributedDatabaseDistributedDatabasePrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_distributed_database_distributed_database_private_endpoint",
	DatasourceClass:        "oci_distributed_database_distributed_database_private_endpoints",
	DatasourceItemsAttr:    "distributed_database_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "distributed_database_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateActive),
	},
}

var exportDistributedDatabaseDistributedAutonomousDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_distributed_database_distributed_autonomous_database",
	DatasourceClass:        "oci_distributed_database_distributed_autonomous_databases",
	DatasourceItemsAttr:    "distributed_autonomous_database_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "distributed_autonomous_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive),
		string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateNeedsAttention),
	},
}

var distributedDatabaseResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDistributedDatabaseDistributedDatabaseHints},
		{TerraformResourceHints: exportDistributedDatabaseDistributedDatabasePrivateEndpointHints},
		{TerraformResourceHints: exportDistributedDatabaseDistributedAutonomousDatabaseHints},
	},
}
