---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_config"
sidebar_current: "docs-oci-resource-managed_kafka-kafka_cluster_config"
description: |-
  Provides the Kafka Cluster Config resource in Oracle Cloud Infrastructure Managed Kafka service
---

# oci_managed_kafka_kafka_cluster_config
This resource provides the Kafka Cluster Config resource in Oracle Cloud Infrastructure Managed Kafka service.

Creates a KafkaClusterConfig.


## Example Usage

```hcl
resource "oci_managed_kafka_kafka_cluster_config" "test_kafka_cluster_config" {
	#Required
	compartment_id = var.compartment_id
	latest_config {
		#Required
		properties = var.kafka_cluster_config_latest_config_properties

		#Optional
		config_id = oci_apm_config_config.test_config.id
		time_created = var.kafka_cluster_config_latest_config_time_created
		version_number = var.kafka_cluster_config_latest_config_version_number
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.kafka_cluster_config_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the KafkaClusterConfig in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `latest_config` - (Required) (Updatable) A shared configuration object used by 0 or more kafka clusters. 
	* `config_id` - (Optional) (Updatable) ID cluster configuration
	* `properties` - (Required) (Updatable) Cluster configuration key-value pairs
	* `time_created` - (Optional) (Updatable) The date and time the KafkaClusterConfigVersion was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
	* `version_number` - (Optional) (Updatable) Version of the cluster configuration


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Kafka Cluster Config
	* `update` - (Defaults to 20 minutes), when updating the Kafka Cluster Config
	* `delete` - (Defaults to 20 minutes), when destroying the Kafka Cluster Config


## Import

KafkaClusterConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config "id"
```

