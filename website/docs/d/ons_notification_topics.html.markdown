---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_notification_topics"
sidebar_current: "docs-oci-datasource-ons-notification_topics"
description: |-
  Provides the list of Notification Topics in Oracle Cloud Infrastructure Ons service
---

# Data Source: oci_ons_notification_topics
This data source provides the list of Notification Topics in Oracle Cloud Infrastructure Ons service.

Lists topics in the specified compartment. 


## Example Usage

```hcl
data "oci_ons_notification_topics" "test_notification_topics" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	id = "${var.notification_topic_id}"
	name = "${var.notification_topic_name}"
	state = "${var.notification_topic_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `id` - (Optional) A filter to only return resources that match the given id exactly. 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) Filter returned list by specified lifecycle state. This parameter is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `notification_topics` - The list of notification_topics.

### NotificationTopic Reference

The following attributes are exported:

* `api_endpoint` - The endpoint for managing topic subscriptions or publishing messages to the topic. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the topic. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the topic. Avoid entering confidential information.
* `etag` - For optimistic concurrency control. See `if-match`. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - The name of the topic. Avoid entering confidential information. 
* `state` - The lifecycle state of the topic.  
* `time_created` - The time the topic was created.
* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic. 

