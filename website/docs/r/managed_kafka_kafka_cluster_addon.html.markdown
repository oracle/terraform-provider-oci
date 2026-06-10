---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_addon"
sidebar_current: "docs-oci-resource-managed_kafka-kafka_cluster_addon"
description: |-
  Provides the Kafka Cluster Addon resource in Oracle Cloud Infrastructure Managed Kafka service
---

# oci_managed_kafka_kafka_cluster_addon
This resource provides the Kafka Cluster Addon resource in Oracle Cloud Infrastructure Managed Kafka service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/kafka/latest/Addon

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/managed_kafka

Installs a KafkaClusterAddon.

## Example Usage

```hcl
resource "oci_managed_kafka_kafka_cluster_addon" "test_kafka_cluster_addon" {
	#Required
	addon_type = var.kafka_cluster_addon_addon_type
	authentication_mechanism = var.kafka_cluster_addon_authentication_mechanism
	kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id
	name = var.kafka_cluster_addon_name
	network_cidrs = var.kafka_cluster_addon_network_cidrs

	#Optional
	description = var.kafka_cluster_addon_description
}
```

## Argument Reference

The following arguments are supported:

* `addon_type` - (Required) (Updatable) This is Addon Type of Oracle Cloud Infrastructure kafka cluster
* `authentication_mechanism` - (Required) Authentication mechanism.
* `description` - (Optional) (Updatable) A brief description of the add on being installed.
* `kafka_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
* `name` - (Required) A unique user-friendly name. Avoid entering confidential information.
* `network_cidrs` - (Required) (Updatable) A list of CIDR's for ingress/egress traffic.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Kafka Cluster Addon
	* `update` - (Defaults to 20 minutes), when updating the Kafka Cluster Addon
	* `delete` - (Defaults to 20 minutes), when destroying the Kafka Cluster Addon


## Import

KafkaClusterAddons can be imported using the `id`, e.g.

```
$ terraform import oci_managed_kafka_kafka_cluster_addon.test_kafka_cluster_addon "kafkaClusters/{kafkaClusterId}/addons/{addonName}" 
```

