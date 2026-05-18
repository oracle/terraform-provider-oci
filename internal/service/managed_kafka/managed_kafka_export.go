package managed_kafka

import (
	"fmt"

	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportManagedKafkaKafkaClusterAddonHints.GetIdFn = getManagedKafkaKafkaClusterAddonId
	tf_export.RegisterCompartmentGraphs("managed_kafka", managedKafkaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getManagedKafkaKafkaClusterAddonId(resource *tf_export.OCIResource) (string, error) {

	addonName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find addonName for ManagedKafka KafkaClusterAddon")
	}
	kafkaClusterId := resource.Parent.Id
	return GetKafkaClusterAddonCompositeId(addonName, kafkaClusterId), nil
}

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

var exportManagedKafkaKafkaClusterAddonHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_managed_kafka_kafka_cluster_addon",
	DatasourceClass:        "oci_managed_kafka_kafka_cluster_addons",
	DatasourceItemsAttr:    "addon_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "kafka_cluster_addon",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_managed_kafka.KafkaClusterAddonLifecycleStateActive),
	},
}

var managedKafkaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportManagedKafkaKafkaClusterConfigHints},
		{TerraformResourceHints: exportManagedKafkaKafkaClusterHints},
	},
	"oci_managed_kafka_kafka_cluster": {
		{
			TerraformResourceHints: exportManagedKafkaKafkaClusterAddonHints,
			DatasourceQueryParams: map[string]string{
				"kafka_cluster_id": "id",
			},
		},
	},
}
