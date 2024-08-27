---
subcategory: "Announcements Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_announcements_service_announcement_subscriptions_filter_group"
sidebar_current: "docs-oci-resource-announcements_service-announcement_subscriptions_filter_group"
description: |-
  Provides the Announcement Subscriptions Filter Group resource in Oracle Cloud Infrastructure Announcements Service service
---

# oci_announcements_service_announcement_subscriptions_filter_group
This resource provides the Announcement Subscriptions Filter Group resource in Oracle Cloud Infrastructure Announcements Service service.

Creates a new filter group in the specified announcement subscription.

This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_announcements_service_announcement_subscriptions_filter_group" "test_announcement_subscriptions_filter_group" {
	#Required
	announcement_subscription_id = oci_announcements_service_announcement_subscription.test_announcement_subscription.id
	filters {
		#Required
		type = var.announcement_subscriptions_filter_group_filters_type
		value = var.announcement_subscriptions_filter_group_filters_value
	}
	name = var.announcement_subscriptions_filter_group_name
}
```

## Argument Reference

The following arguments are supported:

* `announcement_subscription_id` - (Required) The OCID of the announcement subscription.
* `filters` - (Required) (Updatable) A list of filters against which the Announcements service will match announcements. You cannot have more than one of any given filter type within a filter group.
	* `type` - (Required) (Updatable) The type of filter. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
	* `value` - (Required) (Updatable) The value of the filter.
* `name` - (Required) The name of the filter group. The name must be unique and it cannot be changed. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `filters` - A list of filters against which the Announcements service matches announcements. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
	* `type` - The type of filter. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
	* `value` - The value of the filter.
* `name` - The name of the group. The name must be unique and it cannot be changed. Avoid entering confidential information.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Announcement Subscriptions Filter Group
	* `update` - (Defaults to 20 minutes), when updating the Announcement Subscriptions Filter Group
	* `delete` - (Defaults to 20 minutes), when destroying the Announcement Subscriptions Filter Group


## Import

AnnouncementSubscriptionsFilterGroups can be imported using the `id`, e.g.

```
$ terraform import oci_announcements_service_announcement_subscriptions_filter_group.test_announcement_subscriptions_filter_group "id"
```

