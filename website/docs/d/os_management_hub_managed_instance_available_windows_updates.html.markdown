---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_available_windows_updates"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_available_windows_updates"
description: |-
  Provides the list of Managed Instance Available Windows Updates in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_available_windows_updates
This data source provides the list of Managed Instance Available Windows Updates in Oracle Cloud Infrastructure Os Management Hub service.

Returns a list of Windows updates that can be installed on the specified managed instance.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_available_windows_updates" "test_managed_instance_available_windows_updates" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	classification_type = var.managed_instance_available_windows_update_classification_type
	compartment_id = var.compartment_id
	display_name = var.managed_instance_available_windows_update_display_name
	display_name_contains = var.managed_instance_available_windows_update_display_name_contains
	is_installable = var.managed_instance_available_windows_update_is_installable
	name = var.managed_instance_available_windows_update_name
}
```

## Argument Reference

The following arguments are supported:

* `classification_type` - (Optional) A filter to return only packages that match the given update classification type.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `is_installable` - (Optional) Indicates if the update can be installed by the OS Management Hub service.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `name` - (Optional) A filter based on the unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.  Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 


## Attributes Reference

The following attributes are exported:

* `available_windows_update_collection` - The list of available_windows_update_collection.

### ManagedInstanceAvailableWindowsUpdate Reference

The following attributes are exported:

* `items` - List of available Windows updates.
	* `installable` - Indicates whether the update can be installed using the service.
	* `is_reboot_required_for_installation` - Indicates whether a reboot is required to complete the installation of this update.
	* `name` - Name of the Windows update.
	* `update_id` - Unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.  Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 
	* `update_type` - The type of Windows update.

