---
subcategory: "Queue"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_queue_queue"
sidebar_current: "docs-oci-datasource-queue-queue"
description: |-
  Provides details about a specific Queue in Oracle Cloud Infrastructure Queue service
---

# Data Source: oci_queue_queue
This data source provides details about a specific Queue resource in Oracle Cloud Infrastructure Queue service.

Gets a queue by identifier.

## Example Usage

```hcl
data "oci_queue_queue" "test_queue" {
	#Required
	queue_id = oci_queue_queue.test_queue.id
}
```

## Argument Reference

The following arguments are supported:

* `queue_id` - (Required) The unique queue identifier.


## Attributes Reference

The following attributes are exported:

* `channel_consumption_limit` - The percentage of allocated queue resources that can be consumed by a single channel. For example, if a queue has a storage limit of 2Gb, and a single channel consumption limit is 0.1 (10%), that means data size of a single channel  can't exceed 200Mb. Consumption limit of 100% (default) means that a single channel can consume up-to all allocated queue's resources.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the queue.
* `custom_encryption_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom encryption key to be used to encrypt messages content.
* `dead_letter_queue_delivery_count` - The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name for the queue. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier for the queue that is immutable on creation.
* `lifecycle_details` - Any additional details about the current state of the queue.
* `messages_endpoint` - The endpoint to use to consume or publish messages in the queue.
* `retention_in_seconds` - The retention period of the messages in the queue, in seconds.
* `state` - The current state of the queue.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the queue was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 
* `time_updated` - The time that the queue was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 
* `timeout_in_seconds` - The default polling timeout of the messages in the queue, in seconds.
* `visibility_in_seconds` - The default visibility timeout of the messages consumed from the queue, in seconds.

