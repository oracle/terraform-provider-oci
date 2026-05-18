---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_addons"
sidebar_current: "docs-oci-datasource-managed_kafka-kafka_cluster_addons"
description: |-
  Provides the list of Kafka Cluster Addons in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_kafka_cluster_addons
This data source provides the list of Kafka Cluster Addons in Oracle Cloud Infrastructure Managed Kafka service.

Gets a list of KafkaClusterAddons.

## Example Usage

```hcl
data "oci_managed_kafka_kafka_cluster_addons" "test_kafka_cluster_addons" {
	#Required
	kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id

	#Optional
	name = var.kafka_cluster_addon_name
	state = var.kafka_cluster_addon_state
}
```

## Argument Reference

The following arguments are supported:

* `kafka_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
* `name` - (Optional) The name to filter on.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `addon_collection` - The list of addon_collection.

### KafkaClusterAddon Reference

The following attributes are exported:

* `addon_type` - The type of addon
* `authentication_mechanism` - Authentication mechanism.
* `bootstrap_url` - The bootstrap url of the kafka cluster.
* `description` - Description of the add on
* `name` - A unique user-friendly name.
* `network_cidrs` - A list of CIDR ranges for ingress/egress traffic.
* `state` - The current state of the KafkaCluster.
* `time_created` - The time the addon was created.
* `time_updated` - The time the addon was updated.

