---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster"
sidebar_current: "docs-oci-datasource-managed_kafka-kafka_cluster"
description: |-
  Provides details about a specific Kafka Cluster in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_kafka_cluster
This data source provides details about a specific Kafka Cluster resource in Oracle Cloud Infrastructure Managed Kafka service.

Gets information about a KafkaCluster.

## Example Usage

```hcl
data "oci_managed_kafka_kafka_cluster" "test_kafka_cluster" {
	#Required
	kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `kafka_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.


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

