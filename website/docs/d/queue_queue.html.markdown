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

Gets a Queue by identifier

## Example Usage

```hcl
data "oci_queue_queue" "test_queue" {
	#Required
	queue_id = oci_queue_queue.test_queue.id
}
```

## Argument Reference

The following arguments are supported:

* `queue_id` - (Required) unique Queue identifier


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

