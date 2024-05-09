---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_update_all_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_update_all_packages_management"
description: |-
  Provides the Managed Instance Group Update All Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_update_all_packages_management
This resource provides the Managed Instance Group Update All Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Updates all packages on each managed instance in the specified managed instance group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_update_all_packages_management" "test_managed_instance_group_update_all_packages_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	update_types = var.managed_instance_group_update_all_packages_management_update_types
	work_request_details {

		#Optional
		description = var.managed_instance_group_update_all_packages_management_work_request_details_description
		display_name = var.managed_instance_group_update_all_packages_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `update_types` - (Optional) The type of updates to be applied.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Update All Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Update All Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Update All Packages Management


## Import

ManagedInstanceGroupUpdateAllPackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_update_all_packages_management.test_managed_instance_group_update_all_packages_management "id"
```

