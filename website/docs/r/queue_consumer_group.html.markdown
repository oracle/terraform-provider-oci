---
subcategory: "Queue"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_queue_consumer_group"
sidebar_current: "docs-oci-resource-queue-consumer_group"
description: |-
  Provides the Consumer Group resource in Oracle Cloud Infrastructure Queue service
---

# oci_queue_consumer_group
This resource provides the Consumer Group resource in Oracle Cloud Infrastructure Queue service.

Creates a new consumer group.


## Example Usage

```hcl
resource "oci_queue_consumer_group" "test_consumer_group" {
	#Required
	display_name = var.consumer_group_display_name
	queue_id = oci_queue_queue.test_queue.id

	#Optional
	consumer_group_filter = var.consumer_group_consumer_group_filter
	dead_letter_queue_delivery_count = var.consumer_group_dead_letter_queue_delivery_count
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_enabled = var.consumer_group_is_enabled
}
```

## Argument Reference

The following arguments are supported:

* `consumer_group_filter` - (Optional) (Updatable) The filter used by the consumer group. Only messages matching the filter will be available by consumers of the group. The primary consumer group cannot have any filter.
* `dead_letter_queue_delivery_count` - (Optional) (Updatable) The number of times a message can be delivered to a consumer before being moved to the dead letter queue.  A value of 0 indicates that the DLQ is not used. If the value isn't specified, it will be using the value defined at the queue level.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - (Required) (Updatable) The user-friendly name of the consumer group.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `is_enabled` - (Optional) (Updatable) Used to enable or disable the consumer group.  An enabled consumer group will have a lifecycle state of ACTIVE, while a disabled will have its state as INACTIVE.
* `queue_id` - (Required) The OCID of the associated queue.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `consumer_group_filter` - The filter used by the consumer group. Only messages matching the filter will be available by consumers of the consumer group. An empty value means that all messages will be available in the group. The primary consumer group cannot have any filter.
* `dead_letter_queue_delivery_count` - The number of times a message can be delivered to a consumer before being moved to the dead letter queue.  A value of 0 indicates that the DLQ is not used. If the value isn't set, it will be using the value defined at the queue level.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - A user-friendly name for the consumer group. It has to be unique within the same queue in a case-insensitive manner. It's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - A unique identifier for the consumer group that is immutable on creation.
* `lifecycle_details` - Any additional details about the current state of the consumer group.
* `queue_id` - The OCID of the associated queue.
* `state` - The current state of the consumer group.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time that the consumer group was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z`
* `time_updated` - The time that the consumer group was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Consumer Group
* `update` - (Defaults to 20 minutes), when updating the Consumer Group
* `delete` - (Defaults to 20 minutes), when destroying the Consumer Group


## Import

ConsumerGroups can be imported using the `id`, e.g.

```
$ terraform import oci_queue_consumer_group.test_consumer_group "id"
```
