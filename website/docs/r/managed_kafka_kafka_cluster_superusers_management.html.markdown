---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_kafka_cluster_superusers_management"
sidebar_current: "docs-oci-resource-managed_kafka-kafka_cluster_superusers_management"
description: |-
  Provides the Kafka Cluster Superusers Management resource in Oracle Cloud Infrastructure Managed Kafka service
---

# oci_managed_kafka_kafka_cluster_superusers_management
This resource provides the Kafka Cluster Superusers Management resource in Oracle Cloud Infrastructure Managed Kafka service.

Adds a SASL superuser to the Kafka cluster. A generated password will be updated to the specified vault.


## Example Usage

```hcl
resource "oci_managed_kafka_kafka_cluster_superusers_management" "test_kafka_cluster_superusers_management" {
	#Required
	kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id
	enable_superuser = var.enable_superuser

	#Optional
	compartment_id = var.compartment_id
	secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the superuser secret.
* `kafka_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
* `secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret that will be populated with the generated superuser password. 
* `enable_superuser` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Kafka Cluster Superusers Management
	* `update` - (Defaults to 20 minutes), when updating the Kafka Cluster Superusers Management
	* `delete` - (Defaults to 20 minutes), when destroying the Kafka Cluster Superusers Management
