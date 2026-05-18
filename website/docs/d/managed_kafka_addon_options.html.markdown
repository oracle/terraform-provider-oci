---
subcategory: "Managed Kafka"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_managed_kafka_addon_options"
sidebar_current: "docs-oci-datasource-managed_kafka-addon_options"
description: |-
  Provides the list of Addon Options in Oracle Cloud Infrastructure Managed Kafka service
---

# Data Source: oci_managed_kafka_addon_options
This data source provides the list of Addon Options in Oracle Cloud Infrastructure Managed Kafka service.

Gets a list of supported KafkaClusterAddons.

## Example Usage

```hcl
data "oci_managed_kafka_addon_options" "test_addon_options" {

	#Optional
	compartment_id = var.compartment_id
	name = var.addon_option_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `name` - (Optional) The name to filter on.


## Attributes Reference

The following attributes are exported:

* `addon_option_collection` - The list of addon_option_collection.

### AddonOption Reference

The following attributes are exported:

* `items` - List of kafka Cluster AddonOptions.
	* `name` - A user-friendly name.
	* `state` - The current state of the KafkaClusterAddon.

