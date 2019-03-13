---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_notification_topic"
sidebar_current: "docs-oci-datasource-ons-notification_topic"
description: |-
  Provides details about a specific Notification Topic in Oracle Cloud Infrastructure Ons service
---

# Data Source: oci_ons_notification_topic
This data source provides details about a specific Notification Topic resource in Oracle Cloud Infrastructure Ons service.

Gets the specified topic's configuration information.


## Example Usage

```hcl
data "oci_ons_notification_topic" "test_notification_topic" {
	#Required
	topic_id = "${oci_ons_notification_topic.test_notification_topic.id}"
}
```

## Argument Reference

The following arguments are supported:

* `topic_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic to retrieve. 


## Attributes Reference

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

