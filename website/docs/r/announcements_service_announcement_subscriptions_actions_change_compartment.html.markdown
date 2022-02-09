---
subcategory: "Announcements Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_announcements_service_announcement_subscriptions_actions_change_compartment"
sidebar_current: "docs-oci-resource-announcements_service-announcement_subscriptions_actions_change_compartment"
description: |-
  Provides the Announcement Subscriptions Actions Change Compartment resource in Oracle Cloud Infrastructure Announcements Service service
---

# oci_announcements_service_announcement_subscriptions_actions_change_compartment
This resource provides the Announcement Subscriptions Actions Change Compartment resource in Oracle Cloud Infrastructure Announcements Service service.

Moves the specified announcement subscription from one compartment to another compartment. When provided, If-Match is checked against ETag values of the resource.

This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_announcements_service_announcement_subscriptions_actions_change_compartment" "test_announcement_subscriptions_actions_change_compartment" {
	#Required
	announcement_subscription_id = oci_announcements_service_announcement_subscription.test_announcement_subscription.id
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `announcement_subscription_id` - (Required) The OCID of the announcement subscription.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment into which you want to move the announcement subscription. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Announcement Subscriptions Actions Change Compartment
	* `update` - (Defaults to 20 minutes), when updating the Announcement Subscriptions Actions Change Compartment
	* `delete` - (Defaults to 20 minutes), when destroying the Announcement Subscriptions Actions Change Compartment


## Import

AnnouncementSubscriptionsActionsChangeCompartment can be imported using the `id`, e.g.

```
$ terraform import oci_announcements_service_announcement_subscriptions_actions_change_compartment.test_announcement_subscriptions_actions_change_compartment "id"
```

