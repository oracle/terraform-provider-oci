---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_announcements"
sidebar_current: "docs-oci-datasource-fleet_apps_management-announcements"
description: |-
  Provides the list of Announcements in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_announcements
This data source provides the list of Announcements in Oracle Cloud Infrastructure Fleet Apps Management service.

Return a list of AnnouncementSummary items

## Example Usage

```hcl
data "oci_fleet_apps_management_announcements" "test_announcements" {

	#Optional
	display_name = var.announcement_display_name
	summary_contains = var.announcement_summary_contains
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `summary_contains` - (Optional) Filter the list with summary contains the given value. 


## Attributes Reference

The following attributes are exported:

* `announcement_collection` - The list of announcement_collection.

### Announcement Reference

The following attributes are exported:

* `items` - List of AnnouncementSummary items
	* `announcement_end` - Date to end displaying annonucement to user
	* `announcement_start` - Date to start displaying announcement to user
	* `compartment_id` - Tenancy OCID
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
	* `details` - Details of the announcement
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the resource.
	* `resource_region` - Associated region
	* `state` - The lifecycle state of the announcement.
	* `summary` - Summary of the announcement
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
	* `type` - Type of announcement
	* `url` - URL to the announcement

