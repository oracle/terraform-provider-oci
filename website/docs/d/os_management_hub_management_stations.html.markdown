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

Lists management stations within the specified compartment. Filter the list against a variety of criteria 
including but not limited to name, status, and location.


## Example Usage

```hcl
data "oci_os_management_hub_management_stations" "test_management_stations" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.management_station_display_name
	display_name_contains = var.management_station_display_name_contains
	id = var.management_station_id
	location = var.management_station_location
	location_not_equal_to = var.management_station_location_not_equal_to
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	state = var.management_station_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station. A filter that returns information about the specified management station.
* `location` - (Optional) A filter to return only resources whose location matches the given value.
* `location_not_equal_to` - (Optional) A filter to return only resources whose location does not match the given value.
* `managed_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.
* `state` - (Optional) A filter that returns information for management stations in the specified state.


## Attributes Reference

The following attributes are exported:

* `management_station_collection` - The list of management_station_collection.

### ManagementStation Reference

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
* `is_auto_config_enabled` - When enabled, the station setup script automatically runs to configure the firewall and SELinux settings on the station.
* `location` - The location of the instance that is acting as the management station.
* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance that is acting as the management station.
* `mirror` - Mirror information used for the management station configuration.
	* `directory` - Path to the data volume on the management station where software source mirrors are stored.
	* `is_sslverify_enabled` - When enabled, the SSL certificate is verified whenever an instance installs or updates a package from a software source that is mirrored on the management station.
	* `port` - Default mirror listening port for http.
	* `sslcert` - Path to the SSL cerfificate.
	* `sslport` - Default mirror listening port for https.
* `mirror_capacity` - A decimal number representing the amount of mirror capacity used by the sync.
* `mirror_package_count` - The total number of all packages within the mirrored software sources.
* `mirror_size` - The total size of all software source mirrors in bytes.
* `mirror_storage_available_size` - Amount of available mirror storage in bytes.
* `mirror_storage_size` - Total mirror storage size in bytes.
* `mirror_sync_status` - Status summary of the mirror sync.
	* `failed` - Total number of software sources that failed to sync.
	* `queued` - Total number of software sources that are queued for sync.
	* `synced` - Total number of software sources that successfully synced.
	* `syncing` - Total number of software sources currently syncing.
	* `unsynced` - Total number of software sources that have not yet been synced.
* `mirror_unique_package_count` - The total number of unique packages within the mirrored software sources on the station. Each package is counted only once, regardless of how many versions it has.
* `overall_percentage` - A decimal number representing the progress of the current mirror sync.
* `overall_state` - Current state of the mirror sync for the management station.
* `peer_management_stations` - A list of other management stations that are behind the same load balancer within a high availability configuration. Stations are identified as peers if they have the same hostname and compartment.
	* `display_name` - User-friendly name for the management station.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
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

