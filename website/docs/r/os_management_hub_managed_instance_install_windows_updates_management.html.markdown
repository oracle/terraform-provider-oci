---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_install_windows_updates_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_install_windows_updates_management"
description: |-
  Provides the Managed Instance Install Windows Updates Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_install_windows_updates_management
This resource provides the Managed Instance Install Windows Updates Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Installs Windows updates on the specified managed instance.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_install_windows_updates_management" "test_managed_instance_install_windows_updates_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	windows_update_name = oci_os_management_hub_windows_update.test_windows_update.name
	windows_update_types = var.managed_instance_install_windows_updates_management_windows_update_types
	work_request_details {

		#Optional
		description = var.managed_instance_install_windows_updates_management_work_request_details_description
		display_name = var.managed_instance_install_windows_updates_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `windows_update_name` - (Optional) The list of Windows update unique identifiers.  Note that this is not an OCID, but is a unique identifier assigned by Microsoft. Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 
* `windows_update_types` - (Optional) The types of Windows updates to be installed.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Install Windows Updates Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Install Windows Updates Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Install Windows Updates Management


## Import

ManagedInstanceInstallWindowsUpdatesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_install_windows_updates_management.test_managed_instance_install_windows_updates_management "id"
```

