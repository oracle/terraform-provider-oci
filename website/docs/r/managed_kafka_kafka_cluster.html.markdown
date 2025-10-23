---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster"
sidebar_current: "docs-oci-resource-managed_kafka-kafka_cluster"
description: |-
  Provides the Kafka Cluster resource in Oracle Cloud Infrastructure Managed Kafka service
---

# oci_managed_kafka_kafka_cluster
This resource provides the Kafka Cluster resource in Oracle Cloud Infrastructure Managed Kafka service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/kafka/latest/KafkaCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/managed_kafka

Creates a KafkaCluster.


## Example Usage

```hcl
resource "oci_managed_kafka_kafka_cluster" "test_kafka_cluster" {
	#Required
	access_subnets {
		#Required
		subnets = var.kafka_cluster_access_subnets_subnets
	}
	broker_shape {
		#Required
		node_count = var.kafka_cluster_broker_shape_node_count
		ocpu_count = var.kafka_cluster_broker_shape_ocpu_count

		#Optional
		storage_size_in_gbs = var.kafka_cluster_broker_shape_storage_size_in_gbs
	}
	cluster_config_id = oci_apm_config_config.test_config.id
	cluster_config_version = var.kafka_cluster_cluster_config_version
	cluster_type = var.kafka_cluster_cluster_type
	compartment_id = var.compartment_id
	coordination_type = var.kafka_cluster_coordination_type
	kafka_version = var.kafka_cluster_kafka_version

	#Optional
	client_certificate_bundle = var.kafka_cluster_client_certificate_bundle
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.kafka_cluster_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `access_subnets` - (Required) (Updatable) Subnets where broker/coordinator VNICs will be created.
	* `subnets` - (Required) (Updatable) Subnets OCIDs
* `broker_shape` - (Required) (Updatable) Configuration of the broker node.
	* `node_count` - (Required) (Updatable) Number of Kafka broker nodes
	* `ocpu_count` - (Required) (Updatable) Number of OCPUs per nodes
	* `storage_size_in_gbs` - (Optional) (Updatable) Size of the storage per nodes.
* `client_certificate_bundle` - (Optional) (Updatable) CA certificate bundle for mTLS broker authentication.
* `cluster_config_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Kafka Cluster configuration object
* `cluster_config_version` - (Required) (Updatable) The version of configuration object
* `cluster_type` - (Required) Type of the cluster to spin up.  DEVELOPMENT - setting that allows to sacrifice HA and spin up cluster on single node PRODUCTION - Minimum allowed broker count is 3 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the KafkaCluster in. 
* `coordination_type` - (Required) (Updatable) Kafka coordination type. Set of available types depends on Kafka version
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `kafka_version` - (Required) Version of Kafka to use to spin up the cluster


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_subnets` - Subnets where broker/coordinator VNICs will be created.
	* `subnets` - Subnets OCIDs
* `broker_shape` - Configuration of the broker node.
	* `node_count` - Number of Kafka broker nodes
	* `ocpu_count` - Number of OCPUs per nodes
	* `storage_size_in_gbs` - Size of the storage per nodes.
* `client_certificate_bundle` - CA certificate bundle for mTLS broker authentication.
* `cluster_config_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Kafka Cluster configuration object
* `cluster_config_version` - The version of configuration object
* `cluster_type` - Type of the cluster to spin up.  DEVELOPMENT - setting that allows to sacrifice HA and spin up cluster on a single node PRODUCTION - Minimum allowed broker count is 3 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `coordination_type` - Kafka coordination type. Set of available types depends on Kafka version
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
* `kafka_bootstrap_urls` - Bootstrap URL that can be used to connect to Kafka
	* `name` - Name of the Kafka listener providing this bootstrap URL
	* `url` - Bootstrap URL
* `kafka_version` - Version of Kafka to use to spin up the cluster
* `lifecycle_details` - A message that describes the current state of the KafkaCluster in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret that contains superuser password.
* `state` - The current state of the KafkaCluster.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the KafkaCluster was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the KafkaCluster was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Kafka Cluster
	* `update` - (Defaults to 20 minutes), when updating the Kafka Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Kafka Cluster


## Import

KafkaClusters can be imported using the `id`, e.g.

```
$ terraform import oci_managed_kafka_kafka_cluster.test_kafka_cluster "id"
```

