---
subcategory: "Announcements Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_announcements_service_announcement_subscriptions"
sidebar_current: "docs-oci-datasource-announcements_service-announcement_subscriptions"
description: |-
  Provides the list of Announcement Subscriptions in Oracle Cloud Infrastructure Announcements Service service
---

# Data Source: oci_announcements_service_announcement_subscriptions
This data source provides the list of Announcement Subscriptions in Oracle Cloud Infrastructure Announcements Service service.

Gets a list of all announcement subscriptions in the specified compartment.

This call is subject to an Announcements limit that applies to the total number of requests across all read or write operations. Announcements might throttle this call to reject an otherwise valid request when the total rate of operations exceeds 20 requests per second for a given user. The service might also throttle this call to reject an otherwise valid request when the total rate of operations exceeds 100 requests per second for a given tenancy.


## Example Usage

```hcl
data "oci_announcements_service_announcement_subscriptions" "test_announcement_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.announcement_subscription_display_name
	id = var.announcement_subscription_id
	state = var.announcement_subscription_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The OCID of the announcement subscription.
* `state` - (Optional) A filter to return only announcement subscriptions that match the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `announcement_subscription_collection` - The list of announcement_subscription_collection.

### AnnouncementSubscription Reference

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

