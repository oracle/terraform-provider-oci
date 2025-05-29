---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_config"
sidebar_current: "docs-oci-datasource-managed_kafka-kafka_cluster_config"
description: |-
  Provides details about a specific Kafka Cluster Config in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_kafka_cluster_config
This data source provides details about a specific Kafka Cluster Config resource in Oracle Cloud Infrastructure Managed Kafka service.

Gets information about a KafkaClusterConfig.

## Example Usage

```hcl
data "oci_managed_kafka_kafka_cluster_config" "test_kafka_cluster_config" {
	#Required
	kafka_cluster_config_id = oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id
}
```

## Argument Reference

The following arguments are supported:

* `kafka_cluster_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
* `latest_config` - A shared configuration object used by 0 or more kafka clusters. 
	* `config_id` - ID cluster configuration
	* `properties` - Cluster configuration key-value pairs
	* `time_created` - The date and time the KafkaClusterConfigVersion was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
	* `version_number` - Version of the cluster configuration
* `lifecycle_details` - A message that describes the current state of the KafkaClusterConfig in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the KafkaClusterConfig.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the KafkaClusterConfig was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the KafkaClusterConfig was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

