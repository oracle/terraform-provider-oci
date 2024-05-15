---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station"
sidebar_current: "docs-oci-resource-os_management_hub-management_station"
description: |-
  Provides the Management Station resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_management_station
This resource provides the Management Station resource in Oracle Cloud Infrastructure Os Management Hub service.

Create a management station. You must provide proxy and mirror configuration information.


## Example Usage

```hcl
resource "oci_os_management_hub_management_station" "test_management_station" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.management_station_display_name
	hostname = var.management_station_hostname
	mirror {
		#Required
		directory = var.management_station_mirror_directory
		port = var.management_station_mirror_port
		sslport = var.management_station_mirror_sslport

		#Optional
		sslcert = var.management_station_mirror_sslcert
	}
	proxy {
		#Required
		is_enabled = var.management_station_proxy_is_enabled

		#Optional
		forward = var.management_station_proxy_forward
		hosts = var.management_station_proxy_hosts
		port = var.management_station_proxy_port
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.management_station_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the management station.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) User-specified description of the management station. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) User-friendly name for the management station. Does not have to be unique and you can change the name later. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `hostname` - (Required) (Updatable) Hostname of the management station.
* `mirror` - (Required) (Updatable) Information used to create the mirror configuration for a management station.
	* `directory` - (Required) (Updatable) Path to the data volume on the management station where software source mirrors are stored.
	* `port` - (Required) (Updatable) Default mirror listening port for http.
	* `sslcert` - (Optional) (Updatable) Path to the SSL cerfificate.
	* `sslport` - (Required) (Updatable) Default mirror listening port for https.
* `proxy` - (Required) (Updatable) Information used to create the proxy configuration for a management station.
	* `forward` - (Optional) (Updatable) The URL the proxy will forward to.
	* `hosts` - (Optional) (Updatable) List of hosts.
	* `is_enabled` - (Required) (Updatable) Indicates if the proxy should be enabled or disabled. Default is enabled.
	* `port` - (Optional) (Updatable) Listening port used for the proxy.
* `refresh_trigger` - (Optional) (Updatable) An optional property when incremented triggers Refresh. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Station
	* `update` - (Defaults to 20 minutes), when updating the Management Station
	* `delete` - (Defaults to 20 minutes), when destroying the Management Station


## Import

ManagementStations can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_management_station.test_management_station "id"
```

