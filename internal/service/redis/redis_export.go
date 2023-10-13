package redis

import (
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("redis", redisResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportRedisRedisClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_redis_redis_cluster",
	DatasourceClass:        "oci_redis_redis_clusters",
	DatasourceItemsAttr:    "redis_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "redis_cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_redis.RedisClusterLifecycleStateActive),
	},
}

var redisResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportRedisRedisClusterHints},
	},
}
