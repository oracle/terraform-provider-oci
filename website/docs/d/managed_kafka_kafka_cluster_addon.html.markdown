---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_addon"
sidebar_current: "docs-oci-datasource-managed_kafka-kafka_cluster_addon"
description: |-
  Provides details about a specific Kafka Cluster Addon in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_kafka_cluster_addon
This data source provides details about a specific Kafka Cluster Addon resource in Oracle Cloud Infrastructure Managed Kafka service.

Gets information about a KafkaClusterAddon.

## Example Usage

```hcl
data "oci_managed_kafka_kafka_cluster_addon" "test_kafka_cluster_addon" {
	#Required
	addon_name = oci_containerengine_addon.test_addon.name
	kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Required) The unique name of the KafkaClusterAddon.
* `kafka_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.


## Attributes Reference

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

