---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_config_version"
sidebar_current: "docs-oci-datasource-managed_kafka-kafka_cluster_config_version"
description: |-
  Provides details about a specific Kafka Cluster Config Version in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_kafka_cluster_config_version
This data source provides details about a specific Kafka Cluster Config Version resource in Oracle Cloud Infrastructure Managed Kafka service.

Gets information about a KafkaClusterConfig version.

## Example Usage

```hcl
data "oci_managed_kafka_kafka_cluster_config_version" "test_kafka_cluster_config_version" {
	#Required
	kafka_cluster_config_id = oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id
	version_number = var.kafka_cluster_config_version_version_number
}
```

## Argument Reference

The following arguments are supported:

* `kafka_cluster_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
* `version_number` - (Required) The versionNumber of the KafkaClusterConfig.


## Attributes Reference

The following attributes are exported:

* `config_id` - ID cluster configuration
* `properties` - Cluster configuration key-value pairs
* `time_created` - The date and time the KafkaClusterConfigVersion was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `version_number` - Version of the cluster configuration

