---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station_mirrors"
sidebar_current: "docs-oci-datasource-os_management_hub-management_station_mirrors"
description: |-
  Provides the list of Management Station Mirrors in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_management_station_mirrors
This data source provides the list of Management Station Mirrors in Oracle Cloud Infrastructure Os Management Hub service.

Lists all software source mirrors associated with a specified management station.


## Example Usage

```hcl
data "oci_os_management_hub_management_station_mirrors" "test_management_station_mirrors" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id

	#Optional
	display_name = var.management_station_mirror_display_name
	display_name_contains = var.management_station_mirror_display_name_contains
	mirror_states = var.management_station_mirror_mirror_states
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
* `mirror_states` - (Optional) List of Mirror state to filter by


## Attributes Reference

The following attributes are exported:

* `mirrors_collection` - The list of mirrors_collection.

### ManagementStationMirror Reference

The following attributes are exported:

* `items` - List of mirrors
	* `arch_type` - The architecture type supported by the software source.
	* `display_name` - Display name of the mirror.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `log` - The current log from the management station plugin.
	* `os_family` - The OS family of the software source.
	* `package_count` - The number of packages within the mirrored software source.
	* `percentage` - A decimal number representing the percentage of the software source that has been synced.
	* `size` - The size the mirrored software source in bytes.
	* `state` - Current state of the software source mirror.
	* `time_last_synced` - Time that the software source was last synced (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
	* `type` - Type of software source.

