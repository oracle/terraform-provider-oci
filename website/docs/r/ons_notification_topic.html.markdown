---
subcategory: "Notifications"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_notification_topic"
sidebar_current: "docs-oci-resource-ons-notification_topic"
description: |-
  Provides the Notification Topic resource in Oracle Cloud Infrastructure Notifications service
---

# oci_ons_notification_topic
This resource provides the Notification Topic resource in Oracle Cloud Infrastructure Notifications service.

Creates a topic in the specified compartment. For general information about topics, see
[Managing Topics and Subscriptions](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the topic to reside.
For information about access control and compartments, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You must specify a display name for the topic.

All Oracle Cloud Infrastructure resources, including topics, get an Oracle-assigned, unique ID called an
Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response. You can also
retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the resource in the
Console. For more information, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.


## Example Usage

```hcl
resource "oci_ons_notification_topic" "test_notification_topic" {
	#Required
	compartment_id = var.compartment_id
	name = var.notification_topic_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.notification_topic_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the topic being created. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `api_endpoint` - The endpoint for managing subscriptions or publishing messages to the topic. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the topic. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the topic.
* `etag` - For optimistic concurrency control. See `if-match`. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - The name of the topic. 
* `short_topic_id` - A unique short topic Id. This is used only for SMS subscriptions. 
* `state` - The lifecycle state of the topic. 
* `time_created` - The time the topic was created.
* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 15 minutes), when creating the Notification Topic
	* `update` - (Defaults to 15 minutes), when updating the Notification Topic
	* `delete` - (Defaults to 150 minutes), when destroying the Notification Topic


## Import

NotificationTopics can be imported using the `id`, e.g.

```
$ terraform import oci_ons_notification_topic.test_notification_topic "id"
```

