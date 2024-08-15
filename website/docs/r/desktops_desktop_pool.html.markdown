---
subcategory: "Desktops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_desktops_desktop_pool"
sidebar_current: "docs-oci-resource-desktops-desktop_pool"
description: |-
  Provides the Desktop Pool resource in Oracle Cloud Infrastructure Desktops service
---

# oci_desktops_desktop_pool
This resource provides the Desktop Pool resource in Oracle Cloud Infrastructure Desktops service.

Creates a desktop pool with the given configuration parameters.

## Example Usage

```hcl
resource "oci_desktops_desktop_pool" "test_desktop_pool" {
	#Required
	are_privileged_users = var.desktop_pool_are_privileged_users
	availability_domain = var.desktop_pool_availability_domain
	availability_policy {
		#Required
		start_schedule {
			#Required
			cron_expression = "0 10 8 ? * 2"
			timezone = "America/Denver"
		}
		stop_schedule {
			#Required
			cron_expression = "0 20 18 ? * 6"
			timezone = "America/Denver"
		}
	}
	compartment_id = var.compartment_id
	contact_details = var.desktop_pool_contact_details
	device_policy {
		#Required
		audio_mode = var.desktop_pool_device_policy_audio_mode
		cdm_mode = var.desktop_pool_device_policy_cdm_mode
		clipboard_mode = var.desktop_pool_device_policy_clipboard_mode
		is_display_enabled = var.desktop_pool_device_policy_is_display_enabled
		is_keyboard_enabled = var.desktop_pool_device_policy_is_keyboard_enabled
		is_pointer_enabled = var.desktop_pool_device_policy_is_pointer_enabled
		is_printing_enabled = var.desktop_pool_device_policy_is_printing_enabled
	}
	display_name = var.desktop_pool_display_name
	image {
		#Required
		image_id = oci_core_image.test_image.id
		image_name = var.desktop_pool_image_image_name
	}
	is_storage_enabled = var.desktop_pool_is_storage_enabled
	maximum_size = var.desktop_pool_maximum_size
	network_configuration {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id
		vcn_id = oci_core_vcn.test_vcn.id
	}
	shape_name = "VM.Standard.E4.Flex"
	standby_size = var.desktop_pool_standby_size
	storage_backup_policy_id = "ocid1.volumebackuppolicy.oc1.xxxxyyyyyzzzz"
	storage_size_in_gbs = var.desktop_pool_storage_size_in_gbs

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.desktop_pool_description
	freeform_tags = {"Department"= "Finance"}
	nsg_ids = var.desktop_pool_nsg_ids
	time_start_scheduled = var.desktop_pool_time_start_scheduled
	time_stop_scheduled = var.desktop_pool_time_stop_scheduled
}
```

## Argument Reference

The following arguments are supported:

* `are_privileged_users` - (Required) Indicates whether desktop pool users have administrative privileges on their desktop.
* `availability_domain` - (Required) The availability domain of the desktop pool.
* `availability_policy` - (Required) (Updatable) Provides the start and stop schedule information for desktop availability of the desktop pool.
	* `start_schedule` - (Required) (Updatable) Provides the schedule information for a desktop.
		* `cron_expression` - (Required) (Updatable) A cron expression describing the desktop's schedule.
		* `timezone` - (Required) (Updatable) The timezone of the desktop's schedule.
	* `stop_schedule` - (Required) (Updatable) Provides the schedule information for a desktop.
		* `cron_expression` - (Required) (Updatable) A cron expression describing the desktop's schedule.
		* `timezone` - (Required) (Updatable) The timezone of the desktop's schedule.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment which will contain the desktop pool.
* `contact_details` - (Required) (Updatable) Contact information of the desktop pool administrator. Avoid entering confidential information. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user friendly description providing additional information about the resource. Avoid entering confidential information. 
* `device_policy` - (Required) (Updatable) Provides the settings for desktop and client device options, such as audio in and out, client drive mapping, and clipboard access. 
	* `audio_mode` - (Required) (Updatable) The audio mode. NONE: No access to the local audio devices is permitted. TODESKTOP: The user may record audio on their desktop.  FROMDESKTOP: The user may play audio on their desktop. FULL: The user may play and record audio on their desktop. 
	* `cdm_mode` - (Required) (Updatable) The client local drive access mode. NONE: No access to local drives permitted. READONLY: The user may read from local drives on their desktop. FULL: The user may read from and write to their local drives on their desktop.  
	* `clipboard_mode` - (Required) (Updatable) The clipboard mode. NONE: No access to the local clipboard is permitted. TODESKTOP: The clipboard can be used to transfer data to the desktop only.  FROMDESKTOP: The clipboard can be used to transfer data from the desktop only. FULL: The clipboard can be used to transfer data to and from the desktop. 
	* `is_display_enabled` - (Required) (Updatable) Indicates whether the display is enabled.
	* `is_keyboard_enabled` - (Required) (Updatable) Indicates whether the keyboard is enabled.
	* `is_pointer_enabled` - (Required) (Updatable) Indicates whether the pointer is enabled.
	* `is_printing_enabled` - (Required) (Updatable) Indicates whether printing is enabled.
* `display_name` - (Required) (Updatable) A user friendly display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `image` - (Required) Provides information about the desktop image.
	* `image_id` - (Required) The OCID of the desktop image.
	* `image_name` - (Required) The name of the desktop image.
* `is_storage_enabled` - (Required) Indicates whether storage is enabled for the desktop pool.
* `maximum_size` - (Required) (Updatable) The maximum number of desktops permitted in the desktop pool.
* `network_configuration` - (Required) Provides information about the network configuration of the desktop pool.
	* `subnet_id` - (Required) The OCID of the subnet to use for the desktop pool.
	* `vcn_id` - (Required) The OCID of the VCN used by the desktop pool.
* `nsg_ids` - (Optional) A list of network security groups for the desktop pool.
* `shape_name` - (Required) The shape of the desktop pool.
* `standby_size` - (Required) (Updatable) The maximum number of standby desktops available in the desktop pool.
* `storage_backup_policy_id` - (Required) The backup policy OCID of the storage.
* `storage_size_in_gbs` - (Required) The size in GBs of the storage for the desktop pool.
* `time_start_scheduled` - (Optional) (Updatable) The start time of the desktop pool.
* `time_stop_scheduled` - (Optional) (Updatable) The stop time of the desktop pool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Desktop Pool
	* `update` - (Defaults to 20 minutes), when updating the Desktop Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Desktop Pool


## Import

DesktopPools can be imported using the `id`, e.g.

```
$ terraform import oci_desktops_desktop_pool.test_desktop_pool "id"
```

