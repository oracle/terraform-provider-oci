package managed_kafka

import (
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("managed_kafka", managedKafkaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportManagedKafkaKafkaClusterConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_managed_kafka_kafka_cluster_config",
	DatasourceClass:        "oci_managed_kafka_kafka_cluster_configs",
	DatasourceItemsAttr:    "kafka_cluster_config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "kafka_cluster_config",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_managed_kafka.KafkaClusterConfigLifecycleStateActive),
	},
}

var exportManagedKafkaKafkaClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_managed_kafka_kafka_cluster",
	DatasourceClass:        "oci_managed_kafka_kafka_clusters",
	DatasourceItemsAttr:    "kafka_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "kafka_cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_managed_kafka.KafkaClusterLifecycleStateActive),
	},
}

var managedKafkaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportManagedKafkaKafkaClusterConfigHints},
		{TerraformResourceHints: exportManagedKafkaKafkaClusterHints},
	},
}
