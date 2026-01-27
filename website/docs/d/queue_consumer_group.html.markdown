---
subcategory: "Queue"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_queue_consumer_group"
sidebar_current: "docs-oci-datasource-queue-consumer_group"
description: |-
  Provides details about a specific Consumer Group in Oracle Cloud Infrastructure Queue service
---

# Data Source: oci_queue_consumer_group
This data source provides details about a specific Consumer Group resource in Oracle Cloud Infrastructure Queue service.

Gets a consumer group by identifier.

## Example Usage

```hcl
data "oci_queue_consumer_group" "test_consumer_group" {
	#Required
	consumer_group_id = oci_queue_consumer_group.test_consumer_group.id
}
```

## Argument Reference

The following arguments are supported:

* `consumer_group_id` - (Required) The unique consumer group identifier.


## Attributes Reference

The following attributes are exported:

* `consumer_group_filter` - The filter used by the consumer group. Only messages matching the filter will be available by consumers of the group. An empty value means that all messages will be available in the group. The primary consumer group cannot have any filter.
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

