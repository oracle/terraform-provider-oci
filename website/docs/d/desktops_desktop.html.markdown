---
subcategory: "Desktops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_desktops_desktop"
sidebar_current: "docs-oci-datasource-desktops-desktop"
description: |-
  Provides details about a specific Desktop in Oracle Cloud Infrastructure Desktops service
---

# Data Source: oci_desktops_desktop
This data source provides details about a specific Desktop resource in Oracle Cloud Infrastructure Desktops service.

Provides information about the desktop with the specified OCID.

## Example Usage

```hcl
data "oci_desktops_desktop" "test_desktop" {
	#Required
	desktop_id = oci_desktops_desktop.test_desktop.id
}
```

## Argument Reference

The following arguments are supported:

* `desktop_id` - (Required) The OCID of the desktop.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
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
* `hosting_options` - Provides information about where a desktop is hosted.
	* `connect_address` - The connection address of the desktop.
	* `image` - Provides information about the desktop image.
		* `image_id` - The OCID of the desktop image.
		* `image_name` - The name of the desktop image.
* `id` - The OCID of the desktop.
* `pool_id` - The OCID of the desktop pool the desktop is a member of.
* `state` - The state of the desktop.
* `time_created` - The date and time the resource was created.
* `user_name` - The owner of the desktop.

