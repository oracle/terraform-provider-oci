---
subcategory: "Announcements Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_announcements_service_announcement_subscription"
sidebar_current: "docs-oci-resource-announcements_service-announcement_subscription"
description: |-
  Provides the Announcement Subscription resource in Oracle Cloud Infrastructure Announcements Service service
---

# oci_announcements_service_announcement_subscription
This resource provides the Announcement Subscription resource in Oracle Cloud Infrastructure Announcements Service service.

Creates a new announcement subscription.

This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_announcements_service_announcement_subscription" "test_announcement_subscription" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.announcement_subscription_display_name
	ons_topic_id = oci_ons_notification_topic.test_notification_topic.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.announcement_subscription_description
	filter_groups {
		#Required
		filters {
			#Required
			type = var.announcement_subscription_filter_groups_filters_type
			value = var.announcement_subscription_filter_groups_filters_value
		}
	}
	freeform_tags = {"bar-key"= "value"}
	preferred_language = var.announcement_subscription_preferred_language
	preferred_time_zone = var.announcement_subscription_preferred_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the announcement subscription. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A description of the announcement subscription. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) A user-friendly name for the announcement subscription. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `filter_groups` - (Optional) A list of filter groups for the announcement subscription. A filter group combines one or more filters that the Announcements service applies to announcements for matching purposes. 
	* `filters` - (Required) A list of filters against which the Announcements service matches announcements. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
		* `type` - (Required) The type of filter. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
		* `value` - (Required) The value of the filter.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `ons_topic_id` - (Required) (Updatable) The OCID of the Notifications service topic that is the target for publishing announcements that match the configured announcement subscription. The caller of the operation needs the ONS_TOPIC_PUBLISH permission for the targeted Notifications service topic. For more information about Notifications permissions, see [Details for Notifications](https://docs.cloud.oracle.com/iaas/Content/Identity/policyreference/notificationpolicyreference.htm). 
* `preferred_language` - (Optional) (Updatable) (For announcement subscriptions with SaaS configured as the platform type or Oracle Fusion Applications as the service, or both, only) The language in which the user prefers to receive emailed announcements. Specify the preference with a value that uses the x-obmcs-human-language format. For example fr-FR.
* `preferred_time_zone` - (Optional) (Updatable) The time zone in which the user prefers to receive announcements. Specify the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the announcement subscription.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the announcement subscription. Avoid entering confidential information.
* `display_name` - A user-friendly name for the announcement subscription. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `filter_groups` - A list of filter groups for the announcement subscription. A filter group is a combination of multiple filters applied to announcements for matching purposes. 
	* `filters` - A list of filters against which the Announcements service matches announcements. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
		* `type` - The type of filter. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
		* `value` - The value of the filter.
	* `name` - The name of the group. The name must be unique and it cannot be changed. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the announcement subscription.
* `lifecycle_details` - A message describing the current lifecycle state in more detail. For example, details might provide required or recommended actions for a resource in a Failed state. 
* `ons_topic_id` - The OCID of the Notifications service topic that is the target for publishing announcements that match the configured announcement subscription. 
* `preferred_language` - (For announcement subscriptions with SaaS configured as the platform type or Oracle Fusion Applications as the service, or both, only) The language in which the user prefers to receive emailed announcements. Specify the preference with a value that uses the x-obmcs-human-language format. For example fr-FR.
* `preferred_time_zone` - The time zone in which the user prefers to receive announcements. Specify the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles
* `state` - The current lifecycle state of the announcement subscription.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the announcement subscription was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time that the announcement subscription was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Announcement Subscription
	* `update` - (Defaults to 20 minutes), when updating the Announcement Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Announcement Subscription


## Import

AnnouncementSubscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_announcements_service_announcement_subscription.test_announcement_subscription "id"
```

