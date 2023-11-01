---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_stations"
sidebar_current: "docs-oci-datasource-os_management_hub-management_stations"
description: |-
  Provides the list of Management Stations in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_management_stations
This data source provides the list of Management Stations in Oracle Cloud Infrastructure Os Management Hub service.

Lists management stations in a compartment.


## Example Usage

```hcl
data "oci_os_management_hub_management_stations" "test_management_stations" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.management_station_display_name
	display_name_contains = var.management_station_display_name_contains
	id = var.management_station_id
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id
	state = var.management_station_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `id` - (Optional) The OCID of the management station.
* `managed_instance_id` - (Optional) The OCID of the managed instance for which to list resources.
* `state` - (Optional) The current lifecycle state for the object.


## Attributes Reference

The following attributes are exported:

* `management_station_collection` - The list of management_station_collection.

### ManagementStation Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the Management Station.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Details describing the ManagementStation config.
* `display_name` - ManagementStation name
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `hostname` - Name of the host
* `id` - OCID for the ManagementStation config
* `managed_instance_id` - OCID for the Instance associated with the Management Station.
* `mirror` - Information for a mirror configuration
	* `directory` - Directory for the mirroring
	* `port` - Default port for the mirror
	* `sslcert` - Local path for the sslcert
	* `sslport` - Default sslport for the mirror
* `mirror_capacity` - A decimal number representing the mirror capacity
* `mirror_sync_status` - Status summary of all repos
	* `failed` - Total of mirrors in 'failed' state
	* `queued` - Total of mirrors in 'queued' state
	* `synced` - Total of mirrors in 'synced' state
	* `syncing` - Total of mirrors in 'syncing' state
	* `unsynced` - Total of mirrors in 'failed' state
* `overall_percentage` - A decimal number representing the completeness percentage
* `overall_state` - Current state of the mirroring
* `profile_id` - OCID of the Profile associated with the Station
* `proxy` - Information for a proxy configuration
	* `forward` - URL that the proxy will forward to
	* `hosts` - List of hosts
	* `is_enabled` - To enable or disable the proxy (default true)
	* `port` - Port that the proxy will use
* `scheduled_job_id` - OCID of the Scheduled Job for mirror sync
* `state` - The current state of the Management Station config.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `total_mirrors` - A decimal number representing the total of repos

