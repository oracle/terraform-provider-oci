---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station"
sidebar_current: "docs-oci-datasource-os_management_hub-management_station"
description: |-
  Provides details about a specific Management Station in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_management_station
This data source provides details about a specific Management Station resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified management station.


## Example Usage

```hcl
data "oci_os_management_hub_management_station" "test_management_station" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
}
```

## Argument Reference

The following arguments are supported:

* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the management station.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the management station.
* `display_name` - A user-friendly name for the management station.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `health` - Overall health information of the management station.
	* `description` - Explanation of the health status.
	* `state` - Overall health status of the management station.
* `hostname` - Hostname of the management station.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance that is acting as the management station.
* `mirror` - Mirror information used for the management station configuration.
	* `directory` - Path to the data volume on the management station where software source mirrors are stored.
	* `port` - Default mirror listening port for http.
	* `sslcert` - Path to the SSL cerfificate.
	* `sslport` - Default mirror listening port for https.
* `mirror_capacity` - A decimal number representing the amount of mirror capacity used by the sync.
* `mirror_sync_status` - Status summary of the mirror sync.
	* `failed` - Total number of software sources that failed to sync.
	* `queued` - Total number of software sources that are queued for sync.
	* `synced` - Total number of software sources that successfully synced.
	* `syncing` - Total number of software sources currently syncing.
	* `unsynced` - Total number of software sources that have not yet been synced.
* `overall_percentage` - A decimal number representing the progress of the current mirror sync.
* `overall_state` - Current state of the mirror sync for the management station.
* `profile_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile used for the management station.
* `proxy` - Proxy information used for the management station configuration.
	* `forward` - The URL the proxy will forward to.
	* `hosts` - List of hosts.
	* `is_enabled` - Indicates if the proxy should be enabled or disabled. Default is enabled.
	* `port` - Listening port used for the proxy.
* `scheduled_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled job for the mirror sync.
* `state` - The current state of the management station.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `total_mirrors` - The number of software sources that the station is mirroring.

