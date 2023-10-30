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

Creates a management station.


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

* `compartment_id` - (Required) The OCID of the tenancy containing the Management Station.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Details describing the Management Station config.
* `display_name` - (Required) (Updatable) Management Station name
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `hostname` - (Required) (Updatable) Name of the host
* `mirror` - (Required) (Updatable) Information for creating a mirror configuration
	* `directory` - (Required) (Updatable) Directory for the mirroring
	* `port` - (Required) (Updatable) Default port for the mirror
	* `sslcert` - (Optional) (Updatable) Local path for the sslcert
	* `sslport` - (Required) (Updatable) Default sslport for the mirror
* `proxy` - (Required) (Updatable) Information for creating a proxy configuration
	* `forward` - (Optional) (Updatable) URL that the proxy will forward to
	* `hosts` - (Optional) (Updatable) List of hosts
	* `is_enabled` - (Required) (Updatable) To enable or disable the proxy (default true)
	* `port` - (Optional) (Updatable) Port that the proxy will use


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

