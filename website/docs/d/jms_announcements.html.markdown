---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_announcements"
sidebar_current: "docs-oci-datasource-jms-announcements"
description: |-
  Provides the list of Announcements in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_announcements
This data source provides the list of Announcements in Oracle Cloud Infrastructure Jms service.

Return a list of AnnouncementSummary items

## Example Usage

```hcl
data "oci_jms_announcements" "test_announcements" {

	#Optional
	summary_contains = var.announcement_summary_contains
	time_end = var.announcement_time_end
	time_start = var.announcement_time_start
}
```

## Argument Reference

The following arguments are supported:

* `summary_contains` - (Optional) Filter the list with summary contains the given value. 
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `announcement_collection` - The list of announcement_collection.

### Announcement Reference

The following attributes are exported:

* `items` - List of AnnouncementSummary items
	* `key` - Unique id of the announcement
	* `summary` - Summary text of the announcement
	* `time_released` - Date time on which the announcement was released
	* `url` - URL to the announcement web page

