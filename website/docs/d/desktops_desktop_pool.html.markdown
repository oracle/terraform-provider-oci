---
subcategory: "Desktops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_desktops_desktop_pool"
sidebar_current: "docs-oci-datasource-desktops-desktop_pool"
description: |-
  Provides details about a specific Desktop Pool in Oracle Cloud Infrastructure Desktops service
---

# Data Source: oci_desktops_desktop_pool
This data source provides details about a specific Desktop Pool resource in Oracle Cloud Infrastructure Desktops service.

Returns information about the desktop pool including all configuration parameters and the current state.


## Example Usage

```hcl
data "oci_desktops_desktop_pool" "test_desktop_pool" {
	#Required
	desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `desktop_pool_id` - (Required) The OCID of the desktop pool.


## Attributes Reference

The following attributes are exported:

* `active_desktops` - The number of active desktops in the desktop pool.
* `are_privileged_users` - Indicates whether desktop pool users have administrative privileges on their desktop.
* `availability_domain` - The availability domain of the desktop pool.
* `availability_policy` - Provides the start and stop schedule information for desktop availability of the desktop pool.
	* `start_schedule` - Provides the schedule information for a desktop.
		* `cron_expression` - A cron expression describing the desktop's schedule.
		* `timezone` - The timezone of the desktop's schedule.
	* `stop_schedule` - Provides the schedule information for a desktop.
		* `cron_expression` - A cron expression describing the desktop's schedule.
		* `timezone` - The timezone of the desktop's schedule.
* `compartment_id` - The OCID of the compartment of the desktop pool.
* `contact_details` - Contact information of the desktop pool administrator. Avoid entering confidential information. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user friendly description providing additional information about the resource. Avoid entering confidential information. 
* `device_policy` - Provides the settings for desktop and client device options, such as audio in and out, client drive mapping, and clipboard access. 
	* `audio_mode` - The audio mode. NONE: No access to the local audio devices is permitted. TODESKTOP: The user may record audio on their desktop.  FROMDESKTOP: The user may play audio on their desktop. FULL: The user may play and record audio on their desktop. 
	* `cdm_mode` - The client local drive access mode. NONE: No access to local drives permitted. READONLY: The user may read from local drives on their desktop. FULL: The user may read from and write to their local drives on their desktop.  
	* `clipboard_mode` - The clipboard mode. NONE: No access to the local clipboard is permitted. TODESKTOP: The clipboard can be used to transfer data to the desktop only.  FROMDESKTOP: The clipboard can be used to transfer data from the desktop only. FULL: The clipboard can be used to transfer data to and from the desktop. 
	* `is_display_enabled` - Indicates whether the display is enabled.
	* `is_keyboard_enabled` - Indicates whether the keyboard is enabled.
	* `is_pointer_enabled` - Indicates whether the pointer is enabled.
	* `is_printing_enabled` - Indicates whether printing is enabled.
* `display_name` - A user friendly display name. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the desktop pool.
* `image` - Provides information about the desktop image.
	* `image_id` - The OCID of the desktop image.
	* `image_name` - The name of the desktop image.
* `is_storage_enabled` - Indicates whether storage is enabled for the desktop pool.
* `maximum_size` - The maximum number of desktops permitted in the desktop pool.
* `network_configuration` - Provides information about the network configuration of the desktop pool.
	* `subnet_id` - The OCID of the subnet to use for the desktop pool.
	* `vcn_id` - The OCID of the VCN used by the desktop pool.
* `nsg_ids` - A list of network security groups for the desktop pool.
* `shape_name` - The shape of the desktop pool.
* `standby_size` - The maximum number of standby desktops available in the desktop pool.
* `state` - The current state of the desktop pool.
* `storage_backup_policy_id` - The backup policy OCID of the storage.
* `storage_size_in_gbs` - The size in GBs of the storage for the desktop pool.
* `time_created` - The date and time the resource was created.
* `time_start_scheduled` - The start time of the desktop pool.
* `time_stop_scheduled` - The stop time of the desktop pool.

