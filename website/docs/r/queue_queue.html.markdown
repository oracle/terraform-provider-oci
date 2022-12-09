---
subcategory: "Queue"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_queue_queue"
sidebar_current: "docs-oci-resource-queue-queue"
description: |-
  Provides the Queue resource in Oracle Cloud Infrastructure Queue service
---

# oci_queue_queue
This resource provides the Queue resource in Oracle Cloud Infrastructure Queue service.

Creates a new Queue.


## Example Usage

```hcl
resource "oci_queue_queue" "test_queue" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.queue_display_name

	#Optional
	custom_encryption_key_id = oci_kms_key.test_key.id
	dead_letter_queue_delivery_count = var.queue_dead_letter_queue_delivery_count
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	retention_in_seconds = var.queue_retention_in_seconds
	timeout_in_seconds = var.queue_timeout_in_seconds
	visibility_in_seconds = var.queue_visibility_in_seconds
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `custom_encryption_key_id` - (Optional) (Updatable) Id of the custom master encryption key which will be used to encrypt messages content
* `dead_letter_queue_delivery_count` - (Optional) (Updatable) The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Queue Identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `retention_in_seconds` - (Optional) The retention period of the messages in the queue, in seconds.
* `timeout_in_seconds` - (Optional) (Updatable) The default polling timeout of the messages in the queue, in seconds.
* `visibility_in_seconds` - (Optional) (Updatable) The default visibility of the messages consumed from the queue.
* `purge_trigger` - (Optional) (Updatable) An optional property when incremented triggers Purge. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `custom_encryption_key_id` - Id of the custom master encryption key which will be used to encrypt messages content
* `dead_letter_queue_delivery_count` - The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Queue Identifier, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `messages_endpoint` - The endpoint to use to consume or publish messages in the queue.
* `retention_in_seconds` - The retention period of the messages in the queue, in seconds.
* `state` - The current state of the Queue.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Queue was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Queue was updated. An RFC3339 formatted datetime string
* `timeout_in_seconds` - The default polling timeout of the messages in the queue, in seconds.
* `visibility_in_seconds` - The default visibility of the messages consumed from the queue.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Queue
	* `update` - (Defaults to 20 minutes), when updating the Queue
	* `delete` - (Defaults to 20 minutes), when destroying the Queue


## Import

Queues can be imported using the `id`, e.g.

```
$ terraform import oci_queue_queue.test_queue "id"
```

